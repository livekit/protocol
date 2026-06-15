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
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	entraid "github.com/redis/go-redis-entraid"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/auth"

	"github.com/livekit/protocol/xtls"

	"github.com/livekit/protocol/logger"
)

var ErrNotConfigured = errors.New("redis is not configured")

type RedisConfig struct {
	Address  string `yaml:"address,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	DB       int    `yaml:"db,omitempty"`
	// Deprecated: use TLS instead of UseTLS
	UseTLS            bool         `yaml:"use_tls,omitempty"`
	TLS               *xtls.Config `yaml:"tls,omitempty"`
	MasterName        string       `yaml:"sentinel_master_name,omitempty"`
	SentinelUsername  string       `yaml:"sentinel_username,omitempty"`
	SentinelPassword  string       `yaml:"sentinel_password,omitempty"`
	SentinelAddresses []string     `yaml:"sentinel_addresses,omitempty"`
	ClusterAddresses  []string     `yaml:"cluster_addresses,omitempty"`
	DialTimeout       int          `yaml:"dial_timeout,omitempty"`
	ReadTimeout       int          `yaml:"read_timeout,omitempty"`
	WriteTimeout      int          `yaml:"write_timeout,omitempty"`
	// for clustererd mode only, number of redirects to follow, defaults to 2
	MaxRedirects *int          `yaml:"max_redirects,omitempty"`
	PoolTimeout  time.Duration `yaml:"pool_timeout,omitempty"`
	PoolSize     int           `yaml:"pool_size,omitempty"`
	AzureEntra   bool          `yaml:"azure_entra,omitempty"`
}

func (r *RedisConfig) IsConfigured() bool {
	if r.Address != "" {
		return true
	}
	if len(r.SentinelAddresses) > 0 {
		return true
	}
	if len(r.ClusterAddresses) > 0 {
		return true
	}
	return false
}

func (r *RedisConfig) GetMaxRedirects() int {
	if r.MaxRedirects != nil {
		return *r.MaxRedirects
	}
	return 2
}

type clientOptions struct {
	streamingCredentialsProvider auth.StreamingCredentialsProvider
}

type Option func(*clientOptions)

func WithStreamingCredentialsProvider(p auth.StreamingCredentialsProvider) Option {
	return func(o *clientOptions) {
		o.streamingCredentialsProvider = p
	}
}

var azureEntraProviderFactory = newAzureEntraCredentialsProvider

func newAzureEntraCredentialsProvider() (auth.StreamingCredentialsProvider, error) {
	return entraid.NewDefaultAzureCredentialsProvider(entraid.DefaultAzureCredentialsProviderOptions{})
}

func buildRedisOptions(conf *RedisConfig, co clientOptions) (*redis.UniversalOptions, error) {
	var tlsConfig *tls.Config
	if conf.TLS != nil && conf.TLS.Enabled {
		var err error
		tlsConfig, err = conf.TLS.ClientTLSConfig()
		if err != nil {
			return nil, err
		}
	} else if conf.UseTLS {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	var rcOptions *redis.UniversalOptions
	if len(conf.SentinelAddresses) > 0 {
		logger.Infow("connecting to redis", "sentinel", true, "addr", conf.SentinelAddresses, "masterName", conf.MasterName)

		// By default DialTimeout set to 2s
		if conf.DialTimeout == 0 {
			conf.DialTimeout = 2000
		}
		// By default ReadTimeout set to 0.2s
		if conf.ReadTimeout == 0 {
			conf.ReadTimeout = 200
		}
		// By default WriteTimeout set to 0.2s
		if conf.WriteTimeout == 0 {
			conf.WriteTimeout = 200
		}

		rcOptions = &redis.UniversalOptions{
			Addrs:            conf.SentinelAddresses,
			SentinelUsername: conf.SentinelUsername,
			SentinelPassword: conf.SentinelPassword,
			MasterName:       conf.MasterName,
			Username:         conf.Username,
			Password:         conf.Password,
			DB:               conf.DB,
			TLSConfig:        tlsConfig,
			DialTimeout:      time.Duration(conf.DialTimeout) * time.Millisecond,
			ReadTimeout:      time.Duration(conf.ReadTimeout) * time.Millisecond,
			WriteTimeout:     time.Duration(conf.WriteTimeout) * time.Millisecond,
			PoolTimeout:      conf.PoolTimeout,
			PoolSize:         conf.PoolSize,
		}
	} else if len(conf.ClusterAddresses) > 0 {
		logger.Infow("connecting to redis", "cluster", true, "addr", conf.ClusterAddresses)
		rcOptions = &redis.UniversalOptions{
			Addrs:         conf.ClusterAddresses,
			Username:      conf.Username,
			Password:      conf.Password,
			DB:            conf.DB,
			TLSConfig:     tlsConfig,
			MaxRedirects:  conf.GetMaxRedirects(),
			PoolTimeout:   conf.PoolTimeout,
			PoolSize:      conf.PoolSize,
			IsClusterMode: true,
		}
	} else {
		logger.Infow("connecting to redis", "simple", true, "addr", conf.Address)
		rcOptions = &redis.UniversalOptions{
			Addrs:       []string{conf.Address},
			Username:    conf.Username,
			Password:    conf.Password,
			DB:          conf.DB,
			TLSConfig:   tlsConfig,
			PoolTimeout: conf.PoolTimeout,
			PoolSize:    conf.PoolSize,
		}
	}

	provider := co.streamingCredentialsProvider
	if provider == nil && conf.AzureEntra {
		p, err := azureEntraProviderFactory()
		if err != nil {
			return nil, fmt.Errorf("unable to create Azure Entra credentials provider: %w", err)
		}
		provider = p
	}
	if provider != nil {
		rcOptions.StreamingCredentialsProvider = provider
	}

	return rcOptions, nil
}

func GetRedisClient(conf *RedisConfig, opts ...Option) (redis.UniversalClient, error) {
	if conf == nil {
		return nil, nil
	}

	if !conf.IsConfigured() {
		return nil, ErrNotConfigured
	}

	var co clientOptions
	for _, opt := range opts {
		opt(&co)
	}

	rcOptions, err := buildRedisOptions(conf, co)
	if err != nil {
		return nil, err
	}

	rc := redis.NewUniversalClient(rcOptions)

	if err := rc.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("unable to connect to redis: %w", err)
	}

	return rc, nil
}
