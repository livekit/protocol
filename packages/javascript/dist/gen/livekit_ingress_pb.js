import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { AudioCodec, TrackInfo, TrackSource, VideoCodec, VideoLayer } from "./livekit_models_pb.js";
var IngressInput = /* @__PURE__ */ ((IngressInput2) => {
  IngressInput2[IngressInput2["RTMP_INPUT"] = 0] = "RTMP_INPUT";
  IngressInput2[IngressInput2["WHIP_INPUT"] = 1] = "WHIP_INPUT";
  IngressInput2[IngressInput2["URL_INPUT"] = 2] = "URL_INPUT";
  return IngressInput2;
})(IngressInput || {});
proto3.util.setEnumType(IngressInput, "livekit.IngressInput", [
  { no: 0, name: "RTMP_INPUT" },
  { no: 1, name: "WHIP_INPUT" },
  { no: 2, name: "URL_INPUT" }
]);
var IngressAudioEncodingPreset = /* @__PURE__ */ ((IngressAudioEncodingPreset2) => {
  IngressAudioEncodingPreset2[IngressAudioEncodingPreset2["OPUS_STEREO_96KBPS"] = 0] = "OPUS_STEREO_96KBPS";
  IngressAudioEncodingPreset2[IngressAudioEncodingPreset2["OPUS_MONO_64KBS"] = 1] = "OPUS_MONO_64KBS";
  return IngressAudioEncodingPreset2;
})(IngressAudioEncodingPreset || {});
proto3.util.setEnumType(IngressAudioEncodingPreset, "livekit.IngressAudioEncodingPreset", [
  { no: 0, name: "OPUS_STEREO_96KBPS" },
  { no: 1, name: "OPUS_MONO_64KBS" }
]);
var IngressVideoEncodingPreset = /* @__PURE__ */ ((IngressVideoEncodingPreset2) => {
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_720P_30FPS_3_LAYERS"] = 0] = "H264_720P_30FPS_3_LAYERS";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_1080P_30FPS_3_LAYERS"] = 1] = "H264_1080P_30FPS_3_LAYERS";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_540P_25FPS_2_LAYERS"] = 2] = "H264_540P_25FPS_2_LAYERS";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_720P_30FPS_1_LAYER"] = 3] = "H264_720P_30FPS_1_LAYER";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_1080P_30FPS_1_LAYER"] = 4] = "H264_1080P_30FPS_1_LAYER";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_720P_30FPS_3_LAYERS_HIGH_MOTION"] = 5] = "H264_720P_30FPS_3_LAYERS_HIGH_MOTION";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_1080P_30FPS_3_LAYERS_HIGH_MOTION"] = 6] = "H264_1080P_30FPS_3_LAYERS_HIGH_MOTION";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_540P_25FPS_2_LAYERS_HIGH_MOTION"] = 7] = "H264_540P_25FPS_2_LAYERS_HIGH_MOTION";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_720P_30FPS_1_LAYER_HIGH_MOTION"] = 8] = "H264_720P_30FPS_1_LAYER_HIGH_MOTION";
  IngressVideoEncodingPreset2[IngressVideoEncodingPreset2["H264_1080P_30FPS_1_LAYER_HIGH_MOTION"] = 9] = "H264_1080P_30FPS_1_LAYER_HIGH_MOTION";
  return IngressVideoEncodingPreset2;
})(IngressVideoEncodingPreset || {});
proto3.util.setEnumType(IngressVideoEncodingPreset, "livekit.IngressVideoEncodingPreset", [
  { no: 0, name: "H264_720P_30FPS_3_LAYERS" },
  { no: 1, name: "H264_1080P_30FPS_3_LAYERS" },
  { no: 2, name: "H264_540P_25FPS_2_LAYERS" },
  { no: 3, name: "H264_720P_30FPS_1_LAYER" },
  { no: 4, name: "H264_1080P_30FPS_1_LAYER" },
  { no: 5, name: "H264_720P_30FPS_3_LAYERS_HIGH_MOTION" },
  { no: 6, name: "H264_1080P_30FPS_3_LAYERS_HIGH_MOTION" },
  { no: 7, name: "H264_540P_25FPS_2_LAYERS_HIGH_MOTION" },
  { no: 8, name: "H264_720P_30FPS_1_LAYER_HIGH_MOTION" },
  { no: 9, name: "H264_1080P_30FPS_1_LAYER_HIGH_MOTION" }
]);
const _CreateIngressRequest = class _CreateIngressRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.IngressInput input_type = 1;
     */
    this.inputType = 0 /* RTMP_INPUT */;
    /**
     * Where to pull media from, only for URL input type
     *
     * @generated from field: string url = 9;
     */
    this.url = "";
    /**
     * User provided identifier for the ingress
     *
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * room to publish to
     *
     * @generated from field: string room_name = 3;
     */
    this.roomName = "";
    /**
     * publish as participant
     *
     * @generated from field: string participant_identity = 4;
     */
    this.participantIdentity = "";
    /**
     * name of publishing participant (used for display only)
     *
     * @generated from field: string participant_name = 5;
     */
    this.participantName = "";
    /**
     * metadata associated with the publishing participant
     *
     * @generated from field: string participant_metadata = 10;
     */
    this.participantMetadata = "";
    /**
     * [depreacted ] whether to pass through the incoming media without transcoding, only compatible with some input types. Use `enable_transcoding` instead.
     *
     * @generated from field: bool bypass_transcoding = 8 [deprecated = true];
     * @deprecated
     */
    this.bypassTranscoding = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateIngressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateIngressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateIngressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateIngressRequest, a, b);
  }
};
_CreateIngressRequest.runtime = proto3;
_CreateIngressRequest.typeName = "livekit.CreateIngressRequest";
_CreateIngressRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "input_type", kind: "enum", T: proto3.getEnumType(IngressInput) },
  {
    no: 9,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "participant_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 10,
    name: "participant_metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "bypass_transcoding",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 11, name: "enable_transcoding", kind: "scalar", T: 8, opt: true },
  { no: 6, name: "audio", kind: "message", T: IngressAudioOptions },
  { no: 7, name: "video", kind: "message", T: IngressVideoOptions }
]);
let CreateIngressRequest = _CreateIngressRequest;
const _IngressAudioOptions = class _IngressAudioOptions extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string name = 1;
     */
    this.name = "";
    /**
     * @generated from field: livekit.TrackSource source = 2;
     */
    this.source = TrackSource.UNKNOWN;
    /**
     * @generated from oneof livekit.IngressAudioOptions.encoding_options
     */
    this.encodingOptions = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _IngressAudioOptions().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _IngressAudioOptions().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _IngressAudioOptions().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_IngressAudioOptions, a, b);
  }
};
_IngressAudioOptions.runtime = proto3;
_IngressAudioOptions.typeName = "livekit.IngressAudioOptions";
_IngressAudioOptions.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "source", kind: "enum", T: proto3.getEnumType(TrackSource) },
  { no: 3, name: "preset", kind: "enum", T: proto3.getEnumType(IngressAudioEncodingPreset), oneof: "encoding_options" },
  { no: 4, name: "options", kind: "message", T: IngressAudioEncodingOptions, oneof: "encoding_options" }
]);
let IngressAudioOptions = _IngressAudioOptions;
const _IngressVideoOptions = class _IngressVideoOptions extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string name = 1;
     */
    this.name = "";
    /**
     * @generated from field: livekit.TrackSource source = 2;
     */
    this.source = TrackSource.UNKNOWN;
    /**
     * @generated from oneof livekit.IngressVideoOptions.encoding_options
     */
    this.encodingOptions = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _IngressVideoOptions().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _IngressVideoOptions().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _IngressVideoOptions().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_IngressVideoOptions, a, b);
  }
};
_IngressVideoOptions.runtime = proto3;
_IngressVideoOptions.typeName = "livekit.IngressVideoOptions";
_IngressVideoOptions.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "source", kind: "enum", T: proto3.getEnumType(TrackSource) },
  { no: 3, name: "preset", kind: "enum", T: proto3.getEnumType(IngressVideoEncodingPreset), oneof: "encoding_options" },
  { no: 4, name: "options", kind: "message", T: IngressVideoEncodingOptions, oneof: "encoding_options" }
]);
let IngressVideoOptions = _IngressVideoOptions;
const _IngressAudioEncodingOptions = class _IngressAudioEncodingOptions extends Message {
  constructor(data) {
    super();
    /**
     * desired audio codec to publish to room
     *
     * @generated from field: livekit.AudioCodec audio_codec = 1;
     */
    this.audioCodec = AudioCodec.DEFAULT_AC;
    /**
     * @generated from field: uint32 bitrate = 2;
     */
    this.bitrate = 0;
    /**
     * @generated from field: bool disable_dtx = 3;
     */
    this.disableDtx = false;
    /**
     * @generated from field: uint32 channels = 4;
     */
    this.channels = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _IngressAudioEncodingOptions().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _IngressAudioEncodingOptions().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _IngressAudioEncodingOptions().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_IngressAudioEncodingOptions, a, b);
  }
};
_IngressAudioEncodingOptions.runtime = proto3;
_IngressAudioEncodingOptions.typeName = "livekit.IngressAudioEncodingOptions";
_IngressAudioEncodingOptions.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "audio_codec", kind: "enum", T: proto3.getEnumType(AudioCodec) },
  {
    no: 2,
    name: "bitrate",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "disable_dtx",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 4,
    name: "channels",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let IngressAudioEncodingOptions = _IngressAudioEncodingOptions;
const _IngressVideoEncodingOptions = class _IngressVideoEncodingOptions extends Message {
  constructor(data) {
    super();
    /**
     * desired codec to publish to room
     *
     * @generated from field: livekit.VideoCodec video_codec = 1;
     */
    this.videoCodec = VideoCodec.DEFAULT_VC;
    /**
     * @generated from field: double frame_rate = 2;
     */
    this.frameRate = 0;
    /**
     * simulcast layers to publish, when empty, should usually be set to layers at 1/2 and 1/4 of the dimensions
     *
     * @generated from field: repeated livekit.VideoLayer layers = 3;
     */
    this.layers = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _IngressVideoEncodingOptions().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _IngressVideoEncodingOptions().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _IngressVideoEncodingOptions().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_IngressVideoEncodingOptions, a, b);
  }
};
_IngressVideoEncodingOptions.runtime = proto3;
_IngressVideoEncodingOptions.typeName = "livekit.IngressVideoEncodingOptions";
_IngressVideoEncodingOptions.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "video_codec", kind: "enum", T: proto3.getEnumType(VideoCodec) },
  {
    no: 2,
    name: "frame_rate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  },
  { no: 3, name: "layers", kind: "message", T: VideoLayer, repeated: true }
]);
let IngressVideoEncodingOptions = _IngressVideoEncodingOptions;
const _IngressInfo = class _IngressInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string ingress_id = 1;
     */
    this.ingressId = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: string stream_key = 3;
     */
    this.streamKey = "";
    /**
     * URL to point the encoder to for push (RTMP, WHIP), or location to pull media from for pull (URL)
     *
     * @generated from field: string url = 4;
     */
    this.url = "";
    /**
     * for RTMP input, it'll be a rtmp:// URL
     * for FILE input, it'll be a http:// URL
     * for SRT input, it'll be a srt:// URL
     *
     * @generated from field: livekit.IngressInput input_type = 5;
     */
    this.inputType = 0 /* RTMP_INPUT */;
    /**
     * @generated from field: bool bypass_transcoding = 13 [deprecated = true];
     * @deprecated
     */
    this.bypassTranscoding = false;
    /**
     * @generated from field: string room_name = 8;
     */
    this.roomName = "";
    /**
     * @generated from field: string participant_identity = 9;
     */
    this.participantIdentity = "";
    /**
     * @generated from field: string participant_name = 10;
     */
    this.participantName = "";
    /**
     * @generated from field: string participant_metadata = 14;
     */
    this.participantMetadata = "";
    /**
     * @generated from field: bool reusable = 11;
     */
    this.reusable = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _IngressInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _IngressInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _IngressInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_IngressInfo, a, b);
  }
};
_IngressInfo.runtime = proto3;
_IngressInfo.typeName = "livekit.IngressInfo";
_IngressInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "ingress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "stream_key",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "input_type", kind: "enum", T: proto3.getEnumType(IngressInput) },
  {
    no: 13,
    name: "bypass_transcoding",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 15, name: "enable_transcoding", kind: "scalar", T: 8, opt: true },
  { no: 6, name: "audio", kind: "message", T: IngressAudioOptions },
  { no: 7, name: "video", kind: "message", T: IngressVideoOptions },
  {
    no: 8,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 10,
    name: "participant_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 14,
    name: "participant_metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 11,
    name: "reusable",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 12, name: "state", kind: "message", T: IngressState }
]);
let IngressInfo = _IngressInfo;
const _IngressState = class _IngressState extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.IngressState.Status status = 1;
     */
    this.status = 0 /* ENDPOINT_INACTIVE */;
    /**
     * Error/non compliance description if any
     *
     * @generated from field: string error = 2;
     */
    this.error = "";
    /**
     * ID of the current/previous room published to
     *
     * @generated from field: string room_id = 5;
     */
    this.roomId = "";
    /**
     * @generated from field: int64 started_at = 7;
     */
    this.startedAt = protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 8;
     */
    this.endedAt = protoInt64.zero;
    /**
     * @generated from field: int64 updated_at = 10;
     */
    this.updatedAt = protoInt64.zero;
    /**
     * @generated from field: string resource_id = 9;
     */
    this.resourceId = "";
    /**
     * @generated from field: repeated livekit.TrackInfo tracks = 6;
     */
    this.tracks = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _IngressState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _IngressState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _IngressState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_IngressState, a, b);
  }
};
_IngressState.runtime = proto3;
_IngressState.typeName = "livekit.IngressState";
_IngressState.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(IngressState_Status) },
  {
    no: 2,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "video", kind: "message", T: InputVideoState },
  { no: 4, name: "audio", kind: "message", T: InputAudioState },
  {
    no: 5,
    name: "room_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 8,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 10,
    name: "updated_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 9,
    name: "resource_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "tracks", kind: "message", T: TrackInfo, repeated: true }
]);
let IngressState = _IngressState;
var IngressState_Status = /* @__PURE__ */ ((IngressState_Status2) => {
  IngressState_Status2[IngressState_Status2["ENDPOINT_INACTIVE"] = 0] = "ENDPOINT_INACTIVE";
  IngressState_Status2[IngressState_Status2["ENDPOINT_BUFFERING"] = 1] = "ENDPOINT_BUFFERING";
  IngressState_Status2[IngressState_Status2["ENDPOINT_PUBLISHING"] = 2] = "ENDPOINT_PUBLISHING";
  IngressState_Status2[IngressState_Status2["ENDPOINT_ERROR"] = 3] = "ENDPOINT_ERROR";
  IngressState_Status2[IngressState_Status2["ENDPOINT_COMPLETE"] = 4] = "ENDPOINT_COMPLETE";
  return IngressState_Status2;
})(IngressState_Status || {});
proto3.util.setEnumType(IngressState_Status, "livekit.IngressState.Status", [
  { no: 0, name: "ENDPOINT_INACTIVE" },
  { no: 1, name: "ENDPOINT_BUFFERING" },
  { no: 2, name: "ENDPOINT_PUBLISHING" },
  { no: 3, name: "ENDPOINT_ERROR" },
  { no: 4, name: "ENDPOINT_COMPLETE" }
]);
const _InputVideoState = class _InputVideoState extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string mime_type = 1;
     */
    this.mimeType = "";
    /**
     * @generated from field: uint32 average_bitrate = 2;
     */
    this.averageBitrate = 0;
    /**
     * @generated from field: uint32 width = 3;
     */
    this.width = 0;
    /**
     * @generated from field: uint32 height = 4;
     */
    this.height = 0;
    /**
     * @generated from field: double framerate = 5;
     */
    this.framerate = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _InputVideoState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _InputVideoState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _InputVideoState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_InputVideoState, a, b);
  }
};
_InputVideoState.runtime = proto3;
_InputVideoState.typeName = "livekit.InputVideoState";
_InputVideoState.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "mime_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "average_bitrate",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "width",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "height",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 5,
    name: "framerate",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  }
]);
let InputVideoState = _InputVideoState;
const _InputAudioState = class _InputAudioState extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string mime_type = 1;
     */
    this.mimeType = "";
    /**
     * @generated from field: uint32 average_bitrate = 2;
     */
    this.averageBitrate = 0;
    /**
     * @generated from field: uint32 channels = 3;
     */
    this.channels = 0;
    /**
     * @generated from field: uint32 sample_rate = 4;
     */
    this.sampleRate = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _InputAudioState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _InputAudioState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _InputAudioState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_InputAudioState, a, b);
  }
};
_InputAudioState.runtime = proto3;
_InputAudioState.typeName = "livekit.InputAudioState";
_InputAudioState.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "mime_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "average_bitrate",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "channels",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "sample_rate",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let InputAudioState = _InputAudioState;
const _UpdateIngressRequest = class _UpdateIngressRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string ingress_id = 1;
     */
    this.ingressId = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: string room_name = 3;
     */
    this.roomName = "";
    /**
     * @generated from field: string participant_identity = 4;
     */
    this.participantIdentity = "";
    /**
     * @generated from field: string participant_name = 5;
     */
    this.participantName = "";
    /**
     * @generated from field: string participant_metadata = 9;
     */
    this.participantMetadata = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateIngressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateIngressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateIngressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateIngressRequest, a, b);
  }
};
_UpdateIngressRequest.runtime = proto3;
_UpdateIngressRequest.typeName = "livekit.UpdateIngressRequest";
_UpdateIngressRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "ingress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "participant_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "participant_metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 8, name: "bypass_transcoding", kind: "scalar", T: 8, opt: true },
  { no: 10, name: "enable_transcoding", kind: "scalar", T: 8, opt: true },
  { no: 6, name: "audio", kind: "message", T: IngressAudioOptions },
  { no: 7, name: "video", kind: "message", T: IngressVideoOptions }
]);
let UpdateIngressRequest = _UpdateIngressRequest;
const _ListIngressRequest = class _ListIngressRequest extends Message {
  constructor(data) {
    super();
    /**
     * when blank, lists all ingress endpoints
     *
     * (optional, filter by room name)
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * (optional, filter by ingress ID)
     *
     * @generated from field: string ingress_id = 2;
     */
    this.ingressId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListIngressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListIngressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListIngressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListIngressRequest, a, b);
  }
};
_ListIngressRequest.runtime = proto3;
_ListIngressRequest.typeName = "livekit.ListIngressRequest";
_ListIngressRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "ingress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let ListIngressRequest = _ListIngressRequest;
const _ListIngressResponse = class _ListIngressResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.IngressInfo items = 1;
     */
    this.items = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListIngressResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListIngressResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListIngressResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListIngressResponse, a, b);
  }
};
_ListIngressResponse.runtime = proto3;
_ListIngressResponse.typeName = "livekit.ListIngressResponse";
_ListIngressResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "items", kind: "message", T: IngressInfo, repeated: true }
]);
let ListIngressResponse = _ListIngressResponse;
const _DeleteIngressRequest = class _DeleteIngressRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string ingress_id = 1;
     */
    this.ingressId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteIngressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteIngressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteIngressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_DeleteIngressRequest, a, b);
  }
};
_DeleteIngressRequest.runtime = proto3;
_DeleteIngressRequest.typeName = "livekit.DeleteIngressRequest";
_DeleteIngressRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "ingress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let DeleteIngressRequest = _DeleteIngressRequest;
export {
  CreateIngressRequest,
  DeleteIngressRequest,
  IngressAudioEncodingOptions,
  IngressAudioEncodingPreset,
  IngressAudioOptions,
  IngressInfo,
  IngressInput,
  IngressState,
  IngressState_Status,
  IngressVideoEncodingOptions,
  IngressVideoEncodingPreset,
  IngressVideoOptions,
  InputAudioState,
  InputVideoState,
  ListIngressRequest,
  ListIngressResponse,
  UpdateIngressRequest
};
//# sourceMappingURL=livekit_ingress_pb.js.map