import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Job } from "./livekit_agent_pb.js";
const _CreateAgentDispatchRequest = class _CreateAgentDispatchRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string agent_name = 1;
     */
    this.agentName = "";
    /**
     * @generated from field: string room = 2;
     */
    this.room = "";
    /**
     * @generated from field: string metadata = 3;
     */
    this.metadata = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateAgentDispatchRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateAgentDispatchRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateAgentDispatchRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateAgentDispatchRequest, a, b);
  }
};
_CreateAgentDispatchRequest.runtime = proto3;
_CreateAgentDispatchRequest.typeName = "livekit.CreateAgentDispatchRequest";
_CreateAgentDispatchRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "agent_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let CreateAgentDispatchRequest = _CreateAgentDispatchRequest;
const _RoomAgentDispatch = class _RoomAgentDispatch extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string agent_name = 1;
     */
    this.agentName = "";
    /**
     * @generated from field: string metadata = 2;
     */
    this.metadata = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomAgentDispatch().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomAgentDispatch().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomAgentDispatch().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RoomAgentDispatch, a, b);
  }
};
_RoomAgentDispatch.runtime = proto3;
_RoomAgentDispatch.typeName = "livekit.RoomAgentDispatch";
_RoomAgentDispatch.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "agent_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let RoomAgentDispatch = _RoomAgentDispatch;
const _DeleteAgentDispatchRequest = class _DeleteAgentDispatchRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string dispatch_id = 1;
     */
    this.dispatchId = "";
    /**
     * @generated from field: string room = 2;
     */
    this.room = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteAgentDispatchRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteAgentDispatchRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteAgentDispatchRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_DeleteAgentDispatchRequest, a, b);
  }
};
_DeleteAgentDispatchRequest.runtime = proto3;
_DeleteAgentDispatchRequest.typeName = "livekit.DeleteAgentDispatchRequest";
_DeleteAgentDispatchRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "dispatch_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let DeleteAgentDispatchRequest = _DeleteAgentDispatchRequest;
const _ListAgentDispatchRequest = class _ListAgentDispatchRequest extends Message {
  constructor(data) {
    super();
    /**
     * if set, only the dispatch whose id is given will be returned
     *
     * @generated from field: string dispatch_id = 1;
     */
    this.dispatchId = "";
    /**
     * name of the room to list agents for. Must be set.
     *
     * @generated from field: string room = 2;
     */
    this.room = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListAgentDispatchRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListAgentDispatchRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListAgentDispatchRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListAgentDispatchRequest, a, b);
  }
};
_ListAgentDispatchRequest.runtime = proto3;
_ListAgentDispatchRequest.typeName = "livekit.ListAgentDispatchRequest";
_ListAgentDispatchRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "dispatch_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let ListAgentDispatchRequest = _ListAgentDispatchRequest;
const _ListAgentDispatchResponse = class _ListAgentDispatchResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.AgentDispatch agent_dispatches = 1;
     */
    this.agentDispatches = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListAgentDispatchResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListAgentDispatchResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListAgentDispatchResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListAgentDispatchResponse, a, b);
  }
};
_ListAgentDispatchResponse.runtime = proto3;
_ListAgentDispatchResponse.typeName = "livekit.ListAgentDispatchResponse";
_ListAgentDispatchResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "agent_dispatches", kind: "message", T: AgentDispatch, repeated: true }
]);
let ListAgentDispatchResponse = _ListAgentDispatchResponse;
const _AgentDispatch = class _AgentDispatch extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string agent_name = 2;
     */
    this.agentName = "";
    /**
     * @generated from field: string room = 3;
     */
    this.room = "";
    /**
     * @generated from field: string metadata = 4;
     */
    this.metadata = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AgentDispatch().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AgentDispatch().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AgentDispatch().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_AgentDispatch, a, b);
  }
};
_AgentDispatch.runtime = proto3;
_AgentDispatch.typeName = "livekit.AgentDispatch";
_AgentDispatch.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "agent_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "state", kind: "message", T: AgentDispatchState }
]);
let AgentDispatch = _AgentDispatch;
const _AgentDispatchState = class _AgentDispatchState extends Message {
  constructor(data) {
    super();
    /**
     * For dispatches of tyoe JT_ROOM, there will be at most 1 job. 
     * For dispatches of type JT_PUBLISHER, there will be 1 per publisher.
     *
     * @generated from field: repeated livekit.Job jobs = 1;
     */
    this.jobs = [];
    /**
     * @generated from field: int64 created_at = 2;
     */
    this.createdAt = protoInt64.zero;
    /**
     * @generated from field: int64 deleted_at = 3;
     */
    this.deletedAt = protoInt64.zero;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AgentDispatchState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AgentDispatchState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AgentDispatchState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_AgentDispatchState, a, b);
  }
};
_AgentDispatchState.runtime = proto3;
_AgentDispatchState.typeName = "livekit.AgentDispatchState";
_AgentDispatchState.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "jobs", kind: "message", T: Job, repeated: true },
  {
    no: 2,
    name: "created_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "deleted_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let AgentDispatchState = _AgentDispatchState;
export {
  AgentDispatch,
  AgentDispatchState,
  CreateAgentDispatchRequest,
  DeleteAgentDispatchRequest,
  ListAgentDispatchRequest,
  ListAgentDispatchResponse,
  RoomAgentDispatch
};
//# sourceMappingURL=livekit_agent_dispatch_pb.js.map