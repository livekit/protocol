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

package redis

import (
	"testing"

	"github.com/redis/go-redis/v9/auth"
	"github.com/stretchr/testify/require"
)

type fakeStreamingProvider struct{}

func (f *fakeStreamingProvider) Subscribe(_ auth.CredentialsListener) (auth.Credentials, auth.UnsubscribeFunc, error) {
	return nil, nil, nil
}

func TestBuildRedisOptions_ClusterWithoutProvider(t *testing.T) {
	opts, err := buildRedisOptions(&RedisConfig{
		ClusterAddresses: []string{"host:10000"},
		Username:         "user",
		Password:         "pass",
		UseTLS:           true,
	}, clientOptions{})
	require.NoError(t, err)
	require.True(t, opts.IsClusterMode)
	require.Nil(t, opts.StreamingCredentialsProvider)
	require.Equal(t, "user", opts.Username)
	require.Equal(t, "pass", opts.Password)
	require.NotNil(t, opts.TLSConfig)
}

func TestBuildRedisOptions_WithStreamingCredentialsProvider(t *testing.T) {
	fake := &fakeStreamingProvider{}
	var co clientOptions
	WithStreamingCredentialsProvider(fake)(&co)

	opts, err := buildRedisOptions(&RedisConfig{
		ClusterAddresses: []string{"host:10000"},
	}, co)
	require.NoError(t, err)
	require.Same(t, fake, opts.StreamingCredentialsProvider)
}

func TestBuildRedisOptions_AzureEntraFlag(t *testing.T) {
	fake := &fakeStreamingProvider{}
	orig := azureEntraProviderFactory
	azureEntraProviderFactory = func() (auth.StreamingCredentialsProvider, error) {
		return fake, nil
	}
	t.Cleanup(func() { azureEntraProviderFactory = orig })

	opts, err := buildRedisOptions(&RedisConfig{
		ClusterAddresses: []string{"host:10000"},
		UseTLS:           true,
		AzureEntra:       true,
	}, clientOptions{})
	require.NoError(t, err)
	require.Same(t, fake, opts.StreamingCredentialsProvider)
}

func TestBuildRedisOptions_ExplicitProviderBeatsAzureEntraFlag(t *testing.T) {
	explicit := &fakeStreamingProvider{}
	azureFromFactory := &fakeStreamingProvider{}
	orig := azureEntraProviderFactory
	azureEntraProviderFactory = func() (auth.StreamingCredentialsProvider, error) {
		return azureFromFactory, nil
	}
	t.Cleanup(func() { azureEntraProviderFactory = orig })

	var co clientOptions
	WithStreamingCredentialsProvider(explicit)(&co)

	opts, err := buildRedisOptions(&RedisConfig{
		ClusterAddresses: []string{"host:10000"},
		AzureEntra:       true,
	}, co)
	require.NoError(t, err)
	require.Same(t, explicit, opts.StreamingCredentialsProvider)
}
