package rpc

import (
	"context"
	"errors"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

func TestErrorCodeLabel(t *testing.T) {
	InitPSRPCStats(prometheus.Labels{})

	info := psrpc.RPCInfo{Service: "svc", Method: "meth"}
	o := PSRPCMetricsObserver{}
	o.OnUnaryRequest(middleware.ClientRole, info, 0, psrpc.NewErrorf(psrpc.Unavailable, "x"), 0, 0)
	o.OnStreamSend(middleware.ClientRole, info, 0, errors.New("bare"), 0)
	o.OnStreamRecv(middleware.ServerRole, info, context.Canceled, 0)
	o.OnMultiRequest(middleware.ClientRole, info, 0, 0, 1, 0, 0)

	families, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	var family *dto.MetricFamily
	for _, f := range families {
		if f.GetName() == "livekit_psrpc_error_total" {
			family = f
			break
		}
	}
	require.NotNil(t, family, "livekit_psrpc_error_total not registered")

	codes := map[string]bool{}
	for _, m := range family.GetMetric() {
		found := false
		for _, l := range m.GetLabel() {
			if l.GetName() == "error_code" {
				require.NotEmpty(t, l.GetValue(), "empty error_code on %v", m.GetLabel())
				codes[l.GetValue()] = true
				found = true
			}
		}
		require.True(t, found, "missing error_code label on %v", m.GetLabel())
	}

	require.True(t, codes["unavailable"], "expected unavailable in %v", codes)
	require.True(t, codes["unknown"], "expected unknown in %v", codes)
}
