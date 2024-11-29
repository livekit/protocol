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
var livekit_egress_pb_exports = {};
__export(livekit_egress_pb_exports, {
  AliOSSUpload: () => AliOSSUpload,
  AutoParticipantEgress: () => AutoParticipantEgress,
  AutoTrackEgress: () => AutoTrackEgress,
  AzureBlobUpload: () => AzureBlobUpload,
  DirectFileOutput: () => DirectFileOutput,
  EgressInfo: () => EgressInfo,
  EgressStatus: () => EgressStatus,
  EncodedFileOutput: () => EncodedFileOutput,
  EncodedFileType: () => EncodedFileType,
  EncodingOptions: () => EncodingOptions,
  EncodingOptionsPreset: () => EncodingOptionsPreset,
  FileInfo: () => FileInfo,
  GCPUpload: () => GCPUpload,
  ImageFileSuffix: () => ImageFileSuffix,
  ImageOutput: () => ImageOutput,
  ImagesInfo: () => ImagesInfo,
  ListEgressRequest: () => ListEgressRequest,
  ListEgressResponse: () => ListEgressResponse,
  ParticipantEgressRequest: () => ParticipantEgressRequest,
  ProxyConfig: () => ProxyConfig,
  RoomCompositeEgressRequest: () => RoomCompositeEgressRequest,
  S3Upload: () => S3Upload,
  SegmentedFileOutput: () => SegmentedFileOutput,
  SegmentedFileProtocol: () => SegmentedFileProtocol,
  SegmentedFileSuffix: () => SegmentedFileSuffix,
  SegmentsInfo: () => SegmentsInfo,
  StopEgressRequest: () => StopEgressRequest,
  StreamInfo: () => StreamInfo,
  StreamInfoList: () => StreamInfoList,
  StreamInfo_Status: () => StreamInfo_Status,
  StreamOutput: () => StreamOutput,
  StreamProtocol: () => StreamProtocol,
  TrackCompositeEgressRequest: () => TrackCompositeEgressRequest,
  TrackEgressRequest: () => TrackEgressRequest,
  UpdateLayoutRequest: () => UpdateLayoutRequest,
  UpdateStreamRequest: () => UpdateStreamRequest,
  WebEgressRequest: () => WebEgressRequest
});
module.exports = __toCommonJS(livekit_egress_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var import_livekit_models_pb = require("./livekit_models_pb.cjs");
var EncodedFileType = /* @__PURE__ */ ((EncodedFileType2) => {
  EncodedFileType2[EncodedFileType2["DEFAULT_FILETYPE"] = 0] = "DEFAULT_FILETYPE";
  EncodedFileType2[EncodedFileType2["MP4"] = 1] = "MP4";
  EncodedFileType2[EncodedFileType2["OGG"] = 2] = "OGG";
  return EncodedFileType2;
})(EncodedFileType || {});
import_protobuf.proto3.util.setEnumType(EncodedFileType, "livekit.EncodedFileType", [
  { no: 0, name: "DEFAULT_FILETYPE" },
  { no: 1, name: "MP4" },
  { no: 2, name: "OGG" }
]);
var SegmentedFileProtocol = /* @__PURE__ */ ((SegmentedFileProtocol2) => {
  SegmentedFileProtocol2[SegmentedFileProtocol2["DEFAULT_SEGMENTED_FILE_PROTOCOL"] = 0] = "DEFAULT_SEGMENTED_FILE_PROTOCOL";
  SegmentedFileProtocol2[SegmentedFileProtocol2["HLS_PROTOCOL"] = 1] = "HLS_PROTOCOL";
  return SegmentedFileProtocol2;
})(SegmentedFileProtocol || {});
import_protobuf.proto3.util.setEnumType(SegmentedFileProtocol, "livekit.SegmentedFileProtocol", [
  { no: 0, name: "DEFAULT_SEGMENTED_FILE_PROTOCOL" },
  { no: 1, name: "HLS_PROTOCOL" }
]);
var SegmentedFileSuffix = /* @__PURE__ */ ((SegmentedFileSuffix2) => {
  SegmentedFileSuffix2[SegmentedFileSuffix2["INDEX"] = 0] = "INDEX";
  SegmentedFileSuffix2[SegmentedFileSuffix2["TIMESTAMP"] = 1] = "TIMESTAMP";
  return SegmentedFileSuffix2;
})(SegmentedFileSuffix || {});
import_protobuf.proto3.util.setEnumType(SegmentedFileSuffix, "livekit.SegmentedFileSuffix", [
  { no: 0, name: "INDEX" },
  { no: 1, name: "TIMESTAMP" }
]);
var ImageFileSuffix = /* @__PURE__ */ ((ImageFileSuffix2) => {
  ImageFileSuffix2[ImageFileSuffix2["IMAGE_SUFFIX_INDEX"] = 0] = "IMAGE_SUFFIX_INDEX";
  ImageFileSuffix2[ImageFileSuffix2["IMAGE_SUFFIX_TIMESTAMP"] = 1] = "IMAGE_SUFFIX_TIMESTAMP";
  return ImageFileSuffix2;
})(ImageFileSuffix || {});
import_protobuf.proto3.util.setEnumType(ImageFileSuffix, "livekit.ImageFileSuffix", [
  { no: 0, name: "IMAGE_SUFFIX_INDEX" },
  { no: 1, name: "IMAGE_SUFFIX_TIMESTAMP" }
]);
var StreamProtocol = /* @__PURE__ */ ((StreamProtocol2) => {
  StreamProtocol2[StreamProtocol2["DEFAULT_PROTOCOL"] = 0] = "DEFAULT_PROTOCOL";
  StreamProtocol2[StreamProtocol2["RTMP"] = 1] = "RTMP";
  StreamProtocol2[StreamProtocol2["SRT"] = 2] = "SRT";
  return StreamProtocol2;
})(StreamProtocol || {});
import_protobuf.proto3.util.setEnumType(StreamProtocol, "livekit.StreamProtocol", [
  { no: 0, name: "DEFAULT_PROTOCOL" },
  { no: 1, name: "RTMP" },
  { no: 2, name: "SRT" }
]);
var EncodingOptionsPreset = /* @__PURE__ */ ((EncodingOptionsPreset2) => {
  EncodingOptionsPreset2[EncodingOptionsPreset2["H264_720P_30"] = 0] = "H264_720P_30";
  EncodingOptionsPreset2[EncodingOptionsPreset2["H264_720P_60"] = 1] = "H264_720P_60";
  EncodingOptionsPreset2[EncodingOptionsPreset2["H264_1080P_30"] = 2] = "H264_1080P_30";
  EncodingOptionsPreset2[EncodingOptionsPreset2["H264_1080P_60"] = 3] = "H264_1080P_60";
  EncodingOptionsPreset2[EncodingOptionsPreset2["PORTRAIT_H264_720P_30"] = 4] = "PORTRAIT_H264_720P_30";
  EncodingOptionsPreset2[EncodingOptionsPreset2["PORTRAIT_H264_720P_60"] = 5] = "PORTRAIT_H264_720P_60";
  EncodingOptionsPreset2[EncodingOptionsPreset2["PORTRAIT_H264_1080P_30"] = 6] = "PORTRAIT_H264_1080P_30";
  EncodingOptionsPreset2[EncodingOptionsPreset2["PORTRAIT_H264_1080P_60"] = 7] = "PORTRAIT_H264_1080P_60";
  return EncodingOptionsPreset2;
})(EncodingOptionsPreset || {});
import_protobuf.proto3.util.setEnumType(EncodingOptionsPreset, "livekit.EncodingOptionsPreset", [
  { no: 0, name: "H264_720P_30" },
  { no: 1, name: "H264_720P_60" },
  { no: 2, name: "H264_1080P_30" },
  { no: 3, name: "H264_1080P_60" },
  { no: 4, name: "PORTRAIT_H264_720P_30" },
  { no: 5, name: "PORTRAIT_H264_720P_60" },
  { no: 6, name: "PORTRAIT_H264_1080P_30" },
  { no: 7, name: "PORTRAIT_H264_1080P_60" }
]);
var EgressStatus = /* @__PURE__ */ ((EgressStatus2) => {
  EgressStatus2[EgressStatus2["EGRESS_STARTING"] = 0] = "EGRESS_STARTING";
  EgressStatus2[EgressStatus2["EGRESS_ACTIVE"] = 1] = "EGRESS_ACTIVE";
  EgressStatus2[EgressStatus2["EGRESS_ENDING"] = 2] = "EGRESS_ENDING";
  EgressStatus2[EgressStatus2["EGRESS_COMPLETE"] = 3] = "EGRESS_COMPLETE";
  EgressStatus2[EgressStatus2["EGRESS_FAILED"] = 4] = "EGRESS_FAILED";
  EgressStatus2[EgressStatus2["EGRESS_ABORTED"] = 5] = "EGRESS_ABORTED";
  EgressStatus2[EgressStatus2["EGRESS_LIMIT_REACHED"] = 6] = "EGRESS_LIMIT_REACHED";
  return EgressStatus2;
})(EgressStatus || {});
import_protobuf.proto3.util.setEnumType(EgressStatus, "livekit.EgressStatus", [
  { no: 0, name: "EGRESS_STARTING" },
  { no: 1, name: "EGRESS_ACTIVE" },
  { no: 2, name: "EGRESS_ENDING" },
  { no: 3, name: "EGRESS_COMPLETE" },
  { no: 4, name: "EGRESS_FAILED" },
  { no: 5, name: "EGRESS_ABORTED" },
  { no: 6, name: "EGRESS_LIMIT_REACHED" }
]);
const _RoomCompositeEgressRequest = class _RoomCompositeEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * (optional)
     *
     * @generated from field: string layout = 2;
     */
    this.layout = "";
    /**
     * (default false)
     *
     * @generated from field: bool audio_only = 3;
     */
    this.audioOnly = false;
    /**
     * (default false)
     *
     * @generated from field: bool video_only = 4;
     */
    this.videoOnly = false;
    /**
     * template base url (default https://recorder.livekit.io)
     *
     * @generated from field: string custom_base_url = 5;
     */
    this.customBaseUrl = "";
    /**
     * deprecated (use _output fields)
     *
     * @generated from oneof livekit.RoomCompositeEgressRequest.output
     */
    this.output = { case: void 0 };
    /**
     * @generated from oneof livekit.RoomCompositeEgressRequest.options
     */
    this.options = { case: void 0 };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 11;
     */
    this.fileOutputs = [];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 12;
     */
    this.streamOutputs = [];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 13;
     */
    this.segmentOutputs = [];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 14;
     */
    this.imageOutputs = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RoomCompositeEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RoomCompositeEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RoomCompositeEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_RoomCompositeEgressRequest, a, b);
  }
};
_RoomCompositeEgressRequest.runtime = import_protobuf.proto3;
_RoomCompositeEgressRequest.typeName = "livekit.RoomCompositeEgressRequest";
_RoomCompositeEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "layout",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "audio_only",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 4,
    name: "video_only",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 5,
    name: "custom_base_url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "file", kind: "message", T: EncodedFileOutput, oneof: "output" },
  { no: 7, name: "stream", kind: "message", T: StreamOutput, oneof: "output" },
  { no: 10, name: "segments", kind: "message", T: SegmentedFileOutput, oneof: "output" },
  { no: 8, name: "preset", kind: "enum", T: import_protobuf.proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
  { no: 9, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
  { no: 11, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
  { no: 12, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
  { no: 13, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
  { no: 14, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true }
]);
let RoomCompositeEgressRequest = _RoomCompositeEgressRequest;
const _WebEgressRequest = class _WebEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string url = 1;
     */
    this.url = "";
    /**
     * @generated from field: bool audio_only = 2;
     */
    this.audioOnly = false;
    /**
     * @generated from field: bool video_only = 3;
     */
    this.videoOnly = false;
    /**
     * @generated from field: bool await_start_signal = 12;
     */
    this.awaitStartSignal = false;
    /**
     * deprecated (use _output fields)
     *
     * @generated from oneof livekit.WebEgressRequest.output
     */
    this.output = { case: void 0 };
    /**
     * @generated from oneof livekit.WebEgressRequest.options
     */
    this.options = { case: void 0 };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 9;
     */
    this.fileOutputs = [];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 10;
     */
    this.streamOutputs = [];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 11;
     */
    this.segmentOutputs = [];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 13;
     */
    this.imageOutputs = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _WebEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _WebEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _WebEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_WebEgressRequest, a, b);
  }
};
_WebEgressRequest.runtime = import_protobuf.proto3;
_WebEgressRequest.typeName = "livekit.WebEgressRequest";
_WebEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "audio_only",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 3,
    name: "video_only",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 12,
    name: "await_start_signal",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 4, name: "file", kind: "message", T: EncodedFileOutput, oneof: "output" },
  { no: 5, name: "stream", kind: "message", T: StreamOutput, oneof: "output" },
  { no: 6, name: "segments", kind: "message", T: SegmentedFileOutput, oneof: "output" },
  { no: 7, name: "preset", kind: "enum", T: import_protobuf.proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
  { no: 8, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
  { no: 9, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
  { no: 10, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
  { no: 11, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
  { no: 13, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true }
]);
let WebEgressRequest = _WebEgressRequest;
const _ParticipantEgressRequest = class _ParticipantEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * required
     *
     * @generated from field: string identity = 2;
     */
    this.identity = "";
    /**
     * (default false)
     *
     * @generated from field: bool screen_share = 3;
     */
    this.screenShare = false;
    /**
     * @generated from oneof livekit.ParticipantEgressRequest.options
     */
    this.options = { case: void 0 };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 6;
     */
    this.fileOutputs = [];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 7;
     */
    this.streamOutputs = [];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 8;
     */
    this.segmentOutputs = [];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 9;
     */
    this.imageOutputs = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ParticipantEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ParticipantEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ParticipantEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ParticipantEgressRequest, a, b);
  }
};
_ParticipantEgressRequest.runtime = import_protobuf.proto3;
_ParticipantEgressRequest.typeName = "livekit.ParticipantEgressRequest";
_ParticipantEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
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
    name: "screen_share",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 4, name: "preset", kind: "enum", T: import_protobuf.proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
  { no: 5, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
  { no: 6, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
  { no: 7, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
  { no: 8, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
  { no: 9, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true }
]);
let ParticipantEgressRequest = _ParticipantEgressRequest;
const _TrackCompositeEgressRequest = class _TrackCompositeEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * (optional)
     *
     * @generated from field: string audio_track_id = 2;
     */
    this.audioTrackId = "";
    /**
     * (optional)
     *
     * @generated from field: string video_track_id = 3;
     */
    this.videoTrackId = "";
    /**
     * deprecated (use _output fields)
     *
     * @generated from oneof livekit.TrackCompositeEgressRequest.output
     */
    this.output = { case: void 0 };
    /**
     * @generated from oneof livekit.TrackCompositeEgressRequest.options
     */
    this.options = { case: void 0 };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 11;
     */
    this.fileOutputs = [];
    /**
     * @generated from field: repeated livekit.StreamOutput stream_outputs = 12;
     */
    this.streamOutputs = [];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 13;
     */
    this.segmentOutputs = [];
    /**
     * @generated from field: repeated livekit.ImageOutput image_outputs = 14;
     */
    this.imageOutputs = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackCompositeEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackCompositeEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackCompositeEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_TrackCompositeEgressRequest, a, b);
  }
};
_TrackCompositeEgressRequest.runtime = import_protobuf.proto3;
_TrackCompositeEgressRequest.typeName = "livekit.TrackCompositeEgressRequest";
_TrackCompositeEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "audio_track_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "video_track_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "file", kind: "message", T: EncodedFileOutput, oneof: "output" },
  { no: 5, name: "stream", kind: "message", T: StreamOutput, oneof: "output" },
  { no: 8, name: "segments", kind: "message", T: SegmentedFileOutput, oneof: "output" },
  { no: 6, name: "preset", kind: "enum", T: import_protobuf.proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
  { no: 7, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
  { no: 11, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
  { no: 12, name: "stream_outputs", kind: "message", T: StreamOutput, repeated: true },
  { no: 13, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true },
  { no: 14, name: "image_outputs", kind: "message", T: ImageOutput, repeated: true }
]);
let TrackCompositeEgressRequest = _TrackCompositeEgressRequest;
const _TrackEgressRequest = class _TrackEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * required
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * required
     *
     * @generated from field: string track_id = 2;
     */
    this.trackId = "";
    /**
     * required
     *
     * @generated from oneof livekit.TrackEgressRequest.output
     */
    this.output = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TrackEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TrackEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TrackEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_TrackEgressRequest, a, b);
  }
};
_TrackEgressRequest.runtime = import_protobuf.proto3;
_TrackEgressRequest.typeName = "livekit.TrackEgressRequest";
_TrackEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "track_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "file", kind: "message", T: DirectFileOutput, oneof: "output" },
  { no: 4, name: "websocket_url", kind: "scalar", T: 9, oneof: "output" }
]);
let TrackEgressRequest = _TrackEgressRequest;
const _EncodedFileOutput = class _EncodedFileOutput extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * (optional)
     *
     * @generated from field: livekit.EncodedFileType file_type = 1;
     */
    this.fileType = 0 /* DEFAULT_FILETYPE */;
    /**
     * see egress docs for templating (default {room_name}-{time})
     *
     * @generated from field: string filepath = 2;
     */
    this.filepath = "";
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 6;
     */
    this.disableManifest = false;
    /**
     * @generated from oneof livekit.EncodedFileOutput.output
     */
    this.output = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _EncodedFileOutput().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _EncodedFileOutput().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _EncodedFileOutput().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_EncodedFileOutput, a, b);
  }
};
_EncodedFileOutput.runtime = import_protobuf.proto3;
_EncodedFileOutput.typeName = "livekit.EncodedFileOutput";
_EncodedFileOutput.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "file_type", kind: "enum", T: import_protobuf.proto3.getEnumType(EncodedFileType) },
  {
    no: 2,
    name: "filepath",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "disable_manifest",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 3, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
  { no: 4, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
  { no: 5, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
  { no: 7, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" }
]);
let EncodedFileOutput = _EncodedFileOutput;
const _SegmentedFileOutput = class _SegmentedFileOutput extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * (optional)
     *
     * @generated from field: livekit.SegmentedFileProtocol protocol = 1;
     */
    this.protocol = 0 /* DEFAULT_SEGMENTED_FILE_PROTOCOL */;
    /**
     * (optional)
     *
     * @generated from field: string filename_prefix = 2;
     */
    this.filenamePrefix = "";
    /**
     * (optional)
     *
     * @generated from field: string playlist_name = 3;
     */
    this.playlistName = "";
    /**
     * (optional, disabled if not provided). Path of a live playlist
     *
     * @generated from field: string live_playlist_name = 11;
     */
    this.livePlaylistName = "";
    /**
     * in seconds (optional)
     *
     * @generated from field: uint32 segment_duration = 4;
     */
    this.segmentDuration = 0;
    /**
     * (optional, default INDEX)
     *
     * @generated from field: livekit.SegmentedFileSuffix filename_suffix = 10;
     */
    this.filenameSuffix = 0 /* INDEX */;
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 8;
     */
    this.disableManifest = false;
    /**
     * required
     *
     * @generated from oneof livekit.SegmentedFileOutput.output
     */
    this.output = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SegmentedFileOutput().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SegmentedFileOutput().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SegmentedFileOutput().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SegmentedFileOutput, a, b);
  }
};
_SegmentedFileOutput.runtime = import_protobuf.proto3;
_SegmentedFileOutput.typeName = "livekit.SegmentedFileOutput";
_SegmentedFileOutput.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "protocol", kind: "enum", T: import_protobuf.proto3.getEnumType(SegmentedFileProtocol) },
  {
    no: 2,
    name: "filename_prefix",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "playlist_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 11,
    name: "live_playlist_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "segment_duration",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 10, name: "filename_suffix", kind: "enum", T: import_protobuf.proto3.getEnumType(SegmentedFileSuffix) },
  {
    no: 8,
    name: "disable_manifest",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 5, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
  { no: 6, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
  { no: 7, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
  { no: 9, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" }
]);
let SegmentedFileOutput = _SegmentedFileOutput;
const _DirectFileOutput = class _DirectFileOutput extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * see egress docs for templating (default {track_id}-{time})
     *
     * @generated from field: string filepath = 1;
     */
    this.filepath = "";
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 5;
     */
    this.disableManifest = false;
    /**
     * @generated from oneof livekit.DirectFileOutput.output
     */
    this.output = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DirectFileOutput().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DirectFileOutput().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DirectFileOutput().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_DirectFileOutput, a, b);
  }
};
_DirectFileOutput.runtime = import_protobuf.proto3;
_DirectFileOutput.typeName = "livekit.DirectFileOutput";
_DirectFileOutput.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "filepath",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "disable_manifest",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 2, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
  { no: 3, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
  { no: 4, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
  { no: 6, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" }
]);
let DirectFileOutput = _DirectFileOutput;
const _ImageOutput = class _ImageOutput extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * in seconds (required)
     *
     * @generated from field: uint32 capture_interval = 1;
     */
    this.captureInterval = 0;
    /**
     * (optional, defaults to track width)
     *
     * @generated from field: int32 width = 2;
     */
    this.width = 0;
    /**
     * (optional, defaults to track height)
     *
     * @generated from field: int32 height = 3;
     */
    this.height = 0;
    /**
     * (optional)
     *
     * @generated from field: string filename_prefix = 4;
     */
    this.filenamePrefix = "";
    /**
     * (optional, default INDEX)
     *
     * @generated from field: livekit.ImageFileSuffix filename_suffix = 5;
     */
    this.filenameSuffix = 0 /* IMAGE_SUFFIX_INDEX */;
    /**
     * (optional)
     *
     * @generated from field: livekit.ImageCodec image_codec = 6;
     */
    this.imageCodec = import_livekit_models_pb.ImageCodec.IC_DEFAULT;
    /**
     * disable upload of manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 7;
     */
    this.disableManifest = false;
    /**
     * required
     *
     * @generated from oneof livekit.ImageOutput.output
     */
    this.output = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ImageOutput().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ImageOutput().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ImageOutput().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ImageOutput, a, b);
  }
};
_ImageOutput.runtime = import_protobuf.proto3;
_ImageOutput.typeName = "livekit.ImageOutput";
_ImageOutput.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "capture_interval",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 2,
    name: "width",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 3,
    name: "height",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 4,
    name: "filename_prefix",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "filename_suffix", kind: "enum", T: import_protobuf.proto3.getEnumType(ImageFileSuffix) },
  { no: 6, name: "image_codec", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.ImageCodec) },
  {
    no: 7,
    name: "disable_manifest",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 8, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
  { no: 9, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
  { no: 10, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
  { no: 11, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" }
]);
let ImageOutput = _ImageOutput;
const _S3Upload = class _S3Upload extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string access_key = 1;
     */
    this.accessKey = "";
    /**
     * @generated from field: string secret = 2;
     */
    this.secret = "";
    /**
     * @generated from field: string session_token = 11;
     */
    this.sessionToken = "";
    /**
     * @generated from field: string region = 3;
     */
    this.region = "";
    /**
     * @generated from field: string endpoint = 4;
     */
    this.endpoint = "";
    /**
     * @generated from field: string bucket = 5;
     */
    this.bucket = "";
    /**
     * @generated from field: bool force_path_style = 6;
     */
    this.forcePathStyle = false;
    /**
     * @generated from field: map<string, string> metadata = 7;
     */
    this.metadata = {};
    /**
     * @generated from field: string tagging = 8;
     */
    this.tagging = "";
    /**
     * Content-Disposition header
     *
     * @generated from field: string content_disposition = 9;
     */
    this.contentDisposition = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _S3Upload().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _S3Upload().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _S3Upload().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_S3Upload, a, b);
  }
};
_S3Upload.runtime = import_protobuf.proto3;
_S3Upload.typeName = "livekit.S3Upload";
_S3Upload.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "access_key",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "secret",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 11,
    name: "session_token",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "endpoint",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "bucket",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "force_path_style",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 7, name: "metadata", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  {
    no: 8,
    name: "tagging",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "content_disposition",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 10, name: "proxy", kind: "message", T: ProxyConfig }
]);
let S3Upload = _S3Upload;
const _GCPUpload = class _GCPUpload extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * service account credentials serialized in JSON "credentials.json"
     *
     * @generated from field: string credentials = 1;
     */
    this.credentials = "";
    /**
     * @generated from field: string bucket = 2;
     */
    this.bucket = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GCPUpload().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GCPUpload().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GCPUpload().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_GCPUpload, a, b);
  }
};
_GCPUpload.runtime = import_protobuf.proto3;
_GCPUpload.typeName = "livekit.GCPUpload";
_GCPUpload.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "credentials",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "bucket",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "proxy", kind: "message", T: ProxyConfig }
]);
let GCPUpload = _GCPUpload;
const _AzureBlobUpload = class _AzureBlobUpload extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string account_name = 1;
     */
    this.accountName = "";
    /**
     * @generated from field: string account_key = 2;
     */
    this.accountKey = "";
    /**
     * @generated from field: string container_name = 3;
     */
    this.containerName = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AzureBlobUpload().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AzureBlobUpload().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AzureBlobUpload().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AzureBlobUpload, a, b);
  }
};
_AzureBlobUpload.runtime = import_protobuf.proto3;
_AzureBlobUpload.typeName = "livekit.AzureBlobUpload";
_AzureBlobUpload.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "account_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "account_key",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "container_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let AzureBlobUpload = _AzureBlobUpload;
const _AliOSSUpload = class _AliOSSUpload extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string access_key = 1;
     */
    this.accessKey = "";
    /**
     * @generated from field: string secret = 2;
     */
    this.secret = "";
    /**
     * @generated from field: string region = 3;
     */
    this.region = "";
    /**
     * @generated from field: string endpoint = 4;
     */
    this.endpoint = "";
    /**
     * @generated from field: string bucket = 5;
     */
    this.bucket = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AliOSSUpload().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AliOSSUpload().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AliOSSUpload().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AliOSSUpload, a, b);
  }
};
_AliOSSUpload.runtime = import_protobuf.proto3;
_AliOSSUpload.typeName = "livekit.AliOSSUpload";
_AliOSSUpload.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "access_key",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "secret",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "region",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "endpoint",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "bucket",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let AliOSSUpload = _AliOSSUpload;
const _ProxyConfig = class _ProxyConfig extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string url = 1;
     */
    this.url = "";
    /**
     * @generated from field: string username = 2;
     */
    this.username = "";
    /**
     * @generated from field: string password = 3;
     */
    this.password = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ProxyConfig().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ProxyConfig().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ProxyConfig().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ProxyConfig, a, b);
  }
};
_ProxyConfig.runtime = import_protobuf.proto3;
_ProxyConfig.typeName = "livekit.ProxyConfig";
_ProxyConfig.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let ProxyConfig = _ProxyConfig;
const _StreamOutput = class _StreamOutput extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * required
     *
     * @generated from field: livekit.StreamProtocol protocol = 1;
     */
    this.protocol = 0 /* DEFAULT_PROTOCOL */;
    /**
     * required
     *
     * @generated from field: repeated string urls = 2;
     */
    this.urls = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StreamOutput().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StreamOutput().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StreamOutput().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_StreamOutput, a, b);
  }
};
_StreamOutput.runtime = import_protobuf.proto3;
_StreamOutput.typeName = "livekit.StreamOutput";
_StreamOutput.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "protocol", kind: "enum", T: import_protobuf.proto3.getEnumType(StreamProtocol) },
  { no: 2, name: "urls", kind: "scalar", T: 9, repeated: true }
]);
let StreamOutput = _StreamOutput;
const _EncodingOptions = class _EncodingOptions extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * (default 1920)
     *
     * @generated from field: int32 width = 1;
     */
    this.width = 0;
    /**
     * (default 1080)
     *
     * @generated from field: int32 height = 2;
     */
    this.height = 0;
    /**
     * (default 24)
     *
     * @generated from field: int32 depth = 3;
     */
    this.depth = 0;
    /**
     * (default 30)
     *
     * @generated from field: int32 framerate = 4;
     */
    this.framerate = 0;
    /**
     * (default OPUS)
     *
     * @generated from field: livekit.AudioCodec audio_codec = 5;
     */
    this.audioCodec = import_livekit_models_pb.AudioCodec.DEFAULT_AC;
    /**
     * (default 128)
     *
     * @generated from field: int32 audio_bitrate = 6;
     */
    this.audioBitrate = 0;
    /**
     * quality setting on audio encoder
     *
     * @generated from field: int32 audio_quality = 11;
     */
    this.audioQuality = 0;
    /**
     * (default 44100)
     *
     * @generated from field: int32 audio_frequency = 7;
     */
    this.audioFrequency = 0;
    /**
     * (default H264_MAIN)
     *
     * @generated from field: livekit.VideoCodec video_codec = 8;
     */
    this.videoCodec = import_livekit_models_pb.VideoCodec.DEFAULT_VC;
    /**
     * (default 4500)
     *
     * @generated from field: int32 video_bitrate = 9;
     */
    this.videoBitrate = 0;
    /**
     * quality setting on video encoder
     *
     * @generated from field: int32 video_quality = 12;
     */
    this.videoQuality = 0;
    /**
     * in seconds (default 4s for streaming, segment duration for segmented output, encoder default for files)
     *
     * @generated from field: double key_frame_interval = 10;
     */
    this.keyFrameInterval = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _EncodingOptions().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _EncodingOptions().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _EncodingOptions().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_EncodingOptions, a, b);
  }
};
_EncodingOptions.runtime = import_protobuf.proto3;
_EncodingOptions.typeName = "livekit.EncodingOptions";
_EncodingOptions.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "width",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 2,
    name: "height",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 3,
    name: "depth",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 4,
    name: "framerate",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 5, name: "audio_codec", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.AudioCodec) },
  {
    no: 6,
    name: "audio_bitrate",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 11,
    name: "audio_quality",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 7,
    name: "audio_frequency",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 8, name: "video_codec", kind: "enum", T: import_protobuf.proto3.getEnumType(import_livekit_models_pb.VideoCodec) },
  {
    no: 9,
    name: "video_bitrate",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 12,
    name: "video_quality",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 10,
    name: "key_frame_interval",
    kind: "scalar",
    T: 1
    /* ScalarType.DOUBLE */
  }
]);
let EncodingOptions = _EncodingOptions;
const _UpdateLayoutRequest = class _UpdateLayoutRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string egress_id = 1;
     */
    this.egressId = "";
    /**
     * @generated from field: string layout = 2;
     */
    this.layout = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateLayoutRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateLayoutRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateLayoutRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UpdateLayoutRequest, a, b);
  }
};
_UpdateLayoutRequest.runtime = import_protobuf.proto3;
_UpdateLayoutRequest.typeName = "livekit.UpdateLayoutRequest";
_UpdateLayoutRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "egress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "layout",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let UpdateLayoutRequest = _UpdateLayoutRequest;
const _UpdateStreamRequest = class _UpdateStreamRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string egress_id = 1;
     */
    this.egressId = "";
    /**
     * @generated from field: repeated string add_output_urls = 2;
     */
    this.addOutputUrls = [];
    /**
     * @generated from field: repeated string remove_output_urls = 3;
     */
    this.removeOutputUrls = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateStreamRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateStreamRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateStreamRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UpdateStreamRequest, a, b);
  }
};
_UpdateStreamRequest.runtime = import_protobuf.proto3;
_UpdateStreamRequest.typeName = "livekit.UpdateStreamRequest";
_UpdateStreamRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "egress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "add_output_urls", kind: "scalar", T: 9, repeated: true },
  { no: 3, name: "remove_output_urls", kind: "scalar", T: 9, repeated: true }
]);
let UpdateStreamRequest = _UpdateStreamRequest;
const _ListEgressRequest = class _ListEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * (optional, filter by room name)
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * (optional, filter by egress ID)
     *
     * @generated from field: string egress_id = 2;
     */
    this.egressId = "";
    /**
     * (optional, list active egress only)
     *
     * @generated from field: bool active = 3;
     */
    this.active = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ListEgressRequest, a, b);
  }
};
_ListEgressRequest.runtime = import_protobuf.proto3;
_ListEgressRequest.typeName = "livekit.ListEgressRequest";
_ListEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "egress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "active",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let ListEgressRequest = _ListEgressRequest;
const _ListEgressResponse = class _ListEgressResponse extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.EgressInfo items = 1;
     */
    this.items = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListEgressResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListEgressResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListEgressResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ListEgressResponse, a, b);
  }
};
_ListEgressResponse.runtime = import_protobuf.proto3;
_ListEgressResponse.typeName = "livekit.ListEgressResponse";
_ListEgressResponse.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "items", kind: "message", T: EgressInfo, repeated: true }
]);
let ListEgressResponse = _ListEgressResponse;
const _StopEgressRequest = class _StopEgressRequest extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string egress_id = 1;
     */
    this.egressId = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StopEgressRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StopEgressRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StopEgressRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_StopEgressRequest, a, b);
  }
};
_StopEgressRequest.runtime = import_protobuf.proto3;
_StopEgressRequest.typeName = "livekit.StopEgressRequest";
_StopEgressRequest.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "egress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let StopEgressRequest = _StopEgressRequest;
const _EgressInfo = class _EgressInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string egress_id = 1;
     */
    this.egressId = "";
    /**
     * @generated from field: string room_id = 2;
     */
    this.roomId = "";
    /**
     * @generated from field: string room_name = 13;
     */
    this.roomName = "";
    /**
     * @generated from field: livekit.EgressStatus status = 3;
     */
    this.status = 0 /* EGRESS_STARTING */;
    /**
     * @generated from field: int64 started_at = 10;
     */
    this.startedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 11;
     */
    this.endedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 updated_at = 18;
     */
    this.updatedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string details = 21;
     */
    this.details = "";
    /**
     * @generated from field: string error = 9;
     */
    this.error = "";
    /**
     * @generated from field: int32 error_code = 22;
     */
    this.errorCode = 0;
    /**
     * @generated from oneof livekit.EgressInfo.request
     */
    this.request = { case: void 0 };
    /**
     * deprecated (use _result fields)
     *
     * @generated from oneof livekit.EgressInfo.result
     */
    this.result = { case: void 0 };
    /**
     * @generated from field: repeated livekit.StreamInfo stream_results = 15;
     */
    this.streamResults = [];
    /**
     * @generated from field: repeated livekit.FileInfo file_results = 16;
     */
    this.fileResults = [];
    /**
     * @generated from field: repeated livekit.SegmentsInfo segment_results = 17;
     */
    this.segmentResults = [];
    /**
     * @generated from field: repeated livekit.ImagesInfo image_results = 20;
     */
    this.imageResults = [];
    /**
     * @generated from field: string manifest_location = 23;
     */
    this.manifestLocation = "";
    /**
     * next ID: 26
     *
     * @generated from field: bool backup_storage_used = 25;
     */
    this.backupStorageUsed = false;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _EgressInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _EgressInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _EgressInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_EgressInfo, a, b);
  }
};
_EgressInfo.runtime = import_protobuf.proto3;
_EgressInfo.typeName = "livekit.EgressInfo";
_EgressInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "egress_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "room_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 13,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "status", kind: "enum", T: import_protobuf.proto3.getEnumType(EgressStatus) },
  {
    no: 10,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 11,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 18,
    name: "updated_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 21,
    name: "details",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 22,
    name: "error_code",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 4, name: "room_composite", kind: "message", T: RoomCompositeEgressRequest, oneof: "request" },
  { no: 14, name: "web", kind: "message", T: WebEgressRequest, oneof: "request" },
  { no: 19, name: "participant", kind: "message", T: ParticipantEgressRequest, oneof: "request" },
  { no: 5, name: "track_composite", kind: "message", T: TrackCompositeEgressRequest, oneof: "request" },
  { no: 6, name: "track", kind: "message", T: TrackEgressRequest, oneof: "request" },
  { no: 7, name: "stream", kind: "message", T: StreamInfoList, oneof: "result" },
  { no: 8, name: "file", kind: "message", T: FileInfo, oneof: "result" },
  { no: 12, name: "segments", kind: "message", T: SegmentsInfo, oneof: "result" },
  { no: 15, name: "stream_results", kind: "message", T: StreamInfo, repeated: true },
  { no: 16, name: "file_results", kind: "message", T: FileInfo, repeated: true },
  { no: 17, name: "segment_results", kind: "message", T: SegmentsInfo, repeated: true },
  { no: 20, name: "image_results", kind: "message", T: ImagesInfo, repeated: true },
  {
    no: 23,
    name: "manifest_location",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 25,
    name: "backup_storage_used",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let EgressInfo = _EgressInfo;
const _StreamInfoList = class _StreamInfoList extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.StreamInfo info = 1;
     */
    this.info = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StreamInfoList().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StreamInfoList().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StreamInfoList().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_StreamInfoList, a, b);
  }
};
_StreamInfoList.runtime = import_protobuf.proto3;
_StreamInfoList.typeName = "livekit.StreamInfoList";
_StreamInfoList.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "info", kind: "message", T: StreamInfo, repeated: true }
]);
let StreamInfoList = _StreamInfoList;
const _StreamInfo = class _StreamInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string url = 1;
     */
    this.url = "";
    /**
     * @generated from field: int64 started_at = 2;
     */
    this.startedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 3;
     */
    this.endedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 duration = 4;
     */
    this.duration = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: livekit.StreamInfo.Status status = 5;
     */
    this.status = 0 /* ACTIVE */;
    /**
     * @generated from field: string error = 6;
     */
    this.error = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StreamInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StreamInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StreamInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_StreamInfo, a, b);
  }
};
_StreamInfo.runtime = import_protobuf.proto3;
_StreamInfo.typeName = "livekit.StreamInfo";
_StreamInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 4,
    name: "duration",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  { no: 5, name: "status", kind: "enum", T: import_protobuf.proto3.getEnumType(StreamInfo_Status) },
  {
    no: 6,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let StreamInfo = _StreamInfo;
var StreamInfo_Status = /* @__PURE__ */ ((StreamInfo_Status2) => {
  StreamInfo_Status2[StreamInfo_Status2["ACTIVE"] = 0] = "ACTIVE";
  StreamInfo_Status2[StreamInfo_Status2["FINISHED"] = 1] = "FINISHED";
  StreamInfo_Status2[StreamInfo_Status2["FAILED"] = 2] = "FAILED";
  return StreamInfo_Status2;
})(StreamInfo_Status || {});
import_protobuf.proto3.util.setEnumType(StreamInfo_Status, "livekit.StreamInfo.Status", [
  { no: 0, name: "ACTIVE" },
  { no: 1, name: "FINISHED" },
  { no: 2, name: "FAILED" }
]);
const _FileInfo = class _FileInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string filename = 1;
     */
    this.filename = "";
    /**
     * @generated from field: int64 started_at = 2;
     */
    this.startedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 3;
     */
    this.endedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 duration = 6;
     */
    this.duration = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 size = 4;
     */
    this.size = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string location = 5;
     */
    this.location = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _FileInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _FileInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _FileInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_FileInfo, a, b);
  }
};
_FileInfo.runtime = import_protobuf.proto3;
_FileInfo.typeName = "livekit.FileInfo";
_FileInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "filename",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 6,
    name: "duration",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 4,
    name: "size",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 5,
    name: "location",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let FileInfo = _FileInfo;
const _SegmentsInfo = class _SegmentsInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string playlist_name = 1;
     */
    this.playlistName = "";
    /**
     * @generated from field: string live_playlist_name = 8;
     */
    this.livePlaylistName = "";
    /**
     * @generated from field: int64 duration = 2;
     */
    this.duration = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 size = 3;
     */
    this.size = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string playlist_location = 4;
     */
    this.playlistLocation = "";
    /**
     * @generated from field: string live_playlist_location = 9;
     */
    this.livePlaylistLocation = "";
    /**
     * @generated from field: int64 segment_count = 5;
     */
    this.segmentCount = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 started_at = 6;
     */
    this.startedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 7;
     */
    this.endedAt = import_protobuf.protoInt64.zero;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SegmentsInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SegmentsInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SegmentsInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_SegmentsInfo, a, b);
  }
};
_SegmentsInfo.runtime = import_protobuf.proto3;
_SegmentsInfo.typeName = "livekit.SegmentsInfo";
_SegmentsInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "playlist_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "live_playlist_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "duration",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "size",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 4,
    name: "playlist_location",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "live_playlist_location",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "segment_count",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 6,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 7,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let SegmentsInfo = _SegmentsInfo;
const _ImagesInfo = class _ImagesInfo extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string filename_prefix = 4;
     */
    this.filenamePrefix = "";
    /**
     * @generated from field: int64 image_count = 1;
     */
    this.imageCount = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 started_at = 2;
     */
    this.startedAt = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 3;
     */
    this.endedAt = import_protobuf.protoInt64.zero;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ImagesInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ImagesInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ImagesInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ImagesInfo, a, b);
  }
};
_ImagesInfo.runtime = import_protobuf.proto3;
_ImagesInfo.typeName = "livekit.ImagesInfo";
_ImagesInfo.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 4,
    name: "filename_prefix",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 1,
    name: "image_count",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 3,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let ImagesInfo = _ImagesInfo;
const _AutoParticipantEgress = class _AutoParticipantEgress extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.AutoParticipantEgress.options
     */
    this.options = { case: void 0 };
    /**
     * @generated from field: repeated livekit.EncodedFileOutput file_outputs = 3;
     */
    this.fileOutputs = [];
    /**
     * @generated from field: repeated livekit.SegmentedFileOutput segment_outputs = 4;
     */
    this.segmentOutputs = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AutoParticipantEgress().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AutoParticipantEgress().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AutoParticipantEgress().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AutoParticipantEgress, a, b);
  }
};
_AutoParticipantEgress.runtime = import_protobuf.proto3;
_AutoParticipantEgress.typeName = "livekit.AutoParticipantEgress";
_AutoParticipantEgress.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "preset", kind: "enum", T: import_protobuf.proto3.getEnumType(EncodingOptionsPreset), oneof: "options" },
  { no: 2, name: "advanced", kind: "message", T: EncodingOptions, oneof: "options" },
  { no: 3, name: "file_outputs", kind: "message", T: EncodedFileOutput, repeated: true },
  { no: 4, name: "segment_outputs", kind: "message", T: SegmentedFileOutput, repeated: true }
]);
let AutoParticipantEgress = _AutoParticipantEgress;
const _AutoTrackEgress = class _AutoTrackEgress extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * see docs for templating (default {track_id}-{time})
     *
     * @generated from field: string filepath = 1;
     */
    this.filepath = "";
    /**
     * disables upload of json manifest file (default false)
     *
     * @generated from field: bool disable_manifest = 5;
     */
    this.disableManifest = false;
    /**
     * @generated from oneof livekit.AutoTrackEgress.output
     */
    this.output = { case: void 0 };
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AutoTrackEgress().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AutoTrackEgress().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AutoTrackEgress().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_AutoTrackEgress, a, b);
  }
};
_AutoTrackEgress.runtime = import_protobuf.proto3;
_AutoTrackEgress.typeName = "livekit.AutoTrackEgress";
_AutoTrackEgress.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "filepath",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "disable_manifest",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 2, name: "s3", kind: "message", T: S3Upload, oneof: "output" },
  { no: 3, name: "gcp", kind: "message", T: GCPUpload, oneof: "output" },
  { no: 4, name: "azure", kind: "message", T: AzureBlobUpload, oneof: "output" },
  { no: 6, name: "aliOSS", kind: "message", T: AliOSSUpload, oneof: "output" }
]);
let AutoTrackEgress = _AutoTrackEgress;
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  AliOSSUpload,
  AutoParticipantEgress,
  AutoTrackEgress,
  AzureBlobUpload,
  DirectFileOutput,
  EgressInfo,
  EgressStatus,
  EncodedFileOutput,
  EncodedFileType,
  EncodingOptions,
  EncodingOptionsPreset,
  FileInfo,
  GCPUpload,
  ImageFileSuffix,
  ImageOutput,
  ImagesInfo,
  ListEgressRequest,
  ListEgressResponse,
  ParticipantEgressRequest,
  ProxyConfig,
  RoomCompositeEgressRequest,
  S3Upload,
  SegmentedFileOutput,
  SegmentedFileProtocol,
  SegmentedFileSuffix,
  SegmentsInfo,
  StopEgressRequest,
  StreamInfo,
  StreamInfoList,
  StreamInfo_Status,
  StreamOutput,
  StreamProtocol,
  TrackCompositeEgressRequest,
  TrackEgressRequest,
  UpdateLayoutRequest,
  UpdateStreamRequest,
  WebEgressRequest
});
//# sourceMappingURL=livekit_egress_pb.cjs.map