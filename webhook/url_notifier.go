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
	"crypto/sha256"
	"encoding/base64"
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/atomic"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

const (
	numWorkers = 10
)

type URLNotifierParams struct {
	HTTPClientParams
	Logger    logger.Logger
	QueueSize int
	URL       string
	APIKey    string
	APISecret string
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
	mu      sync.RWMutex
	params  URLNotifierParams
	client  *retryablehttp.Client
	dropped atomic.Int32
	pool    core.QueuePool
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
		OnDropped:    func() { n.dropped.Inc() },
	})
	return n
}

func (n *URLNotifier) SetKeys(apiKey, apiSecret string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.params.APIKey = apiKey
	n.params.APISecret = apiSecret
}

func (n *URLNotifier) QueueNotify(event *livekit.WebhookEvent) error {
	enqueuedAt := time.Now()

	n.pool.Submit(n.eventKey(event), func() {
		fields := logFields(event)
		fields = append(fields,
			"url", n.params.URL,
			"queueDuration", time.Since(enqueuedAt),
		)
		sentStart := time.Now()
		err := n.send(event)
		fields = append(fields, "sendDuration", time.Since(sentStart))
		if err != nil {
			n.params.Logger.Warnw("failed to send webhook", err, fields...)
			n.dropped.Add(event.NumDropped + 1)
		} else {
			n.params.Logger.Infow("sent webhook", fields...)
		}
	})
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

func logFields(event *livekit.WebhookEvent) []interface{} {
	fields := make([]interface{}, 0, 20)
	fields = append(fields,
		"event", event.Event,
		"id", event.Id,
		"webhookTime", event.CreatedAt,
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
