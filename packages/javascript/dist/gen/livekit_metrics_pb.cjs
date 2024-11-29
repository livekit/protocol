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
var livekit_metrics_pb_exports = {};
__export(livekit_metrics_pb_exports, {
  EventMetric: () => EventMetric,
  MetricLabel: () => MetricLabel,
  MetricSample: () => MetricSample,
  MetricsBatch: () => MetricsBatch,
  TimeSeriesMetric: () => TimeSeriesMetric
});
module.exports = __toCommonJS(livekit_metrics_pb_exports);
var import_protobuf = require("@bufbuild/protobuf");
var MetricLabel = /* @__PURE__ */ ((MetricLabel2) => {
  MetricLabel2[MetricLabel2["AGENTS_LLM_TTFT"] = 0] = "AGENTS_LLM_TTFT";
  MetricLabel2[MetricLabel2["AGENTS_STT_TTFT"] = 1] = "AGENTS_STT_TTFT";
  MetricLabel2[MetricLabel2["AGENTS_TTS_TTFB"] = 2] = "AGENTS_TTS_TTFB";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_SUBSCRIBER_FREEZE_COUNT"] = 3] = "CLIENT_VIDEO_SUBSCRIBER_FREEZE_COUNT";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_SUBSCRIBER_TOTAL_FREEZE_DURATION"] = 4] = "CLIENT_VIDEO_SUBSCRIBER_TOTAL_FREEZE_DURATION";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_SUBSCRIBER_PAUSE_COUNT"] = 5] = "CLIENT_VIDEO_SUBSCRIBER_PAUSE_COUNT";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_SUBSCRIBER_TOTAL_PAUSES_DURATION"] = 6] = "CLIENT_VIDEO_SUBSCRIBER_TOTAL_PAUSES_DURATION";
  MetricLabel2[MetricLabel2["CLIENT_AUDIO_SUBSCRIBER_CONCEALED_SAMPLES"] = 7] = "CLIENT_AUDIO_SUBSCRIBER_CONCEALED_SAMPLES";
  MetricLabel2[MetricLabel2["CLIENT_AUDIO_SUBSCRIBER_SILENT_CONCEALED_SAMPLES"] = 8] = "CLIENT_AUDIO_SUBSCRIBER_SILENT_CONCEALED_SAMPLES";
  MetricLabel2[MetricLabel2["CLIENT_AUDIO_SUBSCRIBER_CONCEALMENT_EVENTS"] = 9] = "CLIENT_AUDIO_SUBSCRIBER_CONCEALMENT_EVENTS";
  MetricLabel2[MetricLabel2["CLIENT_AUDIO_SUBSCRIBER_INTERRUPTION_COUNT"] = 10] = "CLIENT_AUDIO_SUBSCRIBER_INTERRUPTION_COUNT";
  MetricLabel2[MetricLabel2["CLIENT_AUDIO_SUBSCRIBER_TOTAL_INTERRUPTION_DURATION"] = 11] = "CLIENT_AUDIO_SUBSCRIBER_TOTAL_INTERRUPTION_DURATION";
  MetricLabel2[MetricLabel2["CLIENT_SUBSCRIBER_JITTER_BUFFER_DELAY"] = 12] = "CLIENT_SUBSCRIBER_JITTER_BUFFER_DELAY";
  MetricLabel2[MetricLabel2["CLIENT_SUBSCRIBER_JITTER_BUFFER_EMITTED_COUNT"] = 13] = "CLIENT_SUBSCRIBER_JITTER_BUFFER_EMITTED_COUNT";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_BANDWIDTH"] = 14] = "CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_BANDWIDTH";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_CPU"] = 15] = "CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_CPU";
  MetricLabel2[MetricLabel2["CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_OTHER"] = 16] = "CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_OTHER";
  MetricLabel2[MetricLabel2["PUBLISHER_RTT"] = 17] = "PUBLISHER_RTT";
  MetricLabel2[MetricLabel2["SERVER_MESH_RTT"] = 18] = "SERVER_MESH_RTT";
  MetricLabel2[MetricLabel2["SUBSCRIBER_RTT"] = 19] = "SUBSCRIBER_RTT";
  MetricLabel2[MetricLabel2["METRIC_LABEL_PREDEFINED_MAX_VALUE"] = 4096] = "METRIC_LABEL_PREDEFINED_MAX_VALUE";
  return MetricLabel2;
})(MetricLabel || {});
import_protobuf.proto3.util.setEnumType(MetricLabel, "livekit.MetricLabel", [
  { no: 0, name: "AGENTS_LLM_TTFT" },
  { no: 1, name: "AGENTS_STT_TTFT" },
  { no: 2, name: "AGENTS_TTS_TTFB" },
  { no: 3, name: "CLIENT_VIDEO_SUBSCRIBER_FREEZE_COUNT" },
  { no: 4, name: "CLIENT_VIDEO_SUBSCRIBER_TOTAL_FREEZE_DURATION" },
  { no: 5, name: "CLIENT_VIDEO_SUBSCRIBER_PAUSE_COUNT" },
  { no: 6, name: "CLIENT_VIDEO_SUBSCRIBER_TOTAL_PAUSES_DURATION" },
  { no: 7, name: "CLIENT_AUDIO_SUBSCRIBER_CONCEALED_SAMPLES" },
  { no: 8, name: "CLIENT_AUDIO_SUBSCRIBER_SILENT_CONCEALED_SAMPLES" },
  { no: 9, name: "CLIENT_AUDIO_SUBSCRIBER_CONCEALMENT_EVENTS" },
  { no: 10, name: "CLIENT_AUDIO_SUBSCRIBER_INTERRUPTION_COUNT" },
  { no: 11, name: "CLIENT_AUDIO_SUBSCRIBER_TOTAL_INTERRUPTION_DURATION" },
  { no: 12, name: "CLIENT_SUBSCRIBER_JITTER_BUFFER_DELAY" },
  { no: 13, name: "CLIENT_SUBSCRIBER_JITTER_BUFFER_EMITTED_COUNT" },
  { no: 14, name: "CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_BANDWIDTH" },
  { no: 15, name: "CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_CPU" },
  { no: 16, name: "CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_OTHER" },
  { no: 17, name: "PUBLISHER_RTT" },
  { no: 18, name: "SERVER_MESH_RTT" },
  { no: 19, name: "SUBSCRIBER_RTT" },
  { no: 4096, name: "METRIC_LABEL_PREDEFINED_MAX_VALUE" }
]);
const _MetricsBatch = class _MetricsBatch extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * time at which this batch is sent based on a monotonic clock (millisecond resolution)
     *
     * @generated from field: int64 timestamp_ms = 1;
     */
    this.timestampMs = import_protobuf.protoInt64.zero;
    /**
     * To avoid repeating string values, we store them in a separate list and reference them by index
     * This is useful for storing participant identities, track names, etc.
     * There is also a predefined list of labels that can be used to reference common metrics.
     * They have reserved indices from 0 to (METRIC_LABEL_PREDEFINED_MAX_VALUE - 1).
     * Indexes pointing at str_data should start from METRIC_LABEL_PREDEFINED_MAX_VALUE, 
     * such that str_data[0] == index of METRIC_LABEL_PREDEFINED_MAX_VALUE.
     *
     * @generated from field: repeated string str_data = 3;
     */
    this.strData = [];
    /**
     * @generated from field: repeated livekit.TimeSeriesMetric time_series = 4;
     */
    this.timeSeries = [];
    /**
     * @generated from field: repeated livekit.EventMetric events = 5;
     */
    this.events = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _MetricsBatch().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _MetricsBatch().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _MetricsBatch().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_MetricsBatch, a, b);
  }
};
_MetricsBatch.runtime = import_protobuf.proto3;
_MetricsBatch.typeName = "livekit.MetricsBatch";
_MetricsBatch.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "timestamp_ms",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  { no: 2, name: "normalized_timestamp", kind: "message", T: import_protobuf.Timestamp },
  { no: 3, name: "str_data", kind: "scalar", T: 9, repeated: true },
  { no: 4, name: "time_series", kind: "message", T: TimeSeriesMetric, repeated: true },
  { no: 5, name: "events", kind: "message", T: EventMetric, repeated: true }
]);
let MetricsBatch = _MetricsBatch;
const _TimeSeriesMetric = class _TimeSeriesMetric extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * Metric name e.g "speech_probablity". The string value is not directly stored in the message, but referenced by index
     * in the `str_data` field of `MetricsBatch`
     *
     * @generated from field: uint32 label = 1;
     */
    this.label = 0;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 participant_identity = 2;
     */
    this.participantIdentity = 0;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 track_sid = 3;
     */
    this.trackSid = 0;
    /**
     * @generated from field: repeated livekit.MetricSample samples = 4;
     */
    this.samples = [];
    /**
     * index into 'str_data'
     *
     * @generated from field: uint32 rid = 5;
     */
    this.rid = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TimeSeriesMetric().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TimeSeriesMetric().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TimeSeriesMetric().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_TimeSeriesMetric, a, b);
  }
};
_TimeSeriesMetric.runtime = import_protobuf.proto3;
_TimeSeriesMetric.typeName = "livekit.TimeSeriesMetric";
_TimeSeriesMetric.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "label",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 2,
    name: "participant_identity",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "track_sid",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 4, name: "samples", kind: "message", T: MetricSample, repeated: true },
  {
    no: 5,
    name: "rid",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let TimeSeriesMetric = _TimeSeriesMetric;
const _MetricSample = class _MetricSample extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * time of metric based on a monotonic clock (in milliseconds)
     *
     * @generated from field: int64 timestamp_ms = 1;
     */
    this.timestampMs = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: float value = 3;
     */
    this.value = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _MetricSample().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _MetricSample().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _MetricSample().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_MetricSample, a, b);
  }
};
_MetricSample.runtime = import_protobuf.proto3;
_MetricSample.typeName = "livekit.MetricSample";
_MetricSample.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "timestamp_ms",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  { no: 2, name: "normalized_timestamp", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 3,
    name: "value",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  }
]);
let MetricSample = _MetricSample;
const _EventMetric = class _EventMetric extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: uint32 label = 1;
     */
    this.label = 0;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 participant_identity = 2;
     */
    this.participantIdentity = 0;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 track_sid = 3;
     */
    this.trackSid = 0;
    /**
     * start time of event based on a monotonic clock (in milliseconds)
     *
     * @generated from field: int64 start_timestamp_ms = 4;
     */
    this.startTimestampMs = import_protobuf.protoInt64.zero;
    /**
     * @generated from field: string metadata = 8;
     */
    this.metadata = "";
    /**
     * index into 'str_data'
     *
     * @generated from field: uint32 rid = 9;
     */
    this.rid = 0;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _EventMetric().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _EventMetric().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _EventMetric().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_EventMetric, a, b);
  }
};
_EventMetric.runtime = import_protobuf.proto3;
_EventMetric.typeName = "livekit.EventMetric";
_EventMetric.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "label",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 2,
    name: "participant_identity",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "track_sid",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 4,
    name: "start_timestamp_ms",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  { no: 5, name: "end_timestamp_ms", kind: "scalar", T: 3, opt: true },
  { no: 6, name: "normalized_start_timestamp", kind: "message", T: import_protobuf.Timestamp },
  { no: 7, name: "normalized_end_timestamp", kind: "message", T: import_protobuf.Timestamp, opt: true },
  {
    no: 8,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "rid",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let EventMetric = _EventMetric;
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  EventMetric,
  MetricLabel,
  MetricSample,
  MetricsBatch,
  TimeSeriesMetric
});
//# sourceMappingURL=livekit_metrics_pb.cjs.map