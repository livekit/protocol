// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	webhookCheckInterval = 100 * time.Millisecond
)

var authProvider = auth.NewSimpleKeyProvider(
	apiKey, apiSecret,
)

var sampleEvents = []*livekit.WebhookEvent{
	{
		Event: EventTrackPublished,
		Participant: &livekit.ParticipantInfo{
			Identity: "test",
		},
		Track: &livekit.TrackInfo{
			Sid: "TR_abcde",
		},
	},
	{
		Event: EventParticipantJoined,
		Participant: &livekit.ParticipantInfo{
			Identity: "test",
		},
		Track: &livekit.TrackInfo{
			Sid: "TR_abcde",
		},
	},
	{
		Event: EventParticipantJoined,
		Participant: &livekit.ParticipantInfo{
			Identity: "test",
		},
		Track: &livekit.TrackInfo{
			Sid: "TR_abcde",
		},
	},
	{
		Event: EventRoomFinished,
		Participant: &livekit.ParticipantInfo{
			Identity: "test",
		},
		Track: &livekit.TrackInfo{
			Sid: "TR_abcde",
		},
	},
	{
		Event: EventRoomStarted,
		Participant: &livekit.ParticipantInfo{
			Identity: "test",
		},
		Track: &livekit.TrackInfo{
			Sid: "TR_abcde",
		},
	},
}

func TestWebHook(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewDefaultNotifier(apiKey, apiSecret, []string{testUrl}, true, nil, nil)
	defer notifier.Stop(true)

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

func TestBatchWebHook(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewBatchedNotifier(context.Background(), apiKey, apiSecret, []string{testUrl}, nil, nil)
	defer notifier.Stop(true)

	t.Run("test events payload", func(t *testing.T) {
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
			decodedEvent, err := ReceiveWebhookEventBatched(r, authProvider)
			require.NoError(t, err)

			require.Equal(t, 3, len(decodedEvent))
			for _, ev := range decodedEvent {
				require.EqualValues(t, event, ev)
			}
		}
		// send 3 times
		require.NoError(t, notifier.QueueNotify(context.Background(), event))
		require.NoError(t, notifier.QueueNotify(context.Background(), event))
		require.NoError(t, notifier.QueueNotify(context.Background(), event))
		wg.Wait()
	})

}

func TestWebHookWithIncludeFilter(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewDefaultNotifier(apiKey, apiSecret, []string{testUrl}, true, []string{EventParticipantJoined, EventTrackPublished}, nil)
	defer notifier.Stop(true)

	t.Run("test event payload", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(3)
		counter := 0
		s.handler = func(r *http.Request) {
			defer wg.Done()
			decodedEvent, err := ReceiveWebhookEvent(r, authProvider)
			require.NoError(t, err)

			require.EqualValues(t, sampleEvents[counter], decodedEvent)
			counter++
		}
		for _, ev := range sampleEvents {
			require.NoError(t, notifier.QueueNotify(context.Background(), ev))
		}
		wg.Wait()
		require.Equal(t, 3, counter)
	})
}

func TestBatchWebHookWithIncludeFilter(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewBatchedNotifier(context.Background(), apiKey, apiSecret, []string{testUrl}, []string{EventParticipantJoined, EventTrackPublished}, nil)
	defer notifier.Stop(true)

	t.Run("test events payload", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		s.handler = func(r *http.Request) {
			defer wg.Done()
			decodedEvent, err := ReceiveWebhookEventBatched(r, authProvider)
			require.NoError(t, err)

			require.Equal(t, 3, len(decodedEvent))
			for idx, ev := range decodedEvent {
				require.EqualValues(t, sampleEvents[idx], ev)
			}
		}
		for _, ev := range sampleEvents {
			require.NoError(t, notifier.QueueNotify(context.Background(), ev))
		}
		wg.Wait()
	})

}

func TestWebHookWithExcludeFilter(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewDefaultNotifier(apiKey, apiSecret, []string{testUrl}, true, nil, []string{EventParticipantJoined, EventTrackPublished})
	defer notifier.Stop(true)

	t.Run("test event payload", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(2)
		counter := 0
		s.handler = func(r *http.Request) {
			defer wg.Done()
			decodedEvent, err := ReceiveWebhookEvent(r, authProvider)
			require.NoError(t, err)

			require.EqualValues(t, sampleEvents[counter+3], decodedEvent)
			counter++
		}
		for _, ev := range sampleEvents {
			require.NoError(t, notifier.QueueNotify(context.Background(), ev))
		}
		wg.Wait()
		require.Equal(t, 2, counter)
	})
}

func TestBatchWebHookWithExcludeFilter(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	notifier := NewBatchedNotifier(context.Background(), apiKey, apiSecret, []string{testUrl}, nil, []string{EventParticipantJoined, EventTrackPublished})
	defer notifier.Stop(true)

	t.Run("test events payload", func(t *testing.T) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		s.handler = func(r *http.Request) {
			defer wg.Done()
			decodedEvent, err := ReceiveWebhookEventBatched(r, authProvider)
			require.NoError(t, err)

			require.Equal(t, 2, len(decodedEvent))
			for idx, ev := range decodedEvent {
				require.EqualValues(t, sampleEvents[idx+3], ev)
			}
		}
		for _, ev := range sampleEvents {
			require.NoError(t, notifier.QueueNotify(context.Background(), ev))
		}
		wg.Wait()
	})

}

func TestURLNotifierDropped(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	t.Run("DropWhenFull = true", func(t *testing.T) {
		urlNotifier := newTestNotifier(true)
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
		for i := 0; i < 10; i++ {
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomStarted})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventParticipantJoined})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomFinished})
		}

		time.Sleep(webhookCheckInterval)

		require.Equal(t, int32(30), totalDropped.Load()+totalReceived.Load())
		// at least one request dropped
		require.Less(t, int32(0), totalDropped.Load())
	})

	t.Run("DropWhenFull = false", func(t *testing.T) {
		urlNotifier := newTestNotifier(false)
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
		for i := 0; i < 10; i++ {
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomStarted})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventParticipantJoined})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomFinished})
		}

		time.Sleep(webhookCheckInterval)

		require.Equal(t, int32(30), totalDropped.Load()+totalReceived.Load())
		// at least one request dropped
		require.Equal(t, int32(0), totalDropped.Load())
	})
}

func TestURLNotifierLifecycle(t *testing.T) {
	s := newServer(testAddr)
	require.NoError(t, s.Start())
	defer s.Stop()

	t.Run("start/stop without use", func(t *testing.T) {
		urlNotifier := newTestNotifier(true)
		urlNotifier.Stop(false)
	})

	t.Run("stop allowing to drain", func(t *testing.T) {
		urlNotifier := newTestNotifier(true)
		numCalled := atomic.Int32{}
		s.handler = func(r *http.Request) {
			numCalled.Inc()
		}
		for i := 0; i < 10; i++ {
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomStarted})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomFinished})
		}
		urlNotifier.Stop(false)
		require.Eventually(t, func() bool { return numCalled.Load() == 20 }, 5*time.Second, webhookCheckInterval)
	})

	t.Run("force stop", func(t *testing.T) {
		urlNotifier := newTestNotifier(true)
		numCalled := atomic.Int32{}
		s.handler = func(r *http.Request) {
			numCalled.Inc()
		}
		for i := 0; i < 10; i++ {
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomStarted})
			_ = urlNotifier.QueueNotify(&livekit.WebhookEvent{Event: EventRoomFinished})
		}
		urlNotifier.Stop(true)
		time.Sleep(time.Second)
		require.Greater(t, int32(20), numCalled.Load())
	})
}

func newTestNotifier(dropWhenFull bool) URLNotifier {
	return NewDefaultURLNotifier(URLNotifierParams{
		QueueSize:    20,
		URL:          testUrl,
		APIKey:       apiKey,
		APISecret:    apiSecret,
		DropWhenFull: dropWhenFull,
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
