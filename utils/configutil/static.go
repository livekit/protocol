// Copyright 2026 LiveKit, Inc.
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

package configutil

import (
	"github.com/livekit/protocol/utils/events"
)

// NewStaticObserver stubs observable config in tests, and backs dev paths that
// must hand back the same *Observer the production path does.
//
// The nil builder is safe only because load is unreachable without a watcher.
func NewStaticObserver[T any](conf *T) *Observer[T] {
	c := &Observer[T]{observers: events.NewObserverList[*T](events.WithBlocking())}
	c.conf.Store(conf)
	return c
}
