import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ClientInfo, PlayoutDelay, ReconnectReason } from "./livekit_models_pb.js";
import { CreateRoomRequest } from "./livekit_room_pb.js";
import { AutoParticipantEgress, AutoTrackEgress } from "./livekit_egress_pb.js";
import { RoomAgentDispatch } from "./livekit_agent_dispatch_pb.js";
/**
 * @generated from enum livekit.NodeType
 */
export declare enum NodeType {
    /**
     * @generated from enum value: SERVER = 0;
     */
    SERVER = 0,
    /**
     * @generated from enum value: CONTROLLER = 1;
     */
    CONTROLLER = 1,
    /**
     * @generated from enum value: MEDIA = 2;
     */
    MEDIA = 2,
    /**
     * @generated from enum value: TURN = 4;
     */
    TURN = 4,
    /**
     * @generated from enum value: SWEEPER = 5;
     */
    SWEEPER = 5,
    /**
     * @generated from enum value: DIRECTOR = 6;
     */
    DIRECTOR = 6
}
/**
 * @generated from enum livekit.NodeState
 */
export declare enum NodeState {
    /**
     * @generated from enum value: STARTING_UP = 0;
     */
    STARTING_UP = 0,
    /**
     * @generated from enum value: SERVING = 1;
     */
    SERVING = 1,
    /**
     * @generated from enum value: SHUTTING_DOWN = 2;
     */
    SHUTTING_DOWN = 2
}
/**
 * @generated from enum livekit.ICECandidateType
 */
export declare enum ICECandidateType {
    /**
     * @generated from enum value: ICT_NONE = 0;
     */
    ICT_NONE = 0,
    /**
     * @generated from enum value: ICT_TCP = 1;
     */
    ICT_TCP = 1,
    /**
     * @generated from enum value: ICT_TLS = 2;
     */
    ICT_TLS = 2
}
/**
 * @generated from message livekit.Node
 */
export declare class Node extends Message<Node> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string ip = 2;
     */
    ip: string;
    /**
     * @generated from field: uint32 num_cpus = 3;
     */
    numCpus: number;
    /**
     * @generated from field: livekit.NodeStats stats = 4;
     */
    stats?: NodeStats;
    /**
     * @generated from field: livekit.NodeType type = 5;
     */
    type: NodeType;
    /**
     * @generated from field: livekit.NodeState state = 6;
     */
    state: NodeState;
    /**
     * @generated from field: string region = 7;
     */
    region: string;
    constructor(data?: PartialMessage<Node>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.Node";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Node;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Node;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Node;
    static equals(a: Node | PlainMessage<Node> | undefined, b: Node | PlainMessage<Node> | undefined): boolean;
}
/**
 * @generated from message livekit.NodeStats
 */
export declare class NodeStats extends Message<NodeStats> {
    /**
     * when server was started
     *
     * @generated from field: int64 started_at = 1;
     */
    startedAt: bigint;
    /**
     * when server last reported its status
     *
     * @generated from field: int64 updated_at = 2;
     */
    updatedAt: bigint;
    /**
     * room
     *
     * @generated from field: int32 num_rooms = 3;
     */
    numRooms: number;
    /**
     * @generated from field: int32 num_clients = 4;
     */
    numClients: number;
    /**
     * @generated from field: int32 num_tracks_in = 5;
     */
    numTracksIn: number;
    /**
     * @generated from field: int32 num_tracks_out = 6;
     */
    numTracksOut: number;
    /**
     * @generated from field: int32 num_track_publish_attempts = 36;
     */
    numTrackPublishAttempts: number;
    /**
     * @generated from field: float track_publish_attempts_per_sec = 37;
     */
    trackPublishAttemptsPerSec: number;
    /**
     * @generated from field: int32 num_track_publish_success = 38;
     */
    numTrackPublishSuccess: number;
    /**
     * @generated from field: float track_publish_success_per_sec = 39;
     */
    trackPublishSuccessPerSec: number;
    /**
     * @generated from field: int32 num_track_subscribe_attempts = 40;
     */
    numTrackSubscribeAttempts: number;
    /**
     * @generated from field: float track_subscribe_attempts_per_sec = 41;
     */
    trackSubscribeAttemptsPerSec: number;
    /**
     * @generated from field: int32 num_track_subscribe_success = 42;
     */
    numTrackSubscribeSuccess: number;
    /**
     * @generated from field: float track_subscribe_success_per_sec = 43;
     */
    trackSubscribeSuccessPerSec: number;
    /**
     * packet
     *
     * @generated from field: uint64 bytes_in = 7;
     */
    bytesIn: bigint;
    /**
     * @generated from field: uint64 bytes_out = 8;
     */
    bytesOut: bigint;
    /**
     * @generated from field: uint64 packets_in = 9;
     */
    packetsIn: bigint;
    /**
     * @generated from field: uint64 packets_out = 10;
     */
    packetsOut: bigint;
    /**
     * @generated from field: uint64 nack_total = 11;
     */
    nackTotal: bigint;
    /**
     * @generated from field: float bytes_in_per_sec = 12;
     */
    bytesInPerSec: number;
    /**
     * @generated from field: float bytes_out_per_sec = 13;
     */
    bytesOutPerSec: number;
    /**
     * @generated from field: float packets_in_per_sec = 14;
     */
    packetsInPerSec: number;
    /**
     * @generated from field: float packets_out_per_sec = 15;
     */
    packetsOutPerSec: number;
    /**
     * @generated from field: float nack_per_sec = 16;
     */
    nackPerSec: number;
    /**
     * system
     *
     * @generated from field: uint32 num_cpus = 17;
     */
    numCpus: number;
    /**
     * @generated from field: float load_avg_last1min = 18;
     */
    loadAvgLast1min: number;
    /**
     * @generated from field: float load_avg_last5min = 19;
     */
    loadAvgLast5min: number;
    /**
     * @generated from field: float load_avg_last15min = 20;
     */
    loadAvgLast15min: number;
    /**
     * @generated from field: float cpu_load = 21;
     */
    cpuLoad: number;
    /**
     * deprecated
     *
     * @generated from field: float memory_load = 33;
     */
    memoryLoad: number;
    /**
     * @generated from field: uint64 memory_total = 34;
     */
    memoryTotal: bigint;
    /**
     * @generated from field: uint64 memory_used = 35;
     */
    memoryUsed: bigint;
    /**
     * @generated from field: uint32 sys_packets_out = 28;
     */
    sysPacketsOut: number;
    /**
     * @generated from field: uint32 sys_packets_dropped = 29;
     */
    sysPacketsDropped: number;
    /**
     * @generated from field: float sys_packets_out_per_sec = 30;
     */
    sysPacketsOutPerSec: number;
    /**
     * @generated from field: float sys_packets_dropped_per_sec = 31;
     */
    sysPacketsDroppedPerSec: number;
    /**
     * @generated from field: float sys_packets_dropped_pct_per_sec = 32;
     */
    sysPacketsDroppedPctPerSec: number;
    /**
     * retransmissions
     *
     * @generated from field: uint64 retransmit_bytes_out = 22;
     */
    retransmitBytesOut: bigint;
    /**
     * @generated from field: uint64 retransmit_packets_out = 23;
     */
    retransmitPacketsOut: bigint;
    /**
     * @generated from field: float retransmit_bytes_out_per_sec = 24;
     */
    retransmitBytesOutPerSec: number;
    /**
     * @generated from field: float retransmit_packets_out_per_sec = 25;
     */
    retransmitPacketsOutPerSec: number;
    /**
     * participant joins
     *
     * @generated from field: uint64 participant_signal_connected = 26;
     */
    participantSignalConnected: bigint;
    /**
     * @generated from field: float participant_signal_connected_per_sec = 27;
     */
    participantSignalConnectedPerSec: number;
    /**
     * @generated from field: uint64 participant_rtc_connected = 44;
     */
    participantRtcConnected: bigint;
    /**
     * @generated from field: float participant_rtc_connected_per_sec = 45;
     */
    participantRtcConnectedPerSec: number;
    /**
     * @generated from field: uint64 participant_rtc_init = 46;
     */
    participantRtcInit: bigint;
    /**
     * @generated from field: float participant_rtc_init_per_sec = 47;
     */
    participantRtcInitPerSec: number;
    /**
     * forward metrics
     *
     * @generated from field: uint32 forward_latency = 48;
     */
    forwardLatency: number;
    /**
     * @generated from field: uint32 forward_jitter = 49;
     */
    forwardJitter: number;
    constructor(data?: PartialMessage<NodeStats>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.NodeStats";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeStats;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeStats;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeStats;
    static equals(a: NodeStats | PlainMessage<NodeStats> | undefined, b: NodeStats | PlainMessage<NodeStats> | undefined): boolean;
}
/**
 * @generated from message livekit.StartSession
 */
export declare class StartSession extends Message<StartSession> {
    /**
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * @generated from field: string connection_id = 3;
     */
    connectionId: string;
    /**
     * if a client is reconnecting (i.e. resume instead of restart)
     *
     * @generated from field: bool reconnect = 4;
     */
    reconnect: boolean;
    /**
     * @generated from field: bool auto_subscribe = 9;
     */
    autoSubscribe: boolean;
    /**
     * @generated from field: bool hidden = 10;
     */
    hidden: boolean;
    /**
     * @generated from field: livekit.ClientInfo client = 11;
     */
    client?: ClientInfo;
    /**
     * @generated from field: bool recorder = 12;
     */
    recorder: boolean;
    /**
     * @generated from field: string name = 13;
     */
    name: string;
    /**
     * A user's ClaimGrants serialized in JSON
     *
     * @generated from field: string grants_json = 14;
     */
    grantsJson: string;
    /**
     * @generated from field: bool adaptive_stream = 15;
     */
    adaptiveStream: boolean;
    /**
     * if reconnect, client will set current sid
     *
     * @generated from field: string participant_id = 16;
     */
    participantId: string;
    /**
     * @generated from field: livekit.ReconnectReason reconnect_reason = 17;
     */
    reconnectReason: ReconnectReason;
    /**
     * @generated from field: optional bool subscriber_allow_pause = 18;
     */
    subscriberAllowPause?: boolean;
    /**
     * @generated from field: bool disable_ice_lite = 19;
     */
    disableIceLite: boolean;
    /**
     * @generated from field: livekit.CreateRoomRequest create_room = 20;
     */
    createRoom?: CreateRoomRequest;
    constructor(data?: PartialMessage<StartSession>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.StartSession";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartSession;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartSession;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartSession;
    static equals(a: StartSession | PlainMessage<StartSession> | undefined, b: StartSession | PlainMessage<StartSession> | undefined): boolean;
}
/**
 * room info that should not be returned to clients
 *
 * @generated from message livekit.RoomInternal
 */
export declare class RoomInternal extends Message<RoomInternal> {
    /**
     * @generated from field: livekit.AutoTrackEgress track_egress = 1;
     */
    trackEgress?: AutoTrackEgress;
    /**
     * @generated from field: livekit.AutoParticipantEgress participant_egress = 2;
     */
    participantEgress?: AutoParticipantEgress;
    /**
     * @generated from field: livekit.PlayoutDelay playout_delay = 3;
     */
    playoutDelay?: PlayoutDelay;
    /**
     * @generated from field: repeated livekit.RoomAgentDispatch agent_dispatches = 5;
     */
    agentDispatches: RoomAgentDispatch[];
    /**
     * @generated from field: bool sync_streams = 4;
     */
    syncStreams: boolean;
    /**
     * @generated from field: bool replay_enabled = 6;
     */
    replayEnabled: boolean;
    constructor(data?: PartialMessage<RoomInternal>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomInternal";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomInternal;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomInternal;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomInternal;
    static equals(a: RoomInternal | PlainMessage<RoomInternal> | undefined, b: RoomInternal | PlainMessage<RoomInternal> | undefined): boolean;
}
/**
 * @generated from message livekit.ICEConfig
 */
export declare class ICEConfig extends Message<ICEConfig> {
    /**
     * @generated from field: livekit.ICECandidateType preference_subscriber = 1;
     */
    preferenceSubscriber: ICECandidateType;
    /**
     * @generated from field: livekit.ICECandidateType preference_publisher = 2;
     */
    preferencePublisher: ICECandidateType;
    constructor(data?: PartialMessage<ICEConfig>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ICEConfig";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ICEConfig;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ICEConfig;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ICEConfig;
    static equals(a: ICEConfig | PlainMessage<ICEConfig> | undefined, b: ICEConfig | PlainMessage<ICEConfig> | undefined): boolean;
}
//# sourceMappingURL=livekit_internal_pb.d.ts.map