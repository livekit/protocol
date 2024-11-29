import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AudioCodec, TrackInfo, TrackSource, VideoCodec, VideoLayer } from "./livekit_models_pb.js";
/**
 * @generated from enum livekit.IngressInput
 */
export declare enum IngressInput {
    /**
     * @generated from enum value: RTMP_INPUT = 0;
     */
    RTMP_INPUT = 0,
    /**
     * @generated from enum value: WHIP_INPUT = 1;
     */
    WHIP_INPUT = 1,
    /**
     * Pull from the provided URL. Only HTTP url are supported, serving either a single media file or a HLS stream
     *
     * @generated from enum value: URL_INPUT = 2;
     */
    URL_INPUT = 2
}
/**
 * @generated from enum livekit.IngressAudioEncodingPreset
 */
export declare enum IngressAudioEncodingPreset {
    /**
     * OPUS, 2 channels, 96kbps
     *
     * @generated from enum value: OPUS_STEREO_96KBPS = 0;
     */
    OPUS_STEREO_96KBPS = 0,
    /**
     * OPUS, 1 channel, 64kbps
     *
     * @generated from enum value: OPUS_MONO_64KBS = 1;
     */
    OPUS_MONO_64KBS = 1
}
/**
 * @generated from enum livekit.IngressVideoEncodingPreset
 */
export declare enum IngressVideoEncodingPreset {
    /**
     * 1280x720,  30fps, 1900kbps main layer, 3 layers total
     *
     * @generated from enum value: H264_720P_30FPS_3_LAYERS = 0;
     */
    H264_720P_30FPS_3_LAYERS = 0,
    /**
     * 1980x1080, 30fps, 3500kbps main layer, 3 layers total
     *
     * @generated from enum value: H264_1080P_30FPS_3_LAYERS = 1;
     */
    H264_1080P_30FPS_3_LAYERS = 1,
    /**
     *  960x540,  25fps, 1000kbps  main layer, 2 layers total
     *
     * @generated from enum value: H264_540P_25FPS_2_LAYERS = 2;
     */
    H264_540P_25FPS_2_LAYERS = 2,
    /**
     * 1280x720,  30fps, 1900kbps, no simulcast
     *
     * @generated from enum value: H264_720P_30FPS_1_LAYER = 3;
     */
    H264_720P_30FPS_1_LAYER = 3,
    /**
     * 1980x1080, 30fps, 3500kbps, no simulcast
     *
     * @generated from enum value: H264_1080P_30FPS_1_LAYER = 4;
     */
    H264_1080P_30FPS_1_LAYER = 4,
    /**
     * 1280x720,  30fps, 2500kbps main layer, 3 layers total, higher bitrate for high motion, harder to encode content
     *
     * @generated from enum value: H264_720P_30FPS_3_LAYERS_HIGH_MOTION = 5;
     */
    H264_720P_30FPS_3_LAYERS_HIGH_MOTION = 5,
    /**
     * 1980x1080, 30fps, 4500kbps main layer, 3 layers total, higher bitrate for high motion, harder to encode content
     *
     * @generated from enum value: H264_1080P_30FPS_3_LAYERS_HIGH_MOTION = 6;
     */
    H264_1080P_30FPS_3_LAYERS_HIGH_MOTION = 6,
    /**
     *  960x540,  25fps, 1300kbps  main layer, 2 layers total, higher bitrate for high motion, harder to encode content
     *
     * @generated from enum value: H264_540P_25FPS_2_LAYERS_HIGH_MOTION = 7;
     */
    H264_540P_25FPS_2_LAYERS_HIGH_MOTION = 7,
    /**
     * 1280x720,  30fps, 2500kbps, no simulcast, higher bitrate for high motion, harder to encode content
     *
     * @generated from enum value: H264_720P_30FPS_1_LAYER_HIGH_MOTION = 8;
     */
    H264_720P_30FPS_1_LAYER_HIGH_MOTION = 8,
    /**
     * 1980x1080, 30fps, 4500kbps, no simulcast, higher bitrate for high motion, harder to encode content
     *
     * @generated from enum value: H264_1080P_30FPS_1_LAYER_HIGH_MOTION = 9;
     */
    H264_1080P_30FPS_1_LAYER_HIGH_MOTION = 9
}
/**
 * @generated from message livekit.CreateIngressRequest
 */
export declare class CreateIngressRequest extends Message<CreateIngressRequest> {
    /**
     * @generated from field: livekit.IngressInput input_type = 1;
     */
    inputType: IngressInput;
    /**
     * Where to pull media from, only for URL input type
     *
     * @generated from field: string url = 9;
     */
    url: string;
    /**
     * User provided identifier for the ingress
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * room to publish to
     *
     * @generated from field: string room_name = 3;
     */
    roomName: string;
    /**
     * publish as participant
     *
     * @generated from field: string participant_identity = 4;
     */
    participantIdentity: string;
    /**
     * name of publishing participant (used for display only)
     *
     * @generated from field: string participant_name = 5;
     */
    participantName: string;
    /**
     * metadata associated with the publishing participant
     *
     * @generated from field: string participant_metadata = 10;
     */
    participantMetadata: string;
    /**
     * [depreacted ] whether to pass through the incoming media without transcoding, only compatible with some input types. Use `enable_transcoding` instead.
     *
     * @generated from field: bool bypass_transcoding = 8 [deprecated = true];
     * @deprecated
     */
    bypassTranscoding: boolean;
    /**
     * Whether to transcode the ingested media. Only WHIP supports disabling transcoding currently. WHIP will default to transcoding disabled. Replaces `bypass_transcoding.
     *
     * @generated from field: optional bool enable_transcoding = 11;
     */
    enableTranscoding?: boolean;
    /**
     * @generated from field: livekit.IngressAudioOptions audio = 6;
     */
    audio?: IngressAudioOptions;
    /**
     * @generated from field: livekit.IngressVideoOptions video = 7;
     */
    video?: IngressVideoOptions;
    constructor(data?: PartialMessage<CreateIngressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateIngressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateIngressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateIngressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateIngressRequest;
    static equals(a: CreateIngressRequest | PlainMessage<CreateIngressRequest> | undefined, b: CreateIngressRequest | PlainMessage<CreateIngressRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.IngressAudioOptions
 */
export declare class IngressAudioOptions extends Message<IngressAudioOptions> {
    /**
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * @generated from field: livekit.TrackSource source = 2;
     */
    source: TrackSource;
    /**
     * @generated from oneof livekit.IngressAudioOptions.encoding_options
     */
    encodingOptions: {
        /**
         * @generated from field: livekit.IngressAudioEncodingPreset preset = 3;
         */
        value: IngressAudioEncodingPreset;
        case: "preset";
    } | {
        /**
         * @generated from field: livekit.IngressAudioEncodingOptions options = 4;
         */
        value: IngressAudioEncodingOptions;
        case: "options";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<IngressAudioOptions>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.IngressAudioOptions";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngressAudioOptions;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngressAudioOptions;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngressAudioOptions;
    static equals(a: IngressAudioOptions | PlainMessage<IngressAudioOptions> | undefined, b: IngressAudioOptions | PlainMessage<IngressAudioOptions> | undefined): boolean;
}
/**
 * @generated from message livekit.IngressVideoOptions
 */
export declare class IngressVideoOptions extends Message<IngressVideoOptions> {
    /**
     * @generated from field: string name = 1;
     */
    name: string;
    /**
     * @generated from field: livekit.TrackSource source = 2;
     */
    source: TrackSource;
    /**
     * @generated from oneof livekit.IngressVideoOptions.encoding_options
     */
    encodingOptions: {
        /**
         * @generated from field: livekit.IngressVideoEncodingPreset preset = 3;
         */
        value: IngressVideoEncodingPreset;
        case: "preset";
    } | {
        /**
         * @generated from field: livekit.IngressVideoEncodingOptions options = 4;
         */
        value: IngressVideoEncodingOptions;
        case: "options";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<IngressVideoOptions>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.IngressVideoOptions";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngressVideoOptions;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngressVideoOptions;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngressVideoOptions;
    static equals(a: IngressVideoOptions | PlainMessage<IngressVideoOptions> | undefined, b: IngressVideoOptions | PlainMessage<IngressVideoOptions> | undefined): boolean;
}
/**
 * @generated from message livekit.IngressAudioEncodingOptions
 */
export declare class IngressAudioEncodingOptions extends Message<IngressAudioEncodingOptions> {
    /**
     * desired audio codec to publish to room
     *
     * @generated from field: livekit.AudioCodec audio_codec = 1;
     */
    audioCodec: AudioCodec;
    /**
     * @generated from field: uint32 bitrate = 2;
     */
    bitrate: number;
    /**
     * @generated from field: bool disable_dtx = 3;
     */
    disableDtx: boolean;
    /**
     * @generated from field: uint32 channels = 4;
     */
    channels: number;
    constructor(data?: PartialMessage<IngressAudioEncodingOptions>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.IngressAudioEncodingOptions";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngressAudioEncodingOptions;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngressAudioEncodingOptions;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngressAudioEncodingOptions;
    static equals(a: IngressAudioEncodingOptions | PlainMessage<IngressAudioEncodingOptions> | undefined, b: IngressAudioEncodingOptions | PlainMessage<IngressAudioEncodingOptions> | undefined): boolean;
}
/**
 * @generated from message livekit.IngressVideoEncodingOptions
 */
export declare class IngressVideoEncodingOptions extends Message<IngressVideoEncodingOptions> {
    /**
     * desired codec to publish to room
     *
     * @generated from field: livekit.VideoCodec video_codec = 1;
     */
    videoCodec: VideoCodec;
    /**
     * @generated from field: double frame_rate = 2;
     */
    frameRate: number;
    /**
     * simulcast layers to publish, when empty, should usually be set to layers at 1/2 and 1/4 of the dimensions
     *
     * @generated from field: repeated livekit.VideoLayer layers = 3;
     */
    layers: VideoLayer[];
    constructor(data?: PartialMessage<IngressVideoEncodingOptions>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.IngressVideoEncodingOptions";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngressVideoEncodingOptions;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngressVideoEncodingOptions;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngressVideoEncodingOptions;
    static equals(a: IngressVideoEncodingOptions | PlainMessage<IngressVideoEncodingOptions> | undefined, b: IngressVideoEncodingOptions | PlainMessage<IngressVideoEncodingOptions> | undefined): boolean;
}
/**
 * @generated from message livekit.IngressInfo
 */
export declare class IngressInfo extends Message<IngressInfo> {
    /**
     * @generated from field: string ingress_id = 1;
     */
    ingressId: string;
    /**
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * @generated from field: string stream_key = 3;
     */
    streamKey: string;
    /**
     * URL to point the encoder to for push (RTMP, WHIP), or location to pull media from for pull (URL)
     *
     * @generated from field: string url = 4;
     */
    url: string;
    /**
     * for RTMP input, it'll be a rtmp:// URL
     * for FILE input, it'll be a http:// URL
     * for SRT input, it'll be a srt:// URL
     *
     * @generated from field: livekit.IngressInput input_type = 5;
     */
    inputType: IngressInput;
    /**
     * @generated from field: bool bypass_transcoding = 13 [deprecated = true];
     * @deprecated
     */
    bypassTranscoding: boolean;
    /**
     * @generated from field: optional bool enable_transcoding = 15;
     */
    enableTranscoding?: boolean;
    /**
     * @generated from field: livekit.IngressAudioOptions audio = 6;
     */
    audio?: IngressAudioOptions;
    /**
     * @generated from field: livekit.IngressVideoOptions video = 7;
     */
    video?: IngressVideoOptions;
    /**
     * @generated from field: string room_name = 8;
     */
    roomName: string;
    /**
     * @generated from field: string participant_identity = 9;
     */
    participantIdentity: string;
    /**
     * @generated from field: string participant_name = 10;
     */
    participantName: string;
    /**
     * @generated from field: string participant_metadata = 14;
     */
    participantMetadata: string;
    /**
     * @generated from field: bool reusable = 11;
     */
    reusable: boolean;
    /**
     * Description of error/stream non compliance and debug info for publisher otherwise (received bitrate, resolution, bandwidth)
     *
     * @generated from field: livekit.IngressState state = 12;
     */
    state?: IngressState;
    constructor(data?: PartialMessage<IngressInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.IngressInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngressInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngressInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngressInfo;
    static equals(a: IngressInfo | PlainMessage<IngressInfo> | undefined, b: IngressInfo | PlainMessage<IngressInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.IngressState
 */
export declare class IngressState extends Message<IngressState> {
    /**
     * @generated from field: livekit.IngressState.Status status = 1;
     */
    status: IngressState_Status;
    /**
     * Error/non compliance description if any
     *
     * @generated from field: string error = 2;
     */
    error: string;
    /**
     * @generated from field: livekit.InputVideoState video = 3;
     */
    video?: InputVideoState;
    /**
     * @generated from field: livekit.InputAudioState audio = 4;
     */
    audio?: InputAudioState;
    /**
     * ID of the current/previous room published to
     *
     * @generated from field: string room_id = 5;
     */
    roomId: string;
    /**
     * @generated from field: int64 started_at = 7;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 8;
     */
    endedAt: bigint;
    /**
     * @generated from field: int64 updated_at = 10;
     */
    updatedAt: bigint;
    /**
     * @generated from field: string resource_id = 9;
     */
    resourceId: string;
    /**
     * @generated from field: repeated livekit.TrackInfo tracks = 6;
     */
    tracks: TrackInfo[];
    constructor(data?: PartialMessage<IngressState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.IngressState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngressState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngressState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngressState;
    static equals(a: IngressState | PlainMessage<IngressState> | undefined, b: IngressState | PlainMessage<IngressState> | undefined): boolean;
}
/**
 * @generated from enum livekit.IngressState.Status
 */
export declare enum IngressState_Status {
    /**
     * @generated from enum value: ENDPOINT_INACTIVE = 0;
     */
    ENDPOINT_INACTIVE = 0,
    /**
     * @generated from enum value: ENDPOINT_BUFFERING = 1;
     */
    ENDPOINT_BUFFERING = 1,
    /**
     * @generated from enum value: ENDPOINT_PUBLISHING = 2;
     */
    ENDPOINT_PUBLISHING = 2,
    /**
     * @generated from enum value: ENDPOINT_ERROR = 3;
     */
    ENDPOINT_ERROR = 3,
    /**
     * @generated from enum value: ENDPOINT_COMPLETE = 4;
     */
    ENDPOINT_COMPLETE = 4
}
/**
 * @generated from message livekit.InputVideoState
 */
export declare class InputVideoState extends Message<InputVideoState> {
    /**
     * @generated from field: string mime_type = 1;
     */
    mimeType: string;
    /**
     * @generated from field: uint32 average_bitrate = 2;
     */
    averageBitrate: number;
    /**
     * @generated from field: uint32 width = 3;
     */
    width: number;
    /**
     * @generated from field: uint32 height = 4;
     */
    height: number;
    /**
     * @generated from field: double framerate = 5;
     */
    framerate: number;
    constructor(data?: PartialMessage<InputVideoState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.InputVideoState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InputVideoState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InputVideoState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InputVideoState;
    static equals(a: InputVideoState | PlainMessage<InputVideoState> | undefined, b: InputVideoState | PlainMessage<InputVideoState> | undefined): boolean;
}
/**
 * @generated from message livekit.InputAudioState
 */
export declare class InputAudioState extends Message<InputAudioState> {
    /**
     * @generated from field: string mime_type = 1;
     */
    mimeType: string;
    /**
     * @generated from field: uint32 average_bitrate = 2;
     */
    averageBitrate: number;
    /**
     * @generated from field: uint32 channels = 3;
     */
    channels: number;
    /**
     * @generated from field: uint32 sample_rate = 4;
     */
    sampleRate: number;
    constructor(data?: PartialMessage<InputAudioState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.InputAudioState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InputAudioState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InputAudioState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InputAudioState;
    static equals(a: InputAudioState | PlainMessage<InputAudioState> | undefined, b: InputAudioState | PlainMessage<InputAudioState> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateIngressRequest
 */
export declare class UpdateIngressRequest extends Message<UpdateIngressRequest> {
    /**
     * @generated from field: string ingress_id = 1;
     */
    ingressId: string;
    /**
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * @generated from field: string room_name = 3;
     */
    roomName: string;
    /**
     * @generated from field: string participant_identity = 4;
     */
    participantIdentity: string;
    /**
     * @generated from field: string participant_name = 5;
     */
    participantName: string;
    /**
     * @generated from field: string participant_metadata = 9;
     */
    participantMetadata: string;
    /**
     * @generated from field: optional bool bypass_transcoding = 8 [deprecated = true];
     * @deprecated
     */
    bypassTranscoding?: boolean;
    /**
     * @generated from field: optional bool enable_transcoding = 10;
     */
    enableTranscoding?: boolean;
    /**
     * @generated from field: livekit.IngressAudioOptions audio = 6;
     */
    audio?: IngressAudioOptions;
    /**
     * @generated from field: livekit.IngressVideoOptions video = 7;
     */
    video?: IngressVideoOptions;
    constructor(data?: PartialMessage<UpdateIngressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateIngressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateIngressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateIngressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateIngressRequest;
    static equals(a: UpdateIngressRequest | PlainMessage<UpdateIngressRequest> | undefined, b: UpdateIngressRequest | PlainMessage<UpdateIngressRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListIngressRequest
 */
export declare class ListIngressRequest extends Message<ListIngressRequest> {
    /**
     * when blank, lists all ingress endpoints
     *
     * (optional, filter by room name)
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * (optional, filter by ingress ID)
     *
     * @generated from field: string ingress_id = 2;
     */
    ingressId: string;
    constructor(data?: PartialMessage<ListIngressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListIngressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListIngressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListIngressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListIngressRequest;
    static equals(a: ListIngressRequest | PlainMessage<ListIngressRequest> | undefined, b: ListIngressRequest | PlainMessage<ListIngressRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListIngressResponse
 */
export declare class ListIngressResponse extends Message<ListIngressResponse> {
    /**
     * @generated from field: repeated livekit.IngressInfo items = 1;
     */
    items: IngressInfo[];
    constructor(data?: PartialMessage<ListIngressResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListIngressResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListIngressResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListIngressResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListIngressResponse;
    static equals(a: ListIngressResponse | PlainMessage<ListIngressResponse> | undefined, b: ListIngressResponse | PlainMessage<ListIngressResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.DeleteIngressRequest
 */
export declare class DeleteIngressRequest extends Message<DeleteIngressRequest> {
    /**
     * @generated from field: string ingress_id = 1;
     */
    ingressId: string;
    constructor(data?: PartialMessage<DeleteIngressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DeleteIngressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteIngressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteIngressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteIngressRequest;
    static equals(a: DeleteIngressRequest | PlainMessage<DeleteIngressRequest> | undefined, b: DeleteIngressRequest | PlainMessage<DeleteIngressRequest> | undefined): boolean;
}
//# sourceMappingURL=livekit_ingress_pb.d.ts.map