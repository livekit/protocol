package webhook

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"sync"
	"time"

	"github.com/frostbyte73/core"
	"github.com/gammazero/deque"
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

var ErrNotifierStopped = errors.New("notifier has already been stopped")

// URLNotifier is a QueuedNotifier that sends a POST request to a Webhook URL.
// It will retry on failure, and will drop events if notification fall too far behind
type URLNotifier struct {
	params    URLNotifierParams
	queue     *deque.Deque[*livekit.WebhookEvent]
	client    *retryablehttp.Client
	draining  atomic.Bool
	dropped   atomic.Int32
	fuse      core.Fuse
	mu        sync.RWMutex
	done      chan struct{}
	jobSignal chan struct{}
}

func NewURLNotifier(params URLNotifierParams) *URLNotifier {
	if params.QueueSize == 0 {
		params.QueueSize = defaultQueueSize
	}
	if params.Logger == nil {
		params.Logger = logger.GetLogger()
	}
	return &URLNotifier{
		params:    params,
		client:    retryablehttp.NewClient(),
		queue:     deque.New[*livekit.WebhookEvent](),
		fuse:      core.NewFuse(),
		jobSignal: make(chan struct{}, 1),
	}
}

func (n *URLNotifier) Start() {
	if n.done != nil {
		return
	}
	n.done = make(chan struct{})
	go n.worker()
}

func (n *URLNotifier) Stop(force bool) {
	if force {
		// triggers immediate closure
		n.fuse.Break()
	} else {
		// closes after current queue is processed
		n.draining.Store(true)
	}

	if !force {
		// wait for current queue to be processed
		<-n.done
	}
}

func (n *URLNotifier) SetKeys(apiKey, apiSecret string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.params.APIKey = apiKey
	n.params.APISecret = apiSecret
}

func (n *URLNotifier) QueueNotify(event *livekit.WebhookEvent) error {
	if n.draining.Load() || n.fuse.IsBroken() {
		return ErrNotifierStopped
	}
	n.mu.Lock()
	n.queue.PushBack(event)
	if n.queue.Len() > n.params.QueueSize {
		n.dropped.Inc()
		n.queue.PopFront()
	}
	n.mu.Unlock()
	select {
	case n.jobSignal <- struct{}{}:
	default:
	}
	return nil
}

func (n *URLNotifier) worker() {
	defer close(n.done)

	waitTicker := time.NewTicker(100 * time.Millisecond)
	for !n.fuse.IsBroken() {
		select {
		case <-waitTicker.C:
		case <-n.fuse.Watch():
			return
		case <-n.jobSignal:
			n.processQueue()
		}
		// when draining, ensure all events are processed and exit
		if n.draining.Load() {
			n.processQueue()
			return
		}
	}
}

func (n *URLNotifier) processQueue() {
	for event := n.nextItem(); event != nil && !n.fuse.IsBroken(); event = n.nextItem() {
		if err := n.send(event); err != nil {
			n.params.Logger.Warnw("failed to send webhook", err, "url", n.params.URL, "event", event.Event)
			n.dropped.Store(event.NumDropped + 1)
		}
	}
}

func (n *URLNotifier) nextItem() *livekit.WebhookEvent {
	n.mu.Lock()
	defer n.mu.Unlock()
	if n.queue.Len() == 0 {
		return nil
	}
	return n.queue.PopFront()
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
