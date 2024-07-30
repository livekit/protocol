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

type URLNotifier interface {
	SetKeys(apiKey, apiSecret string)
	QueueNotify(event *livekit.WebhookEvent) error
	Stop(force bool)
}

type URLNotifierParams struct {
	Logger       logger.Logger
	QueueSize    int
	DropWhenFull bool
	URL          string
	APIKey       string
	APISecret    string
}

const defaultQueueSize = 100

// DefaultURLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It will retry on failure, and will drop events if notification fall too far behind
type DefaultURLNotifier struct {
	mu      sync.RWMutex
	params  URLNotifierParams
	client  *retryablehttp.Client
	dropped atomic.Int32
	worker  core.QueueWorker
}

func NewDefaultURLNotifier(params URLNotifierParams) URLNotifier {
	if params.QueueSize == 0 {
		params.QueueSize = defaultQueueSize
	}
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}

	n := &DefaultURLNotifier{
		params: params,
		client: retryablehttp.NewClient(),
	}
	n.client.Logger = &logAdapter{}
	n.worker = core.NewQueueWorker(core.QueueWorkerParams{
		QueueSize:    params.QueueSize,
		DropWhenFull: params.DropWhenFull,
		OnDropped:    func() { n.dropped.Inc() },
	})
	return n
}

func (n *DefaultURLNotifier) SetKeys(apiKey, apiSecret string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.params.APIKey = apiKey
	n.params.APISecret = apiSecret
}

func (n *DefaultURLNotifier) QueueNotify(event *livekit.WebhookEvent) error {
	n.worker.Submit(func() {
		if err := n.send(event); err != nil {
			n.params.Logger.Warnw("failed to send webhook", err, "url", n.params.URL, "event", event.Event)
			n.dropped.Add(event.NumDropped + 1)
		} else {
			n.params.Logger.Infow("sent webhook", "url", n.params.URL, "event", event.Event, "eventDetails", logger.Proto(event))
		}
	})
	return nil
}

func (n *DefaultURLNotifier) Stop(force bool) {
	if force {
		n.worker.Kill()
	} else {
		n.worker.Drain()
	}
}

func (n *DefaultURLNotifier) send(event *livekit.WebhookEvent) error {
	// set dropped count
	event.NumDropped = n.dropped.Swap(0)
	event.DequeuedAt = time.Now().Unix()
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
