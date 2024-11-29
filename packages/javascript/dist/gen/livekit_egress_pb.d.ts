import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AudioCodec, ImageCodec, VideoCodec } from "./livekit_models_pb.js";
/**
 * @generated from enum livekit.EncodedFileType
 */
export declare enum EncodedFileType {
    /**
     * file type chosen based on codecs
     *
     * @generated from enum value: DEFAULT_FILETYPE = 0;
     */
    DEFAULT_FILETYPE = 0,
    /**
     * @generated from enum value: MP4 = 1;
     */
    MP4 = 1,
    /**
     * @generated from enum value: OGG = 2;
     */
    OGG = 2
}
/**
 * @generated from enum livekit.SegmentedFileProtocol
 */
export declare enum SegmentedFileProtocol {
    /**
     * @generated from enum value: DEFAULT_SEGMENTED_FILE_PROTOCOL = 0;
     */
    DEFAULT_SEGMENTED_FILE_PROTOCOL = 0,
    /**
     * @generated from enum value: HLS_PROTOCOL = 1;
     */
    HLS_PROTOCOL = 1
}
/**
 * @generated from enum livekit.SegmentedFileSuffix
 */
export declare enum SegmentedFileSuffix {
    /**
     * @generated from enum value: INDEX = 0;
     */
    INDEX = 0,
    /**
     * @generated from enum value: TIMESTAMP = 1;
     */
    TIMESTAMP = 1
}
/**
 * @generated from enum livekit.ImageFileSuffix
 */
export declare enum ImageFileSuffix {
    /**
     * @generated from enum value: IMAGE_SUFFIX_INDEX = 0;
     */
    IMAGE_SUFFIX_INDEX = 0,
    /**
     * @generated from enum value: IMAGE_SUFFIX_TIMESTAMP = 1;
     */
    IMAGE_SUFFIX_TIMESTAMP = 1
}
/**
 * @generated from enum livekit.StreamProtocol
 */
export declare enum StreamProtocol {
    /**
     * protocol chosen based on urls
     *
     * @generated from enum value: DEFAULT_PROTOCOL = 0;
     */
    DEFAULT_PROTOCOL = 0,
    /**
     * @generated from enum value: RTMP = 1;
     */
    RTMP = 1,
    /**
     * @generated from enum value: SRT = 2;
     */
    SRT = 2
}
/**
 * @generated from enum livekit.EncodingOptionsPreset
 */
export declare enum EncodingOptionsPreset {
    /**
     *  1280x720, 30fps, 3000kpbs, H.264_MAIN / OPUS
     *
     * @generated from enum value: H264_720P_30 = 0;
     */
    H264_720P_30 = 0,
    /**
     *  1280x720, 60fps, 4500kbps, H.264_MAIN / OPUS
     *
     * @generated from enum value: H264_720P_60 = 1;
     */
    H264_720P_60 = 1,
    /**
     * 1920x1080, 30fps, 4500kbps, H.264_MAIN / OPUS
     *
     * @generated from enum value: H264_1080P_30 = 2;
     */
    H264_1080P_30 = 2,
    /**
     * 1920x1080, 60fps, 6000kbps, H.264_MAIN / OPUS
     *
     * @generated from enum value: H264_1080P_60 = 3;
     */
    H264_1080P_60 = 3,
    /**
     *  720x1280, 30fps, 3000kpbs, H.264_MAIN / OPUS
     *
     * @generated from enum value: PORTRAIT_H264_720P_30 = 4;
     */
    PORTRAIT_H264_720P_30 = 4,
    /**
     *  720x1280, 60fps, 4500kbps, H.264_MAIN / OPUS
     *
     * @generated from enum value: PORTRAIT_H264_720P_60 = 5;
     */
    PORTRAIT_H264_720P_60 = 5,
    /**
     * 1080x1920, 30fps, 4500kbps, H.264_MAIN / OPUS
     *
     * @generated from enum value: PORTRAIT_H264_1080P_30 = 6;
     */
    PORTRAIT_H264_1080P_30 = 6,
    /**
     * 1080x1920, 60fps, 6000kbps, H.264_MAIN / OPUS
     *
     * @generated from enum value: PORTRAIT_H264_1080P_60 = 7;
     */
    PORTRAIT_H264_1080P_60 = 7
}
/**
 * @generated from enum livekit.EgressStatus
 */
export declare enum EgressStatus {
    /**
     * @generated from enum value: EGRESS_STARTING = 0;
     */
    EGRESS_STARTING = 0,
    /**
     * @generated from enum value: EGRESS_ACTIVE = 1;
     */
    EGRESS_ACTIVE = 1,
    /**
     * @generated from enum value: EGRESS_ENDING = 2;
     */
    EGRESS_ENDING = 2,
    /**
     * @generated from enum value: EGRESS_COMPLETE = 3;
     */
    EGRESS_COMPLETE = 3,
    /**
     * @generated from enum value: EGRESS_FAILED = 4;
     */
    EGRESS_FAILED = 4,
    /**
     * @generated from enum value: EGRESS_ABORTED = 5;
     */
    EGRESS_ABORTED = 5,
    /**
     * @generated from enum value: EGRESS_LIMIT_REACHED = 6;
     */
    EGRESS_LIMIT_REACHED = 6
}
/**
 * composite using a web browser
 *
 * @generated from message livekit.RoomCompositeEgressRequest
 */
export declare class RoomCompositeEgressRequest extends Message<RoomCompositeEgressRequest> {
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * (optional)
     *
     * @generated from field: string layout = 2;
     */
    layout: string;
    /**
     * (default false)
     *
     * @generated from field: bool audio_only = 3;
     */
    audioOnly: boolean;
    /**
     * (default false)
     *
     * @generated from field: bool video_only = 4;
     */
    videoOnly: boolean;
    /**
     * template base url (default https://recorder.livekit.io)
     *
     * @generated from field: string custom_base_url = 5;
     */
    customBaseUrl: string;
    /**
     * deprecated (use _output fields)
     *
     * @generated from oneof livekit.RoomCompositeEgressRequest.output
     */
    output: {
        /**
         * @generated from field: livekit.EncodedFileOutput file = 6 [deprecated = true];
         * @deprecated
         */
        value: EncodedFileOutput;
        case: "file";
    } | {
        /**
         * @generated from field: livekit.StreamOutput stream = 7 [deprecated = true];
         * @deprecated
         */
        value: StreamOutput;
        case: "stream";
    } | {
        /**
         * @generated from field: livekit.SegmentedFileOutput segments = 10 [deprecated = true];
         * @deprecated
         */
        value: SegmentedFileOutput;
        case: "segments";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from oneof livekit.RoomCompositeEgressRequest.options
     */
    options: {
        /**
         * (default H264_720P_30)
         *
         * @generated from field: livekit.EncodingOptionsPreset preset = 8;
         */
        value: EncodingOptionsPreset;
        case: "preset";
    } | {
        /**
         * (optional)
         *
         * @generated from field: livekit.EncodingOptions advanced = 9;
         */
        value: EncodingOptions;
        case: "advanced";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 11;
     */
    fileOutputs: EncodedFileOutput[];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 12;
     */
    streamOutputs: StreamOutput[];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 13;
     */
    segmentOutputs: SegmentedFileOutput[];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 14;
     */
    imageOutputs: ImageOutput[];
    constructor(data?: PartialMessage<RoomCompositeEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomCompositeEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomCompositeEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomCompositeEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomCompositeEgressRequest;
    static equals(a: RoomCompositeEgressRequest | PlainMessage<RoomCompositeEgressRequest> | undefined, b: RoomCompositeEgressRequest | PlainMessage<RoomCompositeEgressRequest> | undefined): boolean;
}
/**
 * record any website
 *
 * @generated from message livekit.WebEgressRequest
 */
export declare class WebEgressRequest extends Message<WebEgressRequest> {
    /**
     * @generated from field: string url = 1;
     */
    url: string;
    /**
     * @generated from field: bool audio_only = 2;
     */
    audioOnly: boolean;
    /**
     * @generated from field: bool video_only = 3;
     */
    videoOnly: boolean;
    /**
     * @generated from field: bool await_start_signal = 12;
     */
    awaitStartSignal: boolean;
    /**
     * deprecated (use _output fields)
     *
     * @generated from oneof livekit.WebEgressRequest.output
     */
    output: {
        /**
         * @generated from field: livekit.EncodedFileOutput file = 4 [deprecated = true];
         * @deprecated
         */
        value: EncodedFileOutput;
        case: "file";
    } | {
        /**
         * @generated from field: livekit.StreamOutput stream = 5 [deprecated = true];
         * @deprecated
         */
        value: StreamOutput;
        case: "stream";
    } | {
        /**
         * @generated from field: livekit.SegmentedFileOutput segments = 6 [deprecated = true];
         * @deprecated
         */
        value: SegmentedFileOutput;
        case: "segments";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from oneof livekit.WebEgressRequest.options
     */
    options: {
        /**
         * @generated from field: livekit.EncodingOptionsPreset preset = 7;
         */
        value: EncodingOptionsPreset;
        case: "preset";
    } | {
        /**
         * @generated from field: livekit.EncodingOptions advanced = 8;
         */
        value: EncodingOptions;
        case: "advanced";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 9;
     */
    fileOutputs: EncodedFileOutput[];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 10;
     */
    streamOutputs: StreamOutput[];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 11;
     */
    segmentOutputs: SegmentedFileOutput[];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 13;
     */
    imageOutputs: ImageOutput[];
    constructor(data?: PartialMessage<WebEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.WebEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WebEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WebEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WebEgressRequest;
    static equals(a: WebEgressRequest | PlainMessage<WebEgressRequest> | undefined, b: WebEgressRequest | PlainMessage<WebEgressRequest> | undefined): boolean;
}
/**
 * record audio and video from a single participant
 *
 * @generated from message livekit.ParticipantEgressRequest
 */
export declare class ParticipantEgressRequest extends Message<ParticipantEgressRequest> {
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * required
     *
     * @generated from field: string identity = 2;
     */
    identity: string;
    /**
     * (default false)
     *
     * @generated from field: bool screen_share = 3;
     */
    screenShare: boolean;
    /**
     * @generated from oneof livekit.ParticipantEgressRequest.options
     */
    options: {
        /**
         * (default H264_720P_30)
         *
         * @generated from field: livekit.EncodingOptionsPreset preset = 4;
         */
        value: EncodingOptionsPreset;
        case: "preset";
    } | {
        /**
         * (optional)
         *
         * @generated from field: livekit.EncodingOptions advanced = 5;
         */
        value: EncodingOptions;
        case: "advanced";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 6;
     */
    fileOutputs: EncodedFileOutput[];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 7;
     */
    streamOutputs: StreamOutput[];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 8;
     */
    segmentOutputs: SegmentedFileOutput[];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 9;
     */
    imageOutputs: ImageOutput[];
    constructor(data?: PartialMessage<ParticipantEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ParticipantEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ParticipantEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ParticipantEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ParticipantEgressRequest;
    static equals(a: ParticipantEgressRequest | PlainMessage<ParticipantEgressRequest> | undefined, b: ParticipantEgressRequest | PlainMessage<ParticipantEgressRequest> | undefined): boolean;
}
/**
 * containerize up to one audio and one video track
 *
 * @generated from message livekit.TrackCompositeEgressRequest
 */
export declare class TrackCompositeEgressRequest extends Message<TrackCompositeEgressRequest> {
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * (optional)
     *
     * @generated from field: string audio_track_id = 2;
     */
    audioTrackId: string;
    /**
     * (optional)
     *
     * @generated from field: string video_track_id = 3;
     */
    videoTrackId: string;
    /**
     * deprecated (use _output fields)
     *
     * @generated from oneof livekit.TrackCompositeEgressRequest.output
     */
    output: {
        /**
         * @generated from field: livekit.EncodedFileOutput file = 4 [deprecated = true];
         * @deprecated
         */
        value: EncodedFileOutput;
        case: "file";
    } | {
        /**
         * @generated from field: livekit.StreamOutput stream = 5 [deprecated = true];
         * @deprecated
         */
        value: StreamOutput;
        case: "stream";
    } | {
        /**
         * @generated from field: livekit.SegmentedFileOutput segments = 8 [deprecated = true];
         * @deprecated
         */
        value: SegmentedFileOutput;
        case: "segments";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from oneof livekit.TrackCompositeEgressRequest.options
     */
    options: {
        /**
         * (default H264_720P_30)
         *
         * @generated from field: livekit.EncodingOptionsPreset preset = 6;
         */
        value: EncodingOptionsPreset;
        case: "preset";
    } | {
        /**
         * (optional)
         *
         * @generated from field: livekit.EncodingOptions advanced = 7;
         */
        value: EncodingOptions;
        case: "advanced";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 11;
     */
    fileOutputs: EncodedFileOutput[];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 12;
     */
    streamOutputs: StreamOutput[];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 13;
     */
    segmentOutputs: SegmentedFileOutput[];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 14;
     */
    imageOutputs: ImageOutput[];
    constructor(data?: PartialMessage<TrackCompositeEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TrackCompositeEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TrackCompositeEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TrackCompositeEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TrackCompositeEgressRequest;
    static equals(a: TrackCompositeEgressRequest | PlainMessage<TrackCompositeEgressRequest> | undefined, b: TrackCompositeEgressRequest | PlainMessage<TrackCompositeEgressRequest> | undefined): boolean;
}
/**
 * record tracks individually, without transcoding
 *
 * @generated from message livekit.TrackEgressRequest
 */
export declare class TrackEgressRequest extends Message<TrackEgressRequest> {
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * required
     *
     * @generated from field: string track_id = 2;
     */
    trackId: string;
    /**
     * required
     *
     * @generated from oneof livekit.TrackEgressRequest.output
     */
    output: {
        /**
         * @generated from field: livekit.DirectFileOutput file = 3;
         */
        value: DirectFileOutput;
        case: "file";
    } | {
        /**
         * @generated from field: string websocket_url = 4;
         */
        value: string;
        case: "websocketUrl";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<TrackEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TrackEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TrackEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TrackEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TrackEgressRequest;
    static equals(a: TrackEgressRequest | PlainMessage<TrackEgressRequest> | undefined, b: TrackEgressRequest | PlainMessage<TrackEgressRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.EncodedFileOutput
 */
export declare class EncodedFileOutput extends Message<EncodedFileOutput> {
    /**
     * (optional)
     *
     * @generated from field: livekit.EncodedFileType file_type = 1;
     */
    fileType: EncodedFileType;
    /**
     * see egress docs for templating (default {room_name}-{time})
     *
     * @generated from field: string filepath = 2;
     */
    filepath: string;
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 6;
     */
    disableManifest: boolean;
    /**
     * @generated from oneof livekit.EncodedFileOutput.output
     */
    output: {
        /**
         * @generated from field: livekit.S3Upload s3 = 3;
         */
        value: S3Upload;
        case: "s3";
    } | {
        /**
         * @generated from field: livekit.GCPUpload gcp = 4;
         */
        value: GCPUpload;
        case: "gcp";
    } | {
        /**
         * @generated from field: livekit.AzureBlobUpload azure = 5;
         */
        value: AzureBlobUpload;
        case: "azure";
    } | {
        /**
         * @generated from field: livekit.AliOSSUpload aliOSS = 7;
         */
        value: AliOSSUpload;
        case: "aliOSS";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<EncodedFileOutput>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.EncodedFileOutput";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EncodedFileOutput;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EncodedFileOutput;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EncodedFileOutput;
    static equals(a: EncodedFileOutput | PlainMessage<EncodedFileOutput> | undefined, b: EncodedFileOutput | PlainMessage<EncodedFileOutput> | undefined): boolean;
}
/**
 * Used to generate HLS segments or other kind of segmented output
 *
 * @generated from message livekit.SegmentedFileOutput
 */
export declare class SegmentedFileOutput extends Message<SegmentedFileOutput> {
    /**
     * (optional)
     *
     * @generated from field: livekit.SegmentedFileProtocol protocol = 1;
     */
    protocol: SegmentedFileProtocol;
    /**
     * (optional)
     *
     * @generated from field: string filename_prefix = 2;
     */
    filenamePrefix: string;
    /**
     * (optional)
     *
     * @generated from field: string playlist_name = 3;
     */
    playlistName: string;
    /**
     * (optional, disabled if not provided). Path of a live playlist
     *
     * @generated from field: string live_playlist_name = 11;
     */
    livePlaylistName: string;
    /**
     * in seconds (optional)
     *
     * @generated from field: uint32 segment_duration = 4;
     */
    segmentDuration: number;
    /**
     * (optional, default INDEX)
     *
     * @generated from field: livekit.SegmentedFileSuffix filename_suffix = 10;
     */
    filenameSuffix: SegmentedFileSuffix;
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 8;
     */
    disableManifest: boolean;
    /**
     * required
     *
     * @generated from oneof livekit.SegmentedFileOutput.output
     */
    output: {
        /**
         * @generated from field: livekit.S3Upload s3 = 5;
         */
        value: S3Upload;
        case: "s3";
    } | {
        /**
         * @generated from field: livekit.GCPUpload gcp = 6;
         */
        value: GCPUpload;
        case: "gcp";
    } | {
        /**
         * @generated from field: livekit.AzureBlobUpload azure = 7;
         */
        value: AzureBlobUpload;
        case: "azure";
    } | {
        /**
         * @generated from field: livekit.AliOSSUpload aliOSS = 9;
         */
        value: AliOSSUpload;
        case: "aliOSS";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<SegmentedFileOutput>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SegmentedFileOutput";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SegmentedFileOutput;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SegmentedFileOutput;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SegmentedFileOutput;
    static equals(a: SegmentedFileOutput | PlainMessage<SegmentedFileOutput> | undefined, b: SegmentedFileOutput | PlainMessage<SegmentedFileOutput> | undefined): boolean;
}
/**
 * @generated from message livekit.DirectFileOutput
 */
export declare class DirectFileOutput extends Message<DirectFileOutput> {
    /**
     * see egress docs for templating (default {track_id}-{time})
     *
     * @generated from field: string filepath = 1;
     */
    filepath: string;
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 5;
     */
    disableManifest: boolean;
    /**
     * @generated from oneof livekit.DirectFileOutput.output
     */
    output: {
        /**
         * @generated from field: livekit.S3Upload s3 = 2;
         */
        value: S3Upload;
        case: "s3";
    } | {
        /**
         * @generated from field: livekit.GCPUpload gcp = 3;
         */
        value: GCPUpload;
        case: "gcp";
    } | {
        /**
         * @generated from field: livekit.AzureBlobUpload azure = 4;
         */
        value: AzureBlobUpload;
        case: "azure";
    } | {
        /**
         * @generated from field: livekit.AliOSSUpload aliOSS = 6;
         */
        value: AliOSSUpload;
        case: "aliOSS";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<DirectFileOutput>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DirectFileOutput";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DirectFileOutput;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DirectFileOutput;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DirectFileOutput;
    static equals(a: DirectFileOutput | PlainMessage<DirectFileOutput> | undefined, b: DirectFileOutput | PlainMessage<DirectFileOutput> | undefined): boolean;
}
/**
 * @generated from message livekit.ImageOutput
 */
export declare class ImageOutput extends Message<ImageOutput> {
    /**
     * in seconds (required)
     *
     * @generated from field: uint32 capture_interval = 1;
     */
    captureInterval: number;
    /**
     * (optional, defaults to track width)
     *
     * @generated from field: int32 width = 2;
     */
    width: number;
    /**
     * (optional, defaults to track height)
     *
     * @generated from field: int32 height = 3;
     */
    height: number;
    /**
     * (optional)
     *
     * @generated from field: string filename_prefix = 4;
     */
    filenamePrefix: string;
    /**
     * (optional, default INDEX)
     *
     * @generated from field: livekit.ImageFileSuffix filename_suffix = 5;
     */
    filenameSuffix: ImageFileSuffix;
    /**
     * (optional)
     *
     * @generated from field: livekit.ImageCodec image_codec = 6;
     */
    imageCodec: ImageCodec;
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 7;
     */
    disableManifest: boolean;
    /**
     * required
     *
     * @generated from oneof livekit.ImageOutput.output
     */
    output: {
        /**
         * @generated from field: livekit.S3Upload s3 = 8;
         */
        value: S3Upload;
        case: "s3";
    } | {
        /**
         * @generated from field: livekit.GCPUpload gcp = 9;
         */
        value: GCPUpload;
        case: "gcp";
    } | {
        /**
         * @generated from field: livekit.AzureBlobUpload azure = 10;
         */
        value: AzureBlobUpload;
        case: "azure";
    } | {
        /**
         * @generated from field: livekit.AliOSSUpload aliOSS = 11;
         */
        value: AliOSSUpload;
        case: "aliOSS";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<ImageOutput>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ImageOutput";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImageOutput;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImageOutput;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImageOutput;
    static equals(a: ImageOutput | PlainMessage<ImageOutput> | undefined, b: ImageOutput | PlainMessage<ImageOutput> | undefined): boolean;
}
/**
 * @generated from message livekit.S3Upload
 */
export declare class S3Upload extends Message<S3Upload> {
    /**
     * @generated from field: string access_key = 1;
     */
    accessKey: string;
    /**
     * @generated from field: string secret = 2;
     */
    secret: string;
    /**
     * @generated from field: string session_token = 11;
     */
    sessionToken: string;
    /**
     * @generated from field: string region = 3;
     */
    region: string;
    /**
     * @generated from field: string endpoint = 4;
     */
    endpoint: string;
    /**
     * @generated from field: string bucket = 5;
     */
    bucket: string;
    /**
     * @generated from field: bool force_path_style = 6;
     */
    forcePathStyle: boolean;
    /**
     * @generated from field: map<string, string> metadata = 7;
     */
    metadata: {
        [key: string]: string;
    };
    /**
     * @generated from field: string tagging = 8;
     */
    tagging: string;
    /**
     * Content-Disposition header
     *
     * @generated from field: string content_disposition = 9;
     */
    contentDisposition: string;
    /**
     * @generated from field: livekit.ProxyConfig proxy = 10;
     */
    proxy?: ProxyConfig;
    constructor(data?: PartialMessage<S3Upload>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.S3Upload";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): S3Upload;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): S3Upload;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): S3Upload;
    static equals(a: S3Upload | PlainMessage<S3Upload> | undefined, b: S3Upload | PlainMessage<S3Upload> | undefined): boolean;
}
/**
 * @generated from message livekit.GCPUpload
 */
export declare class GCPUpload extends Message<GCPUpload> {
    /**
     * service account credentials serialized in JSON "credentials.json"
     *
     * @generated from field: string credentials = 1;
     */
    credentials: string;
    /**
     * @generated from field: string bucket = 2;
     */
    bucket: string;
    /**
     * @generated from field: livekit.ProxyConfig proxy = 3;
     */
    proxy?: ProxyConfig;
    constructor(data?: PartialMessage<GCPUpload>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.GCPUpload";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GCPUpload;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GCPUpload;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GCPUpload;
    static equals(a: GCPUpload | PlainMessage<GCPUpload> | undefined, b: GCPUpload | PlainMessage<GCPUpload> | undefined): boolean;
}
/**
 * @generated from message livekit.AzureBlobUpload
 */
export declare class AzureBlobUpload extends Message<AzureBlobUpload> {
    /**
     * @generated from field: string account_name = 1;
     */
    accountName: string;
    /**
     * @generated from field: string account_key = 2;
     */
    accountKey: string;
    /**
     * @generated from field: string container_name = 3;
     */
    containerName: string;
    constructor(data?: PartialMessage<AzureBlobUpload>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AzureBlobUpload";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AzureBlobUpload;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AzureBlobUpload;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AzureBlobUpload;
    static equals(a: AzureBlobUpload | PlainMessage<AzureBlobUpload> | undefined, b: AzureBlobUpload | PlainMessage<AzureBlobUpload> | undefined): boolean;
}
/**
 * @generated from message livekit.AliOSSUpload
 */
export declare class AliOSSUpload extends Message<AliOSSUpload> {
    /**
     * @generated from field: string access_key = 1;
     */
    accessKey: string;
    /**
     * @generated from field: string secret = 2;
     */
    secret: string;
    /**
     * @generated from field: string region = 3;
     */
    region: string;
    /**
     * @generated from field: string endpoint = 4;
     */
    endpoint: string;
    /**
     * @generated from field: string bucket = 5;
     */
    bucket: string;
    constructor(data?: PartialMessage<AliOSSUpload>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AliOSSUpload";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AliOSSUpload;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AliOSSUpload;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AliOSSUpload;
    static equals(a: AliOSSUpload | PlainMessage<AliOSSUpload> | undefined, b: AliOSSUpload | PlainMessage<AliOSSUpload> | undefined): boolean;
}
/**
 * @generated from message livekit.ProxyConfig
 */
export declare class ProxyConfig extends Message<ProxyConfig> {
    /**
     * @generated from field: string url = 1;
     */
    url: string;
    /**
     * @generated from field: string username = 2;
     */
    username: string;
    /**
     * @generated from field: string password = 3;
     */
    password: string;
    constructor(data?: PartialMessage<ProxyConfig>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ProxyConfig";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProxyConfig;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProxyConfig;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProxyConfig;
    static equals(a: ProxyConfig | PlainMessage<ProxyConfig> | undefined, b: ProxyConfig | PlainMessage<ProxyConfig> | undefined): boolean;
}
/**
 * @generated from message livekit.StreamOutput
 */
export declare class StreamOutput extends Message<StreamOutput> {
    /**
     * required
     *
     * @generated from field: livekit.StreamProtocol protocol = 1;
     */
    protocol: StreamProtocol;
    /**
     * required
     *
     * @generated from field: repeated string urls = 2;
     */
    urls: string[];
    constructor(data?: PartialMessage<StreamOutput>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.StreamOutput";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamOutput;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamOutput;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamOutput;
    static equals(a: StreamOutput | PlainMessage<StreamOutput> | undefined, b: StreamOutput | PlainMessage<StreamOutput> | undefined): boolean;
}
/**
 * @generated from message livekit.EncodingOptions
 */
export declare class EncodingOptions extends Message<EncodingOptions> {
    /**
     * (default 1920)
     *
     * @generated from field: int32 width = 1;
     */
    width: number;
    /**
     * (default 1080)
     *
     * @generated from field: int32 height = 2;
     */
    height: number;
    /**
     * (default 24)
     *
     * @generated from field: int32 depth = 3;
     */
    depth: number;
    /**
     * (default 30)
     *
     * @generated from field: int32 framerate = 4;
     */
    framerate: number;
    /**
     * (default OPUS)
     *
     * @generated from field: livekit.AudioCodec audio_codec = 5;
     */
    audioCodec: AudioCodec;
    /**
     * (default 128)
     *
     * @generated from field: int32 audio_bitrate = 6;
     */
    audioBitrate: number;
    /**
     * quality setting on audio encoder
     *
     * @generated from field: int32 audio_quality = 11;
     */
    audioQuality: number;
    /**
     * (default 44100)
     *
     * @generated from field: int32 audio_frequency = 7;
     */
    audioFrequency: number;
    /**
     * (default H264_MAIN)
     *
     * @generated from field: livekit.VideoCodec video_codec = 8;
     */
    videoCodec: VideoCodec;
    /**
     * (default 4500)
     *
     * @generated from field: int32 video_bitrate = 9;
     */
    videoBitrate: number;
    /**
     * quality setting on video encoder
     *
     * @generated from field: int32 video_quality = 12;
     */
    videoQuality: number;
    /**
     * in seconds (default 4s for streaming, segment duration for segmented output, encoder default for files)
     *
     * @generated from field: double key_frame_interval = 10;
     */
    keyFrameInterval: number;
    constructor(data?: PartialMessage<EncodingOptions>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.EncodingOptions";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EncodingOptions;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EncodingOptions;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EncodingOptions;
    static equals(a: EncodingOptions | PlainMessage<EncodingOptions> | undefined, b: EncodingOptions | PlainMessage<EncodingOptions> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateLayoutRequest
 */
export declare class UpdateLayoutRequest extends Message<UpdateLayoutRequest> {
    /**
     * @generated from field: string egress_id = 1;
     */
    egressId: string;
    /**
     * @generated from field: string layout = 2;
     */
    layout: string;
    constructor(data?: PartialMessage<UpdateLayoutRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateLayoutRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateLayoutRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateLayoutRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateLayoutRequest;
    static equals(a: UpdateLayoutRequest | PlainMessage<UpdateLayoutRequest> | undefined, b: UpdateLayoutRequest | PlainMessage<UpdateLayoutRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateStreamRequest
 */
export declare class UpdateStreamRequest extends Message<UpdateStreamRequest> {
    /**
     * @generated from field: string egress_id = 1;
     */
    egressId: string;
    /**
     * @generated from field: repeated string add_output_urls = 2;
     */
    addOutputUrls: string[];
    /**
     * @generated from field: repeated string remove_output_urls = 3;
     */
    removeOutputUrls: string[];
    constructor(data?: PartialMessage<UpdateStreamRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateStreamRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateStreamRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateStreamRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateStreamRequest;
    static equals(a: UpdateStreamRequest | PlainMessage<UpdateStreamRequest> | undefined, b: UpdateStreamRequest | PlainMessage<UpdateStreamRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListEgressRequest
 */
export declare class ListEgressRequest extends Message<ListEgressRequest> {
    /**
     * (optional, filter by room name)
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * (optional, filter by egress ID)
     *
     * @generated from field: string egress_id = 2;
     */
    egressId: string;
    /**
     * (optional, list active egress only)
     *
     * @generated from field: bool active = 3;
     */
    active: boolean;
    constructor(data?: PartialMessage<ListEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEgressRequest;
    static equals(a: ListEgressRequest | PlainMessage<ListEgressRequest> | undefined, b: ListEgressRequest | PlainMessage<ListEgressRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListEgressResponse
 */
export declare class ListEgressResponse extends Message<ListEgressResponse> {
    /**
     * @generated from field: repeated livekit.EgressInfo items = 1;
     */
    items: EgressInfo[];
    constructor(data?: PartialMessage<ListEgressResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListEgressResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEgressResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEgressResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEgressResponse;
    static equals(a: ListEgressResponse | PlainMessage<ListEgressResponse> | undefined, b: ListEgressResponse | PlainMessage<ListEgressResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.StopEgressRequest
 */
export declare class StopEgressRequest extends Message<StopEgressRequest> {
    /**
     * @generated from field: string egress_id = 1;
     */
    egressId: string;
    constructor(data?: PartialMessage<StopEgressRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.StopEgressRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StopEgressRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StopEgressRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StopEgressRequest;
    static equals(a: StopEgressRequest | PlainMessage<StopEgressRequest> | undefined, b: StopEgressRequest | PlainMessage<StopEgressRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.EgressInfo
 */
export declare class EgressInfo extends Message<EgressInfo> {
    /**
     * @generated from field: string egress_id = 1;
     */
    egressId: string;
    /**
     * @generated from field: string room_id = 2;
     */
    roomId: string;
    /**
     * @generated from field: string room_name = 13;
     */
    roomName: string;
    /**
     * @generated from field: livekit.EgressStatus status = 3;
     */
    status: EgressStatus;
    /**
     * @generated from field: int64 started_at = 10;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 11;
     */
    endedAt: bigint;
    /**
     * @generated from field: int64 updated_at = 18;
     */
    updatedAt: bigint;
    /**
     * @generated from field: string details = 21;
     */
    details: string;
    /**
     * @generated from field: string error = 9;
     */
    error: string;
    /**
     * @generated from field: int32 error_code = 22;
     */
    errorCode: number;
    /**
     * @generated from oneof livekit.EgressInfo.request
     */
    request: {
        /**
         * @generated from field: livekit.RoomCompositeEgressRequest room_composite = 4;
         */
        value: RoomCompositeEgressRequest;
        case: "roomComposite";
    } | {
        /**
         * @generated from field: livekit.WebEgressRequest web = 14;
         */
        value: WebEgressRequest;
        case: "web";
    } | {
        /**
         * @generated from field: livekit.ParticipantEgressRequest participant = 19;
         */
        value: ParticipantEgressRequest;
        case: "participant";
    } | {
        /**
         * @generated from field: livekit.TrackCompositeEgressRequest track_composite = 5;
         */
        value: TrackCompositeEgressRequest;
        case: "trackComposite";
    } | {
        /**
         * @generated from field: livekit.TrackEgressRequest track = 6;
         */
        value: TrackEgressRequest;
        case: "track";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * deprecated (use _result fields)
     *
     * @generated from oneof livekit.EgressInfo.result
     */
    result: {
        /**
         * @generated from field: livekit.StreamInfoList stream = 7 [deprecated = true];
         * @deprecated
         */
        value: StreamInfoList;
        case: "stream";
    } | {
        /**
         * @generated from field: livekit.FileInfo file = 8 [deprecated = true];
         * @deprecated
         */
        value: FileInfo;
        case: "file";
    } | {
        /**
         * @generated from field: livekit.SegmentsInfo segments = 12 [deprecated = true];
         * @deprecated
         */
        value: SegmentsInfo;
        case: "segments";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.StreamInfo stream_results = 15;
     */
    streamResults: StreamInfo[];
    /**
     * @generated from field: repeated livekit.FileInfo file_results = 16;
     */
    fileResults: FileInfo[];
    /**
     * @generated from field: repeated livekit.SegmentsInfo segment_results = 17;
     */
    segmentResults: SegmentsInfo[];
    /**
     * @generated from field: repeated livekit.ImagesInfo image_results = 20;
     */
    imageResults: ImagesInfo[];
    /**
     * @generated from field: string manifest_location = 23;
     */
    manifestLocation: string;
    /**
     * next ID: 26
     *
     * @generated from field: bool backup_storage_used = 25;
     */
    backupStorageUsed: boolean;
    constructor(data?: PartialMessage<EgressInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.EgressInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EgressInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EgressInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EgressInfo;
    static equals(a: EgressInfo | PlainMessage<EgressInfo> | undefined, b: EgressInfo | PlainMessage<EgressInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.StreamInfoList
 * @deprecated
 */
export declare class StreamInfoList extends Message<StreamInfoList> {
    /**
     * @generated from field: repeated livekit.StreamInfo info = 1;
     */
    info: StreamInfo[];
    constructor(data?: PartialMessage<StreamInfoList>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.StreamInfoList";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamInfoList;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamInfoList;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamInfoList;
    static equals(a: StreamInfoList | PlainMessage<StreamInfoList> | undefined, b: StreamInfoList | PlainMessage<StreamInfoList> | undefined): boolean;
}
/**
 * @generated from message livekit.StreamInfo
 */
export declare class StreamInfo extends Message<StreamInfo> {
    /**
     * @generated from field: string url = 1;
     */
    url: string;
    /**
     * @generated from field: int64 started_at = 2;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 3;
     */
    endedAt: bigint;
    /**
     * @generated from field: int64 duration = 4;
     */
    duration: bigint;
    /**
     * @generated from field: livekit.StreamInfo.Status status = 5;
     */
    status: StreamInfo_Status;
    /**
     * @generated from field: string error = 6;
     */
    error: string;
    constructor(data?: PartialMessage<StreamInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.StreamInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamInfo;
    static equals(a: StreamInfo | PlainMessage<StreamInfo> | undefined, b: StreamInfo | PlainMessage<StreamInfo> | undefined): boolean;
}
/**
 * @generated from enum livekit.StreamInfo.Status
 */
export declare enum StreamInfo_Status {
    /**
     * @generated from enum value: ACTIVE = 0;
     */
    ACTIVE = 0,
    /**
     * @generated from enum value: FINISHED = 1;
     */
    FINISHED = 1,
    /**
     * @generated from enum value: FAILED = 2;
     */
    FAILED = 2
}
/**
 * @generated from message livekit.FileInfo
 */
export declare class FileInfo extends Message<FileInfo> {
    /**
     * @generated from field: string filename = 1;
     */
    filename: string;
    /**
     * @generated from field: int64 started_at = 2;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 3;
     */
    endedAt: bigint;
    /**
     * @generated from field: int64 duration = 6;
     */
    duration: bigint;
    /**
     * @generated from field: int64 size = 4;
     */
    size: bigint;
    /**
     * @generated from field: string location = 5;
     */
    location: string;
    constructor(data?: PartialMessage<FileInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.FileInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FileInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FileInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FileInfo;
    static equals(a: FileInfo | PlainMessage<FileInfo> | undefined, b: FileInfo | PlainMessage<FileInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.SegmentsInfo
 */
export declare class SegmentsInfo extends Message<SegmentsInfo> {
    /**
     * @generated from field: string playlist_name = 1;
     */
    playlistName: string;
    /**
     * @generated from field: string live_playlist_name = 8;
     */
    livePlaylistName: string;
    /**
     * @generated from field: int64 duration = 2;
     */
    duration: bigint;
    /**
     * @generated from field: int64 size = 3;
     */
    size: bigint;
    /**
     * @generated from field: string playlist_location = 4;
     */
    playlistLocation: string;
    /**
     * @generated from field: string live_playlist_location = 9;
     */
    livePlaylistLocation: string;
    /**
     * @generated from field: int64 segment_count = 5;
     */
    segmentCount: bigint;
    /**
     * @generated from field: int64 started_at = 6;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 7;
     */
    endedAt: bigint;
    constructor(data?: PartialMessage<SegmentsInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SegmentsInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SegmentsInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SegmentsInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SegmentsInfo;
    static equals(a: SegmentsInfo | PlainMessage<SegmentsInfo> | undefined, b: SegmentsInfo | PlainMessage<SegmentsInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.ImagesInfo
 */
export declare class ImagesInfo extends Message<ImagesInfo> {
    /**
     * @generated from field: string filename_prefix = 4;
     */
    filenamePrefix: string;
    /**
     * @generated from field: int64 image_count = 1;
     */
    imageCount: bigint;
    /**
     * @generated from field: int64 started_at = 2;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 3;
     */
    endedAt: bigint;
    constructor(data?: PartialMessage<ImagesInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ImagesInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImagesInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImagesInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImagesInfo;
    static equals(a: ImagesInfo | PlainMessage<ImagesInfo> | undefined, b: ImagesInfo | PlainMessage<ImagesInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.AutoParticipantEgress
 */
export declare class AutoParticipantEgress extends Message<AutoParticipantEgress> {
    /**
     * @generated from oneof livekit.AutoParticipantEgress.options
     */
    options: {
        /**
         * (default H264_720P_30)
         *
         * @generated from field: livekit.EncodingOptionsPreset preset = 1;
         */
        value: EncodingOptionsPreset;
        case: "preset";
    } | {
        /**
         * (optional)
         *
         * @generated from field: livekit.EncodingOptions advanced = 2;
         */
        value: EncodingOptions;
        case: "advanced";
    } | {
        case: undefined;
        value?: undefined;
    };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 3;
     */
    fileOutputs: EncodedFileOutput[];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 4;
     */
    segmentOutputs: SegmentedFileOutput[];
    constructor(data?: PartialMessage<AutoParticipantEgress>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AutoParticipantEgress";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AutoParticipantEgress;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AutoParticipantEgress;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AutoParticipantEgress;
    static equals(a: AutoParticipantEgress | PlainMessage<AutoParticipantEgress> | undefined, b: AutoParticipantEgress | PlainMessage<AutoParticipantEgress> | undefined): boolean;
}
/**
 * @generated from message livekit.AutoTrackEgress
 */
export declare class AutoTrackEgress extends Message<AutoTrackEgress> {
    /**
     * see docs for templating (default {track_id}-{time})
     *
     * @generated from field: string filepath = 1;
     */
    filepath: string;
    /**
     * disables upload of json manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 5;
     */
    disableManifest: boolean;
    /**
     * @generated from oneof livekit.AutoTrackEgress.output
     */
    output: {
        /**
         * @generated from field: livekit.S3Upload s3 = 2;
         */
        value: S3Upload;
        case: "s3";
    } | {
        /**
         * @generated from field: livekit.GCPUpload gcp = 3;
         */
        value: GCPUpload;
        case: "gcp";
    } | {
        /**
         * @generated from field: livekit.AzureBlobUpload azure = 4;
         */
        value: AzureBlobUpload;
        case: "azure";
    } | {
        /**
         * @generated from field: livekit.AliOSSUpload aliOSS = 6;
         */
        value: AliOSSUpload;
        case: "aliOSS";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<AutoTrackEgress>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AutoTrackEgress";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AutoTrackEgress;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AutoTrackEgress;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AutoTrackEgress;
    static equals(a: AutoTrackEgress | PlainMessage<AutoTrackEgress> | undefined, b: AutoTrackEgress | PlainMessage<AutoTrackEgress> | undefined): boolean;
}
//# sourceMappingURL=livekit_egress_pb.d.ts.map