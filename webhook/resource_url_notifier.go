// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/hashicorp/go-retryablehttp"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
	"github.com/livekit/protocol/utils/protojson"
)

const (
	defaultTimeout = 5 * time.Minute

	// an endpoint that accepts and never answers must not hold a queue worker forever
	defaultClientTimeout = 10 * time.Second
	// only the status code of a response is used
	maxResponseHeaderBytes = 64 << 10
)

var (
	errClosed = errors.New("notifier is closed")
	errNoKey  = errors.New("no singing key or secret was provided")
)

type ResourceURLNotifierConfig struct {
	MaxAge   time.Duration `yaml:"max_age,omitempty"`
	MaxDepth int           `yaml:"max_depth,omitempty"`
}

var DefaultResourceURLNotifierConfig = ResourceURLNotifierConfig{
	MaxAge:   30 * time.Second,
	MaxDepth: 200,
}

type poster interface {
	Process(
		ctx context.Context,
		queuedAt time.Time,
		event *livekit.WebhookEvent,
		params *ResourceURLNotifierParams,
		qLen int,
	)
}

type resourceQueueInfo struct {
	*resourceQueue
	key string
	tqi *utils.TimeoutQueueItem[*resourceQueueInfo]
}

type ResourceURLNotifierParams struct {
	HTTPClientParams
	Logger     logger.Logger
	Timeout    time.Duration
	Config     ResourceURLNotifierConfig
	URL        string
	APIKey     string
	APISecret  string
	FieldsHook func(whi *livekit.WebhookInfo)
	EventKey   func(event *livekit.WebhookEvent) string
	FilterParams
}

// ResourceURLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It queues up events per resource (could be egress, ingress, room, participant, track, etc.)
// to avoid blocking events of one resource blocking another resource's event(s).
// It will retry on failure, and will drop events if notification fall too far behind,
// either in age or queue depth.
type ResourceURLNotifier struct {
	mu            sync.RWMutex
	params        ResourceURLNotifierParams
	client        *retryablehttp.Client
	processedHook func(ctx context.Context, whi *livekit.WebhookInfo)

	resourceQueues            map[string]*resourceQueueInfo
	resourceQueueTimeoutQueue utils.TimeoutQueue[*resourceQueueInfo]

	filter *filter

	closed core.Fuse
}

func NewResourceURLNotifier(params ResourceURLNotifierParams) *ResourceURLNotifier {
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}

	if params.Timeout == 0 {
		params.Timeout = defaultTimeout
	}
	if params.Config.MaxAge == 0 {
		params.Config.MaxAge = DefaultResourceURLNotifierConfig.MaxAge
	}
	if params.Config.MaxDepth == 0 {
		params.Config.MaxDepth = DefaultResourceURLNotifierConfig.MaxDepth
	}
	if params.UserAgent == "" {
		params.UserAgent = DefaultUserAgent
	}

	r := &ResourceURLNotifier{
		params:         params,
		client:         newHTTPClient(params.HTTPClientParams),
		resourceQueues: make(map[string]*resourceQueueInfo),
		filter:         newFilter(params.FilterParams),
	}

	go r.sweeper()
	return r
}

func newHTTPClient(params HTTPClientParams) *retryablehttp.Client {
	rhc := retryablehttp.NewClient()
	if params.RetryWaitMin > 0 {
		rhc.RetryWaitMin = params.RetryWaitMin
	}
	if params.RetryWaitMax > 0 {
		rhc.RetryWaitMax = params.RetryWaitMax
	}
	if params.MaxRetries > 0 {
		rhc.RetryMax = params.MaxRetries
	}
	rhc.HTTPClient.Timeout = defaultClientTimeout
	if params.ClientTimeout > 0 {
		rhc.HTTPClient.Timeout = params.ClientTimeout
	}
	rhc.Backoff = cappedBackoff
	rhc.ErrorHandler = redactedErrorHandler
	rhc.Logger = &logAdapter{}

	tr, ok := rhc.HTTPClient.Transport.(*http.Transport)
	if !ok || tr == nil {
		tr = http.DefaultTransport.(*http.Transport).Clone()
	}
	if params.ForceIPv4 || params.BlockPrivateAddresses {
		dialer := &net.Dialer{}
		if params.BlockPrivateAddresses {
			dialer.Control = GuardWebhookDial
		}
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			if params.ForceIPv4 {
				network = "tcp4"
			}
			return dialer.DialContext(ctx, network, addr)
		}
	}
	if params.BlockPrivateAddresses {
		tr.Proxy = nil
		// A redirect is reported as the receiver's answer, not followed.
		rhc.HTTPClient.CheckRedirect = func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	tr.MaxResponseHeaderBytes = maxResponseHeaderBytes
	rhc.HTTPClient.Transport = tr
	return rhc
}

// the webhook receiver sets Retry-After, so cap the wait at RetryWaitMax
func cappedBackoff(waitMin, waitMax time.Duration, attemptNum int, resp *http.Response) time.Duration {
	return min(retryablehttp.DefaultBackoff(waitMin, waitMax, attemptNum, resp), waitMax)
}

// the default handler repeats the request URL, which may carry the receiver's credentials
func redactedErrorHandler(resp *http.Response, err error, attempts int) (*http.Response, error) {
	if resp != nil {
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
		_ = resp.Body.Close()
	}
	if err == nil {
		return nil, fmt.Errorf("giving up after %d attempt(s)", attempts)
	}
	return nil, fmt.Errorf("giving up after %d attempt(s): %w", attempts, redactURLError(err))
}

func redactURLError(err error) error {
	if ue, ok := err.(*url.Error); ok {
		return &url.Error{Op: ue.Op, URL: redactURL(ue.URL), Err: ue.Err}
	}
	return err
}

// masks userinfo and query values; works on the raw string so an unparsable URL is masked too
func redactURL(rawURL string) string {
	if scheme, rest, ok := strings.Cut(rawURL, "://"); ok {
		authority := rest
		if end := strings.IndexAny(rest, "/?#"); end >= 0 {
			authority = rest[:end]
		}
		if at := strings.LastIndexByte(authority, '@'); at >= 0 {
			rawURL = scheme + "://{...}" + rest[at:]
		}
	}
	return utils.RedactUrlQueryValues(rawURL)
}

func (r *ResourceURLNotifier) SetKeys(apiKey, apiSecret string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.params.APIKey = apiKey
	r.params.APISecret = apiSecret
}

func (r *ResourceURLNotifier) SetFilter(params FilterParams) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.filter.SetFilter(params)
}

func (r *ResourceURLNotifier) IsAllowed(event string) bool {
	return r.filter.IsAllowed(event)
}

func (r *ResourceURLNotifier) RegisterProcessedHook(hook func(ctx context.Context, whi *livekit.WebhookInfo)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.processedHook = hook
}

func (r *ResourceURLNotifier) getProcessedHook() func(ctx context.Context, whi *livekit.WebhookInfo) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.processedHook
}

func (r *ResourceURLNotifier) QueueNotify(ctx context.Context, event *livekit.WebhookEvent, opts ...NotifyOption) error {
	if !r.filter.IsAllowed(event.Event) {
		return nil
	}

	if r.closed.IsBroken() {
		return errClosed
	}

	var key string
	if r.params.EventKey != nil {
		key = r.params.EventKey(event)
	} else {
		key = EventKey(event)
	}

	p := &NotifyParams{}
	for _, o := range opts {
		o(p)
	}

	r.mu.Lock()
	// copy the parameters
	params := r.params
	r.mu.Unlock()

	if len(p.ExtraWebhooks) > 1 {
		return fmt.Errorf("more than 1 extra webhook url unexpected")
	}
	if len(p.ExtraWebhooks) == 1 {
		params.URL = p.ExtraWebhooks[0].Url
		if p.ExtraWebhooks[0].SigningKey != "" {
			params.APIKey = p.ExtraWebhooks[0].SigningKey
		}
	}

	if p.Secret != "" {
		params.APISecret = p.Secret
	}

	if params.APIKey == "" || params.APISecret == "" {
		return errNoKey
	}

	r.mu.Lock()
	rqi := r.resourceQueues[key]
	if rqi == nil || !r.resourceQueueTimeoutQueue.Reset(rqi.tqi) {
		rq := newResourceQueue(resourceQueueParams{
			MaxDepth: params.Config.MaxDepth,
			Poster:   r,
		})
		rqi = &resourceQueueInfo{resourceQueue: rq, key: key}
		rqi.tqi = &utils.TimeoutQueueItem[*resourceQueueInfo]{Value: rqi}
		r.resourceQueueTimeoutQueue.Reset(rqi.tqi)
		r.resourceQueues[key] = rqi
	}
	r.mu.Unlock()

	qLen, err := rqi.resourceQueue.Enqueue(ctx, event, &params)
	if err != nil {
		fields := logFields(event, params.URL)
		fields = append(fields, "reason", err)
		params.Logger.Infow("dropped webhook", fields...)
		IncDispatchDrop(err.Error())

		if ph := r.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				time.Time{},
				0,
				time.Time{},
				0,
				params.URL,
				true,
				nil,
			)
			if params.FieldsHook != nil {
				params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
	} else {
		RecordQueueLength(qLen)
	}
	return err
}

func (r *ResourceURLNotifier) Stop(force bool) {
	r.closed.Break()

	r.mu.Lock()
	resourceQueues := r.resourceQueues
	r.resourceQueues = make(map[string]*resourceQueueInfo)
	r.mu.Unlock()

	for _, rq := range resourceQueues {
		rq.Stop(force)
	}
	if !force {
		// wait for the queued events to be sent, as URLNotifier does when draining
		for _, rq := range resourceQueues {
			<-rq.Done()
		}
	}
}

// poster interface
func (r *ResourceURLNotifier) Process(
	ctx context.Context,
	queuedAt time.Time,
	event *livekit.WebhookEvent,
	params *ResourceURLNotifierParams,
	qLen int,
) {
	fields := logFields(event, params.URL)

	queueDuration := time.Since(queuedAt)
	fields = append(fields, "queueDuration", queueDuration, "qLen", qLen)

	if queueDuration > params.Config.MaxAge {
		fields = append(fields, "reason", "age")
		params.Logger.Infow("dropped webhook", fields...)
		IncDispatchDrop("age")

		if ph := r.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				queuedAt,
				queueDuration,
				time.Time{},
				0,
				params.URL,
				true,
				nil,
			)
			if params.FieldsHook != nil {
				params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
		return
	}

	sendStart := time.Now()
	res, err := r.send(event, params)
	sendDuration := time.Since(sendStart)
	fields = append(fields, "sendDuration", sendDuration)
	if res != nil {
		fields = append(fields, "status", http.StatusText(res.StatusCode), "statusCode", res.StatusCode)
	}
	if err != nil {
		params.Logger.Warnw("failed to send webhook", err, fields...)
		IncDispatchFailure()
	} else {
		params.Logger.Infow("sent webhook", fields...)
		IncDispatchSuccess()
	}
	if ph := r.getProcessedHook(); ph != nil {
		whi := webhookInfo(
			event,
			queuedAt,
			queueDuration,
			sendStart,
			sendDuration,
			params.URL,
			false,
			err,
		)
		if params.FieldsHook != nil {
			params.FieldsHook(whi)
		}
		ph(ctx, whi)
	}
}

func (r *ResourceURLNotifier) send(event *livekit.WebhookEvent, params *ResourceURLNotifierParams) (*http.Response, error) {
	encoded, err := protojson.Marshal(event)
	if err != nil {
		return nil, err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	apiKey := params.APIKey
	apiSecret := params.APISecret

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return nil, err
	}
	req, err := retryablehttp.NewRequest("POST", params.URL, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return nil, redactURLError(err)
	}
	req.Header.Set(authHeader, token)
	req.Header.Set(contentTypeHeader, ContentType)
	req.Header.Set(userAgentHeader, params.UserAgent)
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()
	return res, nil
}

func (r *ResourceURLNotifier) sweeper() {
	ticker := time.NewTicker(r.params.Timeout / 2)
	defer ticker.Stop()

	for {
		select {
		case <-r.closed.Watch():
			return

		case <-ticker.C:
			for it := r.resourceQueueTimeoutQueue.IterateRemoveAfter(r.params.Timeout); it.Next(); {
				rqi := it.Item().Value

				r.mu.Lock()
				if r.resourceQueues[rqi.key] == rqi {
					delete(r.resourceQueues, rqi.key)
				}
				r.mu.Unlock()

				rqi.Stop(false)
			}
		}
	}
}
