package utils

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/eapache/channels"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"
)

const lockExpiration = time.Second * 5

var (
	PromMessageBusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "livekit",
			Subsystem: "messagebus",
			Name:      "messages",
		},
		[]string{"type", "status"},
	)
)

func init() {
	prometheus.MustRegister(PromMessageBusCounter)
}

type MessageBus interface {
	Subscribe(ctx context.Context, channel string) (PubSub, error)
	// like subscribe, but ensuring only a single instance gets to process the message
	SubscribeQueue(ctx context.Context, channel string) (PubSub, error)
	Publish(ctx context.Context, channel string, msg proto.Message) error
}

type PubSub interface {
	Channel() <-chan interface{}
	Payload(msg interface{}) []byte
	Close() error
}

type RedisMessageBus struct {
	rc *redis.Client
}

func NewRedisMessageBus(rc *redis.Client) MessageBus {
	return &RedisMessageBus{rc: rc}
}

func (r *RedisMessageBus) Lock(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return r.rc.SetNX(ctx, key, rand.Int(), expiration).Result()
}

func (r *RedisMessageBus) Subscribe(ctx context.Context, channel string) (PubSub, error) {
	ps := r.rc.Subscribe(ctx, channel)
	return &RedisPubSub{
		ps:   ps,
		c:    channels.Wrap(ps.Channel()).Out(),
		done: make(chan struct{}, 1),
	}, nil
}

func (r *RedisMessageBus) SubscribeQueue(ctx context.Context, channel string) (PubSub, error) {
	sub := r.rc.Subscribe(ctx, channel)
	c := make(chan *redis.Message, 100) // same chan size as redis pubsub
	ps := &RedisPubSub{
		ps:   sub,
		c:    channels.Wrap(c).Out(),
		done: make(chan struct{}, 1),
	}

	go func() {
		for {
			select {
			case <-ps.done:
				return
			case msg := <-sub.Channel():
				sha := sha256.Sum256([]byte(msg.Payload))
				hash := base64.StdEncoding.EncodeToString(sha[:])
				acquired, _ := r.Lock(ctx, hash, lockExpiration)
				if acquired {
					PromMessageBusCounter.WithLabelValues("in", "success").Add(1)
					c <- msg
				}
			}
		}
	}()

	return ps, nil
}

func (r *RedisMessageBus) Publish(ctx context.Context, channel string, msg proto.Message) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		PromMessageBusCounter.WithLabelValues("out", "failure").Add(1)
		return err
	}

	err = r.rc.Publish(ctx, channel, b).Err()
	if err == nil {
		PromMessageBusCounter.WithLabelValues("out", "success").Add(1)
	} else {
		PromMessageBusCounter.WithLabelValues("out", "failure").Add(1)
	}

	return err
}

type RedisPubSub struct {
	ps   *redis.PubSub
	c    <-chan interface{}
	done chan struct{}
}

func (r *RedisPubSub) Channel() <-chan interface{} {
	return r.c
}

func (r *RedisPubSub) Payload(msg interface{}) []byte {
	return []byte(msg.(*redis.Message).Payload)
}

func (r *RedisPubSub) Close() error {
	r.done <- struct{}{}
	return r.ps.Close()
}
