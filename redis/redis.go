package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/abdulhaseeb08/protocol/logger"
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
	ClusterAddresses  []string `yaml:"cluster_addresses"`
}

func GetRedisClient(conf *RedisConfig) (redis.UniversalClient, error) {
	if conf == nil {
		return nil, nil
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
			Addrs:     conf.ClusterAddresses,
			Username:  conf.Username,
			Password:  conf.Password,
			DB:        conf.DB,
			TLSConfig: tlsConfig,
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
	fmt.Println("Address: ", conf.Address)
	rc = redis.NewUniversalClient(rcOptions)

	if err := rc.Ping(context.Background()).Err(); err != nil {
		err = errors.Wrap(err, "unable to connect to redis")
		return nil, err
	}

	return rc, nil
}

func GetRedisClientTest(conf *RedisConfig) (redis.UniversalClient, error) {
	if conf == nil {
		return nil, nil
	}

	var rcOptions *redis.UniversalOptions
	var rc redis.UniversalClient

	rcOptions = &redis.UniversalOptions{
		Addrs: []string{"localhost:6379"},
	}
	rc = redis.NewUniversalClient(rcOptions)

	if err := rc.Ping(context.Background()).Err(); err != nil {
		fmt.Println(rcOptions.Addrs)
		err = errors.Wrap(err, "unable to connect to redis")
		return nil, err
	}

	return rc, nil
}
