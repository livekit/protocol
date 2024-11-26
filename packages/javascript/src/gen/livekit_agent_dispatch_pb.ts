// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file livekit_agent_dispatch.proto (package livekit, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Job } from "./livekit_agent_pb.js";

/**
 * @generated from message livekit.CreateAgentDispatchRequest
 */
export class CreateAgentDispatchRequest extends Message<CreateAgentDispatchRequest> {
  /**
   * @generated from field: string agent_name = 1;
   */
  agentName = "";

  /**
   * @generated from field: string room = 2;
   */
  room = "";

  /**
   * @generated from field: string metadata = 3;
   */
  metadata = "";

  constructor(data?: PartialMessage<CreateAgentDispatchRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.CreateAgentDispatchRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "agent_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "metadata", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAgentDispatchRequest {
    return new CreateAgentDispatchRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAgentDispatchRequest {
    return new CreateAgentDispatchRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAgentDispatchRequest {
    return new CreateAgentDispatchRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateAgentDispatchRequest | PlainMessage<CreateAgentDispatchRequest> | undefined, b: CreateAgentDispatchRequest | PlainMessage<CreateAgentDispatchRequest> | undefined): boolean {
    return proto3.util.equals(CreateAgentDispatchRequest, a, b);
  }
}

/**
 * @generated from message livekit.RoomAgentDispatch
 */
export class RoomAgentDispatch extends Message<RoomAgentDispatch> {
  /**
   * @generated from field: string agent_name = 1;
   */
  agentName = "";

  /**
   * @generated from field: string metadata = 2;
   */
  metadata = "";

  constructor(data?: PartialMessage<RoomAgentDispatch>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.RoomAgentDispatch";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "agent_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "metadata", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomAgentDispatch {
    return new RoomAgentDispatch().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomAgentDispatch {
    return new RoomAgentDispatch().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomAgentDispatch {
    return new RoomAgentDispatch().fromJsonString(jsonString, options);
  }

  static equals(a: RoomAgentDispatch | PlainMessage<RoomAgentDispatch> | undefined, b: RoomAgentDispatch | PlainMessage<RoomAgentDispatch> | undefined): boolean {
    return proto3.util.equals(RoomAgentDispatch, a, b);
  }
}

/**
 * @generated from message livekit.DeleteAgentDispatchRequest
 */
export class DeleteAgentDispatchRequest extends Message<DeleteAgentDispatchRequest> {
  /**
   * @generated from field: string dispatch_id = 1;
   */
  dispatchId = "";

  /**
   * @generated from field: string room = 2;
   */
  room = "";

  constructor(data?: PartialMessage<DeleteAgentDispatchRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.DeleteAgentDispatchRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dispatch_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteAgentDispatchRequest {
    return new DeleteAgentDispatchRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteAgentDispatchRequest {
    return new DeleteAgentDispatchRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteAgentDispatchRequest {
    return new DeleteAgentDispatchRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteAgentDispatchRequest | PlainMessage<DeleteAgentDispatchRequest> | undefined, b: DeleteAgentDispatchRequest | PlainMessage<DeleteAgentDispatchRequest> | undefined): boolean {
    return proto3.util.equals(DeleteAgentDispatchRequest, a, b);
  }
}

/**
 * @generated from message livekit.ListAgentDispatchRequest
 */
export class ListAgentDispatchRequest extends Message<ListAgentDispatchRequest> {
  /**
   * if set, only the dispatch whose id is given will be returned
   *
   * @generated from field: string dispatch_id = 1;
   */
  dispatchId = "";

  /**
   * name of the room to list agents for. Must be set.
   *
   * @generated from field: string room = 2;
   */
  room = "";

  constructor(data?: PartialMessage<ListAgentDispatchRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.ListAgentDispatchRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dispatch_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "room", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAgentDispatchRequest {
    return new ListAgentDispatchRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAgentDispatchRequest {
    return new ListAgentDispatchRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAgentDispatchRequest {
    return new ListAgentDispatchRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAgentDispatchRequest | PlainMessage<ListAgentDispatchRequest> | undefined, b: ListAgentDispatchRequest | PlainMessage<ListAgentDispatchRequest> | undefined): boolean {
    return proto3.util.equals(ListAgentDispatchRequest, a, b);
  }
}

/**
 * @generated from message livekit.ListAgentDispatchResponse
 */
export class ListAgentDispatchResponse extends Message<ListAgentDispatchResponse> {
  /**
   * @generated from field: repeated livekit.AgentDispatch agent_dispatches = 1;
   */
  agentDispatches: AgentDispatch[] = [];

  constructor(data?: PartialMessage<ListAgentDispatchResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.ListAgentDispatchResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "agent_dispatches", kind: "message", T: AgentDispatch, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAgentDispatchResponse {
    return new ListAgentDispatchResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAgentDispatchResponse {
    return new ListAgentDispatchResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAgentDispatchResponse {
    return new ListAgentDispatchResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAgentDispatchResponse | PlainMessage<ListAgentDispatchResponse> | undefined, b: ListAgentDispatchResponse | PlainMessage<ListAgentDispatchResponse> | undefined): boolean {
    return proto3.util.equals(ListAgentDispatchResponse, a, b);
  }
}

/**
 * @generated from message livekit.AgentDispatch
 */
export class AgentDispatch extends Message<AgentDispatch> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string agent_name = 2;
   */
  agentName = "";

  /**
   * @generated from field: string room = 3;
   */
  room = "";

  /**
   * @generated from field: string metadata = 4;
   */
  metadata = "";

  /**
   * @generated from field: livekit.AgentDispatchState state = 5;
   */
  state?: AgentDispatchState;

  constructor(data?: PartialMessage<AgentDispatch>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.AgentDispatch";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "agent_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "room", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "metadata", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "state", kind: "message", T: AgentDispatchState },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AgentDispatch {
    return new AgentDispatch().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AgentDispatch {
    return new AgentDispatch().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AgentDispatch {
    return new AgentDispatch().fromJsonString(jsonString, options);
  }

  static equals(a: AgentDispatch | PlainMessage<AgentDispatch> | undefined, b: AgentDispatch | PlainMessage<AgentDispatch> | undefined): boolean {
    return proto3.util.equals(AgentDispatch, a, b);
  }
}

/**
 * @generated from message livekit.AgentDispatchState
 */
export class AgentDispatchState extends Message<AgentDispatchState> {
  /**
   * For dispatches of tyoe JT_ROOM, there will be at most 1 job. 
   * For dispatches of type JT_PUBLISHER, there will be 1 per publisher.
   *
   * @generated from field: repeated livekit.Job jobs = 1;
   */
  jobs: Job[] = [];

  /**
   * @generated from field: int64 created_at = 2;
   */
  createdAt = protoInt64.zero;

  /**
   * @generated from field: int64 deleted_at = 3;
   */
  deletedAt = protoInt64.zero;

  constructor(data?: PartialMessage<AgentDispatchState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "livekit.AgentDispatchState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "jobs", kind: "message", T: Job, repeated: true },
    { no: 2, name: "created_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "deleted_at", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AgentDispatchState {
    return new AgentDispatchState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AgentDispatchState {
    return new AgentDispatchState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AgentDispatchState {
    return new AgentDispatchState().fromJsonString(jsonString, options);
  }

  static equals(a: AgentDispatchState | PlainMessage<AgentDispatchState> | undefined, b: AgentDispatchState | PlainMessage<AgentDispatchState> | undefined): boolean {
    return proto3.util.equals(AgentDispatchState, a, b);
  }
}

