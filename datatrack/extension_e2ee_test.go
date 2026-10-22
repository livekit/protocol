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

func TestExtensionE2EE(t *testing.T) {
	iv := [ExtensionE2EEIVLength]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	extE2EE := NewExtensionE2EE(0xfa, iv)

	expectedExt := Extension{
		id:   ExtensionE2EEID,
		data: []byte{0xfa, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	}
	ext, err := extE2EE.Marshal()
	require.NoError(t, err)
	require.Equal(t, expectedExt, ext)

	var unmarshaled ExtensionE2EE
	err = unmarshaled.Unmarshal(ext)
	require.NoError(t, err)
	assert.Equal(t, uint8(0xfa), unmarshaled.KeyIndex())
	assert.Equal(t, iv, unmarshaled.IV())
}
