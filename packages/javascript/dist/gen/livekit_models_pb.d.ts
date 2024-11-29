import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { MetricsBatch } from "./livekit_metrics_pb.js";
/**
 * @generated from enum livekit.AudioCodec
 */
export declare enum AudioCodec {
    /**
     * @generated from enum value: DEFAULT_AC = 0;
     */
    DEFAULT_AC = 0,
    /**
     * @generated from enum value: OPUS = 1;
     */
    OPUS = 1,
    /**
     * @generated from enum value: AAC = 2;
     */
    AAC = 2
}
/**
 * @generated from enum livekit.VideoCodec
 */
export declare enum VideoCodec {
    /**
     * @generated from enum value: DEFAULT_VC = 0;
     */
    DEFAULT_VC = 0,
    /**
     * @generated from enum value: H264_BASELINE = 1;
     */
    H264_BASELINE = 1,
    /**
     * @generated from enum value: H264_MAIN = 2;
     */
    H264_MAIN = 2,
    /**
     * @generated from enum value: H264_HIGH = 3;
     */
    H264_HIGH = 3,
    /**
     * @generated from enum value: VP8 = 4;
     */
    VP8 = 4
}
/**
 * @generated from enum livekit.ImageCodec
 */
export declare enum ImageCodec {
    /**
     * @generated from enum value: IC_DEFAULT = 0;
     */
    IC_DEFAULT = 0,
    /**
     * @generated from enum value: IC_JPEG = 1;
     */
    IC_JPEG = 1
}
/**
 * @generated from enum livekit.TrackType
 */
export declare enum TrackType {
    /**
     * @generated from enum value: AUDIO = 0;
     */
    AUDIO = 0,
    /**
     * @generated from enum value: VIDEO = 1;
     */
    VIDEO = 1,
    /**
     * @generated from enum value: DATA = 2;
     */
    DATA = 2
}
/**
 * @generated from enum livekit.TrackSource
 */
export declare enum TrackSource {
    /**
     * @generated from enum value: UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from enum value: CAMERA = 1;
     */
    CAMERA = 1,
    /**
     * @generated from enum value: MICROPHONE = 2;
     */
    MICROPHONE = 2,
    /**
     * @generated from enum value: SCREEN_SHARE = 3;
     */
    SCREEN_SHARE = 3,
    /**
     * @generated from enum value: SCREEN_SHARE_AUDIO = 4;
     */
    SCREEN_SHARE_AUDIO = 4
}
/**
 * @generated from enum livekit.VideoQuality
 */
export declare enum VideoQuality {
    /**
     * @generated from enum value: LOW = 0;
     */
    LOW = 0,
    /**
     * @generated from enum value: MEDIUM = 1;
     */
    MEDIUM = 1,
    /**
     * @generated from enum value: HIGH = 2;
     */
    HIGH = 2,
    /**
     * @generated from enum value: OFF = 3;
     */
    OFF = 3
}
/**
 * @generated from enum livekit.ConnectionQuality
 */
export declare enum ConnectionQuality {
    /**
     * @generated from enum value: POOR = 0;
     */
    POOR = 0,
    /**
     * @generated from enum value: GOOD = 1;
     */
    GOOD = 1,
    /**
     * @generated from enum value: EXCELLENT = 2;
     */
    EXCELLENT = 2,
    /**
     * @generated from enum value: LOST = 3;
     */
    LOST = 3
}
/**
 * @generated from enum livekit.ClientConfigSetting
 */
export declare enum ClientConfigSetting {
    /**
     * @generated from enum value: UNSET = 0;
     */
    UNSET = 0,
    /**
     * @generated from enum value: DISABLED = 1;
     */
    DISABLED = 1,
    /**
     * @generated from enum value: ENABLED = 2;
     */
    ENABLED = 2
}
/**
 * @generated from enum livekit.DisconnectReason
 */
export declare enum DisconnectReason {
    /**
     * @generated from enum value: UNKNOWN_REASON = 0;
     */
    UNKNOWN_REASON = 0,
    /**
     * the client initiated the disconnect
     *
     * @generated from enum value: CLIENT_INITIATED = 1;
     */
    CLIENT_INITIATED = 1,
    /**
     * another participant with the same identity has joined the room
     *
     * @generated from enum value: DUPLICATE_IDENTITY = 2;
     */
    DUPLICATE_IDENTITY = 2,
    /**
     * the server instance is shutting down
     *
     * @generated from enum value: SERVER_SHUTDOWN = 3;
     */
    SERVER_SHUTDOWN = 3,
    /**
     * RoomService.RemoveParticipant was called
     *
     * @generated from enum value: PARTICIPANT_REMOVED = 4;
     */
    PARTICIPANT_REMOVED = 4,
    /**
     * RoomService.DeleteRoom was called
     *
     * @generated from enum value: ROOM_DELETED = 5;
     */
    ROOM_DELETED = 5,
    /**
     * the client is attempting to resume a session, but server is not aware of it
     *
     * @generated from enum value: STATE_MISMATCH = 6;
     */
    STATE_MISMATCH = 6,
    /**
     * client was unable to connect fully
     *
     * @generated from enum value: JOIN_FAILURE = 7;
     */
    JOIN_FAILURE = 7,
    /**
     * Cloud-only, the server requested Participant to migrate the connection elsewhere
     *
     * @generated from enum value: MIGRATION = 8;
     */
    MIGRATION = 8,
    /**
     * the signal websocket was closed unexpectedly
     *
     * @generated from enum value: SIGNAL_CLOSE = 9;
     */
    SIGNAL_CLOSE = 9,
    /**
     * the room was closed, due to all Standard and Ingress participants having left
     *
     * @generated from enum value: ROOM_CLOSED = 10;
     */
    ROOM_CLOSED = 10,
    /**
     * SIP callee did not respond in time
     *
     * @generated from enum value: USER_UNAVAILABLE = 11;
     */
    USER_UNAVAILABLE = 11,
    /**
     * SIP callee rejected the call (busy)
     *
     * @generated from enum value: USER_REJECTED = 12;
     */
    USER_REJECTED = 12,
    /**
     * SIP protocol failure or unexpected response
     *
     * @generated from enum value: SIP_TRUNK_FAILURE = 13;
     */
    SIP_TRUNK_FAILURE = 13
}
/**
 * @generated from enum livekit.ReconnectReason
 */
export declare enum ReconnectReason {
    /**
     * @generated from enum value: RR_UNKNOWN = 0;
     */
    RR_UNKNOWN = 0,
    /**
     * @generated from enum value: RR_SIGNAL_DISCONNECTED = 1;
     */
    RR_SIGNAL_DISCONNECTED = 1,
    /**
     * @generated from enum value: RR_PUBLISHER_FAILED = 2;
     */
    RR_PUBLISHER_FAILED = 2,
    /**
     * @generated from enum value: RR_SUBSCRIBER_FAILED = 3;
     */
    RR_SUBSCRIBER_FAILED = 3,
    /**
     * @generated from enum value: RR_SWITCH_CANDIDATE = 4;
     */
    RR_SWITCH_CANDIDATE = 4
}
/**
 * @generated from enum livekit.SubscriptionError
 */
export declare enum SubscriptionError {
    /**
     * @generated from enum value: SE_UNKNOWN = 0;
     */
    SE_UNKNOWN = 0,
    /**
     * @generated from enum value: SE_CODEC_UNSUPPORTED = 1;
     */
    SE_CODEC_UNSUPPORTED = 1,
    /**
     * @generated from enum value: SE_TRACK_NOTFOUND = 2;
     */
    SE_TRACK_NOTFOUND = 2
}
/**
 * @generated from enum livekit.AudioTrackFeature
 */
export declare enum AudioTrackFeature {
    /**
     * @generated from enum value: TF_STEREO = 0;
     */
    TF_STEREO = 0,
    /**
     * @generated from enum value: TF_NO_DTX = 1;
     */
    TF_NO_DTX = 1,
    /**
     * @generated from enum value: TF_AUTO_GAIN_CONTROL = 2;
     */
    TF_AUTO_GAIN_CONTROL = 2,
    /**
     * @generated from enum value: TF_ECHO_CANCELLATION = 3;
     */
    TF_ECHO_CANCELLATION = 3,
    /**
     * @generated from enum value: TF_NOISE_SUPPRESSION = 4;
     */
    TF_NOISE_SUPPRESSION = 4,
    /**
     * @generated from enum value: TF_ENHANCED_NOISE_CANCELLATION = 5;
     */
    TF_ENHANCED_NOISE_CANCELLATION = 5
}
/**
 * @generated from message livekit.Room
 */
export declare class Room extends Message<Room> {
    /**
     * @generated from field: string sid = 1;
     */
    sid: string;
    /**
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * @generated from field: uint32 empty_timeout = 3;
     */
    emptyTimeout: number;
    /**
     * @generated from field: uint32 departure_timeout = 14;
     */
    departureTimeout: number;
    /**
     * @generated from field: uint32 max_participants = 4;
     */
    maxParticipants: number;
    /**
     * @generated from field: int64 creation_time = 5;
     */
    creationTime: bigint;
    /**
     * @generated from field: string turn_password = 6;
     */
    turnPassword: string;
    /**
     * @generated from field: repeated livekit.Codec enabled_codecs = 7;
     */
    enabledCodecs: Codec[];
    /**
     * @generated from field: string metadata = 8;
     */
    metadata: string;
    /**
     * @generated from field: uint32 num_participants = 9;
     */
    numParticipants: number;
    /**
     * @generated from field: uint32 num_publishers = 11;
     */
    numPublishers: number;
    /**
     * @generated from field: bool active_recording = 10;
     */
    activeRecording: boolean;
    /**
     * @generated from field: livekit.TimedVersion version = 13;
     */
    version?: TimedVersion;
    constructor(data?: PartialMessage<Room>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.Room";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Room;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Room;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Room;
    static equals(a: Room | PlainMessage<Room> | undefined, b: Room | PlainMessage<Room> | undefined): boolean;
}
/**
 * @generated from message livekit.Codec
 */
export declare class Codec extends Message<Codec> {
    /**
     * @generated from field: string mime = 1;
     */
    mime: string;
    /**
     * @generated from field: string fmtp_line = 2;
     */
    fmtpLine: string;
    constructor(data?: PartialMessage<Codec>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.Codec";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Codec;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Codec;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Codec;
    static equals(a: Codec | PlainMessage<Codec> | undefined, b: Codec | PlainMessage<Codec> | undefined): boolean;
}
/**
 * @generated from message livekit.PlayoutDelay
 */
export declare class PlayoutDelay extends Message<PlayoutDelay> {
    /**
     * @generated from field: bool enabled = 1;
     */
    enabled: boolean;
    /**
     * @generated from field: uint32 min = 2;
     */
    min: number;
    /**
     * @generated from field: uint32 max = 3;
     */
    max: number;
    constructor(data?: PartialMessage<PlayoutDelay>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.PlayoutDelay";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PlayoutDelay;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PlayoutDelay;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PlayoutDelay;
    static equals(a: PlayoutDelay | PlainMessage<PlayoutDelay> | undefined, b: PlayoutDelay | PlainMessage<PlayoutDelay> | undefined): boolean;
}
/**
 * @generated from message livekit.ParticipantPermission
 */
export declare class ParticipantPermission extends Message<ParticipantPermission> {
    /**
     * allow participant to subscribe to other tracks in the room
     *
     * @generated from field: bool can_subscribe = 1;
     */
    canSubscribe: boolean;
    /**
     * allow participant to publish new tracks to room
     *
     * @generated from field: bool can_publish = 2;
     */
    canPublish: boolean;
    /**
     * allow participant to publish data
     *
     * @generated from field: bool can_publish_data = 3;
     */
    canPublishData: boolean;
    /**
     * sources that are allowed to be published
     *
     * @generated from field: repeated livekit.TrackSource can_publish_sources = 9;
     */
    canPublishSources: TrackSource[];
    /**
     * indicates that it's hidden to others
     *
     * @generated from field: bool hidden = 7;
     */
    hidden: boolean;
    /**
     * indicates it's a recorder instance
     * deprecated: use ParticipantInfo.kind instead
     *
     * @generated from field: bool recorder = 8 [deprecated = true];
     * @deprecated
     */
    recorder: boolean;
    /**
     * indicates that participant can update own metadata and attributes
     *
     * @generated from field: bool can_update_metadata = 10;
     */
    canUpdateMetadata: boolean;
    /**
     * indicates that participant is an agent
     * deprecated: use ParticipantInfo.kind instead
     *
     * @generated from field: bool agent = 11 [deprecated = true];
     * @deprecated
     */
    agent: boolean;
    /**
     * if a participant can subscribe to metrics
     *
     * @generated from field: bool can_subscribe_metrics = 12;
     */
    canSubscribeMetrics: boolean;
    constructor(data?: PartialMessage<ParticipantPermission>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ParticipantPermission";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ParticipantPermission;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ParticipantPermission;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ParticipantPermission;
    static equals(a: ParticipantPermission | PlainMessage<ParticipantPermission> | undefined, b: ParticipantPermission | PlainMessage<ParticipantPermission> | undefined): boolean;
}
/**
 * @generated from message livekit.ParticipantInfo
 */
export declare class ParticipantInfo extends Message<ParticipantInfo> {
    /**
     * @generated from field: string sid = 1;
     */
    sid: string;
    /**
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * @generated from field: livekit.ParticipantInfo.State state = 3;
     */
    state: ParticipantInfo_State;
    /**
     * @generated from field: repeated livekit.TrackInfo tracks = 4;
     */
    tracks: TrackInfo[];
    /**
     * @generated from field: string metadata = 5;
     */
    metadata: string;
    /**
     * timestamp when participant joined room, in seconds
     *
     * @generated from field: int64 joined_at = 6;
     */
    joinedAt: bigint;
    /**
     * @generated from field: string name = 9;
     */
    name: string;
    /**
     * @generated from field: uint32 version = 10;
     */
    version: number;
    /**
     * @generated from field: livekit.ParticipantPermission permission = 11;
     */
    permission?: ParticipantPermission;
    /**
     * @generated from field: string region = 12;
     */
    region: string;
    /**
     * indicates the participant has an active publisher connection
     * and can publish to the server
     *
     * @generated from field: bool is_publisher = 13;
     */
    isPublisher: boolean;
    /**
     * @generated from field: livekit.ParticipantInfo.Kind kind = 14;
     */
    kind: ParticipantInfo_Kind;
    /**
     * @generated from field: map<string, string> attributes = 15;
     */
    attributes: {
        [key: string]: string;
    };
    /**
     * @generated from field: livekit.DisconnectReason disconnect_reason = 16;
     */
    disconnectReason: DisconnectReason;
    constructor(data?: PartialMessage<ParticipantInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ParticipantInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ParticipantInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ParticipantInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ParticipantInfo;
    static equals(a: ParticipantInfo | PlainMessage<ParticipantInfo> | undefined, b: ParticipantInfo | PlainMessage<ParticipantInfo> | undefined): boolean;
}
/**
 * @generated from enum livekit.ParticipantInfo.State
 */
export declare enum ParticipantInfo_State {
    /**
     * websocket' connected, but not offered yet
     *
     * @generated from enum value: JOINING = 0;
     */
    JOINING = 0,
    /**
     * server received client offer
     *
     * @generated from enum value: JOINED = 1;
     */
    JOINED = 1,
    /**
     * ICE connectivity established
     *
     * @generated from enum value: ACTIVE = 2;
     */
    ACTIVE = 2,
    /**
     * WS disconnected
     *
     * @generated from enum value: DISCONNECTED = 3;
     */
    DISCONNECTED = 3
}
/**
 * @generated from enum livekit.ParticipantInfo.Kind
 */
export declare enum ParticipantInfo_Kind {
    /**
     * standard participants, e.g. web clients
     *
     * @generated from enum value: STANDARD = 0;
     */
    STANDARD = 0,
    /**
     * only ingests streams
     *
     * @generated from enum value: INGRESS = 1;
     */
    INGRESS = 1,
    /**
     * only consumes streams
     *
     * @generated from enum value: EGRESS = 2;
     */
    EGRESS = 2,
    /**
     * SIP participants
     *
     * @generated from enum value: SIP = 3;
     */
    SIP = 3,
    /**
     * LiveKit agents
     *
     * @generated from enum value: AGENT = 4;
     */
    AGENT = 4
}
/**
 * @generated from message livekit.Encryption
 */
export declare class Encryption extends Message<Encryption> {
    constructor(data?: PartialMessage<Encryption>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.Encryption";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Encryption;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Encryption;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Encryption;
    static equals(a: Encryption | PlainMessage<Encryption> | undefined, b: Encryption | PlainMessage<Encryption> | undefined): boolean;
}
/**
 * @generated from enum livekit.Encryption.Type
 */
export declare enum Encryption_Type {
    /**
     * @generated from enum value: NONE = 0;
     */
    NONE = 0,
    /**
     * @generated from enum value: GCM = 1;
     */
    GCM = 1,
    /**
     * @generated from enum value: CUSTOM = 2;
     */
    CUSTOM = 2
}
/**
 * @generated from message livekit.SimulcastCodecInfo
 */
export declare class SimulcastCodecInfo extends Message<SimulcastCodecInfo> {
    /**
     * @generated from field: string mime_type = 1;
     */
    mimeType: string;
    /**
     * @generated from field: string mid = 2;
     */
    mid: string;
    /**
     * @generated from field: string cid = 3;
     */
    cid: string;
    /**
     * @generated from field: repeated livekit.VideoLayer layers = 4;
     */
    layers: VideoLayer[];
    constructor(data?: PartialMessage<SimulcastCodecInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SimulcastCodecInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SimulcastCodecInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SimulcastCodecInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SimulcastCodecInfo;
    static equals(a: SimulcastCodecInfo | PlainMessage<SimulcastCodecInfo> | undefined, b: SimulcastCodecInfo | PlainMessage<SimulcastCodecInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.TrackInfo
 */
export declare class TrackInfo extends Message<TrackInfo> {
    /**
     * @generated from field: string sid = 1;
     */
    sid: string;
    /**
     * @generated from field: livekit.TrackType type = 2;
     */
    type: TrackType;
    /**
     * @generated from field: string name = 3;
     */
    name: string;
    /**
     * @generated from field: bool muted = 4;
     */
    muted: boolean;
    /**
     * original width of video (unset for audio)
     * clients may receive a lower resolution version with simulcast
     *
     * @generated from field: uint32 width = 5;
     */
    width: number;
    /**
     * original height of video (unset for audio)
     *
     * @generated from field: uint32 height = 6;
     */
    height: number;
    /**
     * true if track is simulcasted
     *
     * @generated from field: bool simulcast = 7;
     */
    simulcast: boolean;
    /**
     * true if DTX (Discontinuous Transmission) is disabled for audio
     *
     * @generated from field: bool disable_dtx = 8;
     */
    disableDtx: boolean;
    /**
     * source of media
     *
     * @generated from field: livekit.TrackSource source = 9;
     */
    source: TrackSource;
    /**
     * @generated from field: repeated livekit.VideoLayer layers = 10;
     */
    layers: VideoLayer[];
    /**
     * mime type of codec
     *
     * @generated from field: string mime_type = 11;
     */
    mimeType: string;
    /**
     * @generated from field: string mid = 12;
     */
    mid: string;
    /**
     * @generated from field: repeated livekit.SimulcastCodecInfo codecs = 13;
     */
    codecs: SimulcastCodecInfo[];
    /**
     * @generated from field: bool stereo = 14;
     */
    stereo: boolean;
    /**
     * true if RED (Redundant Encoding) is disabled for audio
     *
     * @generated from field: bool disable_red = 15;
     */
    disableRed: boolean;
    /**
     * @generated from field: livekit.Encryption.Type encryption = 16;
     */
    encryption: Encryption_Type;
    /**
     * @generated from field: string stream = 17;
     */
    stream: string;
    /**
     * @generated from field: livekit.TimedVersion version = 18;
     */
    version?: TimedVersion;
    /**
     * @generated from field: repeated livekit.AudioTrackFeature audio_features = 19;
     */
    audioFeatures: AudioTrackFeature[];
    constructor(data?: PartialMessage<TrackInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TrackInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TrackInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TrackInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TrackInfo;
    static equals(a: TrackInfo | PlainMessage<TrackInfo> | undefined, b: TrackInfo | PlainMessage<TrackInfo> | undefined): boolean;
}
/**
 * provide information about available spatial layers
 *
 * @generated from message livekit.VideoLayer
 */
export declare class VideoLayer extends Message<VideoLayer> {
    /**
     * for tracks with a single layer, this should be HIGH
     *
     * @generated from field: livekit.VideoQuality quality = 1;
     */
    quality: VideoQuality;
    /**
     * @generated from field: uint32 width = 2;
     */
    width: number;
    /**
     * @generated from field: uint32 height = 3;
     */
    height: number;
    /**
     * target bitrate in bit per second (bps), server will measure actual
     *
     * @generated from field: uint32 bitrate = 4;
     */
    bitrate: number;
    /**
     * @generated from field: uint32 ssrc = 5;
     */
    ssrc: number;
    constructor(data?: PartialMessage<VideoLayer>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.VideoLayer";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VideoLayer;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VideoLayer;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VideoLayer;
    static equals(a: VideoLayer | PlainMessage<VideoLayer> | undefined, b: VideoLayer | PlainMessage<VideoLayer> | undefined): boolean;
}
/**
 * new DataPacket API
 *
 * @generated from message livekit.DataPacket
 */
export declare class DataPacket extends Message<DataPacket> {
    /**
     * @generated from field: livekit.DataPacket.Kind kind = 1 [deprecated = true];
     * @deprecated
     */
    kind: DataPacket_Kind;
    /**
     * participant identity of user that sent the message
     *
     * @generated from field: string participant_identity = 4;
     */
    participantIdentity: string;
    /**
     * identities of participants who will receive the message (sent to all by default)
     *
     * @generated from field: repeated string destination_identities = 5;
     */
    destinationIdentities: string[];
    /**
     * @generated from oneof livekit.DataPacket.value
     */
    value: {
        /**
         * @generated from field: livekit.UserPacket user = 2;
         */
        value: UserPacket;
        case: "user";
    } | {
        /**
         * @generated from field: livekit.ActiveSpeakerUpdate speaker = 3 [deprecated = true];
         * @deprecated
         */
        value: ActiveSpeakerUpdate;
        case: "speaker";
    } | {
        /**
         * @generated from field: livekit.SipDTMF sip_dtmf = 6;
         */
        value: SipDTMF;
        case: "sipDtmf";
    } | {
        /**
         * @generated from field: livekit.Transcription transcription = 7;
         */
        value: Transcription;
        case: "transcription";
    } | {
        /**
         * @generated from field: livekit.MetricsBatch metrics = 8;
         */
        value: MetricsBatch;
        case: "metrics";
    } | {
        /**
         * @generated from field: livekit.ChatMessage chat_message = 9;
         */
        value: ChatMessage;
        case: "chatMessage";
    } | {
        /**
         * @generated from field: livekit.RpcRequest rpc_request = 10;
         */
        value: RpcRequest;
        case: "rpcRequest";
    } | {
        /**
         * @generated from field: livekit.RpcAck rpc_ack = 11;
         */
        value: RpcAck;
        case: "rpcAck";
    } | {
        /**
         * @generated from field: livekit.RpcResponse rpc_response = 12;
         */
        value: RpcResponse;
        case: "rpcResponse";
    } | {
        /**
         * @generated from field: livekit.DataStream.Header stream_header = 13;
         */
        value: DataStream_Header;
        case: "streamHeader";
    } | {
        /**
         * @generated from field: livekit.DataStream.Chunk stream_chunk = 14;
         */
        value: DataStream_Chunk;
        case: "streamChunk";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<DataPacket>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DataPacket";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataPacket;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataPacket;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataPacket;
    static equals(a: DataPacket | PlainMessage<DataPacket> | undefined, b: DataPacket | PlainMessage<DataPacket> | undefined): boolean;
}
/**
 * @generated from enum livekit.DataPacket.Kind
 */
export declare enum DataPacket_Kind {
    /**
     * @generated from enum value: RELIABLE = 0;
     */
    RELIABLE = 0,
    /**
     * @generated from enum value: LOSSY = 1;
     */
    LOSSY = 1
}
/**
 * @generated from message livekit.ActiveSpeakerUpdate
 */
export declare class ActiveSpeakerUpdate extends Message<ActiveSpeakerUpdate> {
    /**
     * @generated from field: repeated livekit.SpeakerInfo speakers = 1;
     */
    speakers: SpeakerInfo[];
    constructor(data?: PartialMessage<ActiveSpeakerUpdate>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ActiveSpeakerUpdate";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ActiveSpeakerUpdate;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ActiveSpeakerUpdate;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ActiveSpeakerUpdate;
    static equals(a: ActiveSpeakerUpdate | PlainMessage<ActiveSpeakerUpdate> | undefined, b: ActiveSpeakerUpdate | PlainMessage<ActiveSpeakerUpdate> | undefined): boolean;
}
/**
 * @generated from message livekit.SpeakerInfo
 */
export declare class SpeakerInfo extends Message<SpeakerInfo> {
    /**
     * @generated from field: string sid = 1;
     */
    sid: string;
    /**
     * audio level, 0-1.0, 1 is loudest
     *
     * @generated from field: float level = 2;
     */
    level: number;
    /**
     * true if speaker is currently active
     *
     * @generated from field: bool active = 3;
     */
    active: boolean;
    constructor(data?: PartialMessage<SpeakerInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SpeakerInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SpeakerInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SpeakerInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SpeakerInfo;
    static equals(a: SpeakerInfo | PlainMessage<SpeakerInfo> | undefined, b: SpeakerInfo | PlainMessage<SpeakerInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.UserPacket
 */
export declare class UserPacket extends Message<UserPacket> {
    /**
     * participant ID of user that sent the message
     *
     * @generated from field: string participant_sid = 1 [deprecated = true];
     * @deprecated
     */
    participantSid: string;
    /**
     * @generated from field: string participant_identity = 5 [deprecated = true];
     * @deprecated
     */
    participantIdentity: string;
    /**
     * user defined payload
     *
     * @generated from field: bytes payload = 2;
     */
    payload: Uint8Array;
    /**
     * the ID of the participants who will receive the message (sent to all by default)
     *
     * @generated from field: repeated string destination_sids = 3 [deprecated = true];
     * @deprecated
     */
    destinationSids: string[];
    /**
     * identities of participants who will receive the message (sent to all by default)
     *
     * @generated from field: repeated string destination_identities = 6 [deprecated = true];
     * @deprecated
     */
    destinationIdentities: string[];
    /**
     * topic under which the message was published
     *
     * @generated from field: optional string topic = 4;
     */
    topic?: string;
    /**
     * Unique ID to indentify the message
     *
     * @generated from field: optional string id = 8;
     */
    id?: string;
    /**
     * start and end time allow relating the message to specific media time
     *
     * @generated from field: optional uint64 start_time = 9;
     */
    startTime?: bigint;
    /**
     * @generated from field: optional uint64 end_time = 10;
     */
    endTime?: bigint;
    constructor(data?: PartialMessage<UserPacket>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UserPacket";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserPacket;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserPacket;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserPacket;
    static equals(a: UserPacket | PlainMessage<UserPacket> | undefined, b: UserPacket | PlainMessage<UserPacket> | undefined): boolean;
}
/**
 * @generated from message livekit.SipDTMF
 */
export declare class SipDTMF extends Message<SipDTMF> {
    /**
     * @generated from field: uint32 code = 3;
     */
    code: number;
    /**
     * @generated from field: string digit = 4;
     */
    digit: string;
    constructor(data?: PartialMessage<SipDTMF>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SipDTMF";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SipDTMF;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SipDTMF;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SipDTMF;
    static equals(a: SipDTMF | PlainMessage<SipDTMF> | undefined, b: SipDTMF | PlainMessage<SipDTMF> | undefined): boolean;
}
/**
 * @generated from message livekit.Transcription
 */
export declare class Transcription extends Message<Transcription> {
    /**
     * Participant that got its speech transcribed
     *
     * @generated from field: string transcribed_participant_identity = 2;
     */
    transcribedParticipantIdentity: string;
    /**
     * @generated from field: string track_id = 3;
     */
    trackId: string;
    /**
     * @generated from field: repeated livekit.TranscriptionSegment segments = 4;
     */
    segments: TranscriptionSegment[];
    constructor(data?: PartialMessage<Transcription>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.Transcription";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transcription;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transcription;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transcription;
    static equals(a: Transcription | PlainMessage<Transcription> | undefined, b: Transcription | PlainMessage<Transcription> | undefined): boolean;
}
/**
 * @generated from message livekit.TranscriptionSegment
 */
export declare class TranscriptionSegment extends Message<TranscriptionSegment> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string text = 2;
     */
    text: string;
    /**
     * @generated from field: uint64 start_time = 3;
     */
    startTime: bigint;
    /**
     * @generated from field: uint64 end_time = 4;
     */
    endTime: bigint;
    /**
     * @generated from field: bool final = 5;
     */
    final: boolean;
    /**
     * @generated from field: string language = 6;
     */
    language: string;
    constructor(data?: PartialMessage<TranscriptionSegment>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TranscriptionSegment";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TranscriptionSegment;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TranscriptionSegment;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TranscriptionSegment;
    static equals(a: TranscriptionSegment | PlainMessage<TranscriptionSegment> | undefined, b: TranscriptionSegment | PlainMessage<TranscriptionSegment> | undefined): boolean;
}
/**
 * @generated from message livekit.ChatMessage
 */
export declare class ChatMessage extends Message<ChatMessage> {
    /**
     * uuid
     *
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: int64 timestamp = 2;
     */
    timestamp: bigint;
    /**
     * populated only if the intent is to edit/update an existing message
     *
     * @generated from field: optional int64 edit_timestamp = 3;
     */
    editTimestamp?: bigint;
    /**
     * @generated from field: string message = 4;
     */
    message: string;
    /**
     * true to remove message
     *
     * @generated from field: bool deleted = 5;
     */
    deleted: boolean;
    /**
     * true if the chat message has been generated by an agent from a participant's audio transcription
     *
     * @generated from field: bool generated = 6;
     */
    generated: boolean;
    constructor(data?: PartialMessage<ChatMessage>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ChatMessage";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChatMessage;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChatMessage;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChatMessage;
    static equals(a: ChatMessage | PlainMessage<ChatMessage> | undefined, b: ChatMessage | PlainMessage<ChatMessage> | undefined): boolean;
}
/**
 * @generated from message livekit.RpcRequest
 */
export declare class RpcRequest extends Message<RpcRequest> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string method = 2;
     */
    method: string;
    /**
     * @generated from field: string payload = 3;
     */
    payload: string;
    /**
     * @generated from field: uint32 response_timeout_ms = 4;
     */
    responseTimeoutMs: number;
    /**
     * @generated from field: uint32 version = 5;
     */
    version: number;
    constructor(data?: PartialMessage<RpcRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RpcRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RpcRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RpcRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RpcRequest;
    static equals(a: RpcRequest | PlainMessage<RpcRequest> | undefined, b: RpcRequest | PlainMessage<RpcRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.RpcAck
 */
export declare class RpcAck extends Message<RpcAck> {
    /**
     * @generated from field: string request_id = 1;
     */
    requestId: string;
    constructor(data?: PartialMessage<RpcAck>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RpcAck";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RpcAck;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RpcAck;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RpcAck;
    static equals(a: RpcAck | PlainMessage<RpcAck> | undefined, b: RpcAck | PlainMessage<RpcAck> | undefined): boolean;
}
/**
 * @generated from message livekit.RpcResponse
 */
export declare class RpcResponse extends Message<RpcResponse> {
    /**
     * @generated from field: string request_id = 1;
     */
    requestId: string;
    /**
     * @generated from oneof livekit.RpcResponse.value
     */
    value: {
        /**
         * @generated from field: string payload = 2;
         */
        value: string;
        case: "payload";
    } | {
        /**
         * @generated from field: livekit.RpcError error = 3;
         */
        value: RpcError;
        case: "error";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<RpcResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RpcResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RpcResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RpcResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RpcResponse;
    static equals(a: RpcResponse | PlainMessage<RpcResponse> | undefined, b: RpcResponse | PlainMessage<RpcResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.RpcError
 */
export declare class RpcError extends Message<RpcError> {
    /**
     * @generated from field: uint32 code = 1;
     */
    code: number;
    /**
     * @generated from field: string message = 2;
     */
    message: string;
    /**
     * @generated from field: string data = 3;
     */
    data: string;
    constructor(data?: PartialMessage<RpcError>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RpcError";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RpcError;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RpcError;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RpcError;
    static equals(a: RpcError | PlainMessage<RpcError> | undefined, b: RpcError | PlainMessage<RpcError> | undefined): boolean;
}
/**
 * @generated from message livekit.ParticipantTracks
 */
export declare class ParticipantTracks extends Message<ParticipantTracks> {
    /**
     * participant ID of participant to whom the tracks belong
     *
     * @generated from field: string participant_sid = 1;
     */
    participantSid: string;
    /**
     * @generated from field: repeated string track_sids = 2;
     */
    trackSids: string[];
    constructor(data?: PartialMessage<ParticipantTracks>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ParticipantTracks";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ParticipantTracks;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ParticipantTracks;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ParticipantTracks;
    static equals(a: ParticipantTracks | PlainMessage<ParticipantTracks> | undefined, b: ParticipantTracks | PlainMessage<ParticipantTracks> | undefined): boolean;
}
/**
 * details about the server
 *
 * @generated from message livekit.ServerInfo
 */
export declare class ServerInfo extends Message<ServerInfo> {
    /**
     * @generated from field: livekit.ServerInfo.Edition edition = 1;
     */
    edition: ServerInfo_Edition;
    /**
     * @generated from field: string version = 2;
     */
    version: string;
    /**
     * @generated from field: int32 protocol = 3;
     */
    protocol: number;
    /**
     * @generated from field: string region = 4;
     */
    region: string;
    /**
     * @generated from field: string node_id = 5;
     */
    nodeId: string;
    /**
     * additional debugging information. sent only if server is in development mode
     *
     * @generated from field: string debug_info = 6;
     */
    debugInfo: string;
    /**
     * @generated from field: int32 agent_protocol = 7;
     */
    agentProtocol: number;
    constructor(data?: PartialMessage<ServerInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ServerInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ServerInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ServerInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ServerInfo;
    static equals(a: ServerInfo | PlainMessage<ServerInfo> | undefined, b: ServerInfo | PlainMessage<ServerInfo> | undefined): boolean;
}
/**
 * @generated from enum livekit.ServerInfo.Edition
 */
export declare enum ServerInfo_Edition {
    /**
     * @generated from enum value: Standard = 0;
     */
    Standard = 0,
    /**
     * @generated from enum value: Cloud = 1;
     */
    Cloud = 1
}
/**
 * details about the client
 *
 * @generated from message livekit.ClientInfo
 */
export declare class ClientInfo extends Message<ClientInfo> {
    /**
     * @generated from field: livekit.ClientInfo.SDK sdk = 1;
     */
    sdk: ClientInfo_SDK;
    /**
     * @generated from field: string version = 2;
     */
    version: string;
    /**
     * @generated from field: int32 protocol = 3;
     */
    protocol: number;
    /**
     * @generated from field: string os = 4;
     */
    os: string;
    /**
     * @generated from field: string os_version = 5;
     */
    osVersion: string;
    /**
     * @generated from field: string device_model = 6;
     */
    deviceModel: string;
    /**
     * @generated from field: string browser = 7;
     */
    browser: string;
    /**
     * @generated from field: string browser_version = 8;
     */
    browserVersion: string;
    /**
     * @generated from field: string address = 9;
     */
    address: string;
    /**
     * wifi, wired, cellular, vpn, empty if not known
     *
     * @generated from field: string network = 10;
     */
    network: string;
    /**
     * comma separated list of additional LiveKit SDKs in use of this client, with versions
     * e.g. "components-js:1.2.3,track-processors-js:1.2.3"
     *
     * @generated from field: string other_sdks = 11;
     */
    otherSdks: string;
    constructor(data?: PartialMessage<ClientInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ClientInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClientInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClientInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClientInfo;
    static equals(a: ClientInfo | PlainMessage<ClientInfo> | undefined, b: ClientInfo | PlainMessage<ClientInfo> | undefined): boolean;
}
/**
 * @generated from enum livekit.ClientInfo.SDK
 */
export declare enum ClientInfo_SDK {
    /**
     * @generated from enum value: UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from enum value: JS = 1;
     */
    JS = 1,
    /**
     * @generated from enum value: SWIFT = 2;
     */
    SWIFT = 2,
    /**
     * @generated from enum value: ANDROID = 3;
     */
    ANDROID = 3,
    /**
     * @generated from enum value: FLUTTER = 4;
     */
    FLUTTER = 4,
    /**
     * @generated from enum value: GO = 5;
     */
    GO = 5,
    /**
     * @generated from enum value: UNITY = 6;
     */
    UNITY = 6,
    /**
     * @generated from enum value: REACT_NATIVE = 7;
     */
    REACT_NATIVE = 7,
    /**
     * @generated from enum value: RUST = 8;
     */
    RUST = 8,
    /**
     * @generated from enum value: PYTHON = 9;
     */
    PYTHON = 9,
    /**
     * @generated from enum value: CPP = 10;
     */
    CPP = 10,
    /**
     * @generated from enum value: UNITY_WEB = 11;
     */
    UNITY_WEB = 11,
    /**
     * @generated from enum value: NODE = 12;
     */
    NODE = 12
}
/**
 * server provided client configuration
 *
 * @generated from message livekit.ClientConfiguration
 */
export declare class ClientConfiguration extends Message<ClientConfiguration> {
    /**
     * @generated from field: livekit.VideoConfiguration video = 1;
     */
    video?: VideoConfiguration;
    /**
     * @generated from field: livekit.VideoConfiguration screen = 2;
     */
    screen?: VideoConfiguration;
    /**
     * @generated from field: livekit.ClientConfigSetting resume_connection = 3;
     */
    resumeConnection: ClientConfigSetting;
    /**
     * @generated from field: livekit.DisabledCodecs disabled_codecs = 4;
     */
    disabledCodecs?: DisabledCodecs;
    /**
     * @generated from field: livekit.ClientConfigSetting force_relay = 5;
     */
    forceRelay: ClientConfigSetting;
    constructor(data?: PartialMessage<ClientConfiguration>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ClientConfiguration";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ClientConfiguration;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ClientConfiguration;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ClientConfiguration;
    static equals(a: ClientConfiguration | PlainMessage<ClientConfiguration> | undefined, b: ClientConfiguration | PlainMessage<ClientConfiguration> | undefined): boolean;
}
/**
 * @generated from message livekit.VideoConfiguration
 */
export declare class VideoConfiguration extends Message<VideoConfiguration> {
    /**
     * @generated from field: livekit.ClientConfigSetting hardware_encoder = 1;
     */
    hardwareEncoder: ClientConfigSetting;
    constructor(data?: PartialMessage<VideoConfiguration>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.VideoConfiguration";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VideoConfiguration;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VideoConfiguration;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VideoConfiguration;
    static equals(a: VideoConfiguration | PlainMessage<VideoConfiguration> | undefined, b: VideoConfiguration | PlainMessage<VideoConfiguration> | undefined): boolean;
}
/**
 * @generated from message livekit.DisabledCodecs
 */
export declare class DisabledCodecs extends Message<DisabledCodecs> {
    /**
     * disabled for both publish and subscribe
     *
     * @generated from field: repeated livekit.Codec codecs = 1;
     */
    codecs: Codec[];
    /**
     * only disable for publish
     *
     * @generated from field: repeated livekit.Codec publish = 2;
     */
    publish: Codec[];
    constructor(data?: PartialMessage<DisabledCodecs>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DisabledCodecs";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DisabledCodecs;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DisabledCodecs;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DisabledCodecs;
    static equals(a: DisabledCodecs | PlainMessage<DisabledCodecs> | undefined, b: DisabledCodecs | PlainMessage<DisabledCodecs> | undefined): boolean;
}
/**
 * @generated from message livekit.RTPDrift
 */
export declare class RTPDrift extends Message<RTPDrift> {
    /**
     * @generated from field: google.protobuf.Timestamp start_time = 1;
     */
    startTime?: Timestamp;
    /**
     * @generated from field: google.protobuf.Timestamp end_time = 2;
     */
    endTime?: Timestamp;
    /**
     * @generated from field: double duration = 3;
     */
    duration: number;
    /**
     * @generated from field: uint64 start_timestamp = 4;
     */
    startTimestamp: bigint;
    /**
     * @generated from field: uint64 end_timestamp = 5;
     */
    endTimestamp: bigint;
    /**
     * @generated from field: uint64 rtp_clock_ticks = 6;
     */
    rtpClockTicks: bigint;
    /**
     * @generated from field: int64 drift_samples = 7;
     */
    driftSamples: bigint;
    /**
     * @generated from field: double drift_ms = 8;
     */
    driftMs: number;
    /**
     * @generated from field: double clock_rate = 9;
     */
    clockRate: number;
    constructor(data?: PartialMessage<RTPDrift>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RTPDrift";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RTPDrift;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RTPDrift;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RTPDrift;
    static equals(a: RTPDrift | PlainMessage<RTPDrift> | undefined, b: RTPDrift | PlainMessage<RTPDrift> | undefined): boolean;
}
/**
 * @generated from message livekit.RTPStats
 */
export declare class RTPStats extends Message<RTPStats> {
    /**
     * @generated from field: google.protobuf.Timestamp start_time = 1;
     */
    startTime?: Timestamp;
    /**
     * @generated from field: google.protobuf.Timestamp end_time = 2;
     */
    endTime?: Timestamp;
    /**
     * @generated from field: double duration = 3;
     */
    duration: number;
    /**
     * @generated from field: uint32 packets = 4;
     */
    packets: number;
    /**
     * @generated from field: double packet_rate = 5;
     */
    packetRate: number;
    /**
     * @generated from field: uint64 bytes = 6;
     */
    bytes: bigint;
    /**
     * @generated from field: uint64 header_bytes = 39;
     */
    headerBytes: bigint;
    /**
     * @generated from field: double bitrate = 7;
     */
    bitrate: number;
    /**
     * @generated from field: uint32 packets_lost = 8;
     */
    packetsLost: number;
    /**
     * @generated from field: double packet_loss_rate = 9;
     */
    packetLossRate: number;
    /**
     * @generated from field: float packet_loss_percentage = 10;
     */
    packetLossPercentage: number;
    /**
     * @generated from field: uint32 packets_duplicate = 11;
     */
    packetsDuplicate: number;
    /**
     * @generated from field: double packet_duplicate_rate = 12;
     */
    packetDuplicateRate: number;
    /**
     * @generated from field: uint64 bytes_duplicate = 13;
     */
    bytesDuplicate: bigint;
    /**
     * @generated from field: uint64 header_bytes_duplicate = 40;
     */
    headerBytesDuplicate: bigint;
    /**
     * @generated from field: double bitrate_duplicate = 14;
     */
    bitrateDuplicate: number;
    /**
     * @generated from field: uint32 packets_padding = 15;
     */
    packetsPadding: number;
    /**
     * @generated from field: double packet_padding_rate = 16;
     */
    packetPaddingRate: number;
    /**
     * @generated from field: uint64 bytes_padding = 17;
     */
    bytesPadding: bigint;
    /**
     * @generated from field: uint64 header_bytes_padding = 41;
     */
    headerBytesPadding: bigint;
    /**
     * @generated from field: double bitrate_padding = 18;
     */
    bitratePadding: number;
    /**
     * @generated from field: uint32 packets_out_of_order = 19;
     */
    packetsOutOfOrder: number;
    /**
     * @generated from field: uint32 frames = 20;
     */
    frames: number;
    /**
     * @generated from field: double frame_rate = 21;
     */
    frameRate: number;
    /**
     * @generated from field: double jitter_current = 22;
     */
    jitterCurrent: number;
    /**
     * @generated from field: double jitter_max = 23;
     */
    jitterMax: number;
    /**
     * @generated from field: map<int32, uint32> gap_histogram = 24;
     */
    gapHistogram: {
        [key: number]: number;
    };
    /**
     * @generated from field: uint32 nacks = 25;
     */
    nacks: number;
    /**
     * @generated from field: uint32 nack_acks = 37;
     */
    nackAcks: number;
    /**
     * @generated from field: uint32 nack_misses = 26;
     */
    nackMisses: number;
    /**
     * @generated from field: uint32 nack_repeated = 38;
     */
    nackRepeated: number;
    /**
     * @generated from field: uint32 plis = 27;
     */
    plis: number;
    /**
     * @generated from field: google.protobuf.Timestamp last_pli = 28;
     */
    lastPli?: Timestamp;
    /**
     * @generated from field: uint32 firs = 29;
     */
    firs: number;
    /**
     * @generated from field: google.protobuf.Timestamp last_fir = 30;
     */
    lastFir?: Timestamp;
    /**
     * @generated from field: uint32 rtt_current = 31;
     */
    rttCurrent: number;
    /**
     * @generated from field: uint32 rtt_max = 32;
     */
    rttMax: number;
    /**
     * @generated from field: uint32 key_frames = 33;
     */
    keyFrames: number;
    /**
     * @generated from field: google.protobuf.Timestamp last_key_frame = 34;
     */
    lastKeyFrame?: Timestamp;
    /**
     * @generated from field: uint32 layer_lock_plis = 35;
     */
    layerLockPlis: number;
    /**
     * @generated from field: google.protobuf.Timestamp last_layer_lock_pli = 36;
     */
    lastLayerLockPli?: Timestamp;
    /**
     * @generated from field: livekit.RTPDrift packet_drift = 44;
     */
    packetDrift?: RTPDrift;
    /**
     * @generated from field: livekit.RTPDrift ntp_report_drift = 45;
     */
    ntpReportDrift?: RTPDrift;
    /**
     * @generated from field: livekit.RTPDrift rebased_report_drift = 46;
     */
    rebasedReportDrift?: RTPDrift;
    /**
     * NEXT_ID: 48
     *
     * @generated from field: livekit.RTPDrift received_report_drift = 47;
     */
    receivedReportDrift?: RTPDrift;
    constructor(data?: PartialMessage<RTPStats>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RTPStats";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RTPStats;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RTPStats;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RTPStats;
    static equals(a: RTPStats | PlainMessage<RTPStats> | undefined, b: RTPStats | PlainMessage<RTPStats> | undefined): boolean;
}
/**
 * @generated from message livekit.RTCPSenderReportState
 */
export declare class RTCPSenderReportState extends Message<RTCPSenderReportState> {
    /**
     * @generated from field: uint32 rtp_timestamp = 1;
     */
    rtpTimestamp: number;
    /**
     * @generated from field: uint64 rtp_timestamp_ext = 2;
     */
    rtpTimestampExt: bigint;
    /**
     * @generated from field: uint64 ntp_timestamp = 3;
     */
    ntpTimestamp: bigint;
    /**
     * time at which this happened
     *
     * @generated from field: int64 at = 4;
     */
    at: bigint;
    /**
     * @generated from field: int64 at_adjusted = 5;
     */
    atAdjusted: bigint;
    /**
     * @generated from field: uint32 packets = 6;
     */
    packets: number;
    /**
     * @generated from field: uint64 octets = 7;
     */
    octets: bigint;
    constructor(data?: PartialMessage<RTCPSenderReportState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RTCPSenderReportState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RTCPSenderReportState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RTCPSenderReportState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RTCPSenderReportState;
    static equals(a: RTCPSenderReportState | PlainMessage<RTCPSenderReportState> | undefined, b: RTCPSenderReportState | PlainMessage<RTCPSenderReportState> | undefined): boolean;
}
/**
 * @generated from message livekit.RTPForwarderState
 */
export declare class RTPForwarderState extends Message<RTPForwarderState> {
    /**
     * @generated from field: bool started = 1;
     */
    started: boolean;
    /**
     * @generated from field: int32 reference_layer_spatial = 2;
     */
    referenceLayerSpatial: number;
    /**
     * @generated from field: int64 pre_start_time = 3;
     */
    preStartTime: bigint;
    /**
     * @generated from field: uint64 ext_first_timestamp = 4;
     */
    extFirstTimestamp: bigint;
    /**
     * @generated from field: uint64 dummy_start_timestamp_offset = 5;
     */
    dummyStartTimestampOffset: bigint;
    /**
     * @generated from field: livekit.RTPMungerState rtp_munger = 6;
     */
    rtpMunger?: RTPMungerState;
    /**
     * @generated from oneof livekit.RTPForwarderState.codec_munger
     */
    codecMunger: {
        /**
         * @generated from field: livekit.VP8MungerState vp8_munger = 7;
         */
        value: VP8MungerState;
        case: "vp8Munger";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.RTCPSenderReportState sender_report_state = 8;
     */
    senderReportState: RTCPSenderReportState[];
    constructor(data?: PartialMessage<RTPForwarderState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RTPForwarderState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RTPForwarderState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RTPForwarderState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RTPForwarderState;
    static equals(a: RTPForwarderState | PlainMessage<RTPForwarderState> | undefined, b: RTPForwarderState | PlainMessage<RTPForwarderState> | undefined): boolean;
}
/**
 * @generated from message livekit.RTPMungerState
 */
export declare class RTPMungerState extends Message<RTPMungerState> {
    /**
     * @generated from field: uint64 ext_last_sequence_number = 1;
     */
    extLastSequenceNumber: bigint;
    /**
     * @generated from field: uint64 ext_second_last_sequence_number = 2;
     */
    extSecondLastSequenceNumber: bigint;
    /**
     * @generated from field: uint64 ext_last_timestamp = 3;
     */
    extLastTimestamp: bigint;
    /**
     * @generated from field: uint64 ext_second_last_timestamp = 4;
     */
    extSecondLastTimestamp: bigint;
    /**
     * @generated from field: bool last_marker = 5;
     */
    lastMarker: boolean;
    /**
     * @generated from field: bool second_last_marker = 6;
     */
    secondLastMarker: boolean;
    constructor(data?: PartialMessage<RTPMungerState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RTPMungerState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RTPMungerState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RTPMungerState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RTPMungerState;
    static equals(a: RTPMungerState | PlainMessage<RTPMungerState> | undefined, b: RTPMungerState | PlainMessage<RTPMungerState> | undefined): boolean;
}
/**
 * @generated from message livekit.VP8MungerState
 */
export declare class VP8MungerState extends Message<VP8MungerState> {
    /**
     * @generated from field: int32 ext_last_picture_id = 1;
     */
    extLastPictureId: number;
    /**
     * @generated from field: bool picture_id_used = 2;
     */
    pictureIdUsed: boolean;
    /**
     * @generated from field: uint32 last_tl0_pic_idx = 3;
     */
    lastTl0PicIdx: number;
    /**
     * @generated from field: bool tl0_pic_idx_used = 4;
     */
    tl0PicIdxUsed: boolean;
    /**
     * @generated from field: bool tid_used = 5;
     */
    tidUsed: boolean;
    /**
     * @generated from field: uint32 last_key_idx = 6;
     */
    lastKeyIdx: number;
    /**
     * @generated from field: bool key_idx_used = 7;
     */
    keyIdxUsed: boolean;
    constructor(data?: PartialMessage<VP8MungerState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.VP8MungerState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): VP8MungerState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): VP8MungerState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): VP8MungerState;
    static equals(a: VP8MungerState | PlainMessage<VP8MungerState> | undefined, b: VP8MungerState | PlainMessage<VP8MungerState> | undefined): boolean;
}
/**
 * @generated from message livekit.TimedVersion
 */
export declare class TimedVersion extends Message<TimedVersion> {
    /**
     * @generated from field: int64 unix_micro = 1;
     */
    unixMicro: bigint;
    /**
     * @generated from field: int32 ticks = 2;
     */
    ticks: number;
    constructor(data?: PartialMessage<TimedVersion>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TimedVersion";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TimedVersion;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TimedVersion;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TimedVersion;
    static equals(a: TimedVersion | PlainMessage<TimedVersion> | undefined, b: TimedVersion | PlainMessage<TimedVersion> | undefined): boolean;
}
/**
 * @generated from message livekit.DataStream
 */
export declare class DataStream extends Message<DataStream> {
    constructor(data?: PartialMessage<DataStream>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DataStream";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataStream;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataStream;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataStream;
    static equals(a: DataStream | PlainMessage<DataStream> | undefined, b: DataStream | PlainMessage<DataStream> | undefined): boolean;
}
/**
 * enum for operation types (specific to TextHeader)
 *
 * @generated from enum livekit.DataStream.OperationType
 */
export declare enum DataStream_OperationType {
    /**
     * @generated from enum value: CREATE = 0;
     */
    CREATE = 0,
    /**
     * @generated from enum value: UPDATE = 1;
     */
    UPDATE = 1,
    /**
     * @generated from enum value: DELETE = 2;
     */
    DELETE = 2,
    /**
     * @generated from enum value: REACTION = 3;
     */
    REACTION = 3
}
/**
 * header properties specific to text streams
 *
 * @generated from message livekit.DataStream.TextHeader
 */
export declare class DataStream_TextHeader extends Message<DataStream_TextHeader> {
    /**
     * @generated from field: livekit.DataStream.OperationType operation_type = 1;
     */
    operationType: DataStream_OperationType;
    /**
     * Optional: Version for updates/edits
     *
     * @generated from field: int32 version = 2;
     */
    version: number;
    /**
     * Optional: Reply to specific message
     *
     * @generated from field: string reply_to_stream_id = 3;
     */
    replyToStreamId: string;
    /**
     * file attachments for text streams
     *
     * @generated from field: repeated string attached_stream_ids = 4;
     */
    attachedStreamIds: string[];
    /**
     * true if the text has been generated by an agent from a participant's audio transcription
     *
     * @generated from field: bool generated = 5;
     */
    generated: boolean;
    constructor(data?: PartialMessage<DataStream_TextHeader>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DataStream.TextHeader";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataStream_TextHeader;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataStream_TextHeader;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataStream_TextHeader;
    static equals(a: DataStream_TextHeader | PlainMessage<DataStream_TextHeader> | undefined, b: DataStream_TextHeader | PlainMessage<DataStream_TextHeader> | undefined): boolean;
}
/**
 * header properties specific to file or image streams
 *
 * @generated from message livekit.DataStream.FileHeader
 */
export declare class DataStream_FileHeader extends Message<DataStream_FileHeader> {
    /**
     * name of the file
     *
     * @generated from field: string file_name = 1;
     */
    fileName: string;
    constructor(data?: PartialMessage<DataStream_FileHeader>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DataStream.FileHeader";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataStream_FileHeader;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataStream_FileHeader;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataStream_FileHeader;
    static equals(a: DataStream_FileHeader | PlainMessage<DataStream_FileHeader> | undefined, b: DataStream_FileHeader | PlainMessage<DataStream_FileHeader> | undefined): boolean;
}
/**
 * main DataStream.Header that contains a oneof for specific headers
 *
 * @generated from message livekit.DataStream.Header
 */
export declare class DataStream_Header extends Message<DataStream_Header> {
    /**
     * unique identifier for this data stream
     *
     * @generated from field: string stream_id = 1;
     */
    streamId: string;
    /**
     * using int64 for Unix timestamp
     *
     * @generated from field: int64 timestamp = 2;
     */
    timestamp: bigint;
    /**
     * @generated from field: string topic = 3;
     */
    topic: string;
    /**
     * @generated from field: string mime_type = 4;
     */
    mimeType: string;
    /**
     * only populated for finite streams, if it's a stream of unknown size this stays empty
     *
     * @generated from field: optional uint64 total_length = 5;
     */
    totalLength?: bigint;
    /**
     * only populated for finite streams, if it's a stream of unknown size this stays empty
     *
     * @generated from field: optional uint64 total_chunks = 6;
     */
    totalChunks?: bigint;
    /**
     * defaults to NONE
     *
     * @generated from field: livekit.Encryption.Type encryption_type = 7;
     */
    encryptionType: Encryption_Type;
    /**
     * user defined extensions map that can carry additional info
     *
     * @generated from field: map<string, string> extensions = 8;
     */
    extensions: {
        [key: string]: string;
    };
    /**
     * oneof to choose between specific header types
     *
     * @generated from oneof livekit.DataStream.Header.content_header
     */
    contentHeader: {
        /**
         * @generated from field: livekit.DataStream.TextHeader text_header = 9;
         */
        value: DataStream_TextHeader;
        case: "textHeader";
    } | {
        /**
         * @generated from field: livekit.DataStream.FileHeader file_header = 10;
         */
        value: DataStream_FileHeader;
        case: "fileHeader";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<DataStream_Header>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DataStream.Header";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataStream_Header;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataStream_Header;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataStream_Header;
    static equals(a: DataStream_Header | PlainMessage<DataStream_Header> | undefined, b: DataStream_Header | PlainMessage<DataStream_Header> | undefined): boolean;
}
/**
 * @generated from message livekit.DataStream.Chunk
 */
export declare class DataStream_Chunk extends Message<DataStream_Chunk> {
    /**
     * unique identifier for this data stream to map it to the correct header
     *
     * @generated from field: string stream_id = 1;
     */
    streamId: string;
    /**
     * @generated from field: uint64 chunk_index = 2;
     */
    chunkIndex: bigint;
    /**
     * content as binary (bytes)
     *
     * @generated from field: bytes content = 3;
     */
    content: Uint8Array;
    /**
     * true only if this is the last chunk of this stream - can also be sent with empty content
     *
     * @generated from field: bool complete = 4;
     */
    complete: boolean;
    /**
     * a version indicating that this chunk_index has been retroactively modified and the original one needs to be replaced
     *
     * @generated from field: int32 version = 5;
     */
    version: number;
    /**
     * optional, initialization vector for AES-GCM encryption
     *
     * @generated from field: optional bytes iv = 6;
     */
    iv?: Uint8Array;
    constructor(data?: PartialMessage<DataStream_Chunk>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DataStream.Chunk";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataStream_Chunk;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataStream_Chunk;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataStream_Chunk;
    static equals(a: DataStream_Chunk | PlainMessage<DataStream_Chunk> | undefined, b: DataStream_Chunk | PlainMessage<DataStream_Chunk> | undefined): boolean;
}
//# sourceMappingURL=livekit_models_pb.d.ts.map