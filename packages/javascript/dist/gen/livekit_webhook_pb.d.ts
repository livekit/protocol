import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ParticipantInfo, Room, TrackInfo } from "./livekit_models_pb.js";
import { EgressInfo } from "./livekit_egress_pb.js";
import { IngressInfo } from "./livekit_ingress_pb.js";
/**
 * @generated from message livekit.WebhookEvent
 */
export declare class WebhookEvent extends Message<WebhookEvent> {
    /**
     * one of room_started, room_finished, participant_joined, participant_left,
     * track_published, track_unpublished, egress_started, egress_updated, egress_ended,
     * ingress_started, ingress_ended
     *
     * @generated from field: string event = 1;
     */
    event: string;
    /**
     * @generated from field: livekit.Room room = 2;
     */
    room?: Room;
    /**
     * set when event is participant_* or track_*
     *
     * @generated from field: livekit.ParticipantInfo participant = 3;
     */
    participant?: ParticipantInfo;
    /**
     * set when event is egress_*
     *
     * @generated from field: livekit.EgressInfo egress_info = 9;
     */
    egressInfo?: EgressInfo;
    /**
     * set when event is ingress_*
     *
     * @generated from field: livekit.IngressInfo ingress_info = 10;
     */
    ingressInfo?: IngressInfo;
    /**
     * set when event is track_*
     *
     * @generated from field: livekit.TrackInfo track = 8;
     */
    track?: TrackInfo;
    /**
     * unique event uuid
     *
     * @generated from field: string id = 6;
     */
    id: string;
    /**
     * timestamp in seconds
     *
     * @generated from field: int64 created_at = 7;
     */
    createdAt: bigint;
    /**
     * @generated from field: int32 num_dropped = 11;
     */
    numDropped: number;
    constructor(data?: PartialMessage<WebhookEvent>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.WebhookEvent";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WebhookEvent;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WebhookEvent;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WebhookEvent;
    static equals(a: WebhookEvent | PlainMessage<WebhookEvent> | undefined, b: WebhookEvent | PlainMessage<WebhookEvent> | undefined): boolean;
}
//# sourceMappingURL=livekit_webhook_pb.d.ts.map