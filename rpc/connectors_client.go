// Copyright 2025 LiveKit, Inc.
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

type ConnectorsClient interface {
	ConnectorsInternalClient
	ConnectorsHandlerClient
}

type connectorsClient struct {
	ConnectorsInternalClient
	ConnectorsHandlerClient
}

func NewConnectorsClient(params ClientParams) (ConnectorsClient, error) {
	if params.Bus == nil {
		return nil, nil
	}
	opts := params.Options()

	internalClient, err := NewConnectorsInternalClient(params.Bus, opts...)
	if err != nil {
		return nil, err
	}

	handlerClient, err := NewConnectorsHandlerClient(params.Bus, opts...)
	if err != nil {
		return nil, err
	}

	return &connectorsClient{
		internalClient,
		handlerClient,
	}, nil
}

func (c *connectorsClient) Close() {
	c.ConnectorsInternalClient.Close()
	c.ConnectorsHandlerClient.Close()
}
