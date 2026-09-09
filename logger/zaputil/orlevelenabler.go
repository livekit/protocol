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

package zaputil

import "go.uber.org/zap/zapcore"

// OrLevelEnabler is enabled at a level if any of its members is. Members must
// be non-nil; build one with NewOrLevelEnabler from sources that may not be.
type OrLevelEnabler []zapcore.LevelEnabler

func (e OrLevelEnabler) Enabled(lvl zapcore.Level) bool {
	for _, enab := range e {
		if enab.Enabled(lvl) {
			return true
		}
	}
	return false
}

// NewOrLevelEnabler combines enabs, ignoring any that are nil. It returns nil
// when every one of them is nil, so callers can tell "nothing to combine" apart
// from an enabler that is never enabled.
func NewOrLevelEnabler(enabs ...zapcore.LevelEnabler) zapcore.LevelEnabler {
	e := make(OrLevelEnabler, 0, len(enabs))
	for _, enab := range enabs {
		if enab != nil {
			e = append(e, enab)
		}
	}
	if len(e) == 0 {
		return nil
	}
	return e
}
