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
	"context"
	"github.com/livekit/protocol/logger"
	"sync"

	"github.com/livekit/protocol/livekit"
)

type QueuedNotifier interface {
	QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error
}

type DefaultNotifier struct {
	urlNotifiers []*URLNotifier
}

func NewDefaultNotifier(apiKey, apiSecret string, urls []string) QueuedNotifier {
	n := &DefaultNotifier{}
	for _, url := range urls {
		u := NewURLNotifier(URLNotifierParams{
			URL:          url,
			Logger:       logger.GetLogger().WithComponent("webhook"),
			APIKey:       apiKey,
			APISecret:    apiSecret,
			DropWhenFull: true,
		})
		n.urlNotifiers = append(n.urlNotifiers, u)
	}
	return n
}

func NewDefaultNotifierByParams(params []URLNotifierParams) QueuedNotifier {
	n := &DefaultNotifier{}
	for _, p := range params {
		if p.Logger == nil {
			p.Logger = logger.GetLogger().WithComponent("webhook")
		}

		u := NewURLNotifier(p)
		n.urlNotifiers = append(n.urlNotifiers, u)
	}
	return n
}

func (n *DefaultNotifier) Stop(force bool) {
	wg := sync.WaitGroup{}
	for _, u := range n.urlNotifiers {
		wg.Add(1)
		go func(u *URLNotifier) {
			defer wg.Done()
			u.Stop(force)
		}(u)
	}
	wg.Wait()
}

func (n *DefaultNotifier) QueueNotify(_ context.Context, event *livekit.WebhookEvent) error {
	for _, u := range n.urlNotifiers {
		if err := u.QueueNotify(event); err != nil {
			return err
		}
	}
	return nil
}
