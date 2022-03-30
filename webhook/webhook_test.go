package webhook_test

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/webhook"
)

const (
	apiKey    = "mykey"
	apiSecret = "mysecret"
)

func TestWebHook(t *testing.T) {
	s := newServer(":8765")
	require.NoError(t, s.Start())
	defer s.Stop()

	authProvider := auth.NewSimpleKeyProvider(
		apiKey, apiSecret,
	)
	notifier := webhook.NewNotifier(apiKey, apiSecret, []string{
		"http://localhost:8765",
	})

	t.Run("test json payload", func(t *testing.T) {
		payload := map[string]interface{}{
			"test": "payload",
			"nested": map[string]interface{}{
				"structure": true,
			},
		}

		wg := sync.WaitGroup{}
		wg.Add(1)
		s.handler = func(r *http.Request) {
			defer wg.Done()
			// receive logic
			data, err := webhook.Receive(r, authProvider)
			require.NoError(t, err)

			var decoded map[string]interface{}
			require.NoError(t, json.Unmarshal(data, &decoded))

			require.EqualValues(t, decoded, payload)
		}

		require.NoError(t, notifier.Notify(context.Background(), payload))
		wg.Wait()
	})

	t.Run("test event payload", func(t *testing.T) {
		event := &livekit.WebhookEvent{
			Event: webhook.EventTrackPublished,
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
			decodedEvent, err := webhook.ReceiveWebhookEvent(r, authProvider)
			require.NoError(t, err)

			require.EqualValues(t, event, decodedEvent)
		}
		require.NoError(t, notifier.Notify(context.Background(), event))
		wg.Wait()
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
