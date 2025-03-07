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
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/hashicorp/go-retryablehttp"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
)

const (
	defaultTimeout  = 5 * time.Minute
	defaultMaxAge   = 5 * time.Second
	defaultMaxDepth = 200
)

var (
	errClosed = errors.New("notifier is closed")
)

type poster interface {
	Process(ctx context.Context, queuedAt time.Time, event *livekit.WebhookEvent)
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
	MaxAge     time.Duration
	MaxDepth   int
	URL        string
	APIKey     string
	APISecret  string
	FieldsHook func(whi *livekit.WebhookInfo)
}

// ResourceURLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It queues up events per resource (could be egress, ingress, room, participant, track, etc.)
//   to avoid blocking events of one resource blocking another resource's event(s).
// It will retry on failure, and will drop events if notification fall too far behind,
//   either in age or queue depth.
type ResourceURLNotifier struct {
	mu            sync.RWMutex
	params        ResourceURLNotifierParams
	client        *retryablehttp.Client
	processedHook func(ctx context.Context, whi *livekit.WebhookInfo)

	resourceQueues            map[string]*resourceQueueInfo
	resourceQueueTimeoutQueue utils.TimeoutQueue[*resourceQueueInfo]
	closed                    core.Fuse
}

func NewResourceURLNotifier(params ResourceURLNotifierParams) *ResourceURLNotifier {
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}

	if params.Timeout == 0 {
		params.Timeout = defaultTimeout
	}
	if params.MaxAge == 0 {
		params.MaxAge = defaultMaxAge
	}
	if params.MaxDepth == 0 {
		params.MaxDepth = defaultMaxDepth
	}

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
	if params.ClientTimeout > 0 {
		rhc.HTTPClient.Timeout = params.ClientTimeout
	}
	rhc.Logger = &logAdapter{}
	r := &ResourceURLNotifier{
		params:         params,
		client:         rhc,
		resourceQueues: make(map[string]*resourceQueueInfo),
	}

	go r.sweeper()
	return r
}

func (r *ResourceURLNotifier) SetKeys(apiKey, apiSecret string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.params.APIKey = apiKey
	r.params.APISecret = apiSecret
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

func (r *ResourceURLNotifier) QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error {
	if r.closed.IsBroken() {
		return errClosed
	}

	key := eventKey(event)

	r.mu.Lock()
	rqi := r.resourceQueues[key]
	if rqi == nil || !r.resourceQueueTimeoutQueue.Reset(rqi.tqi) {
		rq := newResourceQueue(resourceQueueParams{
			MaxDepth: r.params.MaxDepth,
			Poster:   r,
		})
		rqi = &resourceQueueInfo{resourceQueue: rq, key: key}
		rqi.tqi = &utils.TimeoutQueueItem[*resourceQueueInfo]{Value: rqi}
		r.resourceQueueTimeoutQueue.Reset(rqi.tqi)
		r.resourceQueues[key] = rqi
	}
	r.mu.Unlock()

	err := rqi.resourceQueue.Enqueue(ctx, event)
	if err != nil {
		fields := logFields(event, r.params.URL)
		fields = append(fields, "reason", err)
		r.params.Logger.Infow("dropped webhook", fields...)

		if ph := r.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				time.Time{},
				0,
				time.Time{},
				0,
				r.params.URL,
				true,
				nil,
			)
			if r.params.FieldsHook != nil {
				r.params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
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
}

// poster interface
func (r *ResourceURLNotifier) Process(ctx context.Context, queuedAt time.Time, event *livekit.WebhookEvent) {
	fields := logFields(event, r.params.URL)

	queueDuration := time.Since(queuedAt)
	fields = append(fields, "queueDuration", queueDuration)

	if queueDuration > r.params.MaxAge {
		fields = append(fields, "reason", "age")
		r.params.Logger.Infow("dropped webhook", fields...)

		if ph := r.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				queuedAt,
				queueDuration,
				time.Time{},
				0,
				r.params.URL,
				true,
				nil,
			)
			if r.params.FieldsHook != nil {
				r.params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
		return
	}

	sendStart := time.Now()
	err := r.send(event)
	sendDuration := time.Since(sendStart)
	fields = append(fields, "sendDuration", sendDuration)
	if err != nil {
		r.params.Logger.Warnw("failed to send webhook", err, fields...)
	} else {
		r.params.Logger.Infow("sent webhook", fields...)
	}
	if ph := r.getProcessedHook(); ph != nil {
		whi := webhookInfo(
			event,
			queuedAt,
			queueDuration,
			sendStart,
			sendDuration,
			r.params.URL,
			false,
			err,
		)
		if r.params.FieldsHook != nil {
			r.params.FieldsHook(whi)
		}
		ph(ctx, whi)
	}
}

func (r *ResourceURLNotifier) send(event *livekit.WebhookEvent) error {
	encoded, err := protojson.Marshal(event)
	if err != nil {
		return err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	r.mu.RLock()
	apiKey := r.params.APIKey
	apiSecret := r.params.APISecret
	r.mu.RUnlock()

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return err
	}
	req, err := retryablehttp.NewRequest("POST", r.params.URL, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return err
	}
	req.Header.Set(authHeader, token)
	// use a custom mime type to ensure signature is checked prior to parsing
	req.Header.Set("content-type", "application/webhook+json")
	res, err := r.client.Do(req)
	if err != nil {
		return err
	}
	_ = res.Body.Close()
	return nil
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
