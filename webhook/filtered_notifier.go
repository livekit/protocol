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
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

type FilteredNotifierParams struct {
	// Events will be used to filter out webhook events. One might want only a subset of events
	// If Events is nil or zero-sized, all events will be sent
	Events []string
	Logger logger.Logger
}

type FilteredNotifier struct {
	logger         logger.Logger
	events         []string // TODO do we really need a map[string]struct{}
	queuedNotifier QueuedNotifier
}

func NewFilteredNotifier(notifier QueuedNotifier, params FilteredNotifierParams) *FilteredNotifier {
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}
	return &FilteredNotifier{
		logger:         params.Logger,
		events:         params.Events,
		queuedNotifier: notifier,
	}
}

func (notifier *FilteredNotifier) QueueNotify(ctx context.Context, event *livekit.WebhookEvent) error {
	if len(notifier.events) == 0 {
		return notifier.queuedNotifier.QueueNotify(ctx, event)
	}

	for _, ev := range notifier.events {
		if ev == event.Event {
			return notifier.queuedNotifier.QueueNotify(ctx, event)
		}
	}

	notifier.logger.Debugw("ignoring event: %s", event.Event)
	return nil
}

func (notifier *FilteredNotifier) Stop(force bool) {
	notifier.queuedNotifier.Stop(force)
}
