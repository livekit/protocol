// Copyright 2026 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package livekit

import (
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type agentDBErrorKind struct {
	applied AgentDB_Applied
	retry   AgentDB_Retry
}

// agentDBErrors is each code by its wire name, and its options, read from the
// descriptor so that the proto stays their only declaration. Read on first
// use: the descriptor is built by the package's init.
var agentDBErrors = sync.OnceValues(func() (map[string]AgentDB_ErrorCode, map[AgentDB_ErrorCode]agentDBErrorKind) {
	names := map[string]AgentDB_ErrorCode{}
	kinds := map[AgentDB_ErrorCode]agentDBErrorKind{}
	values := AgentDB_UNKNOWN.Descriptor().Values()
	for i := range values.Len() {
		v := values.Get(i)
		code := AgentDB_ErrorCode(v.Number())
		names[strings.ToLower(string(v.Name()))] = code
		kinds[code] = agentDBErrorKind{
			applied: proto.GetExtension(v.Options(), E_AgentDB_Applied).(AgentDB_Applied),
			retry:   proto.GetExtension(v.Options(), E_AgentDB_Retry).(AgentDB_Retry),
		}
	}
	return names, kinds
})

// ParseAgentDBErrorCode reads a code as it travels, its name in lowercase.
// Anything else is UNKNOWN.
func ParseAgentDBErrorCode(name string) AgentDB_ErrorCode {
	names, _ := agentDBErrors()
	return names[name]
}

// AgentDBErrorCodeOf is the code a management service's gRPC status carries.
// OK and a status code past the known ones are UNKNOWN.
func AgentDBErrorCodeOf(c codes.Code) AgentDB_ErrorCode {
	if c == codes.OK || c == codes.Unknown || c > codes.Unauthenticated {
		return AgentDB_UNKNOWN
	}
	return AgentDB_ErrorCode(c)
}

// Wire is the code as it travels: its name in lowercase.
func (c AgentDB_ErrorCode) Wire() string {
	if c.Descriptor().Values().ByNumber(protoreflect.EnumNumber(c)) == nil {
		return "unknown"
	}
	return strings.ToLower(c.String())
}

// Applied is whether a request answered with c may have taken effect.
func (c AgentDB_ErrorCode) Applied() AgentDB_Applied {
	_, kinds := agentDBErrors()
	return kinds[c].applied
}

// Retry is whether the same request may succeed if sent again, and where.
func (c AgentDB_ErrorCode) Retry() AgentDB_Retry {
	_, kinds := agentDBErrors()
	return kinds[c].retry
}
