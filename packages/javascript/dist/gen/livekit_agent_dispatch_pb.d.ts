import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Job } from "./livekit_agent_pb.js";
/**
 * @generated from message livekit.CreateAgentDispatchRequest
 */
export declare class CreateAgentDispatchRequest extends Message<CreateAgentDispatchRequest> {
    /**
     * @generated from field: string agent_name = 1;
     */
    agentName: string;
    /**
     * @generated from field: string room = 2;
     */
    room: string;
    /**
     * @generated from field: string metadata = 3;
     */
    metadata: string;
    constructor(data?: PartialMessage<CreateAgentDispatchRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateAgentDispatchRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAgentDispatchRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAgentDispatchRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAgentDispatchRequest;
    static equals(a: CreateAgentDispatchRequest | PlainMessage<CreateAgentDispatchRequest> | undefined, b: CreateAgentDispatchRequest | PlainMessage<CreateAgentDispatchRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.RoomAgentDispatch
 */
export declare class RoomAgentDispatch extends Message<RoomAgentDispatch> {
    /**
     * @generated from field: string agent_name = 1;
     */
    agentName: string;
    /**
     * @generated from field: string metadata = 2;
     */
    metadata: string;
    constructor(data?: PartialMessage<RoomAgentDispatch>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RoomAgentDispatch";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RoomAgentDispatch;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RoomAgentDispatch;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RoomAgentDispatch;
    static equals(a: RoomAgentDispatch | PlainMessage<RoomAgentDispatch> | undefined, b: RoomAgentDispatch | PlainMessage<RoomAgentDispatch> | undefined): boolean;
}
/**
 * @generated from message livekit.DeleteAgentDispatchRequest
 */
export declare class DeleteAgentDispatchRequest extends Message<DeleteAgentDispatchRequest> {
    /**
     * @generated from field: string dispatch_id = 1;
     */
    dispatchId: string;
    /**
     * @generated from field: string room = 2;
     */
    room: string;
    constructor(data?: PartialMessage<DeleteAgentDispatchRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DeleteAgentDispatchRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteAgentDispatchRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteAgentDispatchRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteAgentDispatchRequest;
    static equals(a: DeleteAgentDispatchRequest | PlainMessage<DeleteAgentDispatchRequest> | undefined, b: DeleteAgentDispatchRequest | PlainMessage<DeleteAgentDispatchRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListAgentDispatchRequest
 */
export declare class ListAgentDispatchRequest extends Message<ListAgentDispatchRequest> {
    /**
     * if set, only the dispatch whose id is given will be returned
     *
     * @generated from field: string dispatch_id = 1;
     */
    dispatchId: string;
    /**
     * name of the room to list agents for. Must be set.
     *
     * @generated from field: string room = 2;
     */
    room: string;
    constructor(data?: PartialMessage<ListAgentDispatchRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListAgentDispatchRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAgentDispatchRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAgentDispatchRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAgentDispatchRequest;
    static equals(a: ListAgentDispatchRequest | PlainMessage<ListAgentDispatchRequest> | undefined, b: ListAgentDispatchRequest | PlainMessage<ListAgentDispatchRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListAgentDispatchResponse
 */
export declare class ListAgentDispatchResponse extends Message<ListAgentDispatchResponse> {
    /**
     * @generated from field: repeated livekit.AgentDispatch agent_dispatches = 1;
     */
    agentDispatches: AgentDispatch[];
    constructor(data?: PartialMessage<ListAgentDispatchResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListAgentDispatchResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAgentDispatchResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAgentDispatchResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAgentDispatchResponse;
    static equals(a: ListAgentDispatchResponse | PlainMessage<ListAgentDispatchResponse> | undefined, b: ListAgentDispatchResponse | PlainMessage<ListAgentDispatchResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.AgentDispatch
 */
export declare class AgentDispatch extends Message<AgentDispatch> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string agent_name = 2;
     */
    agentName: string;
    /**
     * @generated from field: string room = 3;
     */
    room: string;
    /**
     * @generated from field: string metadata = 4;
     */
    metadata: string;
    /**
     * @generated from field: livekit.AgentDispatchState state = 5;
     */
    state?: AgentDispatchState;
    constructor(data?: PartialMessage<AgentDispatch>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AgentDispatch";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AgentDispatch;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AgentDispatch;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AgentDispatch;
    static equals(a: AgentDispatch | PlainMessage<AgentDispatch> | undefined, b: AgentDispatch | PlainMessage<AgentDispatch> | undefined): boolean;
}
/**
 * @generated from message livekit.AgentDispatchState
 */
export declare class AgentDispatchState extends Message<AgentDispatchState> {
    /**
     * For dispatches of tyoe JT_ROOM, there will be at most 1 job.
     * For dispatches of type JT_PUBLISHER, there will be 1 per publisher.
     *
     * @generated from field: repeated livekit.Job jobs = 1;
     */
    jobs: Job[];
    /**
     * @generated from field: int64 created_at = 2;
     */
    createdAt: bigint;
    /**
     * @generated from field: int64 deleted_at = 3;
     */
    deletedAt: bigint;
    constructor(data?: PartialMessage<AgentDispatchState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AgentDispatchState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AgentDispatchState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AgentDispatchState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AgentDispatchState;
    static equals(a: AgentDispatchState | PlainMessage<AgentDispatchState> | undefined, b: AgentDispatchState | PlainMessage<AgentDispatchState> | undefined): boolean;
}
//# sourceMappingURL=livekit_agent_dispatch_pb.d.ts.map