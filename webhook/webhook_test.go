package webhook

import (
	"context"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
)

const (
	apiKey               = "mykey"
	apiSecret            = "mysecret"
	testAddr             = ":8765"
	testUrl              = "http://localhost:8765"
	webhookCheckInterval = 500 * time.Millisecond
)

var authProvider = auth.NewSimpleKeyProvider(
	apiKey, apiSecret,
)

func TestWebHook(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewDefaultNotifier(apiKey, apiSecret, []string{testUrl})

	t.Run("test event payload", func(t *testing.T) {
		event := &livekit.WebhookEvent{
			Event: EventTrackPublished,
			Participant: &livekit.ParticipantInfo{
				Identity: "test",
			},
			Track: &livekit.TrackInfo{
				Sid: "TR_abcde",
			},
		}

		wg := sync.WaitGroup{}
		wg.Add(1)
		s.handler = func(r *http.Request) {
			defer wg.Done()
			decodedEvent, err := ReceiveWebhookEvent(r, authProvider)
			require.NoError(t, err)

			require.EqualValues(t, event, decodedEvent)
		}
		require.NoError(t, notifier.QueueNotify(context.Background(), event))
		wg.Wait()
	})

}

func TestURLNotifierDropped(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	urlNotifier := newTestNotifier()
	defer urlNotifier.Stop(true)
	totalDropped := atomic.Int32{}
	totalReceived := atomic.Int32{}
	s.handler = func(r *http.Request) {
		decodedEvent, err := ReceiveWebhookEvent(r, authProvider)
		require.NoError(t, err)
		totalReceived.Inc()
		totalDropped.Add(decodedEvent.NumDropped)
	}
	// send multiple notifications
	_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomStarted})
	_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventParticipantJoined})
	_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomFinished})

	time.Sleep(webhookCheckInterval)

	require.Equal(t, int32(3), totalDropped.Load()+totalReceived.Load())
	// at least one request dropped
	require.Greater(t, totalDropped.Load(), int32(0))
}

func TestURLNotifierLifecycle(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	t.Run("start/stop without use", func(t *testing.T) {
		urlNotifier := newTestNotifier()
		urlNotifier.Stop(false)
	})

	t.Run("stop allowing to drain", func(t *testing.T) {
		urlNotifier := newTestNotifier()
		numCalled := atomic.Int32{}
		s.handler = func(r *http.Request) {
			numCalled.Inc()
		}
		require.NoError(t, urlNotifier.QueueNotify(&livekit.WebhookEvent{
			Event: EventRoomStarted,
		}))
		urlNotifier.Stop(false)
		time.Sleep(webhookCheckInterval)
		require.Equal(t, int32(1), numCalled.Load())
	})

	t.Run("force stop", func(t *testing.T) {
		urlNotifier := newTestNotifier()
		numCalled := atomic.Int32{}
		s.handler = func(r *http.Request) {
			numCalled.Inc()
		}
		go func() {
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{
				Event: EventRoomStarted,
			})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{
				Event: EventRoomStarted,
			})
		}()
		urlNotifier.Stop(true)
		time.Sleep(webhookCheckInterval)
		require.Less(t, numCalled.Load(), int32(2))
	})
}

func newTestNotifier() *URLNotifier {
	return NewURLNotifier(URLNotifierParams{
		QueueSize: 1,
		URL:       testUrl,
		APIKey:    apiKey,
		APISecret: apiSecret,
	})
}

type testServer struct {
	handler func(r *http.Request)
	server  *http.Server
}

func newServer(addr string) *testServer {
	s := &testServer{}
	s.server = &http.Server{
		Addr:    addr,
		Handler: s,
	}
	return s
}

func (s *testServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.handler != nil {
		s.handler(r)
	}
}

func (s *testServer) Start() error {
	l, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return err
	}
	go s.server.Serve(l)
	return nil
}

func (s *testServer) Stop() {
	_ = s.server.Shutdown(context.Background())
}
