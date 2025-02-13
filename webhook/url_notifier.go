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
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/atomic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

const (
	numWorkers = 10
)

type URLNotifierParams struct {
	HTTPClientParams
	Logger     logger.Logger
	QueueSize  int
	URL        string
	APIKey     string
	APISecret  string
	FieldsHook func(whi *livekit.WebhookInfo)
}

type HTTPClientParams struct {
	RetryWaitMin  time.Duration
	RetryWaitMax  time.Duration
	MaxRetries    int
	ClientTimeout time.Duration
}

const defaultQueueSize = 100

// URLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It will retry on failure, and will drop events if notification fall too far behind
type URLNotifier struct {
	mu            sync.RWMutex
	params        URLNotifierParams
	client        *retryablehttp.Client
	dropped       atomic.Int32
	pool          core.QueuePool
	processedHook func(ctx context.Context, whi *livekit.WebhookInfo)
}

func NewURLNotifier(params URLNotifierParams) *URLNotifier {
	if params.QueueSize == 0 {
		params.QueueSize = defaultQueueSize
	}
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
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
	n := &URLNotifier{
		params: params,
		client: rhc,
	}
	n.client.Logger = &logAdapter{}

	n.pool = core.NewQueuePool(numWorkers, core.QueueWorkerParams{
		QueueSize:    params.QueueSize,
		DropWhenFull: true,
	})
	return n
}

func (n *URLNotifier) SetKeys(apiKey, apiSecret string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.params.APIKey = apiKey
	n.params.APISecret = apiSecret
}

func (n *URLNotifier) RegisterProcessedHook(hook func(ctx context.Context, whi *livekit.WebhookInfo)) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.processedHook = hook
}

func (n *URLNotifier) getProcessedHook() func(ctx context.Context, whi *livekit.WebhookInfo) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.processedHook
}

func (n *URLNotifier) QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error {
	enqueuedAt := time.Now()

	if !n.pool.Submit(n.eventKey(event), func() {
		fields := logFields(event, n.params.URL)

		queueDuration := time.Since(enqueuedAt)
		fields = append(fields, "queueDuration", queueDuration)

		sendStart := time.Now()
		err := n.send(event)
		sendDuration := time.Since(sendStart)
		fields = append(fields, "sendDuration", sendDuration)
		if err != nil {
			n.params.Logger.Warnw("failed to send webhook", err, fields...)
			n.dropped.Add(event.NumDropped + 1)
		} else {
			n.params.Logger.Infow("sent webhook", fields...)
		}
		if ph := n.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				enqueuedAt,
				queueDuration,
				sendStart,
				sendDuration,
				n.params.URL,
				false,
				err,
			)
			if n.params.FieldsHook != nil {
				n.params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
	}) {
		n.dropped.Inc()

		fields := logFields(event, n.params.URL)
		n.params.Logger.Infow("dropped webhook", fields...)

		if ph := n.getProcessedHook(); ph != nil {
			whi := webhookInfo(
				event,
				time.Time{},
				0,
				time.Time{},
				0,
				n.params.URL,
				true,
				nil,
			)
			if n.params.FieldsHook != nil {
				n.params.FieldsHook(whi)
			}
			ph(ctx, whi)
		}
	}
	return nil
}

func (c *URLNotifier) eventKey(event *livekit.WebhookEvent) string {
	if event.EgressInfo != nil {
		return event.EgressInfo.EgressId
	}
	if event.IngressInfo != nil {
		return event.IngressInfo.IngressId
	}
	if event.Room != nil {
		return event.Room.Name
	}
	if event.Participant != nil {
		return event.Participant.Identity
	}
	if event.Track != nil {
		return event.Track.Sid
	}
	return "default"
}

func (n *URLNotifier) Stop(force bool) {
	if force {
		n.pool.Kill()
	} else {
		n.pool.Drain()
	}
}

func (n *URLNotifier) send(event *livekit.WebhookEvent) error {
	// set dropped count
	event.NumDropped = n.dropped.Swap(0)
	encoded, err := protojson.Marshal(event)
	if err != nil {
		return err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	n.mu.RLock()
	apiKey := n.params.APIKey
	apiSecret := n.params.APISecret
	n.mu.RUnlock()

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return err
	}
	r, err := retryablehttp.NewRequest("POST", n.params.URL, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return err
	}
	r.Header.Set(authHeader, token)
	// use a custom mime type to ensure signature is checked prior to parsing
	r.Header.Set("content-type", "application/webhook+json")
	res, err := n.client.Do(r)
	if err != nil {
		return err
	}
	_ = res.Body.Close()
	return nil
}

type logAdapter struct{}

func (l *logAdapter) Printf(string, ...interface{}) {}

func logFields(event *livekit.WebhookEvent, url string) []interface{} {
	fields := make([]interface{}, 0, 20)
	fields = append(fields,
		"event", event.Event,
		"id", event.Id,
		"webhookTime", event.CreatedAt,
		"url", url,
	)

	if event.Room != nil {
		fields = append(fields,
			"room", event.Room.Name,
			"roomID", event.Room.Sid,
		)
	}
	if event.Participant != nil {
		fields = append(fields,
			"participant", event.Participant.Identity,
			"pID", event.Participant.Sid,
		)
	}
	if event.Track != nil {
		fields = append(fields,
			"trackID", event.Track.Sid,
		)
	}
	if event.EgressInfo != nil {
		fields = append(fields,
			"egressID", event.EgressInfo.EgressId,
			"status", event.EgressInfo.Status,
		)
		if event.EgressInfo.Error != "" {
			fields = append(fields, "error", event.EgressInfo.Error)
		}
	}
	if event.IngressInfo != nil {
		fields = append(fields,
			"ingressID", event.IngressInfo.IngressId,
		)
		if event.IngressInfo.State != nil {
			fields = append(fields, "status", event.IngressInfo.State.Status)
			if event.IngressInfo.State.Error != "" {
				fields = append(fields, "error", event.IngressInfo.State.Error)
			}
		}
	}
	return fields
}

func webhookInfo(
	event *livekit.WebhookEvent,
	queuedAt time.Time,
	queueDuration time.Duration,
	sentAt time.Time,
	sendDuration time.Duration,
	url string,
	isDropped bool,
	sendError error,
) *livekit.WebhookInfo {
	whi := &livekit.WebhookInfo{
		EventId:         event.Id,
		Event:           event.Event,
		CreatedAt:       timestamppb.New(time.Unix(event.CreatedAt, 0)),
		QueuedAt:        timestamppb.New(queuedAt),
		QueueDurationNs: queueDuration.Nanoseconds(),
		SentAt:          timestamppb.New(sentAt),
		SendDurationNs:  sendDuration.Nanoseconds(),
		Url:             url,
		NumDropped:      event.NumDropped,
		IsDropped:       isDropped,
	}
	if !queuedAt.IsZero() {
		whi.QueuedAt = timestamppb.New(queuedAt)
	}
	if !sentAt.IsZero() {
		whi.SentAt = timestamppb.New(sentAt)
	}
	if event.Room != nil {
		whi.RoomName = event.Room.Name
		whi.RoomId = event.Room.Sid
	}
	if event.Participant != nil {
		whi.ParticipantIdentity = event.Participant.Identity
		whi.ParticipantId = event.Participant.Sid
	}
	if event.Track != nil {
		whi.TrackId = event.Track.Sid
	}
	if event.EgressInfo != nil {
		whi.EgressId = event.EgressInfo.EgressId
		whi.ServiceStatus = event.EgressInfo.Status.String()
		if event.EgressInfo.Error != "" {
			whi.ServiceErrorCode = event.EgressInfo.ErrorCode
			whi.ServiceError = event.EgressInfo.Error
		}
	}
	if event.IngressInfo != nil {
		whi.IngressId = event.IngressInfo.IngressId
		if event.IngressInfo.State != nil {
			whi.ServiceStatus = event.IngressInfo.State.Status.String()
			if event.IngressInfo.State.Error != "" {
				whi.ServiceError = event.IngressInfo.State.Error
			}
		}
	}
	if sendError != nil {
		whi.SendError = sendError.Error()
	}
	return whi
}
