package redis

import (
	"context"
	"crypto/tls"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	UseTLS   bool   `yaml:"use_tls"`
}

func GetRedisClient(conf *RedisConfig) (*redis.Client, error) {
	if conf == nil {
		return nil, nil
	}

	var tlsConfig *tls.Config
	if conf.UseTLS {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
	rc := redis.NewClient(&redis.Options{
		Addr:      conf.Address,
		Username:  conf.Username,
		Password:  conf.Password,
		DB:        conf.DB,
		TLSConfig: tlsConfig,
	})

	err := rc.Ping(context.Background()).Err()
	return rc, err
}
