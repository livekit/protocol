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

const (
	ExtensionE2EEIVLength = 12

	// TODO: once the collision is fully resolved (see #1825), the source of truth for
	// extension IDs will become the Protobuf enum, `DataTrackExtensionID`, at which point
	// this const will be removed.
	ExtensionE2EEID uint8 = 1
)

type ExtensionE2EE struct {
	keyIndex uint8
	iv       [ExtensionE2EEIVLength]byte
}

func NewExtensionE2EE(keyIndex uint8, iv [ExtensionE2EEIVLength]byte) *ExtensionE2EE {
	return &ExtensionE2EE{keyIndex: keyIndex, iv: iv}
}

func (e *ExtensionE2EE) KeyIndex() uint8 {
	return e.keyIndex
}

func (e *ExtensionE2EE) IV() [ExtensionE2EEIVLength]byte {
	return e.iv
}

func (e *ExtensionE2EE) Marshal() (Extension, error) {
	data := make([]byte, 1+len(e.iv))
	data[0] = e.keyIndex
	copy(data[1:], e.iv[:])
	return Extension{
		id:   ExtensionE2EEID,
		data: data,
	}, nil
}

func (e *ExtensionE2EE) Unmarshal(ext Extension) error {
	if ext.id != ExtensionE2EEID {
		return ErrExtensionInvalidID
	}

	if len(ext.data) < 1+len(e.iv) {
		return ErrExtensionDataTooShort
	}

	e.keyIndex = ext.data[0]
	copy(e.iv[:], ext.data[1:])
	return nil
}
