package redis

import (
	"context"
	"crypto/tls"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/livekit/protocol/logger"
)

var ErrNotConfigured = errors.New("Redis is not configured")

type RedisConfig struct {
	Address           string   `yaml:"address"`
	Username          string   `yaml:"username"`
	Password          string   `yaml:"password"`
	DB                int      `yaml:"db"`
	UseTLS            bool     `yaml:"use_tls"`
	MasterName        string   `yaml:"sentinel_master_name"`
	SentinelUsername  string   `yaml:"sentinel_username"`
	SentinelPassword  string   `yaml:"sentinel_password"`
	SentinelAddresses []string `yaml:"sentinel_addresses"`
	ClusterAddresses  []string `yaml:"cluster_addresses"`
	// for clustererd mode only, number of redirects to follow, defaults to 2
	MaxRedirects *int `yaml:"max_redirects"`
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

func GetRedisClient(conf *RedisConfig) (redis.UniversalClient, error) {
	if conf == nil {
		return nil, nil
	}

	if !conf.IsConfigured() {
		return nil, ErrNotConfigured
	}

	var rcOptions *redis.UniversalOptions
	var rc redis.UniversalClient
	var tlsConfig *tls.Config

	if conf.UseTLS {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	if len(conf.SentinelAddresses) > 0 {
		logger.Infow("connecting to redis", "sentinel", true, "addr", conf.SentinelAddresses, "masterName", conf.MasterName)
		rcOptions = &redis.UniversalOptions{
			Addrs:            conf.SentinelAddresses,
			SentinelUsername: conf.SentinelUsername,
			SentinelPassword: conf.SentinelPassword,
			MasterName:       conf.MasterName,
			Username:         conf.Username,
			Password:         conf.Password,
			DB:               conf.DB,
			TLSConfig:        tlsConfig,
		}
	} else if len(conf.ClusterAddresses) > 0 {
		logger.Infow("connecting to redis", "cluster", true, "addr", conf.ClusterAddresses)
		rcOptions = &redis.UniversalOptions{
			Addrs:        conf.ClusterAddresses,
			Username:     conf.Username,
			Password:     conf.Password,
			DB:           conf.DB,
			TLSConfig:    tlsConfig,
			MaxRedirects: conf.GetMaxRedirects(),
		}
	} else {
		logger.Infow("connecting to redis", "simple", true, "addr", conf.Address)
		rcOptions = &redis.UniversalOptions{
			Addrs:     []string{conf.Address},
			Username:  conf.Username,
			Password:  conf.Password,
			DB:        conf.DB,
			TLSConfig: tlsConfig,
		}
	}
	rc = redis.NewUniversalClient(rcOptions)

	if err := rc.Ping(context.Background()).Err(); err != nil {
		err = errors.Wrap(err, "unable to connect to redis")
		return nil, err
	}

	return rc, nil
}
