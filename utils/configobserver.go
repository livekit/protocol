// Copyright 2024 LiveKit, Inc.
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

package utils

import "github.com/livekit/protocol/utils/configutil"

type ConfigBuilder[T any] = configutil.Builder[T]

type ConfigDefaulter[T any] = configutil.Defaulter[T]

type ConfigObserver[T any] = configutil.Observer[T]

func NewConfigObserver[T any](path string, builder ConfigBuilder[T]) (*ConfigObserver[T], *T, error) {
	return configutil.NewObserver(path, builder)
}
