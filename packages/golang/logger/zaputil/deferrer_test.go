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

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestDeferredLogger(t *testing.T) {
	c := &testCore{Core: zap.NewExample().Core()}
	d, resolve := NewDeferrer()
	dc := NewDeferredValueCore(c, d)
	s := zap.New(dc).Sugar()

	s.Infow("test")
	require.Equal(t, 0, c.WriteCount())

	s.With("a", "1").Infow("test")
	require.Equal(t, 0, c.WriteCount())

	resolve("b", "2")
	require.Equal(t, 2, c.WriteCount())

	s.With("c", "3").Infow("test")
	require.Equal(t, 3, c.WriteCount())
}
