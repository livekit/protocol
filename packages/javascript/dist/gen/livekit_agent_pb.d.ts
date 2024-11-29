import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { ParticipantInfo, ParticipantPermission, Room, ServerInfo } from "./livekit_models_pb.js";
/**
 * @generated from enum livekit.JobType
 */
export declare enum JobType {
    /**
     * @generated from enum value: JT_ROOM = 0;
     */
    JT_ROOM = 0,
    /**
     * @generated from enum value: JT_PUBLISHER = 1;
     */
    JT_PUBLISHER = 1
}
/**
 * @generated from enum livekit.WorkerStatus
 */
export declare enum WorkerStatus {
    /**
     * @generated from enum value: WS_AVAILABLE = 0;
     */
    WS_AVAILABLE = 0,
    /**
     * @generated from enum value: WS_FULL = 1;
     */
    WS_FULL = 1
}
/**
 * @generated from enum livekit.JobStatus
 */
export declare enum JobStatus {
    /**
     * @generated from enum value: JS_PENDING = 0;
     */
    JS_PENDING = 0,
    /**
     * @generated from enum value: JS_RUNNING = 1;
     */
    JS_RUNNING = 1,
    /**
     * @generated from enum value: JS_SUCCESS = 2;
     */
    JS_SUCCESS = 2,
    /**
     * @generated from enum value: JS_FAILED = 3;
     */
    JS_FAILED = 3
}
/**
 * @generated from message livekit.Job
 */
export declare class Job extends Message<Job> {
    /**
     * @generated from field: string id = 1;
     */
    id: string;
    /**
     * @generated from field: string dispatch_id = 9;
     */
    dispatchId: string;
    /**
     * @generated from field: livekit.JobType type = 2;
     */
    type: JobType;
    /**
     * @generated from field: livekit.Room room = 3;
     */
    room?: Room;
    /**
     * @generated from field: optional livekit.ParticipantInfo participant = 4;
     */
    participant?: ParticipantInfo;
    /**
     * @generated from field: string namespace = 5 [deprecated = true];
     * @deprecated
     */
    namespace: string;
    /**
     * @generated from field: string metadata = 6;
     */
    metadata: string;
    /**
     * @generated from field: string agent_name = 7;
     */
    agentName: string;
    /**
     * @generated from field: livekit.JobState state = 8;
     */
    state?: JobState;
    constructor(data?: PartialMessage<Job>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.Job";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Job;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Job;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Job;
    static equals(a: Job | PlainMessage<Job> | undefined, b: Job | PlainMessage<Job> | undefined): boolean;
}
/**
 * @generated from message livekit.JobState
 */
export declare class JobState extends Message<JobState> {
    /**
     * @generated from field: livekit.JobStatus status = 1;
     */
    status: JobStatus;
    /**
     * @generated from field: string error = 2;
     */
    error: string;
    /**
     * @generated from field: int64 started_at = 3;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 4;
     */
    endedAt: bigint;
    /**
     * @generated from field: int64 updated_at = 5;
     */
    updatedAt: bigint;
    /**
     * @generated from field: string participant_identity = 6;
     */
    participantIdentity: string;
    constructor(data?: PartialMessage<JobState>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.JobState";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JobState;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JobState;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JobState;
    static equals(a: JobState | PlainMessage<JobState> | undefined, b: JobState | PlainMessage<JobState> | undefined): boolean;
}
/**
 * from Worker to Server
 *
 * @generated from message livekit.WorkerMessage
 */
export declare class WorkerMessage extends Message<WorkerMessage> {
    /**
     * @generated from oneof livekit.WorkerMessage.message
     */
    message: {
        /**
         * agent workers need to register themselves with the server first
         *
         * @generated from field: livekit.RegisterWorkerRequest register = 1;
         */
        value: RegisterWorkerRequest;
        case: "register";
    } | {
        /**
         * worker confirms to server that it's available for a job, or declines it
         *
         * @generated from field: livekit.AvailabilityResponse availability = 2;
         */
        value: AvailabilityResponse;
        case: "availability";
    } | {
        /**
         * worker can update its status to the server, including taking itself out of the pool
         *
         * @generated from field: livekit.UpdateWorkerStatus update_worker = 3;
         */
        value: UpdateWorkerStatus;
        case: "updateWorker";
    } | {
        /**
         * job can send status updates to the server, useful for tracking progress
         *
         * @generated from field: livekit.UpdateJobStatus update_job = 4;
         */
        value: UpdateJobStatus;
        case: "updateJob";
    } | {
        /**
         * @generated from field: livekit.WorkerPing ping = 5;
         */
        value: WorkerPing;
        case: "ping";
    } | {
        /**
         * @generated from field: livekit.SimulateJobRequest simulate_job = 6;
         */
        value: SimulateJobRequest;
        case: "simulateJob";
    } | {
        /**
         * @generated from field: livekit.MigrateJobRequest migrate_job = 7;
         */
        value: MigrateJobRequest;
        case: "migrateJob";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<WorkerMessage>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.WorkerMessage";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkerMessage;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkerMessage;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkerMessage;
    static equals(a: WorkerMessage | PlainMessage<WorkerMessage> | undefined, b: WorkerMessage | PlainMessage<WorkerMessage> | undefined): boolean;
}
/**
 * from Server to Worker
 *
 * @generated from message livekit.ServerMessage
 */
export declare class ServerMessage extends Message<ServerMessage> {
    /**
     * @generated from oneof livekit.ServerMessage.message
     */
    message: {
        /**
         * server confirms the registration, from this moment on, the worker is considered active
         *
         * @generated from field: livekit.RegisterWorkerResponse register = 1;
         */
        value: RegisterWorkerResponse;
        case: "register";
    } | {
        /**
         * server asks worker to confirm availability for a job
         *
         * @generated from field: livekit.AvailabilityRequest availability = 2;
         */
        value: AvailabilityRequest;
        case: "availability";
    } | {
        /**
         * @generated from field: livekit.JobAssignment assignment = 3;
         */
        value: JobAssignment;
        case: "assignment";
    } | {
        /**
         * @generated from field: livekit.JobTermination termination = 5;
         */
        value: JobTermination;
        case: "termination";
    } | {
        /**
         * @generated from field: livekit.WorkerPong pong = 4;
         */
        value: WorkerPong;
        case: "pong";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<ServerMessage>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ServerMessage";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ServerMessage;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ServerMessage;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ServerMessage;
    static equals(a: ServerMessage | PlainMessage<ServerMessage> | undefined, b: ServerMessage | PlainMessage<ServerMessage> | undefined): boolean;
}
/**
 * @generated from message livekit.SimulateJobRequest
 */
export declare class SimulateJobRequest extends Message<SimulateJobRequest> {
    /**
     * @generated from field: livekit.JobType type = 1;
     */
    type: JobType;
    /**
     * @generated from field: livekit.Room room = 2;
     */
    room?: Room;
    /**
     * @generated from field: livekit.ParticipantInfo participant = 3;
     */
    participant?: ParticipantInfo;
    constructor(data?: PartialMessage<SimulateJobRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SimulateJobRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SimulateJobRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SimulateJobRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SimulateJobRequest;
    static equals(a: SimulateJobRequest | PlainMessage<SimulateJobRequest> | undefined, b: SimulateJobRequest | PlainMessage<SimulateJobRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.WorkerPing
 */
export declare class WorkerPing extends Message<WorkerPing> {
    /**
     * @generated from field: int64 timestamp = 1;
     */
    timestamp: bigint;
    constructor(data?: PartialMessage<WorkerPing>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.WorkerPing";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkerPing;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkerPing;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkerPing;
    static equals(a: WorkerPing | PlainMessage<WorkerPing> | undefined, b: WorkerPing | PlainMessage<WorkerPing> | undefined): boolean;
}
/**
 * @generated from message livekit.WorkerPong
 */
export declare class WorkerPong extends Message<WorkerPong> {
    /**
     * @generated from field: int64 last_timestamp = 1;
     */
    lastTimestamp: bigint;
    /**
     * @generated from field: int64 timestamp = 2;
     */
    timestamp: bigint;
    constructor(data?: PartialMessage<WorkerPong>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.WorkerPong";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkerPong;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkerPong;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkerPong;
    static equals(a: WorkerPong | PlainMessage<WorkerPong> | undefined, b: WorkerPong | PlainMessage<WorkerPong> | undefined): boolean;
}
/**
 * @generated from message livekit.RegisterWorkerRequest
 */
export declare class RegisterWorkerRequest extends Message<RegisterWorkerRequest> {
    /**
     * @generated from field: livekit.JobType type = 1;
     */
    type: JobType;
    /**
     * @generated from field: string agent_name = 8;
     */
    agentName: string;
    /**
     * string worker_id = 2;
     *
     * @generated from field: string version = 3;
     */
    version: string;
    /**
     * string name = 4 [deprecated = true];
     *
     * @generated from field: uint32 ping_interval = 5;
     */
    pingInterval: number;
    /**
     * @generated from field: optional string namespace = 6;
     */
    namespace?: string;
    /**
     * @generated from field: livekit.ParticipantPermission allowed_permissions = 7;
     */
    allowedPermissions?: ParticipantPermission;
    constructor(data?: PartialMessage<RegisterWorkerRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RegisterWorkerRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterWorkerRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterWorkerRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterWorkerRequest;
    static equals(a: RegisterWorkerRequest | PlainMessage<RegisterWorkerRequest> | undefined, b: RegisterWorkerRequest | PlainMessage<RegisterWorkerRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.RegisterWorkerResponse
 */
export declare class RegisterWorkerResponse extends Message<RegisterWorkerResponse> {
    /**
     * @generated from field: string worker_id = 1;
     */
    workerId: string;
    /**
     * @generated from field: livekit.ServerInfo server_info = 3;
     */
    serverInfo?: ServerInfo;
    constructor(data?: PartialMessage<RegisterWorkerResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.RegisterWorkerResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterWorkerResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterWorkerResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterWorkerResponse;
    static equals(a: RegisterWorkerResponse | PlainMessage<RegisterWorkerResponse> | undefined, b: RegisterWorkerResponse | PlainMessage<RegisterWorkerResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.MigrateJobRequest
 */
export declare class MigrateJobRequest extends Message<MigrateJobRequest> {
    /**
     * string job_id = 1 [deprecated = true];
     *
     * @generated from field: repeated string job_ids = 2;
     */
    jobIds: string[];
    constructor(data?: PartialMessage<MigrateJobRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.MigrateJobRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MigrateJobRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MigrateJobRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MigrateJobRequest;
    static equals(a: MigrateJobRequest | PlainMessage<MigrateJobRequest> | undefined, b: MigrateJobRequest | PlainMessage<MigrateJobRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.AvailabilityRequest
 */
export declare class AvailabilityRequest extends Message<AvailabilityRequest> {
    /**
     * @generated from field: livekit.Job job = 1;
     */
    job?: Job;
    /**
     * True when the job was previously assigned to another worker but has been
     * migrated due to different reasons (e.g. worker failure, job migration)
     *
     * @generated from field: bool resuming = 2;
     */
    resuming: boolean;
    constructor(data?: PartialMessage<AvailabilityRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AvailabilityRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AvailabilityRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AvailabilityRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AvailabilityRequest;
    static equals(a: AvailabilityRequest | PlainMessage<AvailabilityRequest> | undefined, b: AvailabilityRequest | PlainMessage<AvailabilityRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.AvailabilityResponse
 */
export declare class AvailabilityResponse extends Message<AvailabilityResponse> {
    /**
     * @generated from field: string job_id = 1;
     */
    jobId: string;
    /**
     * @generated from field: bool available = 2;
     */
    available: boolean;
    /**
     * @generated from field: bool supports_resume = 3;
     */
    supportsResume: boolean;
    /**
     * @generated from field: string participant_name = 4;
     */
    participantName: string;
    /**
     * @generated from field: string participant_identity = 5;
     */
    participantIdentity: string;
    /**
     * @generated from field: string participant_metadata = 6;
     */
    participantMetadata: string;
    /**
     * @generated from field: map<string, string> participant_attributes = 7;
     */
    participantAttributes: {
        [key: string]: string;
    };
    constructor(data?: PartialMessage<AvailabilityResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.AvailabilityResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AvailabilityResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AvailabilityResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AvailabilityResponse;
    static equals(a: AvailabilityResponse | PlainMessage<AvailabilityResponse> | undefined, b: AvailabilityResponse | PlainMessage<AvailabilityResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateJobStatus
 */
export declare class UpdateJobStatus extends Message<UpdateJobStatus> {
    /**
     * @generated from field: string job_id = 1;
     */
    jobId: string;
    /**
     * The worker can indicate the job end by either specifying SUCCESS or FAILED
     *
     * @generated from field: livekit.JobStatus status = 2;
     */
    status: JobStatus;
    /**
     * metadata shown on the dashboard, useful for debugging
     *
     * @generated from field: string error = 3;
     */
    error: string;
    constructor(data?: PartialMessage<UpdateJobStatus>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateJobStatus";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateJobStatus;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateJobStatus;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateJobStatus;
    static equals(a: UpdateJobStatus | PlainMessage<UpdateJobStatus> | undefined, b: UpdateJobStatus | PlainMessage<UpdateJobStatus> | undefined): boolean;
}
/**
 * @generated from message livekit.UpdateWorkerStatus
 */
export declare class UpdateWorkerStatus extends Message<UpdateWorkerStatus> {
    /**
     * @generated from field: optional livekit.WorkerStatus status = 1;
     */
    status?: WorkerStatus;
    /**
     * optional string metadata = 2 [deprecated=true];
     *
     * @generated from field: float load = 3;
     */
    load: number;
    /**
     * @generated from field: uint32 job_count = 4;
     */
    jobCount: number;
    constructor(data?: PartialMessage<UpdateWorkerStatus>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.UpdateWorkerStatus";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateWorkerStatus;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateWorkerStatus;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateWorkerStatus;
    static equals(a: UpdateWorkerStatus | PlainMessage<UpdateWorkerStatus> | undefined, b: UpdateWorkerStatus | PlainMessage<UpdateWorkerStatus> | undefined): boolean;
}
/**
 * @generated from message livekit.JobAssignment
 */
export declare class JobAssignment extends Message<JobAssignment> {
    /**
     * @generated from field: livekit.Job job = 1;
     */
    job?: Job;
    /**
     * @generated from field: optional string url = 2;
     */
    url?: string;
    /**
     * @generated from field: string token = 3;
     */
    token: string;
    constructor(data?: PartialMessage<JobAssignment>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.JobAssignment";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JobAssignment;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JobAssignment;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JobAssignment;
    static equals(a: JobAssignment | PlainMessage<JobAssignment> | undefined, b: JobAssignment | PlainMessage<JobAssignment> | undefined): boolean;
}
/**
 * @generated from message livekit.JobTermination
 */
export declare class JobTermination extends Message<JobTermination> {
    /**
     * @generated from field: string job_id = 1;
     */
    jobId: string;
    constructor(data?: PartialMessage<JobTermination>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.JobTermination";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JobTermination;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JobTermination;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JobTermination;
    static equals(a: JobTermination | PlainMessage<JobTermination> | undefined, b: JobTermination | PlainMessage<JobTermination> | undefined): boolean;
}
//# sourceMappingURL=livekit_agent_pb.d.ts.map