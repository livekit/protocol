package redis

import (
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/livekit/protocol/logger"
)

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
}

func GetRedisClient(conf *RedisConfig) (*redis.Client, error) {
	if conf == nil {
		return nil, nil
	}

	var rc *redis.Client
	var tlsConfig *tls.Config

	if conf.UseTLS {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	if len(conf.SentinelAddresses) > 0 {
		logger.Infow("connecting to redis", "sentinel", true, "addr", conf.SentinelAddresses, "masterName", conf.MasterName)
		rcOptions := &redis.FailoverOptions{
			SentinelAddrs:    conf.SentinelAddresses,
			SentinelUsername: conf.SentinelUsername,
			SentinelPassword: conf.SentinelPassword,
			MasterName:       conf.MasterName,
			Username:         conf.Username,
			Password:         conf.Password,
			DB:               conf.DB,
			TLSConfig:        tlsConfig,
		}
		rc = redis.NewFailoverClient(rcOptions)
	} else {
		logger.Infow("connecting to redis", "sentinel", false, "addr", conf.Address)
		rcOptions := &redis.Options{
			Addr:      conf.Address,
			Username:  conf.Username,
			Password:  conf.Password,
			DB:        conf.DB,
			TLSConfig: tlsConfig,
		}
		rc = redis.NewClient(rcOptions)
	}

	if err := rc.Ping(context.Background()).Err(); err != nil {
		err = errors.Wrap(err, "unable to connect to redis")
		return nil, err
	}

	return rc, nil
}
