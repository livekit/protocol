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

package configutil

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// scalarField stands in for the config types LiveKit services embed in the
// struct trees Observer decodes -- rtcconfig.PortRange, rtcconfig.NodeIP,
// featureflags rules -- which parse themselves from a single yaml scalar.
//
// It uses the func-based unmarshal signature on purpose. Observer decodes with
// go.yaml.in/yaml/v3 while many callers still marshal and decode elsewhere with
// gopkg.in/yaml.v3, and a UnmarshalYAML(*yaml.Node) method is only recognized by
// the package that owns the Node type. The other package silently skips it and
// decodes reflectively instead, so a type bound to the wrong package either
// errors with "cannot unmarshal !!str into ..." or comes back zero.
type scalarField struct {
	Raw string
}

func (s *scalarField) UnmarshalYAML(unmarshal func(any) error) error {
	var v any
	if err := unmarshal(&v); err != nil {
		return err
	}
	s.Raw = fmt.Sprint(v)
	return nil
}

type customConfig struct {
	Name   string      `yaml:"name"`
	Scalar scalarField `yaml:"scalar"`
	Nested struct {
		Scalar scalarField `yaml:"scalar"`
	} `yaml:"nested"`
	Ptr *scalarField `yaml:"ptr"`
}

type customConfigBuilder struct{}

func (customConfigBuilder) New() (*customConfig, error) { return &customConfig{}, nil }

// TestObserverHonorsCustomUnmarshallers guards the yaml package Observer decodes
// with. Types reached through the config tree must still have their custom
// unmarshallers invoked, at the top level, nested, and behind a pointer.
//
// If this breaks after changing observer.go's yaml import, the fix is not to
// change the import back in isolation: every custom unmarshaller in every
// service config tree has to agree on the package, or use the func-based
// signature, which no package owns.
func TestObserverHonorsCustomUnmarshallers(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	require.NoError(t, os.WriteFile(path, []byte(`
name: svc
scalar: 7881-7981
nested:
  scalar: 1.2.3.4
ptr: plain
`), 0o600))

	_, conf, err := NewObserver[customConfig](path, customConfigBuilder{})
	require.NoError(t, err)

	require.Equal(t, "svc", conf.Name)
	require.Equal(t, "7881-7981", conf.Scalar.Raw, "top-level custom unmarshaller skipped")
	require.Equal(t, "1.2.3.4", conf.Nested.Scalar.Raw, "nested custom unmarshaller skipped")
	require.NotNil(t, conf.Ptr)
	require.Equal(t, "plain", conf.Ptr.Raw, "pointer custom unmarshaller skipped")
}

// TestObserverDecodesBareIntScalar covers a scalar that resolves to a non-string
// yaml type. Custom unmarshallers must not assume the decoded value is a string.
func TestObserverDecodesBareIntScalar(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	require.NoError(t, os.WriteFile(path, []byte("scalar: 7881\n"), 0o600))

	_, conf, err := NewObserver[customConfig](path, customConfigBuilder{})
	require.NoError(t, err)
	require.Equal(t, "7881", conf.Scalar.Raw)
}
