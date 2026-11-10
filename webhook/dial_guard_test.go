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
	"net/http"
	"net/http/httptest"
	"net/netip"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/require"
)

func TestCheckWebhookAddr(t *testing.T) {
	for _, addr := range []string{"93.184.215.14", "8.8.8.8", "2606:4700::1111", "::ffff:8.8.8.8"} {
		require.NoError(t, checkWebhookAddr(netip.MustParseAddr(addr)), addr)
	}
	for _, addr := range []string{
		"127.0.0.1", "127.1.2.3", "::1", // loopback
		"10.1.2.3", "172.16.0.1", "192.168.1.1", "fd00::1", // private
		"169.254.169.254", "fe80::1", // link-local, cloud metadata
		"0.0.0.0", "0.1.2.3", "::", // unspecified, "this network"
		"224.0.0.1", "ff02::1", "255.255.255.255", // multicast, broadcast
		"100.64.0.1", "198.18.0.1", "240.0.0.1", // CGNAT, benchmarking, reserved
		"::ffff:127.0.0.1", "::ffff:10.0.0.1", "::ffff:169.254.169.254", // IPv4-mapped
	} {
		require.ErrorIs(t, checkWebhookAddr(netip.MustParseAddr(addr)), ErrWebhookTargetBlocked, addr)
	}
}

func TestCheckWebhookHost(t *testing.T) {
	for _, host := range []string{"example.com", "93.184.215.14", "2606:4700::1111"} {
		require.NoError(t, CheckWebhookHost(host), host)
	}
	for _, host := range []string{
		"localhost", "LOCALHOST.", "api.localhost",
		"127.0.0.1", "10.0.0.1", "169.254.169.254",
		"::1", "::ffff:10.0.0.1", "0.0.0.0",
	} {
		require.ErrorIs(t, CheckWebhookHost(host), ErrWebhookTargetBlocked, host)
	}
}

func TestBlockPrivateAddresses(t *testing.T) {
	var hits atomic.Int32
	receiver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
	}))
	t.Cleanup(receiver.Close)

	u, err := url.Parse(receiver.URL)
	require.NoError(t, err)
	byIP := receiver.URL
	byName := "http://localhost:" + u.Port()

	post := func(t *testing.T, params HTTPClientParams, target string) error {
		params.RetryWaitMin = time.Millisecond
		params.RetryWaitMax = time.Millisecond
		req, err := retryablehttp.NewRequest("POST", target, nil)
		require.NoError(t, err)
		res, err := newHTTPClient(params).Do(req)
		if err != nil {
			return err
		}
		_ = res.Body.Close()
		return nil
	}

	t.Run("delivers to private addresses when off", func(t *testing.T) {
		hits.Store(0)
		require.NoError(t, post(t, HTTPClientParams{}, byIP))
		require.EqualValues(t, 1, hits.Load())
	})

	for name, params := range map[string]HTTPClientParams{
		"refuses private addresses":                {BlockPrivateAddresses: true},
		"refuses private addresses with ForceIPv4": {BlockPrivateAddresses: true, ForceIPv4: true},
	} {
		t.Run(name, func(t *testing.T) {
			hits.Store(0)
			// The dial check covers stored URLs saved before the URL check, and
			// hostnames that resolve to a blocked address, like localhost here.
			for _, target := range []string{byIP, byName} {
				require.ErrorIs(t, post(t, params, target), ErrWebhookTargetBlocked, target)
			}
			require.Zero(t, hits.Load())
		})
	}
}
