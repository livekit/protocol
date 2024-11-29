import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { RoomAgentDispatch } from "./livekit_agent_dispatch_pb.js";
import { AutoParticipantEgress, AutoTrackEgress, RoomCompositeEgressRequest } from "./livekit_egress_pb.js";
import { DataPacket_Kind, ParticipantInfo, ParticipantPermission, ParticipantTracks, Room, TrackInfo } from "./livekit_models_pb.js";
/**
 * @generated from message livekit.CreateRoomRequest
 */
export declare class CreateRoomRequest extends Message<CreateRoomRequest> {
    /**
     * name of the room
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * configuration to use for this room parameters. Setting parameters below override the config defaults.
     *
     * @generated from field: string room_preset = 12;
     */
    roomPreset: string;
    /**
     * number of seconds to keep the room open if no one joins
     *
     * @generated from field: uint32 empty_timeout = 2;
     */
    emptyTimeout: number;
    /**
     * number of seconds to keep the room open after everyone leaves
     *
     * @generated from field: uint32 departure_timeout = 10;
     */
    departureTimeout: number;
    /**
     * limit number of participants that can be in a room
     *
     * @generated from field: uint32 max_participants = 3;
     */
    maxParticipants: number;
    /**
     * override the node room is allocated to, for debugging
     *
     * @generated from field: string node_id = 4;
     */
    nodeId: string;
    /**
     * metadata of room
     *
     * @generated from field: string metadata = 5;
     */
    metadata: string;
    /**
     * auto-egress configurations
     *
     * @generated from field: livekit.RoomEgress egress = 6;
     */
    egress?: RoomEgress;
    /**
     * playout delay of subscriber
     *
     * @generated from field: uint32 min_playout_delay = 7;
     */
    minPlayoutDelay: number;
    /**
     * @generated from field: uint32 max_playout_delay = 8;
     */
    maxPlayoutDelay: number;
    /**
     * improves A/V sync when playout_delay set to a value larger than 200ms. It will disables transceiver re-use
     * so not recommended for rooms with frequent subscription changes
     *
     * @generated from field: bool sync_streams = 9;
     */
    syncStreams: boolean;
    /**
     * replay
     *
     * @generated from field: bool replay_enabled = 13;
     */
    replayEnabled: boolean;
    /**
     * Define agents that should be dispatched to this room
     *
     * NEXT-ID: 15
     *
     * @generated from field: repeated livekit.RoomAgentDispatch agents = 14;
     */
    agents: RoomAgentDispatch[];
    constructor(data?: PartialMessage<CreateRoomRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateRoomRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRoomRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRoomRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRoomRequest;
    static equals(a: CreateRoomRequest | PlainMessage<CreateRoomRequest> | undefined, b: CreateRoomRequest | PlainMessage<CreateRoomRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.RoomEgress
 */
export declare class RoomEgress extends Message<RoomEgress> {
    /**
     * @generated from field: livekit.RoomCompositeEgressRequest room = 1;
     */
    room?: RoomCompositeEgressRequest;
    /**
     * @generated from field: livekit.AutoParticipantEgress participant = 3;
     */
    participant?: AutoParticipantEgress;
    /**
     * @generated from field: livekit.AutoTrackEgress tracks = 2;
     */
    tracks?: AutoTrackEgress;
    constructor(data?: PartialMessage<RoomEgress>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomEgress";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomEgress;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomEgress;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomEgress;
    static equals(a: RoomEgress | PlainMessage<RoomEgress> | undefined, b: RoomEgress | PlainMessage<RoomEgress> | undefined): boolean;
}
/**
 * @generated from message livekit.RoomAgent
 */
export declare class RoomAgent extends Message<RoomAgent> {
    /**
     * @generated from field: repeated livekit.RoomAgentDispatch dispatches = 1;
     */
    dispatches: RoomAgentDispatch[];
    constructor(data?: PartialMessage<RoomAgent>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomAgent";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomAgent;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomAgent;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomAgent;
    static equals(a: RoomAgent | PlainMessage<RoomAgent> | undefined, b: RoomAgent | PlainMessage<RoomAgent> | undefined): boolean;
}
/**
 * @generated from message livekit.ListRoomsRequest
 */
export declare class ListRoomsRequest extends Message<ListRoomsRequest> {
    /**
     * when set, will only return rooms with name match
     *
     * @generated from field: repeated string names = 1;
     */
    names: string[];
    constructor(data?: PartialMessage<ListRoomsRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListRoomsRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRoomsRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRoomsRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRoomsRequest;
    static equals(a: ListRoomsRequest | PlainMessage<ListRoomsRequest> | undefined, b: ListRoomsRequest | PlainMessage<ListRoomsRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListRoomsResponse
 */
export declare class ListRoomsResponse extends Message<ListRoomsResponse> {
    /**
     * @generated from field: repeated livekit.Room rooms = 1;
     */
    rooms: Room[];
    constructor(data?: PartialMessage<ListRoomsResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListRoomsResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRoomsResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRoomsResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRoomsResponse;
    static equals(a: ListRoomsResponse | PlainMessage<ListRoomsResponse> | undefined, b: ListRoomsResponse | PlainMessage<ListRoomsResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.DeleteRoomRequest
 */
export declare class DeleteRoomRequest extends Message<DeleteRoomRequest> {
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    room: string;
    constructor(data?: PartialMessage<DeleteRoomRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DeleteRoomRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRoomRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRoomRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRoomRequest;
    static equals(a: DeleteRoomRequest | PlainMessage<DeleteRoomRequest> | undefined, b: DeleteRoomRequest | PlainMessage<DeleteRoomRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.DeleteRoomResponse
 */
export declare class DeleteRoomResponse extends Message<DeleteRoomResponse> {
    constructor(data?: PartialMessage<DeleteRoomResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DeleteRoomResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRoomResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRoomResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRoomResponse;
    static equals(a: DeleteRoomResponse | PlainMessage<DeleteRoomResponse> | undefined, b: DeleteRoomResponse | PlainMessage<DeleteRoomResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.ListParticipantsRequest
 */
export declare class ListParticipantsRequest extends Message<ListParticipantsRequest> {
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    room: string;
    constructor(data?: PartialMessage<ListParticipantsRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListParticipantsRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListParticipantsRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListParticipantsRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListParticipantsRequest;
    static equals(a: ListParticipantsRequest | PlainMessage<ListParticipantsRequest> | undefined, b: ListParticipantsRequest | PlainMessage<ListParticipantsRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListParticipantsResponse
 */
export declare class ListParticipantsResponse extends Message<ListParticipantsResponse> {
    /**
     * @generated from field: repeated livekit.ParticipantInfo participants = 1;
     */
    participants: ParticipantInfo[];
    constructor(data?: PartialMessage<ListParticipantsResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListParticipantsResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListParticipantsResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListParticipantsResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListParticipantsResponse;
    static equals(a: ListParticipantsResponse | PlainMessage<ListParticipantsResponse> | undefined, b: ListParticipantsResponse | PlainMessage<ListParticipantsResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.RoomParticipantIdentity
 */
export declare class RoomParticipantIdentity extends Message<RoomParticipantIdentity> {
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    room: string;
    /**
     * identity of the participant
     *
     * @generated from field: string identity = 2;
     */
    identity: string;
    constructor(data?: PartialMessage<RoomParticipantIdentity>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomParticipantIdentity";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomParticipantIdentity;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomParticipantIdentity;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomParticipantIdentity;
    static equals(a: RoomParticipantIdentity | PlainMessage<RoomParticipantIdentity> | undefined, b: RoomParticipantIdentity | PlainMessage<RoomParticipantIdentity> | undefined): boolean;
}
/**
 * @generated from message livekit.RemoveParticipantResponse
 */
export declare class RemoveParticipantResponse extends Message<RemoveParticipantResponse> {
    constructor(data?: PartialMessage<RemoveParticipantResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RemoveParticipantResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveParticipantResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveParticipantResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveParticipantResponse;
    static equals(a: RemoveParticipantResponse | PlainMessage<RemoveParticipantResponse> | undefined, b: RemoveParticipantResponse | PlainMessage<RemoveParticipantResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.MuteRoomTrackRequest
 */
export declare class MuteRoomTrackRequest extends Message<MuteRoomTrackRequest> {
    /**
     * name of the room
     *
     * @generated from field: string room = 1;
     */
    room: string;
    /**
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * sid of the track to mute
     *
     * @generated from field: string track_sid = 3;
     */
    trackSid: string;
    /**
     * set to true to mute, false to unmute
     *
     * @generated from field: bool muted = 4;
     */
    muted: boolean;
    constructor(data?: PartialMessage<MuteRoomTrackRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.MuteRoomTrackRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MuteRoomTrackRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MuteRoomTrackRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MuteRoomTrackRequest;
    static equals(a: MuteRoomTrackRequest | PlainMessage<MuteRoomTrackRequest> | undefined, b: MuteRoomTrackRequest | PlainMessage<MuteRoomTrackRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.MuteRoomTrackResponse
 */
export declare class MuteRoomTrackResponse extends Message<MuteRoomTrackResponse> {
    /**
     * @generated from field: livekit.TrackInfo track = 1;
     */
    track?: TrackInfo;
    constructor(data?: PartialMessage<MuteRoomTrackResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.MuteRoomTrackResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MuteRoomTrackResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MuteRoomTrackResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MuteRoomTrackResponse;
    static equals(a: MuteRoomTrackResponse | PlainMessage<MuteRoomTrackResponse> | undefined, b: MuteRoomTrackResponse | PlainMessage<MuteRoomTrackResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateParticipantRequest
 */
export declare class UpdateParticipantRequest extends Message<UpdateParticipantRequest> {
    /**
     * @generated from field: string room = 1;
     */
    room: string;
    /**
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * metadata to update. skipping updates if left empty
     *
     * @generated from field: string metadata = 3;
     */
    metadata: string;
    /**
     * set to update the participant's permissions
     *
     * @generated from field: livekit.ParticipantPermission permission = 4;
     */
    permission?: ParticipantPermission;
    /**
     * display name to update
     *
     * @generated from field: string name = 5;
     */
    name: string;
    /**
     * attributes to update. it only updates attributes that have been set
     * to delete attributes, set the value to an empty string
     *
     * @generated from field: map<string, string> attributes = 6;
     */
    attributes: {
        [key: string]: string;
    };
    constructor(data?: PartialMessage<UpdateParticipantRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateParticipantRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateParticipantRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateParticipantRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateParticipantRequest;
    static equals(a: UpdateParticipantRequest | PlainMessage<UpdateParticipantRequest> | undefined, b: UpdateParticipantRequest | PlainMessage<UpdateParticipantRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateSubscriptionsRequest
 */
export declare class UpdateSubscriptionsRequest extends Message<UpdateSubscriptionsRequest> {
    /**
     * @generated from field: string room = 1;
     */
    room: string;
    /**
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * list of sids of tracks
     *
     * @generated from field: repeated string track_sids = 3;
     */
    trackSids: string[];
    /**
     * set to true to subscribe, false to unsubscribe from tracks
     *
     * @generated from field: bool subscribe = 4;
     */
    subscribe: boolean;
    /**
     * list of participants and their tracks
     *
     * @generated from field: repeated livekit.ParticipantTracks participant_tracks = 5;
     */
    participantTracks: ParticipantTracks[];
    constructor(data?: PartialMessage<UpdateSubscriptionsRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateSubscriptionsRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateSubscriptionsRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateSubscriptionsRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateSubscriptionsRequest;
    static equals(a: UpdateSubscriptionsRequest | PlainMessage<UpdateSubscriptionsRequest> | undefined, b: UpdateSubscriptionsRequest | PlainMessage<UpdateSubscriptionsRequest> | undefined): boolean;
}
/**
 * empty for now
 *
 * @generated from message livekit.UpdateSubscriptionsResponse
 */
export declare class UpdateSubscriptionsResponse extends Message<UpdateSubscriptionsResponse> {
    constructor(data?: PartialMessage<UpdateSubscriptionsResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateSubscriptionsResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateSubscriptionsResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateSubscriptionsResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateSubscriptionsResponse;
    static equals(a: UpdateSubscriptionsResponse | PlainMessage<UpdateSubscriptionsResponse> | undefined, b: UpdateSubscriptionsResponse | PlainMessage<UpdateSubscriptionsResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.SendDataRequest
 */
export declare class SendDataRequest extends Message<SendDataRequest> {
    /**
     * @generated from field: string room = 1;
     */
    room: string;
    /**
     * @generated from field: bytes data = 2;
     */
    data: Uint8Array;
    /**
     * @generated from field: livekit.DataPacket.Kind kind = 3;
     */
    kind: DataPacket_Kind;
    /**
     * mark deprecated
     *
     * @generated from field: repeated string destination_sids = 4 [deprecated = true];
     * @deprecated
     */
    destinationSids: string[];
    /**
     * when set, only forward to these identities
     *
     * @generated from field: repeated string destination_identities = 6;
     */
    destinationIdentities: string[];
    /**
     * @generated from field: optional string topic = 5;
     */
    topic?: string;
    constructor(data?: PartialMessage<SendDataRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SendDataRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendDataRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendDataRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendDataRequest;
    static equals(a: SendDataRequest | PlainMessage<SendDataRequest> | undefined, b: SendDataRequest | PlainMessage<SendDataRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SendDataResponse
 */
export declare class SendDataResponse extends Message<SendDataResponse> {
    constructor(data?: PartialMessage<SendDataResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SendDataResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendDataResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendDataResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendDataResponse;
    static equals(a: SendDataResponse | PlainMessage<SendDataResponse> | undefined, b: SendDataResponse | PlainMessage<SendDataResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateRoomMetadataRequest
 */
export declare class UpdateRoomMetadataRequest extends Message<UpdateRoomMetadataRequest> {
    /**
     * @generated from field: string room = 1;
     */
    room: string;
    /**
     * metadata to update. skipping updates if left empty
     *
     * @generated from field: string metadata = 2;
     */
    metadata: string;
    constructor(data?: PartialMessage<UpdateRoomMetadataRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateRoomMetadataRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRoomMetadataRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRoomMetadataRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRoomMetadataRequest;
    static equals(a: UpdateRoomMetadataRequest | PlainMessage<UpdateRoomMetadataRequest> | undefined, b: UpdateRoomMetadataRequest | PlainMessage<UpdateRoomMetadataRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.RoomConfiguration
 */
export declare class RoomConfiguration extends Message<RoomConfiguration> {
    /**
     * Used as ID, must be unique
     *
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * number of seconds to keep the room open if no one joins
     *
     * @generated from field: uint32 empty_timeout = 2;
     */
    emptyTimeout: number;
    /**
     * number of seconds to keep the room open after everyone leaves
     *
     * @generated from field: uint32 departure_timeout = 3;
     */
    departureTimeout: number;
    /**
     * limit number of participants that can be in a room, excluding Egress and Ingress participants
     *
     * @generated from field: uint32 max_participants = 4;
     */
    maxParticipants: number;
    /**
     * egress
     *
     * @generated from field: livekit.RoomEgress egress = 5;
     */
    egress?: RoomEgress;
    /**
     * playout delay of subscriber
     *
     * @generated from field: uint32 min_playout_delay = 7;
     */
    minPlayoutDelay: number;
    /**
     * @generated from field: uint32 max_playout_delay = 8;
     */
    maxPlayoutDelay: number;
    /**
     * improves A/V sync when playout_delay set to a value larger than 200ms. It will disables transceiver re-use
     * so not recommended for rooms with frequent subscription changes
     *
     * @generated from field: bool sync_streams = 9;
     */
    syncStreams: boolean;
    /**
     * Define agents that should be dispatched to this room
     *
     * @generated from field: repeated livekit.RoomAgentDispatch agents = 10;
     */
    agents: RoomAgentDispatch[];
    constructor(data?: PartialMessage<RoomConfiguration>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomConfiguration";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomConfiguration;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomConfiguration;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomConfiguration;
    static equals(a: RoomConfiguration | PlainMessage<RoomConfiguration> | undefined, b: RoomConfiguration | PlainMessage<RoomConfiguration> | undefined): boolean;
}
//# sourceMappingURL=livekit_room_pb.d.ts.map