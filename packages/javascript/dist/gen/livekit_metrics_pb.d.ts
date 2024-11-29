import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
/**
 * index from [0: MAX_LABEL_PREDEFINED_MAX_VALUE) are for predefined labels (`MetricLabel`)
 *
 * @generated from enum livekit.MetricLabel
 */
export declare enum MetricLabel {
    /**
     * time to first token from LLM
     *
     * @generated from enum value: AGENTS_LLM_TTFT = 0;
     */
    AGENTS_LLM_TTFT = 0,
    /**
     * time to final transcription
     *
     * @generated from enum value: AGENTS_STT_TTFT = 1;
     */
    AGENTS_STT_TTFT = 1,
    /**
     * time to first byte
     *
     * @generated from enum value: AGENTS_TTS_TTFB = 2;
     */
    AGENTS_TTS_TTFB = 2,
    /**
     * Number of video freezes
     *
     * @generated from enum value: CLIENT_VIDEO_SUBSCRIBER_FREEZE_COUNT = 3;
     */
    CLIENT_VIDEO_SUBSCRIBER_FREEZE_COUNT = 3,
    /**
     * total duration of freezes
     *
     * @generated from enum value: CLIENT_VIDEO_SUBSCRIBER_TOTAL_FREEZE_DURATION = 4;
     */
    CLIENT_VIDEO_SUBSCRIBER_TOTAL_FREEZE_DURATION = 4,
    /**
     * number of video pauses
     *
     * @generated from enum value: CLIENT_VIDEO_SUBSCRIBER_PAUSE_COUNT = 5;
     */
    CLIENT_VIDEO_SUBSCRIBER_PAUSE_COUNT = 5,
    /**
     * total duration of pauses
     *
     * @generated from enum value: CLIENT_VIDEO_SUBSCRIBER_TOTAL_PAUSES_DURATION = 6;
     */
    CLIENT_VIDEO_SUBSCRIBER_TOTAL_PAUSES_DURATION = 6,
    /**
     * number of concealed (synthesized) audio samples
     *
     * @generated from enum value: CLIENT_AUDIO_SUBSCRIBER_CONCEALED_SAMPLES = 7;
     */
    CLIENT_AUDIO_SUBSCRIBER_CONCEALED_SAMPLES = 7,
    /**
     * number of silent concealed samples
     *
     * @generated from enum value: CLIENT_AUDIO_SUBSCRIBER_SILENT_CONCEALED_SAMPLES = 8;
     */
    CLIENT_AUDIO_SUBSCRIBER_SILENT_CONCEALED_SAMPLES = 8,
    /**
     * number of concealment events
     *
     * @generated from enum value: CLIENT_AUDIO_SUBSCRIBER_CONCEALMENT_EVENTS = 9;
     */
    CLIENT_AUDIO_SUBSCRIBER_CONCEALMENT_EVENTS = 9,
    /**
     * number of interruptions
     *
     * @generated from enum value: CLIENT_AUDIO_SUBSCRIBER_INTERRUPTION_COUNT = 10;
     */
    CLIENT_AUDIO_SUBSCRIBER_INTERRUPTION_COUNT = 10,
    /**
     * total duration of interruptions
     *
     * @generated from enum value: CLIENT_AUDIO_SUBSCRIBER_TOTAL_INTERRUPTION_DURATION = 11;
     */
    CLIENT_AUDIO_SUBSCRIBER_TOTAL_INTERRUPTION_DURATION = 11,
    /**
     * total time spent in jitter buffer
     *
     * @generated from enum value: CLIENT_SUBSCRIBER_JITTER_BUFFER_DELAY = 12;
     */
    CLIENT_SUBSCRIBER_JITTER_BUFFER_DELAY = 12,
    /**
     * total time spent in jitter buffer
     *
     * @generated from enum value: CLIENT_SUBSCRIBER_JITTER_BUFFER_EMITTED_COUNT = 13;
     */
    CLIENT_SUBSCRIBER_JITTER_BUFFER_EMITTED_COUNT = 13,
    /**
     * total duration spent in bandwidth quality limitation
     *
     * @generated from enum value: CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_BANDWIDTH = 14;
     */
    CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_BANDWIDTH = 14,
    /**
     * total duration spent in cpu quality limitation
     *
     * @generated from enum value: CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_CPU = 15;
     */
    CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_CPU = 15,
    /**
     * total duration spent in other quality limitation
     *
     * @generated from enum value: CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_OTHER = 16;
     */
    CLIENT_VIDEO_PUBLISHER_QUALITY_LIMITATION_DURATION_OTHER = 16,
    /**
     * Publisher RTT (participant -> server)
     *
     * @generated from enum value: PUBLISHER_RTT = 17;
     */
    PUBLISHER_RTT = 17,
    /**
     * RTT between publisher node and subscriber node (could involve intermedia node(s))
     *
     * @generated from enum value: SERVER_MESH_RTT = 18;
     */
    SERVER_MESH_RTT = 18,
    /**
     * Subscribe RTT (server -> participant)
     *
     * @generated from enum value: SUBSCRIBER_RTT = 19;
     */
    SUBSCRIBER_RTT = 19,
    /**
     * @generated from enum value: METRIC_LABEL_PREDEFINED_MAX_VALUE = 4096;
     */
    METRIC_LABEL_PREDEFINED_MAX_VALUE = 4096
}
/**
 * @generated from message livekit.MetricsBatch
 */
export declare class MetricsBatch extends Message<MetricsBatch> {
    /**
     * time at which this batch is sent based on a monotonic clock (millisecond resolution)
     *
     * @generated from field: int64 timestamp_ms = 1;
     */
    timestampMs: bigint;
    /**
     * @generated from field: google.protobuf.Timestamp normalized_timestamp = 2;
     */
    normalizedTimestamp?: Timestamp;
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
    strData: string[];
    /**
     * @generated from field: repeated livekit.TimeSeriesMetric time_series = 4;
     */
    timeSeries: TimeSeriesMetric[];
    /**
     * @generated from field: repeated livekit.EventMetric events = 5;
     */
    events: EventMetric[];
    constructor(data?: PartialMessage<MetricsBatch>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.MetricsBatch";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetricsBatch;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetricsBatch;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetricsBatch;
    static equals(a: MetricsBatch | PlainMessage<MetricsBatch> | undefined, b: MetricsBatch | PlainMessage<MetricsBatch> | undefined): boolean;
}
/**
 * @generated from message livekit.TimeSeriesMetric
 */
export declare class TimeSeriesMetric extends Message<TimeSeriesMetric> {
    /**
     * Metric name e.g "speech_probablity". The string value is not directly stored in the message, but referenced by index
     * in the `str_data` field of `MetricsBatch`
     *
     * @generated from field: uint32 label = 1;
     */
    label: number;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 participant_identity = 2;
     */
    participantIdentity: number;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 track_sid = 3;
     */
    trackSid: number;
    /**
     * @generated from field: repeated livekit.MetricSample samples = 4;
     */
    samples: MetricSample[];
    /**
     * index into 'str_data'
     *
     * @generated from field: uint32 rid = 5;
     */
    rid: number;
    constructor(data?: PartialMessage<TimeSeriesMetric>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TimeSeriesMetric";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TimeSeriesMetric;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TimeSeriesMetric;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TimeSeriesMetric;
    static equals(a: TimeSeriesMetric | PlainMessage<TimeSeriesMetric> | undefined, b: TimeSeriesMetric | PlainMessage<TimeSeriesMetric> | undefined): boolean;
}
/**
 * @generated from message livekit.MetricSample
 */
export declare class MetricSample extends Message<MetricSample> {
    /**
     * time of metric based on a monotonic clock (in milliseconds)
     *
     * @generated from field: int64 timestamp_ms = 1;
     */
    timestampMs: bigint;
    /**
     * @generated from field: google.protobuf.Timestamp normalized_timestamp = 2;
     */
    normalizedTimestamp?: Timestamp;
    /**
     * @generated from field: float value = 3;
     */
    value: number;
    constructor(data?: PartialMessage<MetricSample>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.MetricSample";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetricSample;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetricSample;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetricSample;
    static equals(a: MetricSample | PlainMessage<MetricSample> | undefined, b: MetricSample | PlainMessage<MetricSample> | undefined): boolean;
}
/**
 * @generated from message livekit.EventMetric
 */
export declare class EventMetric extends Message<EventMetric> {
    /**
     * @generated from field: uint32 label = 1;
     */
    label: number;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 participant_identity = 2;
     */
    participantIdentity: number;
    /**
     * index into `str_data`
     *
     * @generated from field: uint32 track_sid = 3;
     */
    trackSid: number;
    /**
     * start time of event based on a monotonic clock (in milliseconds)
     *
     * @generated from field: int64 start_timestamp_ms = 4;
     */
    startTimestampMs: bigint;
    /**
     * end time of event based on a monotonic clock (in milliseconds), if needed
     *
     * @generated from field: optional int64 end_timestamp_ms = 5;
     */
    endTimestampMs?: bigint;
    /**
     * @generated from field: google.protobuf.Timestamp normalized_start_timestamp = 6;
     */
    normalizedStartTimestamp?: Timestamp;
    /**
     * @generated from field: optional google.protobuf.Timestamp normalized_end_timestamp = 7;
     */
    normalizedEndTimestamp?: Timestamp;
    /**
     * @generated from field: string metadata = 8;
     */
    metadata: string;
    /**
     * index into 'str_data'
     *
     * @generated from field: uint32 rid = 9;
     */
    rid: number;
    constructor(data?: PartialMessage<EventMetric>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.EventMetric";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventMetric;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventMetric;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventMetric;
    static equals(a: EventMetric | PlainMessage<EventMetric> | undefined, b: EventMetric | PlainMessage<EventMetric> | undefined): boolean;
}
//# sourceMappingURL=livekit_metrics_pb.d.ts.map