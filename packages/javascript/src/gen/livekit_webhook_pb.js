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

// @generated by protoc-gen-es v1.10.0 with parameter "target=dts+js"
// @generated from file livekit_webhook.proto (package livekit, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ParticipantInfo, Room, TrackInfo } from "./livekit_models_pb.js";
import { EgressInfo } from "./livekit_egress_pb.js";
import { IngressInfo } from "./livekit_ingress_pb.js";

/**
 * @generated from message livekit.WebhookEvent
 */
export const WebhookEvent = /*@__PURE__*/ proto3.makeMessageType(
  "livekit.WebhookEvent",
  () => [
    { no: 1, name: "event", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room", kind: "message", T: Room },
    { no: 3, name: "participant", kind: "message", T: ParticipantInfo },
    { no: 9, name: "egress_info", kind: "message", T: EgressInfo },
    { no: 10, name: "ingress_info", kind: "message", T: IngressInfo },
    { no: 8, name: "track", kind: "message", T: TrackInfo },
    { no: 6, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "created_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 11, name: "num_dropped", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

