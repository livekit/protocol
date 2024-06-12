package webhook

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"google.golang.org/protobuf/encoding/protojson"
	"sync"
	"time"
)

const (
	DefaultBatchSendInterval = 100 * time.Millisecond
	DefaultMaxBatchSize      = 10000
)

type BatchURLNotifierParams struct {
	Logger    logger.Logger
	URL       string
	Interval  time.Duration
	MaxSize   int
	APIKey    string
	APISecret string
}

type BatchURLNotifier struct {
	cancelFunc context.CancelFunc
	client     *retryablehttp.Client
	mu         sync.RWMutex
	params     BatchURLNotifierParams
	batch      []*livekit.WebhookEvent
	dropped    int // it's operated inside a mutex scope so no need for atomic type
}

func NewBatchURLNotifier(ctx context.Context, params BatchURLNotifierParams) URLNotifier {
	if params.Interval == 0 {
		params.Interval = DefaultBatchSendInterval
	}
	if params.MaxSize == 0 {
		params.MaxSize = DefaultMaxBatchSize
	}

	ctx, cancel := context.WithCancel(ctx)
	notifier := &BatchURLNotifier{
		cancelFunc: cancel,
		params:     params,
		client:     retryablehttp.NewClient(),
	}

	go notifier.runner(ctx)

	return notifier
}

func (b *BatchURLNotifier) runner(ctx context.Context) {
	ticker := time.NewTicker(b.params.Interval)
	for {
		select {
		case <-ticker.C:
			b.mu.Lock()
			b.sendBatch()
			b.mu.Unlock()
		case <-ctx.Done():
			return
		}
	}
}

func (b *BatchURLNotifier) sendBatch() {
	if len(b.batch) == 0 {
		return
	}
	raw := &livekit.BatchedWebhookEvents{
		Events:     b.batch,
		NumDropped: int32(b.dropped),
		DequeuedAt: time.Now().Unix(),
	}
	defer func() {
		b.batch = nil
	}()
	b.dropped = 0

	encoded, err := protojson.Marshal(raw)
	if err != nil {
		b.params.Logger.Warnw("Failed to marshal event", err)
		b.dropped += len(b.batch)
		return
	}

	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])
	at := auth.NewAccessToken(b.params.APIKey, b.params.APISecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		b.params.Logger.Warnw("Failed to generate jwt token", err)
		b.dropped += len(b.batch)
		return
	}

	req, err := retryablehttp.NewRequest("POST", b.params.URL, bytes.NewReader(encoded))
	if err != nil {
		b.params.Logger.Warnw("Failed to create http req", err)
		b.dropped += len(b.batch)
		return
	}

	req.Header.Set(authHeader, token)
	req.Header.Set("batched", "true")
	req.Header.Set("content-type", "application/webhook+json")
	resp, err := b.client.Do(req)
	if err != nil {
		b.params.Logger.Errorw("Failed to send request", err)
		b.dropped += len(b.batch)
		return
	}
	_ = resp.Body.Close()

	return
}

func (b *BatchURLNotifier) SetKeys(apiKey, apiSecret string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.params.APIKey = apiKey
	b.params.APISecret = apiSecret
}

func (b *BatchURLNotifier) QueueNotify(event *livekit.WebhookEvent) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.batch = append(b.batch, event)

	if len(b.batch) >= b.params.MaxSize {
		b.sendBatch()
	}

	return nil
}

func (b *BatchURLNotifier) Stop(force bool) {
	b.cancelFunc()
	if !force {
		b.mu.Lock()
		b.sendBatch()
		b.mu.Unlock()
	}
}
