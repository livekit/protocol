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

package datatrack

import (
	"encoding/binary"

	"github.com/livekit/protocol/livekit"
)

type ExtensionUserTimestamp struct {
	timestamp uint64
}

func NewExtensionUserTimestamp(timestamp uint64) *ExtensionUserTimestamp {
	return &ExtensionUserTimestamp{timestamp: timestamp}
}

func (e *ExtensionUserTimestamp) Timestamp() uint64 {
	return e.timestamp
}

func (e *ExtensionUserTimestamp) Marshal() (Extension, error) {
	data := make([]byte, binary.Size(e.timestamp))
	binary.BigEndian.PutUint64(data, e.timestamp)
	return Extension{
		id:   uint8(livekit.DataTrackExtensionID_DTEI_USER_TIMESTAMP),
		data: data,
	}, nil
}

func (e *ExtensionUserTimestamp) Unmarshal(ext Extension) error {
	if ext.id != uint8(livekit.DataTrackExtensionID_DTEI_USER_TIMESTAMP) {
		return ErrExtensionInvalidID
	}

	if len(ext.data) < binary.Size(e.timestamp) {
		return ErrExtensionDataTooShort
	}

	e.timestamp = binary.BigEndian.Uint64(ext.data)
	return nil
}
