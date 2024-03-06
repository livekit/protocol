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

package utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestMarshalUnmarshalGuid(t *testing.T) {
	id0 := livekit.TrackID(NewGuid(TrackPrefix))
	b0 := MarshalGuid(id0)
	id1 := UnmarshalGuid[livekit.TrackID](b0)
	b1 := MarshalGuid(id1)
	require.EqualValues(t, id0, id1)
	require.EqualValues(t, b0, b1)
}
