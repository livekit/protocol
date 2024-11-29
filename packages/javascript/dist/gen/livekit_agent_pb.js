import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { ParticipantInfo, ParticipantPermission, Room, ServerInfo } from "./livekit_models_pb.js";
var JobType = /* @__PURE__ */ ((JobType2) => {
  JobType2[JobType2["JT_ROOM"] = 0] = "JT_ROOM";
  JobType2[JobType2["JT_PUBLISHER"] = 1] = "JT_PUBLISHER";
  return JobType2;
})(JobType || {});
proto3.util.setEnumType(JobType, "livekit.JobType", [
  { no: 0, name: "JT_ROOM" },
  { no: 1, name: "JT_PUBLISHER" }
]);
var WorkerStatus = /* @__PURE__ */ ((WorkerStatus2) => {
  WorkerStatus2[WorkerStatus2["WS_AVAILABLE"] = 0] = "WS_AVAILABLE";
  WorkerStatus2[WorkerStatus2["WS_FULL"] = 1] = "WS_FULL";
  return WorkerStatus2;
})(WorkerStatus || {});
proto3.util.setEnumType(WorkerStatus, "livekit.WorkerStatus", [
  { no: 0, name: "WS_AVAILABLE" },
  { no: 1, name: "WS_FULL" }
]);
var JobStatus = /* @__PURE__ */ ((JobStatus2) => {
  JobStatus2[JobStatus2["JS_PENDING"] = 0] = "JS_PENDING";
  JobStatus2[JobStatus2["JS_RUNNING"] = 1] = "JS_RUNNING";
  JobStatus2[JobStatus2["JS_SUCCESS"] = 2] = "JS_SUCCESS";
  JobStatus2[JobStatus2["JS_FAILED"] = 3] = "JS_FAILED";
  return JobStatus2;
})(JobStatus || {});
proto3.util.setEnumType(JobStatus, "livekit.JobStatus", [
  { no: 0, name: "JS_PENDING" },
  { no: 1, name: "JS_RUNNING" },
  { no: 2, name: "JS_SUCCESS" },
  { no: 3, name: "JS_FAILED" }
]);
const _Job = class _Job extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string dispatch_id = 9;
     */
    this.dispatchId = "";
    /**
     * @generated from field: livekit.JobType type = 2;
     */
    this.type = 0 /* JT_ROOM */;
    /**
     * @generated from field: string namespace = 5 [deprecated = true];
     * @deprecated
     */
    this.namespace = "";
    /**
     * @generated from field: string metadata = 6;
     */
    this.metadata = "";
    /**
     * @generated from field: string agent_name = 7;
     */
    this.agentName = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Job().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Job().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Job().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_Job, a, b);
  }
};
_Job.runtime = proto3;
_Job.typeName = "livekit.Job";
_Job.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "dispatch_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(JobType) },
  { no: 3, name: "room", kind: "message", T: Room },
  { no: 4, name: "participant", kind: "message", T: ParticipantInfo, opt: true },
  {
    no: 5,
    name: "namespace",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "agent_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 8, name: "state", kind: "message", T: JobState }
]);
let Job = _Job;
const _JobState = class _JobState extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.JobStatus status = 1;
     */
    this.status = 0 /* JS_PENDING */;
    /**
     * @generated from field: string error = 2;
     */
    this.error = "";
    /**
     * @generated from field: int64 started_at = 3;
     */
    this.startedAt = protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 4;
     */
    this.endedAt = protoInt64.zero;
    /**
     * @generated from field: int64 updated_at = 5;
     */
    this.updatedAt = protoInt64.zero;
    /**
     * @generated from field: string participant_identity = 6;
     */
    this.participantIdentity = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _JobState().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _JobState().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _JobState().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_JobState, a, b);
  }
};
_JobState.runtime = proto3;
_JobState.typeName = "livekit.JobState";
_JobState.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(JobStatus) },
  {
    no: 2,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "started_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 4,
    name: "ended_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 5,
    name: "updated_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 6,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let JobState = _JobState;
const _WorkerMessage = class _WorkerMessage extends Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.WorkerMessage.message
     */
    this.message = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _WorkerMessage().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _WorkerMessage().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _WorkerMessage().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_WorkerMessage, a, b);
  }
};
_WorkerMessage.runtime = proto3;
_WorkerMessage.typeName = "livekit.WorkerMessage";
_WorkerMessage.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "register", kind: "message", T: RegisterWorkerRequest, oneof: "message" },
  { no: 2, name: "availability", kind: "message", T: AvailabilityResponse, oneof: "message" },
  { no: 3, name: "update_worker", kind: "message", T: UpdateWorkerStatus, oneof: "message" },
  { no: 4, name: "update_job", kind: "message", T: UpdateJobStatus, oneof: "message" },
  { no: 5, name: "ping", kind: "message", T: WorkerPing, oneof: "message" },
  { no: 6, name: "simulate_job", kind: "message", T: SimulateJobRequest, oneof: "message" },
  { no: 7, name: "migrate_job", kind: "message", T: MigrateJobRequest, oneof: "message" }
]);
let WorkerMessage = _WorkerMessage;
const _ServerMessage = class _ServerMessage extends Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.ServerMessage.message
     */
    this.message = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ServerMessage().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ServerMessage().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ServerMessage().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ServerMessage, a, b);
  }
};
_ServerMessage.runtime = proto3;
_ServerMessage.typeName = "livekit.ServerMessage";
_ServerMessage.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "register", kind: "message", T: RegisterWorkerResponse, oneof: "message" },
  { no: 2, name: "availability", kind: "message", T: AvailabilityRequest, oneof: "message" },
  { no: 3, name: "assignment", kind: "message", T: JobAssignment, oneof: "message" },
  { no: 5, name: "termination", kind: "message", T: JobTermination, oneof: "message" },
  { no: 4, name: "pong", kind: "message", T: WorkerPong, oneof: "message" }
]);
let ServerMessage = _ServerMessage;
const _SimulateJobRequest = class _SimulateJobRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.JobType type = 1;
     */
    this.type = 0 /* JT_ROOM */;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SimulateJobRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SimulateJobRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SimulateJobRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SimulateJobRequest, a, b);
  }
};
_SimulateJobRequest.runtime = proto3;
_SimulateJobRequest.typeName = "livekit.SimulateJobRequest";
_SimulateJobRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(JobType) },
  { no: 2, name: "room", kind: "message", T: Room },
  { no: 3, name: "participant", kind: "message", T: ParticipantInfo }
]);
let SimulateJobRequest = _SimulateJobRequest;
const _WorkerPing = class _WorkerPing extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: int64 timestamp = 1;
     */
    this.timestamp = protoInt64.zero;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _WorkerPing().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _WorkerPing().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _WorkerPing().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_WorkerPing, a, b);
  }
};
_WorkerPing.runtime = proto3;
_WorkerPing.typeName = "livekit.WorkerPing";
_WorkerPing.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let WorkerPing = _WorkerPing;
const _WorkerPong = class _WorkerPong extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: int64 last_timestamp = 1;
     */
    this.lastTimestamp = protoInt64.zero;
    /**
     * @generated from field: int64 timestamp = 2;
     */
    this.timestamp = protoInt64.zero;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _WorkerPong().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _WorkerPong().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _WorkerPong().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_WorkerPong, a, b);
  }
};
_WorkerPong.runtime = proto3;
_WorkerPong.typeName = "livekit.WorkerPong";
_WorkerPong.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "last_timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "timestamp",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  }
]);
let WorkerPong = _WorkerPong;
const _RegisterWorkerRequest = class _RegisterWorkerRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: livekit.JobType type = 1;
     */
    this.type = 0 /* JT_ROOM */;
    /**
     * @generated from field: string agent_name = 8;
     */
    this.agentName = "";
    /**
     * string worker_id = 2;
     *
     * @generated from field: string version = 3;
     */
    this.version = "";
    /**
     * string name = 4 [deprecated = true];
     *
     * @generated from field: uint32 ping_interval = 5;
     */
    this.pingInterval = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RegisterWorkerRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RegisterWorkerRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RegisterWorkerRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RegisterWorkerRequest, a, b);
  }
};
_RegisterWorkerRequest.runtime = proto3;
_RegisterWorkerRequest.typeName = "livekit.RegisterWorkerRequest";
_RegisterWorkerRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(JobType) },
  {
    no: 8,
    name: "agent_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "version",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "ping_interval",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 6, name: "namespace", kind: "scalar", T: 9, opt: true },
  { no: 7, name: "allowed_permissions", kind: "message", T: ParticipantPermission }
]);
let RegisterWorkerRequest = _RegisterWorkerRequest;
const _RegisterWorkerResponse = class _RegisterWorkerResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string worker_id = 1;
     */
    this.workerId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _RegisterWorkerResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _RegisterWorkerResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _RegisterWorkerResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_RegisterWorkerResponse, a, b);
  }
};
_RegisterWorkerResponse.runtime = proto3;
_RegisterWorkerResponse.typeName = "livekit.RegisterWorkerResponse";
_RegisterWorkerResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "worker_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "server_info", kind: "message", T: ServerInfo }
]);
let RegisterWorkerResponse = _RegisterWorkerResponse;
const _MigrateJobRequest = class _MigrateJobRequest extends Message {
  constructor(data) {
    super();
    /**
     * string job_id = 1 [deprecated = true];
     *
     * @generated from field: repeated string job_ids = 2;
     */
    this.jobIds = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _MigrateJobRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _MigrateJobRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _MigrateJobRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_MigrateJobRequest, a, b);
  }
};
_MigrateJobRequest.runtime = proto3;
_MigrateJobRequest.typeName = "livekit.MigrateJobRequest";
_MigrateJobRequest.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "job_ids", kind: "scalar", T: 9, repeated: true }
]);
let MigrateJobRequest = _MigrateJobRequest;
const _AvailabilityRequest = class _AvailabilityRequest extends Message {
  constructor(data) {
    super();
    /**
     * True when the job was previously assigned to another worker but has been
     * migrated due to different reasons (e.g. worker failure, job migration)
     *
     * @generated from field: bool resuming = 2;
     */
    this.resuming = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AvailabilityRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AvailabilityRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AvailabilityRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_AvailabilityRequest, a, b);
  }
};
_AvailabilityRequest.runtime = proto3;
_AvailabilityRequest.typeName = "livekit.AvailabilityRequest";
_AvailabilityRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "job", kind: "message", T: Job },
  {
    no: 2,
    name: "resuming",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let AvailabilityRequest = _AvailabilityRequest;
const _AvailabilityResponse = class _AvailabilityResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string job_id = 1;
     */
    this.jobId = "";
    /**
     * @generated from field: bool available = 2;
     */
    this.available = false;
    /**
     * @generated from field: bool supports_resume = 3;
     */
    this.supportsResume = false;
    /**
     * @generated from field: string participant_name = 4;
     */
    this.participantName = "";
    /**
     * @generated from field: string participant_identity = 5;
     */
    this.participantIdentity = "";
    /**
     * @generated from field: string participant_metadata = 6;
     */
    this.participantMetadata = "";
    /**
     * @generated from field: map<string, string> participant_attributes = 7;
     */
    this.participantAttributes = {};
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AvailabilityResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AvailabilityResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AvailabilityResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_AvailabilityResponse, a, b);
  }
};
_AvailabilityResponse.runtime = proto3;
_AvailabilityResponse.typeName = "livekit.AvailabilityResponse";
_AvailabilityResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "job_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "available",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 3,
    name: "supports_resume",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 4,
    name: "participant_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "participant_metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 7, name: "participant_attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } }
]);
let AvailabilityResponse = _AvailabilityResponse;
const _UpdateJobStatus = class _UpdateJobStatus extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string job_id = 1;
     */
    this.jobId = "";
    /**
     * The worker can indicate the job end by either specifying SUCCESS or FAILED
     *
     * @generated from field: livekit.JobStatus status = 2;
     */
    this.status = 0 /* JS_PENDING */;
    /**
     * metadata shown on the dashboard, useful for debugging
     *
     * @generated from field: string error = 3;
     */
    this.error = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateJobStatus().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateJobStatus().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateJobStatus().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateJobStatus, a, b);
  }
};
_UpdateJobStatus.runtime = proto3;
_UpdateJobStatus.typeName = "livekit.UpdateJobStatus";
_UpdateJobStatus.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "job_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(JobStatus) },
  {
    no: 3,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let UpdateJobStatus = _UpdateJobStatus;
const _UpdateWorkerStatus = class _UpdateWorkerStatus extends Message {
  constructor(data) {
    super();
    /**
     * optional string metadata = 2 [deprecated=true];
     *
     * @generated from field: float load = 3;
     */
    this.load = 0;
    /**
     * @generated from field: uint32 job_count = 4;
     */
    this.jobCount = 0;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateWorkerStatus().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateWorkerStatus().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateWorkerStatus().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_UpdateWorkerStatus, a, b);
  }
};
_UpdateWorkerStatus.runtime = proto3;
_UpdateWorkerStatus.typeName = "livekit.UpdateWorkerStatus";
_UpdateWorkerStatus.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "status", kind: "enum", T: proto3.getEnumType(WorkerStatus), opt: true },
  {
    no: 3,
    name: "load",
    kind: "scalar",
    T: 2
    /* ScalarType.FLOAT */
  },
  {
    no: 4,
    name: "job_count",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
let UpdateWorkerStatus = _UpdateWorkerStatus;
const _JobAssignment = class _JobAssignment extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string token = 3;
     */
    this.token = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _JobAssignment().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _JobAssignment().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _JobAssignment().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_JobAssignment, a, b);
  }
};
_JobAssignment.runtime = proto3;
_JobAssignment.typeName = "livekit.JobAssignment";
_JobAssignment.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "job", kind: "message", T: Job },
  { no: 2, name: "url", kind: "scalar", T: 9, opt: true },
  {
    no: 3,
    name: "token",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let JobAssignment = _JobAssignment;
const _JobTermination = class _JobTermination extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string job_id = 1;
     */
    this.jobId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _JobTermination().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _JobTermination().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _JobTermination().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_JobTermination, a, b);
  }
};
_JobTermination.runtime = proto3;
_JobTermination.typeName = "livekit.JobTermination";
_JobTermination.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "job_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let JobTermination = _JobTermination;
export {
  AvailabilityRequest,
  AvailabilityResponse,
  Job,
  JobAssignment,
  JobState,
  JobStatus,
  JobTermination,
  JobType,
  MigrateJobRequest,
  RegisterWorkerRequest,
  RegisterWorkerResponse,
  ServerMessage,
  SimulateJobRequest,
  UpdateJobStatus,
  UpdateWorkerStatus,
  WorkerMessage,
  WorkerPing,
  WorkerPong,
  WorkerStatus
};
//# sourceMappingURL=livekit_agent_pb.js.map