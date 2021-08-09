package webhook_test

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"testing"

	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/webhook"
	"github.com/stretchr/testify/require"
)

func TestWebHook(t *testing.T) {
	s := newServer(":8765")
	require.NoError(t, s.Start())
	defer s.Stop()

	authProvider := auth.NewFileBasedKeyProviderFromMap(map[string]string{
		"mykey": "mysecret",
	})

	payload := map[string]interface{}{
		"test": "payload",
		"nested": map[string]interface{}{
			"structure": true,
		},
	}

	s.handler = func(r *http.Request) {
		// receive logic
		data, err := webhook.Receive(r, authProvider)
		require.NoError(t, err)

		var decoded map[string]interface{}
		require.NoError(t, json.Unmarshal(data, &decoded))

		require.EqualValues(t, decoded, payload)
	}

	notifier := webhook.NewNotifier("mykey", "mysecret", []string{
		"http://localhost:8765",
	})
	require.NoError(t, notifier.Notify(payload))
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
