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

import "slices"

type filterParams struct {
	Includes []string
	Excludes []string
}

type filter struct {
	params filterParams
}

func newFilter(params filterParams) *filter {
	return &filter{
		params: params,
	}
}

func (f *filter) IsAllowed(event string) bool {
	// includes get higher precendence than excludes
	if slices.Contains(f.params.Includes, event) {
		return true
	}

	return !slices.Contains(f.params.Excludes, event)
}
