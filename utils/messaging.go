package utils

import (
	"context"
	"math/rand"
	"time"

	"github.com/eapache/channels"
	"github.com/go-redis/redis/v8"
)

type MessageBus interface {
	Lock(ctx context.Context, key string, expiration time.Duration) (acquired bool, err error)
	Subscribe(ctx context.Context, channel string) (PubSub, error)
	Publish(ctx context.Context, channel string, msg interface{}) error
}

type PubSub interface {
	Channel() <-chan interface{}
	Payload(msg interface{}) []byte
	Close() error
}

type RedisMessageBus struct {
	rc *redis.Client
}

func NewRedisMessageBus(rc *redis.Client) *RedisMessageBus {
	return &RedisMessageBus{rc: rc}
}

func (r *RedisMessageBus) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return r.rc.SetNX(ctx, key, rand.Int(), expiration).Result()
}

func (r *RedisMessageBus) Subscribe(ctx context.Context, channel string) (PubSub, error) {
	ps := r.rc.Subscribe(ctx, channel)
	return &RedisPubSub{
		ps: ps,
		c:  channels.Wrap(ps.Channel()).Out(),
	}, nil
}

func (r *RedisMessageBus) Publish(ctx context.Context, channel string, message interface{}) error {
	return r.rc.Publish(ctx, channel, message).Err()
}

type RedisPubSub struct {
	ps *redis.PubSub
	c  <-chan interface{}
}

func (r *RedisPubSub) Channel() <-chan interface{} {
	return r.c
}

func (r *RedisPubSub) Payload(msg interface{}) []byte {
	return []byte(msg.(*redis.Message).Payload)
}

func (r *RedisPubSub) Close() error {
	return r.ps.Close()
}
