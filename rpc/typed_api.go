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

package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils/must"
	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/middleware"
)

type PSRPCConfig struct {
	MaxAttempts int           `yaml:"max_attempts,omitempty"`
	Timeout     time.Duration `yaml:"timeout,omitempty"`
	Backoff     time.Duration `yaml:"backoff,omitempty"`
	BufferSize  int           `yaml:"buffer_size,omitempty"`
}

var DefaultPSRPCConfig = PSRPCConfig{
	MaxAttempts: 3,
	Timeout:     3 * time.Second,
	Backoff:     2 * time.Second,
	BufferSize:  1000,
}

type ClientParams struct {
	PSRPCConfig
	Bus      psrpc.MessageBus
	Logger   logger.Logger
	Observer middleware.MetricsObserver
}

func NewClientParams(
	config PSRPCConfig,
	bus psrpc.MessageBus,
	logger logger.Logger,
	observer middleware.MetricsObserver,
) ClientParams {
	return ClientParams{
		PSRPCConfig: config,
		Bus:         bus,
		Logger:      logger,
		Observer:    observer,
	}
}

func (p *ClientParams) Options() []psrpc.ClientOption {
	opts := make([]psrpc.ClientOption, 0, 4)
	if p.BufferSize != 0 {
		opts = append(opts, psrpc.WithClientChannelSize(p.BufferSize))
	}
	if p.Observer != nil {
		opts = append(opts, middleware.WithClientMetrics(p.Observer))
	}
	if p.Logger != nil {
		opts = append(opts, WithClientLogger(p.Logger))
	}
	if p.MaxAttempts != 0 || p.Timeout != 0 || p.Backoff != 0 {
		opts = append(opts, middleware.WithRPCRetries(middleware.RetryOptions{
			MaxAttempts: p.MaxAttempts,
			Timeout:     p.Timeout,
			Backoff:     p.Backoff,
		}))
	}
	return opts
}

func (p *ClientParams) Args() (psrpc.MessageBus, psrpc.ClientOption) {
	return p.Bus, psrpc.WithClientOptions(p.Options()...)
}

func WithServerObservability(logger logger.Logger) psrpc.ServerOption {
	return psrpc.WithServerOptions(
		middleware.WithServerMetrics(PSRPCMetricsObserver{}),
		WithServerLogger(logger),
	)
}

func WithDefaultServerOptions(psrpcConfig PSRPCConfig, logger logger.Logger) psrpc.ServerOption {
	return psrpc.WithServerOptions(
		psrpc.WithServerChannelSize(psrpcConfig.BufferSize),
		WithServerObservability(logger),
	)
}

func WithClientObservability(logger logger.Logger) psrpc.ClientOption {
	return psrpc.WithClientOptions(
		middleware.WithClientMetrics(PSRPCMetricsObserver{}),
		WithClientLogger(logger),
	)
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type TypedSignalClient = SignalClient[livekit.NodeID]
type TypedSignalServer = SignalServer[livekit.NodeID]

func NewTypedSignalClient(nodeID livekit.NodeID, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedSignalClient, error) {
	return NewSignalClient[livekit.NodeID](bus, psrpc.WithClientOptions(opts...), psrpc.WithClientID(string(nodeID)))
}

func NewTypedSignalServer(nodeID livekit.NodeID, svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedSignalServer, error) {
	return NewSignalServer[livekit.NodeID](svc, bus, psrpc.WithServerOptions(opts...), psrpc.WithServerID(string(nodeID)))
}

type TypedSignalv2Client = Signalv2Client[livekit.NodeID]
type TypedSignalv2Server = Signalv2Server[livekit.NodeID]

func NewTypedSignalv2Client(nodeID livekit.NodeID, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedSignalv2Client, error) {
	return NewSignalv2Client[livekit.NodeID](bus, psrpc.WithClientOptions(opts...), psrpc.WithClientID(string(nodeID)))
}

func NewTypedSignalv2Server(nodeID livekit.NodeID, svc Signalv2ServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedSignalv2Server, error) {
	return NewSignalv2Server[livekit.NodeID](svc, bus, psrpc.WithServerOptions(opts...), psrpc.WithServerID(string(nodeID)))
}

type TypedRoomManagerClient = RoomManagerClient[livekit.NodeID]
type TypedRoomManagerServer = RoomManagerServer[livekit.NodeID]

func NewTypedRoomManagerClient(bus psrpc.MessageBus, opts ...psrpc.ClientOption) (TypedRoomManagerClient, error) {
	return NewRoomManagerClient[livekit.NodeID](bus, opts...)
}

func NewTypedRoomManagerServer(svc RoomManagerServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedRoomManagerServer, error) {
	return NewRoomManagerServer[livekit.NodeID](svc, bus, opts...)
}

type ParticipantTopic string
type RoomTopic string

func FormatParticipantTopic(roomName livekit.RoomName, identity livekit.ParticipantIdentity) ParticipantTopic {
	return ParticipantTopic(fmt.Sprintf("%s_%s", roomName, identity))
}

func FormatRoomTopic(roomName livekit.RoomName) RoomTopic {
	return RoomTopic(roomName)
}

type topicFormatter struct{}

func NewTopicFormatter() TopicFormatter {
	return topicFormatter{}
}

func (f topicFormatter) ParticipantTopic(ctx context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity) ParticipantTopic {
	return FormatParticipantTopic(roomName, identity)
}

func (f topicFormatter) RoomTopic(ctx context.Context, roomName livekit.RoomName) RoomTopic {
	return FormatRoomTopic(roomName)
}

type TopicFormatter interface {
	ParticipantTopic(ctx context.Context, roomName livekit.RoomName, identity livekit.ParticipantIdentity) ParticipantTopic
	RoomTopic(ctx context.Context, roomName livekit.RoomName) RoomTopic
}

//counterfeiter:generate . TypedRoomClient
type TypedRoomClient = RoomClient[RoomTopic]
type TypedRoomServer = RoomServer[RoomTopic]

func NewTypedRoomClient(params ClientParams) (TypedRoomClient, error) {
	return NewRoomClient[RoomTopic](params.Args())
}

func NewTypedRoomServer(svc RoomServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedRoomServer, error) {
	return NewRoomServer[RoomTopic](svc, bus, opts...)
}

//counterfeiter:generate . TypedParticipantClient
type TypedParticipantClient = ParticipantClient[ParticipantTopic]
type TypedParticipantServer = ParticipantServer[ParticipantTopic]

func NewTypedParticipantClient(params ClientParams) (TypedParticipantClient, error) {
	return NewParticipantClient[ParticipantTopic](params.Args())
}

func NewTypedParticipantServer(svc ParticipantServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedParticipantServer, error) {
	return NewParticipantServer[ParticipantTopic](svc, bus, opts...)
}

//counterfeiter:generate . TypedSignalv2ParticipantClient
type TypedSignalv2ParticipantClient = Signalv2ParticipantClient[ParticipantTopic]
type TypedSignalv2ParticipantServer = Signalv2ParticipantServer[ParticipantTopic]

func NewTypedSignalv2ParticipantClient(params ClientParams) (TypedSignalv2ParticipantClient, error) {
	return NewSignalv2ParticipantClient[ParticipantTopic](params.Args())
}

func NewTypedSignalv2ParticipantServer(svc Signalv2ParticipantServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedSignalv2ParticipantServer, error) {
	return NewSignalv2ParticipantServer[ParticipantTopic](svc, bus, opts...)
}

//counterfeiter:generate . TypedWHIPParticipantClient
type TypedWHIPParticipantClient = WHIPParticipantClient[ParticipantTopic]
type TypedWHIPParticipantServer = WHIPParticipantServer[ParticipantTopic]

func NewTypedWHIPParticipantClient(params ClientParams) (TypedWHIPParticipantClient, error) {
	return NewWHIPParticipantClient[ParticipantTopic](params.Args())
}

func NewTypedWHIPParticipantServer(svc WHIPParticipantServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedWHIPParticipantServer, error) {
	return NewWHIPParticipantServer[ParticipantTopic](svc, bus, opts...)
}

//counterfeiter:generate . TypedAgentDispatchInternalClient
type TypedAgentDispatchInternalClient = AgentDispatchInternalClient[RoomTopic]
type TypedAgentDispatchInternalServer = AgentDispatchInternalServer[RoomTopic]

func NewTypedAgentDispatchInternalClient(params ClientParams) (TypedAgentDispatchInternalClient, error) {
	return NewAgentDispatchInternalClient[RoomTopic](params.Args())
}

func NewTypedAgentDispatchInternalServer(svc AgentDispatchInternalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (TypedAgentDispatchInternalServer, error) {
	return NewAgentDispatchInternalServer[RoomTopic](svc, bus, opts...)
}

//counterfeiter:generate . KeepalivePubSub
type KeepalivePubSub interface {
	KeepaliveClient[livekit.NodeID]
	KeepaliveServer[livekit.NodeID]
}

type keepalivePubSub struct {
	KeepaliveClient[livekit.NodeID]
	KeepaliveServer[livekit.NodeID]
}

func NewKeepalivePubSub(params ClientParams) (KeepalivePubSub, error) {
	client, err := NewKeepaliveClient[livekit.NodeID](params.Args())
	if err != nil {
		return nil, err
	}
	server := must.Get(NewKeepaliveServer[livekit.NodeID](nil, params.Bus))
	return &keepalivePubSub{client, server}, nil
}
