package rpc

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/livekit/psrpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

// lateClaimBus holds back the claims a SIP instance publishes, the way a pod's first publish after
// an idle stretch waits on a fresh Redis connection.
type lateClaimBus struct {
	psrpc.MessageBus
	delay time.Duration
}

func (b lateClaimBus) Publish(ctx context.Context, channel psrpc.Channel, msg proto.Message) error {
	if strings.HasSuffix(channel.Legacy, "|CLAIM") {
		time.Sleep(b.delay)
	}
	return b.MessageBus.Publish(ctx, channel, msg)
}

type answeringSIPServer struct{}

func (answeringSIPServer) CreateSIPParticipant(context.Context, *InternalCreateSIPParticipantRequest) (*InternalCreateSIPParticipantResponse, error) {
	return &InternalCreateSIPParticipantResponse{ParticipantId: "PA_late"}, nil
}

func (answeringSIPServer) TransferSIPParticipant(context.Context, *InternalTransferSIPParticipantRequest) (*InternalTransferSIPParticipantResponse, error) {
	return &InternalTransferSIPParticipantResponse{}, nil
}

func TestSIPClientWaitsForALateClaim(t *testing.T) {
	const claimDelay = 1500 * time.Millisecond
	require.Greater(t, claimDelay, psrpc.DefaultAffinityTimeout)
	require.Less(t, claimDelay, sipClaimTimeout)

	bus := psrpc.NewLocalMessageBus()
	srv, err := NewSIPInternalServer(answeringSIPServer{}, lateClaimBus{bus, claimDelay})
	require.NoError(t, err)
	t.Cleanup(srv.Kill)
	require.NoError(t, srv.RegisterCreateSIPParticipantTopic(""))

	t.Run("psrpc's default window gives up first", func(t *testing.T) {
		client, err := NewSIPInternalClient(bus)
		require.NoError(t, err)
		_, err = client.CreateSIPParticipant(context.Background(), "", &InternalCreateSIPParticipantRequest{})
		require.ErrorIs(t, err, psrpc.ErrNoResponse)
	})

	t.Run("the SIP client waits for the claim", func(t *testing.T) {
		client, err := NewSIPClient(bus)
		require.NoError(t, err)
		// livekit-server passes the dial's own timeout, so only the claim window limits the call.
		// Under psrpc's 3s request default the claim window and the whole request end together.
		res, err := client.CreateSIPParticipant(context.Background(), "", &InternalCreateSIPParticipantRequest{},
			psrpc.WithRequestTimeout(time.Minute))
		require.NoError(t, err)
		require.Equal(t, "PA_late", res.ParticipantId)
	})

	t.Run("a caller's own window still wins", func(t *testing.T) {
		client, err := NewSIPClientWithParams(ClientParams{
			Bus:           bus,
			ClientOptions: []psrpc.ClientOption{psrpc.WithClientSelectTimeout(psrpc.DefaultAffinityTimeout)},
		})
		require.NoError(t, err)
		_, err = client.CreateSIPParticipant(context.Background(), "", &InternalCreateSIPParticipantRequest{})
		require.ErrorIs(t, err, psrpc.ErrNoResponse)
	})
}
