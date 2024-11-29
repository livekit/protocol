import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { ClientInfo, ParticipantInfo, ParticipantInfo_State, ReconnectReason, Room, RTPStats, TrackInfo, VideoQuality } from "./livekit_models_pb.js";
import { EgressInfo } from "./livekit_egress_pb.js";
import { IngressInfo } from "./livekit_ingress_pb.js";
import { SIPCallInfo, SIPDispatchRuleInfo, SIPInboundTrunkInfo, SIPOutboundTrunkInfo } from "./livekit_sip_pb.js";
/**
 * @generated from enum livekit.StreamType
 */
export declare enum StreamType {
    /**
     * @generated from enum value: UPSTREAM = 0;
     */
    UPSTREAM = 0,
    /**
     * @generated from enum value: DOWNSTREAM = 1;
     */
    DOWNSTREAM = 1
}
/**
 * @generated from enum livekit.AnalyticsEventType
 */
export declare enum AnalyticsEventType {
    /**
     * @generated from enum value: ROOM_CREATED = 0;
     */
    ROOM_CREATED = 0,
    /**
     * @generated from enum value: ROOM_ENDED = 1;
     */
    ROOM_ENDED = 1,
    /**
     * @generated from enum value: PARTICIPANT_JOINED = 2;
     */
    PARTICIPANT_JOINED = 2,
    /**
     * @generated from enum value: PARTICIPANT_LEFT = 3;
     */
    PARTICIPANT_LEFT = 3,
    /**
     * @generated from enum value: TRACK_PUBLISHED = 4;
     */
    TRACK_PUBLISHED = 4,
    /**
     * @generated from enum value: TRACK_PUBLISH_REQUESTED = 20;
     */
    TRACK_PUBLISH_REQUESTED = 20,
    /**
     * @generated from enum value: TRACK_UNPUBLISHED = 5;
     */
    TRACK_UNPUBLISHED = 5,
    /**
     * @generated from enum value: TRACK_SUBSCRIBED = 6;
     */
    TRACK_SUBSCRIBED = 6,
    /**
     * @generated from enum value: TRACK_SUBSCRIBE_REQUESTED = 21;
     */
    TRACK_SUBSCRIBE_REQUESTED = 21,
    /**
     * @generated from enum value: TRACK_SUBSCRIBE_FAILED = 25;
     */
    TRACK_SUBSCRIBE_FAILED = 25,
    /**
     * @generated from enum value: TRACK_UNSUBSCRIBED = 7;
     */
    TRACK_UNSUBSCRIBED = 7,
    /**
     * @generated from enum value: TRACK_PUBLISHED_UPDATE = 10;
     */
    TRACK_PUBLISHED_UPDATE = 10,
    /**
     * @generated from enum value: TRACK_MUTED = 23;
     */
    TRACK_MUTED = 23,
    /**
     * @generated from enum value: TRACK_UNMUTED = 24;
     */
    TRACK_UNMUTED = 24,
    /**
     * @generated from enum value: TRACK_PUBLISH_STATS = 26;
     */
    TRACK_PUBLISH_STATS = 26,
    /**
     * @generated from enum value: TRACK_SUBSCRIBE_STATS = 27;
     */
    TRACK_SUBSCRIBE_STATS = 27,
    /**
     * @generated from enum value: PARTICIPANT_ACTIVE = 11;
     */
    PARTICIPANT_ACTIVE = 11,
    /**
     * @generated from enum value: PARTICIPANT_RESUMED = 22;
     */
    PARTICIPANT_RESUMED = 22,
    /**
     * @generated from enum value: EGRESS_STARTED = 12;
     */
    EGRESS_STARTED = 12,
    /**
     * @generated from enum value: EGRESS_ENDED = 13;
     */
    EGRESS_ENDED = 13,
    /**
     * @generated from enum value: EGRESS_UPDATED = 28;
     */
    EGRESS_UPDATED = 28,
    /**
     * @generated from enum value: TRACK_MAX_SUBSCRIBED_VIDEO_QUALITY = 14;
     */
    TRACK_MAX_SUBSCRIBED_VIDEO_QUALITY = 14,
    /**
     * @generated from enum value: RECONNECTED = 15;
     */
    RECONNECTED = 15,
    /**
     * @generated from enum value: INGRESS_CREATED = 18;
     */
    INGRESS_CREATED = 18,
    /**
     * @generated from enum value: INGRESS_DELETED = 19;
     */
    INGRESS_DELETED = 19,
    /**
     * @generated from enum value: INGRESS_STARTED = 16;
     */
    INGRESS_STARTED = 16,
    /**
     * @generated from enum value: INGRESS_ENDED = 17;
     */
    INGRESS_ENDED = 17,
    /**
     * @generated from enum value: INGRESS_UPDATED = 29;
     */
    INGRESS_UPDATED = 29,
    /**
     * @generated from enum value: SIP_INBOUND_TRUNK_CREATED = 30;
     */
    SIP_INBOUND_TRUNK_CREATED = 30,
    /**
     * @generated from enum value: SIP_INBOUND_TRUNK_DELETED = 31;
     */
    SIP_INBOUND_TRUNK_DELETED = 31,
    /**
     * @generated from enum value: SIP_OUTBOUND_TRUNK_CREATED = 32;
     */
    SIP_OUTBOUND_TRUNK_CREATED = 32,
    /**
     * @generated from enum value: SIP_OUTBOUND_TRUNK_DELETED = 33;
     */
    SIP_OUTBOUND_TRUNK_DELETED = 33,
    /**
     * @generated from enum value: SIP_DISPATCH_RULE_CREATED = 34;
     */
    SIP_DISPATCH_RULE_CREATED = 34,
    /**
     * @generated from enum value: SIP_DISPATCH_RULE_DELETED = 35;
     */
    SIP_DISPATCH_RULE_DELETED = 35,
    /**
     * @generated from enum value: SIP_PARTICIPANT_CREATED = 36;
     */
    SIP_PARTICIPANT_CREATED = 36,
    /**
     * @generated from enum value: SIP_CALL_INCOMING = 37;
     */
    SIP_CALL_INCOMING = 37,
    /**
     * @generated from enum value: SIP_CALL_STARTED = 38;
     */
    SIP_CALL_STARTED = 38,
    /**
     * @generated from enum value: SIP_CALL_ENDED = 39;
     */
    SIP_CALL_ENDED = 39
}
/**
 * @generated from message livekit.AnalyticsVideoLayer
 */
export declare class AnalyticsVideoLayer extends Message<AnalyticsVideoLayer> {
    /**
     * @generated from field: int32 layer = 1;
     */
    layer: number;
    /**
     * @generated from field: uint32 packets = 2;
     */
    packets: number;
    /**
     * @generated from field: uint64 bytes = 3;
     */
    bytes: bigint;
    /**
     * @generated from field: uint32 frames = 4;
     */
    frames: number;
    constructor(data?: PartialMessage<AnalyticsVideoLayer>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsVideoLayer";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsVideoLayer;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsVideoLayer;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsVideoLayer;
    static equals(a: AnalyticsVideoLayer | PlainMessage<AnalyticsVideoLayer> | undefined, b: AnalyticsVideoLayer | PlainMessage<AnalyticsVideoLayer> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsStream
 */
export declare class AnalyticsStream extends Message<AnalyticsStream> {
    /**
     * @generated from field: uint32 ssrc = 1;
     */
    ssrc: number;
    /**
     * @generated from field: uint32 primary_packets = 2;
     */
    primaryPackets: number;
    /**
     * @generated from field: uint64 primary_bytes = 3;
     */
    primaryBytes: bigint;
    /**
     * @generated from field: uint32 retransmit_packets = 4;
     */
    retransmitPackets: number;
    /**
     * @generated from field: uint64 retransmit_bytes = 5;
     */
    retransmitBytes: bigint;
    /**
     * @generated from field: uint32 padding_packets = 6;
     */
    paddingPackets: number;
    /**
     * @generated from field: uint64 padding_bytes = 7;
     */
    paddingBytes: bigint;
    /**
     * @generated from field: uint32 packets_lost = 8;
     */
    packetsLost: number;
    /**
     * @generated from field: uint32 frames = 9;
     */
    frames: number;
    /**
     * @generated from field: uint32 rtt = 10;
     */
    rtt: number;
    /**
     * @generated from field: uint32 jitter = 11;
     */
    jitter: number;
    /**
     * @generated from field: uint32 nacks = 12;
     */
    nacks: number;
    /**
     * @generated from field: uint32 plis = 13;
     */
    plis: number;
    /**
     * @generated from field: uint32 firs = 14;
     */
    firs: number;
    /**
     * @generated from field: repeated livekit.AnalyticsVideoLayer video_layers = 15;
     */
    videoLayers: AnalyticsVideoLayer[];
    /**
     * @generated from field: google.protobuf.Timestamp start_time = 17;
     */
    startTime?: Timestamp;
    /**
     * @generated from field: google.protobuf.Timestamp end_time = 18;
     */
    endTime?: Timestamp;
    /**
     * @generated from field: uint32 packets_out_of_order = 19;
     */
    packetsOutOfOrder: number;
    constructor(data?: PartialMessage<AnalyticsStream>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsStream";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsStream;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsStream;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsStream;
    static equals(a: AnalyticsStream | PlainMessage<AnalyticsStream> | undefined, b: AnalyticsStream | PlainMessage<AnalyticsStream> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsStat
 */
export declare class AnalyticsStat extends Message<AnalyticsStat> {
    /**
     * unique id for this stat
     *
     * @generated from field: string id = 14;
     */
    id: string;
    /**
     * @generated from field: string analytics_key = 1;
     */
    analyticsKey: string;
    /**
     * @generated from field: livekit.StreamType kind = 2;
     */
    kind: StreamType;
    /**
     * @generated from field: google.protobuf.Timestamp time_stamp = 3;
     */
    timeStamp?: Timestamp;
    /**
     * @generated from field: string node = 4;
     */
    node: string;
    /**
     * @generated from field: string room_id = 5;
     */
    roomId: string;
    /**
     * @generated from field: string room_name = 6;
     */
    roomName: string;
    /**
     * @generated from field: string participant_id = 7;
     */
    participantId: string;
    /**
     * @generated from field: string track_id = 8;
     */
    trackId: string;
    /**
     * average score
     *
     * @generated from field: float score = 9;
     */
    score: number;
    /**
     * @generated from field: repeated livekit.AnalyticsStream streams = 10;
     */
    streams: AnalyticsStream[];
    /**
     * @generated from field: string mime = 11;
     */
    mime: string;
    /**
     * @generated from field: float min_score = 12;
     */
    minScore: number;
    /**
     * @generated from field: float median_score = 13;
     */
    medianScore: number;
    constructor(data?: PartialMessage<AnalyticsStat>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsStat";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsStat;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsStat;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsStat;
    static equals(a: AnalyticsStat | PlainMessage<AnalyticsStat> | undefined, b: AnalyticsStat | PlainMessage<AnalyticsStat> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsStats
 */
export declare class AnalyticsStats extends Message<AnalyticsStats> {
    /**
     * @generated from field: repeated livekit.AnalyticsStat stats = 1;
     */
    stats: AnalyticsStat[];
    constructor(data?: PartialMessage<AnalyticsStats>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsStats";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsStats;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsStats;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsStats;
    static equals(a: AnalyticsStats | PlainMessage<AnalyticsStats> | undefined, b: AnalyticsStats | PlainMessage<AnalyticsStats> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsClientMeta
 */
export declare class AnalyticsClientMeta extends Message<AnalyticsClientMeta> {
    /**
     * @generated from field: string region = 1;
     */
    region: string;
    /**
     * @generated from field: string node = 2;
     */
    node: string;
    /**
     * @generated from field: string client_addr = 3;
     */
    clientAddr: string;
    /**
     * @generated from field: uint32 client_connect_time = 4;
     */
    clientConnectTime: number;
    /**
     * udp, tcp, turn
     *
     * @generated from field: string connection_type = 5;
     */
    connectionType: string;
    /**
     * @generated from field: livekit.ReconnectReason reconnect_reason = 6;
     */
    reconnectReason: ReconnectReason;
    /**
     * @generated from field: optional string geo_hash = 7;
     */
    geoHash?: string;
    /**
     * @generated from field: optional string country = 8;
     */
    country?: string;
    /**
     * @generated from field: optional uint32 isp_asn = 9;
     */
    ispAsn?: number;
    constructor(data?: PartialMessage<AnalyticsClientMeta>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsClientMeta";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsClientMeta;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsClientMeta;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsClientMeta;
    static equals(a: AnalyticsClientMeta | PlainMessage<AnalyticsClientMeta> | undefined, b: AnalyticsClientMeta | PlainMessage<AnalyticsClientMeta> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsEvent
 */
export declare class AnalyticsEvent extends Message<AnalyticsEvent> {
    /**
     * unique id for this event
     *
     * @generated from field: string id = 25;
     */
    id: string;
    /**
     * @generated from field: livekit.AnalyticsEventType type = 1;
     */
    type: AnalyticsEventType;
    /**
     * @generated from field: google.protobuf.Timestamp timestamp = 2;
     */
    timestamp?: Timestamp;
    /**
     * @generated from field: string room_id = 3;
     */
    roomId: string;
    /**
     * @generated from field: livekit.Room room = 4;
     */
    room?: Room;
    /**
     * @generated from field: string participant_id = 5;
     */
    participantId: string;
    /**
     * @generated from field: livekit.ParticipantInfo participant = 6;
     */
    participant?: ParticipantInfo;
    /**
     * @generated from field: string track_id = 7;
     */
    trackId: string;
    /**
     * @generated from field: livekit.TrackInfo track = 8;
     */
    track?: TrackInfo;
    /**
     * @generated from field: string analytics_key = 10;
     */
    analyticsKey: string;
    /**
     * @generated from field: livekit.ClientInfo client_info = 11;
     */
    clientInfo?: ClientInfo;
    /**
     * @generated from field: livekit.AnalyticsClientMeta client_meta = 12;
     */
    clientMeta?: AnalyticsClientMeta;
    /**
     * @generated from field: string egress_id = 13;
     */
    egressId: string;
    /**
     * @generated from field: string ingress_id = 19;
     */
    ingressId: string;
    /**
     * @generated from field: livekit.VideoQuality max_subscribed_video_quality = 14;
     */
    maxSubscribedVideoQuality: VideoQuality;
    /**
     * @generated from field: livekit.ParticipantInfo publisher = 15;
     */
    publisher?: ParticipantInfo;
    /**
     * @generated from field: string mime = 16;
     */
    mime: string;
    /**
     * @generated from field: livekit.EgressInfo egress = 17;
     */
    egress?: EgressInfo;
    /**
     * @generated from field: livekit.IngressInfo ingress = 18;
     */
    ingress?: IngressInfo;
    /**
     * @generated from field: string error = 20;
     */
    error: string;
    /**
     * @generated from field: livekit.RTPStats rtp_stats = 21;
     */
    rtpStats?: RTPStats;
    /**
     * @generated from field: int32 video_layer = 22;
     */
    videoLayer: number;
    /**
     * @generated from field: string node_id = 24;
     */
    nodeId: string;
    /**
     * @generated from field: string sip_call_id = 26;
     */
    sipCallId: string;
    /**
     * @generated from field: livekit.SIPCallInfo sip_call = 27;
     */
    sipCall?: SIPCallInfo;
    /**
     * @generated from field: string sip_trunk_id = 28;
     */
    sipTrunkId: string;
    /**
     * @generated from field: livekit.SIPInboundTrunkInfo sip_inbound_trunk = 29;
     */
    sipInboundTrunk?: SIPInboundTrunkInfo;
    /**
     * @generated from field: livekit.SIPOutboundTrunkInfo sip_outbound_trunk = 30;
     */
    sipOutboundTrunk?: SIPOutboundTrunkInfo;
    /**
     * @generated from field: string sip_dispatch_rule_id = 31;
     */
    sipDispatchRuleId: string;
    /**
     * @generated from field: livekit.SIPDispatchRuleInfo sip_dispatch_rule = 32;
     */
    sipDispatchRule?: SIPDispatchRuleInfo;
    constructor(data?: PartialMessage<AnalyticsEvent>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsEvent";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsEvent;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsEvent;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsEvent;
    static equals(a: AnalyticsEvent | PlainMessage<AnalyticsEvent> | undefined, b: AnalyticsEvent | PlainMessage<AnalyticsEvent> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsEvents
 */
export declare class AnalyticsEvents extends Message<AnalyticsEvents> {
    /**
     * @generated from field: repeated livekit.AnalyticsEvent events = 1;
     */
    events: AnalyticsEvent[];
    constructor(data?: PartialMessage<AnalyticsEvents>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsEvents";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsEvents;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsEvents;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsEvents;
    static equals(a: AnalyticsEvents | PlainMessage<AnalyticsEvents> | undefined, b: AnalyticsEvents | PlainMessage<AnalyticsEvents> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsRoomParticipant
 */
export declare class AnalyticsRoomParticipant extends Message<AnalyticsRoomParticipant> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * @generated from field: livekit.ParticipantInfo.State state = 4;
     */
    state: ParticipantInfo_State;
    /**
     * @generated from field: google.protobuf.Timestamp joined_at = 5;
     */
    joinedAt?: Timestamp;
    constructor(data?: PartialMessage<AnalyticsRoomParticipant>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsRoomParticipant";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsRoomParticipant;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsRoomParticipant;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsRoomParticipant;
    static equals(a: AnalyticsRoomParticipant | PlainMessage<AnalyticsRoomParticipant> | undefined, b: AnalyticsRoomParticipant | PlainMessage<AnalyticsRoomParticipant> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsRoom
 */
export declare class AnalyticsRoom extends Message<AnalyticsRoom> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * @generated from field: string project_id = 5;
     */
    projectId: string;
    /**
     * @generated from field: google.protobuf.Timestamp created_at = 3;
     */
    createdAt?: Timestamp;
    /**
     * @generated from field: repeated livekit.AnalyticsRoomParticipant participants = 4;
     */
    participants: AnalyticsRoomParticipant[];
    constructor(data?: PartialMessage<AnalyticsRoom>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsRoom";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsRoom;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsRoom;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsRoom;
    static equals(a: AnalyticsRoom | PlainMessage<AnalyticsRoom> | undefined, b: AnalyticsRoom | PlainMessage<AnalyticsRoom> | undefined): boolean;
}
/**
 * @generated from message livekit.AnalyticsNodeRooms
 */
export declare class AnalyticsNodeRooms extends Message<AnalyticsNodeRooms> {
    /**
     * @generated from field: string node_id = 1;
     */
    nodeId: string;
    /**
     * @generated from field: uint64 sequence_number = 2;
     */
    sequenceNumber: bigint;
    /**
     * @generated from field: google.protobuf.Timestamp timestamp = 3;
     */
    timestamp?: Timestamp;
    /**
     * @generated from field: repeated livekit.AnalyticsRoom rooms = 4;
     */
    rooms: AnalyticsRoom[];
    constructor(data?: PartialMessage<AnalyticsNodeRooms>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AnalyticsNodeRooms";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyticsNodeRooms;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyticsNodeRooms;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyticsNodeRooms;
    static equals(a: AnalyticsNodeRooms | PlainMessage<AnalyticsNodeRooms> | undefined, b: AnalyticsNodeRooms | PlainMessage<AnalyticsNodeRooms> | undefined): boolean;
}
//# sourceMappingURL=livekit_analytics_pb.d.ts.map