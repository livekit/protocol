package webhook

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/atomic"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
)

type URLNotifierParams struct {
	Logger    logger.Logger
	QueueSize int
	URL       string
	APIKey    string
	APISecret string
}

const defaultQueueSize = 100

// URLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It will retry on failure, and will drop events if notification fall too far behind
type URLNotifier struct {
	mu      sync.RWMutex
	params  URLNotifierParams
	client  *retryablehttp.Client
	dropped atomic.Int32
	worker  core.QueueWorker
}

func NewURLNotifier(params URLNotifierParams) *URLNotifier {
	if params.QueueSize == 0 {
		params.QueueSize = defaultQueueSize
	}
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}

	n := &URLNotifier{
		params: params,
		client: retryablehttp.NewClient(),
	}
	n.worker = core.NewQueueWorker(core.QueueWorkerParams{
		QueueSize:    params.QueueSize,
		DropWhenFull: true,
		OnDropped:    func() { n.dropped.Inc() },
	})
	return n
}

func (n *URLNotifier) SetKeys(apiKey, apiSecret string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.params.APIKey = apiKey
	n.params.APISecret = apiSecret
}

func (n *URLNotifier) QueueNotify(event *livekit.WebhookEvent) error {
	n.worker.Submit(func() {
		if err := n.send(event); err != nil {
			n.params.Logger.Warnw("failed to send webhook", err, "url", n.params.URL, "event", event.Event)
			n.dropped.Add(event.NumDropped + 1)
		} else {
			n.params.Logger.Infow("sent webhook", "url", n.params.URL, "event", event.Event)
		}
	})
	return nil
}

func (n *URLNotifier) Stop(force bool) {
	if force {
		n.worker.Kill()
	} else {
		n.worker.Drain()
	}
}

func (n *URLNotifier) send(event *livekit.WebhookEvent) error {
	// set dropped count
	event.NumDropped = n.dropped.Swap(0)
	encoded, err := protojson.Marshal(event)
	if err != nil {
		return err
	}
	// sign payload
	sum := sha256.Sum256(encoded)
	b64 := base64.StdEncoding.EncodeToString(sum[:])

	n.mu.RLock()
	apiKey := n.params.APIKey
	apiSecret := n.params.APISecret
	n.mu.RUnlock()

	at := auth.NewAccessToken(apiKey, apiSecret).
		SetValidFor(5 * time.Minute).
		SetSha256(b64)
	token, err := at.ToJWT()
	if err != nil {
		return err
	}
	r, err := retryablehttp.NewRequest("POST", n.params.URL, bytes.NewReader(encoded))
	if err != nil {
		// ignore and continue
		return err
	}
	r.Header.Set(authHeader, token)
	// use a custom mime type to ensure signature is checked prior to parsing
	r.Header.Set("content-type", "application/webhook+json")
	res, err := n.client.Do(r)
	if err != nil {
		return err
	}
	_ = res.Body.Close()
	return nil
}
