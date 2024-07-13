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
	Stop(force bool)
}

type NotifierParams struct {
	ApiKey        string
	ApiSecret     string
	Urls          []string
	IncludeEvents []string // when IncludeEvents is not empty ExcludeEvents will be ignored
	ExcludeEvents []string // needs IncludeEvents to be empty, otherwise it won't take effect
	Batched       bool     // when Batched is true, DropWhenFull is ignored
	DropWhenFull  bool     // only works when Batched is disabled
}

func NewNotifier(ctx context.Context, params NotifierParams) QueuedNotifier {
	if len(params.IncludeEvents) > 0 {
		params.ExcludeEvents = nil
	}
	if params.Batched {
		return NewBatchedNotifier(ctx, params.ApiKey, params.ApiSecret, params.Urls, params.IncludeEvents, params.ExcludeEvents)
	} else {
		return NewDefaultNotifier(params.ApiKey, params.ApiSecret, params.Urls, params.DropWhenFull, params.IncludeEvents, params.ExcludeEvents)
	}
}

type DefaultNotifier struct {
	urlNotifiers   []URLNotifier
	includedEvents []string
	excludedEvents []string
}

func NewBatchedNotifier(ctx context.Context, apiKey, apiSecret string, urls []string, includedEvents []string, excludedEvents []string) QueuedNotifier {
	n := &DefaultNotifier{
		includedEvents: includedEvents,
		excludedEvents: excludedEvents,
	}
	for _, url := range urls {
		u := NewBatchURLNotifier(ctx, BatchURLNotifierParams{
			Logger:    logger.GetLogger().WithComponent("webhook"),
			URL:       url,
			APIKey:    apiKey,
			APISecret: apiSecret,
		})
		n.urlNotifiers = append(n.urlNotifiers, u)
	}
	return n
}

func NewDefaultNotifier(apiKey, apiSecret string, urls []string, dropWhenFull bool, includedEvents []string, excludedEvents []string) QueuedNotifier {
	n := &DefaultNotifier{
		includedEvents: includedEvents,
		excludedEvents: excludedEvents,
	}
	for _, url := range urls {
		u := NewDefaultURLNotifier(URLNotifierParams{
			URL:          url,
			Logger:       logger.GetLogger().WithComponent("webhook"),
			APIKey:       apiKey,
			APISecret:    apiSecret,
			DropWhenFull: dropWhenFull,
		})
		n.urlNotifiers = append(n.urlNotifiers, u)
	}
	return n
}

func (n *DefaultNotifier) Stop(force bool) {
	wg := sync.WaitGroup{}
	for _, u := range n.urlNotifiers {
		wg.Add(1)
		go func(u URLNotifier) {
			defer wg.Done()
			u.Stop(force)
		}(u)
	}
	wg.Wait()
}

func (n *DefaultNotifier) QueueNotify(_ context.Context, event *livekit.WebhookEvent) error {
	for _, ev := range n.includedEvents {
		if event.Event == ev {
			return n.queueNotify(event)
		}
	}
	if len(n.includedEvents) > 0 {
		return nil
	}

	for _, ev := range n.excludedEvents {
		if event.Event == ev {
			return nil
		}
	}
	return n.queueNotify(event)
}

func (n *DefaultNotifier) queueNotify(event *livekit.WebhookEvent) error {
	for _, u := range n.urlNotifiers {
		if err := u.QueueNotify(event); err != nil {
			return err
		}
	}
	return nil
}
