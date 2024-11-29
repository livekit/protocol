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
var livekit_room_pb_exports = {};
__export(livekit_room_pb_exports, {
  CreateRoomRequest: () => CreateRoomRequest,
  DeleteRoomRequest: () => DeleteRoomRequest,
  DeleteRoomResponse: () => DeleteRoomResponse,
  ListParticipantsRequest: () => ListParticipantsRequest,
  ListParticipantsResponse: () => ListParticipantsResponse,
  ListRoomsRequest: () => ListRoomsRequest,
  ListRoomsResponse: () => ListRoomsResponse,
  MuteRoomTrackRequest: () => MuteRoomTrackRequest,
  MuteRoomTrackResponse: () => MuteRoomTrackResponse,
  RemoveParticipantResponse: () => RemoveParticipantResponse,
  RoomAgent: () => RoomAgent,
  RoomConfiguration: () => RoomConfiguration,
  RoomEgress: () => RoomEgress,
  RoomParticipantIdentity: () => RoomParticipantIdentity,
  SendDataRequest: () => SendDataRequest,
  SendDataResponse: () => SendDataResponse,
  UpdateParticipantRequest: () => UpdateParticipantRequest,
  UpdateRoomMetadataRequest: () => UpdateRoomMetadataRequest,
  UpdateSubscriptionsRequest: () => UpdateSubscriptionsRequest,
  UpdateSubscriptionsResponse: () => UpdateSubscriptionsResponse
});
module.exports = __toCommonJS(livekit_room_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var import_livekit_agent_dispatch_pb = require("./livekit_agent_dispatch_pb.cjs");
var import_livekit_egress_pb = require("./livekit_egress_pb.cjs");
var import_livekit_models_pb = require("./livekit_models_pb.cjs");
const _CreateRoomRequest = class _CreateRoomRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * name of the room
     *
     * @generated from field: string name = 1;
     */
    this.name = "";
    /**
     * configuration to use for this room parameters. Setting parameters below override the config defaults.
     *
     * @generated from field: string room_preset = 12;
     */
    this.roomPreset = "";
    /**
     * number of seconds to keep the room open if no one joins
     *
     * @generated from field: uint32 empty_timeout = 2;
     */
    this.emptyTimeout = 0;
    /**
     * number of seconds to keep the room open after everyone leaves
     *
     * @generated from field: uint32 departure_timeout = 10;
     */
    this.departureTimeout = 0;
    /**
     * limit number of participants that can be in a room
     *
     * @generated from field: uint32 max_participants = 3;
     */
    this.maxParticipants = 0;
    /**
     * override the node room is allocated to, for debugging
     *
     * @generated from field: string node_id = 4;
     */
    this.nodeId = "";
    /**
     * metadata of room
     *
     * @generated from field: string metadata = 5;
     */
    this.metadata = "";
    /**
     * playout delay of subscriber
     *
     * @generated from field: uint32 min_playout_delay = 7;
     */
    this.minPlayoutDelay = 0;
    /**
     * @generated from field: uint32 max_playout_delay = 8;
     */
    this.maxPlayoutDelay = 0;
    /**
     * improves A/V sync when playout_delay set to a value larger than 200ms. It will disables transceiver re-use
     * so not recommended for rooms with frequent subscription changes
     *
     * @generated from field: bool sync_streams = 9;
     */
    this.syncStreams = false;
    /**
     * replay
     *
     * @generated from field: bool replay_enabled = 13;
     */
    this.replayEnabled = false;
    /**
     * Define agents that should be dispatched to this room
     *
     * NEXT-ID: 15
     *
     * @generated from field: repeated livekit.RoomAgentDispatch agents = 14;
     */
    this.agents = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateRoomRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateRoomRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateRoomRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_CreateRoomRequest, a, b);
  }
};
_CreateRoomRequest.runtime = import_protobuf.proto3;
_CreateRoomRequest.typeName = "livekit.CreateRoomRequest";
_CreateRoomRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 12,
    name: "room_preset",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "empty_timeout",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 10,
    name: "departure_timeout",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "max_participants",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "node_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "egress", kind: "message", T: RoomEgress },
  {
    no: 7,
    name: "min_playout_delay",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 8,
    name: "max_playout_delay",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 9,
    name: "sync_streams",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 13,
    name: "replay_enabled",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 14, name: "agents", kind: "message", T: import_livekit_agent_dispatch_pb.RoomAgentDispatch, repeated: true }
]);
let CreateRoomRequest = _CreateRoomRequest;
const _RoomEgress = class _RoomEgress extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomEgress().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomEgress().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomEgress().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RoomEgress, a, b);
  }
};
_RoomEgress.runtime = import_protobuf.proto3;
_RoomEgress.typeName = "livekit.RoomEgress";
_RoomEgress.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "room", kind: "message", T: import_livekit_egress_pb.RoomCompositeEgressRequest },
  { no: 3, name: "participant", kind: "message", T: import_livekit_egress_pb.AutoParticipantEgress },
  { no: 2, name: "tracks", kind: "message", T: import_livekit_egress_pb.AutoTrackEgress }
]);
let RoomEgress = _RoomEgress;
const _RoomAgent = class _RoomAgent extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.RoomAgentDispatch dispatches = 1;
     */
    this.dispatches = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomAgent().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomAgent().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomAgent().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RoomAgent, a, b);
  }
};
_RoomAgent.runtime = import_protobuf.proto3;
_RoomAgent.typeName = "livekit.RoomAgent";
_RoomAgent.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "dispatches", kind: "message", T: import_livekit_agent_dispatch_pb.RoomAgentDispatch, repeated: true }
]);
let RoomAgent = _RoomAgent;
const _ListRoomsRequest = class _ListRoomsRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * when set, will only return rooms with name match
     *
     * @generated from field: repeated string names = 1;
     */
    this.names = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListRoomsRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListRoomsRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListRoomsRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ListRoomsRequest, a, b);
  }
};
_ListRoomsRequest.runtime = import_protobuf.proto3;
_ListRoomsRequest.typeName = "livekit.ListRoomsRequest";
_ListRoomsRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "names", kind: "scalar", T: 9, repeated: true }
]);
let ListRoomsRequest = _ListRoomsRequest;
const _ListRoomsResponse = class _ListRoomsResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.Room rooms = 1;
     */
    this.rooms = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListRoomsResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListRoomsResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListRoomsResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ListRoomsResponse, a, b);
  }
};
_ListRoomsResponse.runtime = import_protobuf.proto3;
_ListRoomsResponse.typeName = "livekit.ListRoomsResponse";
_ListRoomsResponse.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "rooms", kind: "message", T: import_livekit_models_pb.Room, repeated: true }
]);
let ListRoomsResponse = _ListRoomsResponse;
const _DeleteRoomRequest = class _DeleteRoomRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    this.room = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteRoomRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteRoomRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteRoomRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DeleteRoomRequest, a, b);
  }
};
_DeleteRoomRequest.runtime = import_protobuf.proto3;
_DeleteRoomRequest.typeName = "livekit.DeleteRoomRequest";
_DeleteRoomRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let DeleteRoomRequest = _DeleteRoomRequest;
const _DeleteRoomResponse = class _DeleteRoomResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteRoomResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteRoomResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteRoomResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DeleteRoomResponse, a, b);
  }
};
_DeleteRoomResponse.runtime = import_protobuf.proto3;
_DeleteRoomResponse.typeName = "livekit.DeleteRoomResponse";
_DeleteRoomResponse.fields = import_protobuf.proto3.util.newFieldList(() => []);
let DeleteRoomResponse = _DeleteRoomResponse;
const _ListParticipantsRequest = class _ListParticipantsRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    this.room = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListParticipantsRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListParticipantsRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListParticipantsRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ListParticipantsRequest, a, b);
  }
};
_ListParticipantsRequest.runtime = import_protobuf.proto3;
_ListParticipantsRequest.typeName = "livekit.ListParticipantsRequest";
_ListParticipantsRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let ListParticipantsRequest = _ListParticipantsRequest;
const _ListParticipantsResponse = class _ListParticipantsResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.ParticipantInfo participants = 1;
     */
    this.participants = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListParticipantsResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListParticipantsResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListParticipantsResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ListParticipantsResponse, a, b);
  }
};
_ListParticipantsResponse.runtime = import_protobuf.proto3;
_ListParticipantsResponse.typeName = "livekit.ListParticipantsResponse";
_ListParticipantsResponse.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "participants", kind: "message", T: import_livekit_models_pb.ParticipantInfo, repeated: true }
]);
let ListParticipantsResponse = _ListParticipantsResponse;
const _RoomParticipantIdentity = class _RoomParticipantIdentity extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    this.room = "";
    /**
     * identity of the participant
     *
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomParticipantIdentity().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomParticipantIdentity().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomParticipantIdentity().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RoomParticipantIdentity, a, b);
  }
};
_RoomParticipantIdentity.runtime = import_protobuf.proto3;
_RoomParticipantIdentity.typeName = "livekit.RoomParticipantIdentity";
_RoomParticipantIdentity.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let RoomParticipantIdentity = _RoomParticipantIdentity;
const _RemoveParticipantResponse = class _RemoveParticipantResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RemoveParticipantResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RemoveParticipantResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RemoveParticipantResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RemoveParticipantResponse, a, b);
  }
};
_RemoveParticipantResponse.runtime = import_protobuf.proto3;
_RemoveParticipantResponse.typeName = "livekit.RemoveParticipantResponse";
_RemoveParticipantResponse.fields = import_protobuf.proto3.util.newFieldList(() => []);
let RemoveParticipantResponse = _RemoveParticipantResponse;
const _MuteRoomTrackRequest = class _MuteRoomTrackRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    this.room = "";
    /**
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * sid of the track to mute
     *
     * @generated from field: string track_sid = 3;
     */
    this.trackSid = "";
    /**
     * set to true to mute, false to unmute
     *
     * @generated from field: bool muted = 4;
     */
    this.muted = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _MuteRoomTrackRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _MuteRoomTrackRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _MuteRoomTrackRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_MuteRoomTrackRequest, a, b);
  }
};
_MuteRoomTrackRequest.runtime = import_protobuf.proto3;
_MuteRoomTrackRequest.typeName = "livekit.MuteRoomTrackRequest";
_MuteRoomTrackRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "track_sid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "muted",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let MuteRoomTrackRequest = _MuteRoomTrackRequest;
const _MuteRoomTrackResponse = class _MuteRoomTrackResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _MuteRoomTrackResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _MuteRoomTrackResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _MuteRoomTrackResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_MuteRoomTrackResponse, a, b);
  }
};
_MuteRoomTrackResponse.runtime = import_protobuf.proto3;
_MuteRoomTrackResponse.typeName = "livekit.MuteRoomTrackResponse";
_MuteRoomTrackResponse.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "track", kind: "message", T: import_livekit_models_pb.TrackInfo }
]);
let MuteRoomTrackResponse = _MuteRoomTrackResponse;
const _UpdateParticipantRequest = class _UpdateParticipantRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string room = 1;
     */
    this.room = "";
    /**
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * metadata to update. skipping updates if left empty
     *
     * @generated from field: string metadata = 3;
     */
    this.metadata = "";
    /**
     * display name to update
     *
     * @generated from field: string name = 5;
     */
    this.name = "";
    /**
     * attributes to update. it only updates attributes that have been set
     * to delete attributes, set the value to an empty string
     *
     * @generated from field: map<string, string> attributes = 6;
     */
    this.attributes = {};
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateParticipantRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateParticipantRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateParticipantRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UpdateParticipantRequest, a, b);
  }
};
_UpdateParticipantRequest.runtime = import_protobuf.proto3;
_UpdateParticipantRequest.typeName = "livekit.UpdateParticipantRequest";
_UpdateParticipantRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "identity",
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
  },
  { no: 4, name: "permission", kind: "message", T: import_livekit_models_pb.ParticipantPermission },
  {
    no: 5,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } }
]);
let UpdateParticipantRequest = _UpdateParticipantRequest;
const _UpdateSubscriptionsRequest = class _UpdateSubscriptionsRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string room = 1;
     */
    this.room = "";
    /**
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * list of sids of tracks
     *
     * @generated from field: repeated string track_sids = 3;
     */
    this.trackSids = [];
    /**
     * set to true to subscribe, false to unsubscribe from tracks
     *
     * @generated from field: bool subscribe = 4;
     */
    this.subscribe = false;
    /**
     * list of participants and their tracks
     *
     * @generated from field: repeated livekit.ParticipantTracks participant_tracks = 5;
     */
    this.participantTracks = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateSubscriptionsRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateSubscriptionsRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateSubscriptionsRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UpdateSubscriptionsRequest, a, b);
  }
};
_UpdateSubscriptionsRequest.runtime = import_protobuf.proto3;
_UpdateSubscriptionsRequest.typeName = "livekit.UpdateSubscriptionsRequest";
_UpdateSubscriptionsRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "track_sids", kind: "scalar", T: 9, repeated: true },
  {
    no: 4,
    name: "subscribe",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 5, name: "participant_tracks", kind: "message", T: import_livekit_models_pb.ParticipantTracks, repeated: true }
]);
let UpdateSubscriptionsRequest = _UpdateSubscriptionsRequest;
const _UpdateSubscriptionsResponse = class _UpdateSubscriptionsResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateSubscriptionsResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateSubscriptionsResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateSubscriptionsResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UpdateSubscriptionsResponse, a, b);
  }
};
_UpdateSubscriptionsResponse.runtime = import_protobuf.proto3;
_UpdateSubscriptionsResponse.typeName = "livekit.UpdateSubscriptionsResponse";
_UpdateSubscriptionsResponse.fields = import_protobuf.proto3.util.newFieldList(() => []);
let UpdateSubscriptionsResponse = _UpdateSubscriptionsResponse;
const _SendDataRequest = class _SendDataRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string room = 1;
     */
    this.room = "";
    /**
     * @generated from field: bytes data = 2;
     */
    this.data = new Uint8Array(0);
    /**
     * @generated from field: livekit.DataPacket.Kind kind = 3;
     */
    this.kind = import_livekit_models_pb.DataPacket_Kind.RELIABLE;
    /**
     * mark deprecated
     *
     * @generated from field: repeated string destination_sids = 4 [deprecated = true];
     * @deprecated
     */
    this.destinationSids = [];
    /**
     * when set, only forward to these identities
     *
     * @generated from field: repeated string destination_identities = 6;
     */
    this.destinationIdentities = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SendDataRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SendDataRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SendDataRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SendDataRequest, a, b);
  }
};
_SendDataRequest.runtime = import_protobuf.proto3;
_SendDataRequest.typeName = "livekit.SendDataRequest";
_SendDataRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "data",
    kind: "scalar",
    T: 12
    /* ScalarType.BYTES */
  },
  { no: 3, name: "kind", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.DataPacket_Kind) },
  { no: 4, name: "destination_sids", kind: "scalar", T: 9, repeated: true },
  { no: 6, name: "destination_identities", kind: "scalar", T: 9, repeated: true },
  { no: 5, name: "topic", kind: "scalar", T: 9, opt: true }
]);
let SendDataRequest = _SendDataRequest;
const _SendDataResponse = class _SendDataResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SendDataResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SendDataResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SendDataResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SendDataResponse, a, b);
  }
};
_SendDataResponse.runtime = import_protobuf.proto3;
_SendDataResponse.typeName = "livekit.SendDataResponse";
_SendDataResponse.fields = import_protobuf.proto3.util.newFieldList(() => []);
let SendDataResponse = _SendDataResponse;
const _UpdateRoomMetadataRequest = class _UpdateRoomMetadataRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string room = 1;
     */
    this.room = "";
    /**
     * metadata to update. skipping updates if left empty
     *
     * @generated from field: string metadata = 2;
     */
    this.metadata = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateRoomMetadataRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateRoomMetadataRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateRoomMetadataRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UpdateRoomMetadataRequest, a, b);
  }
};
_UpdateRoomMetadataRequest.runtime = import_protobuf.proto3;
_UpdateRoomMetadataRequest.typeName = "livekit.UpdateRoomMetadataRequest";
_UpdateRoomMetadataRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room",
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
let UpdateRoomMetadataRequest = _UpdateRoomMetadataRequest;
const _RoomConfiguration = class _RoomConfiguration extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * Used as ID, must be unique
     *
     * @generated from field: string name = 1;
     */
    this.name = "";
    /**
     * number of seconds to keep the room open if no one joins
     *
     * @generated from field: uint32 empty_timeout = 2;
     */
    this.emptyTimeout = 0;
    /**
     * number of seconds to keep the room open after everyone leaves
     *
     * @generated from field: uint32 departure_timeout = 3;
     */
    this.departureTimeout = 0;
    /**
     * limit number of participants that can be in a room, excluding Egress and Ingress participants
     *
     * @generated from field: uint32 max_participants = 4;
     */
    this.maxParticipants = 0;
    /**
     * playout delay of subscriber
     *
     * @generated from field: uint32 min_playout_delay = 7;
     */
    this.minPlayoutDelay = 0;
    /**
     * @generated from field: uint32 max_playout_delay = 8;
     */
    this.maxPlayoutDelay = 0;
    /**
     * improves A/V sync when playout_delay set to a value larger than 200ms. It will disables transceiver re-use
     * so not recommended for rooms with frequent subscription changes
     *
     * @generated from field: bool sync_streams = 9;
     */
    this.syncStreams = false;
    /**
     * Define agents that should be dispatched to this room
     *
     * @generated from field: repeated livekit.RoomAgentDispatch agents = 10;
     */
    this.agents = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomConfiguration().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomConfiguration().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomConfiguration().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RoomConfiguration, a, b);
  }
};
_RoomConfiguration.runtime = import_protobuf.proto3;
_RoomConfiguration.typeName = "livekit.RoomConfiguration";
_RoomConfiguration.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "empty_timeout",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "departure_timeout",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "max_participants",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 5, name: "egress", kind: "message", T: RoomEgress },
  {
    no: 7,
    name: "min_playout_delay",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 8,
    name: "max_playout_delay",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 9,
    name: "sync_streams",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 10, name: "agents", kind: "message", T: import_livekit_agent_dispatch_pb.RoomAgentDispatch, repeated: true }
]);
let RoomConfiguration = _RoomConfiguration;
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  CreateRoomRequest,
  DeleteRoomRequest,
  DeleteRoomResponse,
  ListParticipantsRequest,
  ListParticipantsResponse,
  ListRoomsRequest,
  ListRoomsResponse,
  MuteRoomTrackRequest,
  MuteRoomTrackResponse,
  RemoveParticipantResponse,
  RoomAgent,
  RoomConfiguration,
  RoomEgress,
  RoomParticipantIdentity,
  SendDataRequest,
  SendDataResponse,
  UpdateParticipantRequest,
  UpdateRoomMetadataRequest,
  UpdateSubscriptionsRequest,
  UpdateSubscriptionsResponse
});
//# sourceMappingURL=livekit_room_pb.cjs.map