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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtensionUserTimestamp(t *testing.T) {
	extUserTimestamp := NewExtensionUserTimestamp(0x4411221111118811)

	expectedExt := Extension{
		id:   ExtensionUserTimestampID,
		data: []byte{0x44, 0x11, 0x22, 0x11, 0x11, 0x11, 0x88, 0x11},
	}
	ext, err := extUserTimestamp.Marshal()
	require.NoError(t, err)
	require.Equal(t, expectedExt, ext)

	var unmarshaled ExtensionUserTimestamp
	err = unmarshaled.Unmarshal(ext)
	require.NoError(t, err)
	assert.Equal(t, uint64(0x4411221111118811), unmarshaled.Timestamp())
}
