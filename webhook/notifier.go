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
	"sync"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

type QueuedNotifier interface {
	QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error
	Stop(force bool)
}

type DefaultNotifier struct {
	queuedNotifiers []QueuedNotifier
}

func NewDefaultNotifier(apiKey, apiSecret string, urls []string) QueuedNotifier {
	n := &DefaultNotifier{}
	for _, url := range urls {
		u := NewURLNotifierWrapper(URLNotifierParams{
			URL:       url,
			Logger:    logger.GetLogger().WithComponent("webhook"),
			APIKey:    apiKey,
			APISecret: apiSecret,
		})
		n.queuedNotifiers = append(n.queuedNotifiers, u)
	}
	return n
}

// NewDefaultNotifierWithFilter takes an events eventsFilter that is shared across all urls.
// if eventsFilter is nil, then all events will be sent. If not only the events specified
// by eventsFilter will be sent and any other event will be ignored with a debug log
// TODO maybe add eventsFilter per url?! but it's not my use case.
func NewDefaultNotifierWithFilter(apiKey, apiSecret string, urls []string, eventsFilter []string) QueuedNotifier {
	n := &DefaultNotifier{}
	for _, url := range urls {
		u := NewFilteredNotifier(NewURLNotifierWrapper(URLNotifierParams{
			URL:       url,
			Logger:    logger.GetLogger().WithComponent("webhook"),
			APIKey:    apiKey,
			APISecret: apiSecret,
		}), FilteredNotifierParams{
			Events: eventsFilter,
			Logger: logger.GetLogger().WithComponent("webhook"),
		})
		n.queuedNotifiers = append(n.queuedNotifiers, u)
	}
	return n
}

func (n *DefaultNotifier) Stop(force bool) {
	wg := sync.WaitGroup{}
	for _, u := range n.queuedNotifiers {
		wg.Add(1)
		go func(qn QueuedNotifier) {
			defer wg.Done()
			qn.Stop(force)
		}(u)
	}
	wg.Wait()
}

func (n *DefaultNotifier) QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error {
	for _, u := range n.queuedNotifiers {
		if err := u.QueueNotify(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
