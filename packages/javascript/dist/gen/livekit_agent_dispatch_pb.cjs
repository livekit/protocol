var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);
var livekit_agent_dispatch_pb_exports = {};
__export(livekit_agent_dispatch_pb_exports, {
  AgentDispatch: () => AgentDispatch,
  AgentDispatchState: () => AgentDispatchState,
  CreateAgentDispatchRequest: () => CreateAgentDispatchRequest,
  DeleteAgentDispatchRequest: () => DeleteAgentDispatchRequest,
  ListAgentDispatchRequest: () => ListAgentDispatchRequest,
  ListAgentDispatchResponse: () => ListAgentDispatchResponse,
  RoomAgentDispatch: () => RoomAgentDispatch
});
module.exports = __toCommonJS(livekit_agent_dispatch_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var import_livekit_agent_pb = require("./livekit_agent_pb.cjs");
const _CreateAgentDispatchRequest = class _CreateAgentDispatchRequest extends import_protobuf.Message {
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
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_CreateAgentDispatchRequest, a, b);
  }
};
_CreateAgentDispatchRequest.runtime = import_protobuf.proto3;
_CreateAgentDispatchRequest.typeName = "livekit.CreateAgentDispatchRequest";
_CreateAgentDispatchRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
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
const _RoomAgentDispatch = class _RoomAgentDispatch extends import_protobuf.Message {
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
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_RoomAgentDispatch, a, b);
  }
};
_RoomAgentDispatch.runtime = import_protobuf.proto3;
_RoomAgentDispatch.typeName = "livekit.RoomAgentDispatch";
_RoomAgentDispatch.fields = import_protobuf.proto3.util.newFieldList(() => [
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
const _DeleteAgentDispatchRequest = class _DeleteAgentDispatchRequest extends import_protobuf.Message {
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
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_DeleteAgentDispatchRequest, a, b);
  }
};
_DeleteAgentDispatchRequest.runtime = import_protobuf.proto3;
_DeleteAgentDispatchRequest.typeName = "livekit.DeleteAgentDispatchRequest";
_DeleteAgentDispatchRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
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
const _ListAgentDispatchRequest = class _ListAgentDispatchRequest extends import_protobuf.Message {
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
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_ListAgentDispatchRequest, a, b);
  }
};
_ListAgentDispatchRequest.runtime = import_protobuf.proto3;
_ListAgentDispatchRequest.typeName = "livekit.ListAgentDispatchRequest";
_ListAgentDispatchRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
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
const _ListAgentDispatchResponse = class _ListAgentDispatchResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.AgentDispatch agent_dispatches = 1;
     */
    this.agentDispatches = [];
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_ListAgentDispatchResponse, a, b);
  }
};
_ListAgentDispatchResponse.runtime = import_protobuf.proto3;
_ListAgentDispatchResponse.typeName = "livekit.ListAgentDispatchResponse";
_ListAgentDispatchResponse.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "agent_dispatches", kind: "message", T: AgentDispatch, repeated: true }
]);
let ListAgentDispatchResponse = _ListAgentDispatchResponse;
const _AgentDispatch = class _AgentDispatch extends import_protobuf.Message {
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
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_AgentDispatch, a, b);
  }
};
_AgentDispatch.runtime = import_protobuf.proto3;
_AgentDispatch.typeName = "livekit.AgentDispatch";
_AgentDispatch.fields = import_protobuf.proto3.util.newFieldList(() => [
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
const _AgentDispatchState = class _AgentDispatchState extends import_protobuf.Message {
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
    this.createdAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 deleted_at = 3;
     */
    this.deletedAt = import_protobuf.protoInt64.zero;
    import_protobuf.proto3.util.initPartial(data, this);
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
    return import_protobuf.proto3.util.equals(_AgentDispatchState, a, b);
  }
};
_AgentDispatchState.runtime = import_protobuf.proto3;
_AgentDispatchState.typeName = "livekit.AgentDispatchState";
_AgentDispatchState.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "jobs", kind: "message", T: import_livekit_agent_pb.Job, repeated: true },
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
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  AgentDispatch,
  AgentDispatchState,
  CreateAgentDispatchRequest,
  DeleteAgentDispatchRequest,
  ListAgentDispatchRequest,
  ListAgentDispatchResponse,
  RoomAgentDispatch
});
//# sourceMappingURL=livekit_agent_dispatch_pb.cjs.map