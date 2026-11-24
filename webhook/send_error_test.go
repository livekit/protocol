// Copyright 2026 LiveKit, Inc.
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
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/netip"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"

	"github.com/livekit/protocol/livekit"
)

func TestRedactSendErrorThroughFieldsHook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer srv.Close()

	redact := func(whi *livekit.WebhookInfo) { whi.SendError = RedactSendError(whi.SendError) }
	require.Equal(t, noAllowedAddressText, deliverOnce(t, srv.URL, redact))
	require.Contains(t, deliverOnce(t, srv.URL, nil), "127.0.0.1")
}

func deliverOnce(t *testing.T, url string, fieldsHook func(*livekit.WebhookInfo)) string {
	InitWebhookStats(prometheus.Labels{})
	n := NewResourceURLNotifier(ResourceURLNotifierParams{
		HTTPClientParams: HTTPClientParams{
			MaxRetries:   1,
			RetryWaitMin: time.Millisecond,
			RetryWaitMax: time.Millisecond,
			DialControl:  BlockInternal,
		},
		URL:        url,
		APIKey:     testAPIKey,
		APISecret:  testAPISecret,
		FieldsHook: fieldsHook,
	})
	infos := make(chan *livekit.WebhookInfo, 1)
	n.RegisterProcessedHook(func(_ context.Context, whi *livekit.WebhookInfo) { infos <- whi })

	event := &livekit.WebhookEvent{Event: EventRoomStarted, Room: &livekit.Room{Name: "room"}}
	require.NoError(t, n.QueueNotify(context.Background(), event))
	n.Stop(false)
	return (<-infos).SendError
}

func TestRedactSendError(t *testing.T) {
	giveUp := func(err error) error {
		return fmt.Errorf("giving up after 2 attempt(s): %w", &url.Error{Op: "Post", URL: "http://foo.ns.svc.cluster.local", Err: err})
	}
	blocked := giveUp(&net.OpError{
		Op:   "dial",
		Net:  "tcp",
		Addr: &net.TCPAddr{IP: net.ParseIP("10.0.12.7"), Port: 443},
		Err:  fmt.Errorf("%w: %s", ErrBlockedAddress, "10.0.12.7"),
	})
	notFound := giveUp(&net.OpError{Op: "dial", Net: "tcp", Err: &net.DNSError{
		Err: "no such host", Name: "foo.ns.svc.cluster.local", Server: "10.96.0.10:53", IsNotFound: true,
	}})
	require.Equal(t, noAllowedAddressText, RedactSendError(blocked.Error()))
	require.Equal(t, noAllowedAddressText, RedactSendError(notFound.Error()))

	for _, uniform := range []string{
		`Post "http://foo.ns.svc.cluster.local": dial tcp 10.0.12.7:443: blocked destination address: 10.0.12.7 (Client.Timeout exceeded while awaiting headers)`,
		`Post "http://foo.ns.svc.cluster.local": dial tcp: lookup foo.ns.svc.cluster.local on 10.96.0.10:53: no such host (Client.Timeout exceeded while awaiting headers)`,
		`Post "http://foo.ns.svc.cluster.local": dial tcp4: address foo.ns.svc.cluster.local: no suitable address found`,
	} {
		require.Equal(t, noAllowedAddressText, RedactSendError(uniform), uniform)
	}

	cases := []struct {
		in, hides, keeps string
	}{
		{`lookup example.com on 10.96.0.10:53: i/o timeout`, "10.96.0.10", "i/o timeout"},
		{`dial tcp [fe80::1%eth0]:443: i/o timeout`, "fe80::1", "i/o timeout"},
		{`read tcp 10.1.2.3:51234->93.184.215.14:443: read: connection reset by peer`, "10.1.2.3", "93.184.215.14"},
		{`dial ip fd00::: i/o timeout`, "fd00::", "i/o timeout"},
		{`at 2026-11-17T20:30:45Z`, "", "20:30:45"},
	}
	for _, c := range cases {
		out := RedactSendError(c.in)
		if c.hides != "" {
			require.NotContains(t, out, c.hides, c.in)
		}
		require.Contains(t, out, c.keeps, c.in)
	}
}

var sendErrorShapes = []func(addr string) string{
	func(a string) string {
		return "dial tcp " + net.JoinHostPort(a, "80") + ": connect: connection refused"
	},
	func(a string) string { return "lookup x.example on " + net.JoinHostPort(a, "53") + ": i/o timeout" },
	func(a string) string {
		return "read tcp " + net.JoinHostPort(a, "51234") + "->" + net.JoinHostPort(a, "443") + ": connection reset by peer"
	},
	func(a string) string { return "dial ip " + a + ": i/o timeout" },
	func(a string) string { return "connect to " + a + "." },
	func(a string) string {
		return "dial tcp " + net.JoinHostPort(a, "443") + ": i/o timeout (Client.Timeout exceeded while awaiting headers)"
	},
}

func FuzzRedactSendError(f *testing.F) {
	for _, s := range []string{
		"10.0.0.1", "127.0.0.1", "169.254.169.254", "8.8.8.8", "93.184.215.14",
		"::1", "fd00::", "fe80::1", "::ffff:10.0.0.1", "::ffff:8.8.8.8", "1::", "2606:4700:4700::1111",
	} {
		raw := netip.MustParseAddr(s).AsSlice()
		for shape := range sendErrorShapes {
			f.Add(raw, uint8(shape))
		}
	}
	f.Fuzz(func(t *testing.T, raw []byte, shape uint8) {
		addr, ok := netip.AddrFromSlice(raw)
		if !ok {
			t.Skip()
		}
		text := addr.String()
		in := sendErrorShapes[int(shape)%len(sendErrorShapes)](text)
		want := in
		if isBlockedIP(raw) {
			want = strings.ReplaceAll(in, text, "internal-address")
		}
		require.Equal(t, want, RedactSendError(in))
	})
}
