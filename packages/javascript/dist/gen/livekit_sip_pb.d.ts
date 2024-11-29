import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3 } from "@bufbuild/protobuf";
import { DisconnectReason } from "./livekit_models_pb.js";
/**
 * @generated from enum livekit.SIPTransport
 */
export declare enum SIPTransport {
    /**
     * @generated from enum value: SIP_TRANSPORT_AUTO = 0;
     */
    SIP_TRANSPORT_AUTO = 0,
    /**
     * @generated from enum value: SIP_TRANSPORT_UDP = 1;
     */
    SIP_TRANSPORT_UDP = 1,
    /**
     * @generated from enum value: SIP_TRANSPORT_TCP = 2;
     */
    SIP_TRANSPORT_TCP = 2,
    /**
     * @generated from enum value: SIP_TRANSPORT_TLS = 3;
     */
    SIP_TRANSPORT_TLS = 3
}
/**
 * @generated from enum livekit.SIPCallStatus
 */
export declare enum SIPCallStatus {
    /**
     * Incoming call is being handled by the SIP service. The SIP participant hasn't joined a LiveKit room yet
     *
     * @generated from enum value: SCS_CALL_INCOMING = 0;
     */
    SCS_CALL_INCOMING = 0,
    /**
     * SIP participant for outgoing call has been created. The SIP outgoing call is being established
     *
     * @generated from enum value: SCS_PARTICIPANT_JOINED = 1;
     */
    SCS_PARTICIPANT_JOINED = 1,
    /**
     * Call is ongoing. SIP participant is active in the LiveKit room
     *
     * @generated from enum value: SCS_ACTIVE = 2;
     */
    SCS_ACTIVE = 2,
    /**
     * Call has ended
     *
     * @generated from enum value: SCS_DISCONNECTED = 3;
     */
    SCS_DISCONNECTED = 3,
    /**
     * Call has ended or never succeeded because of an error
     *
     * @generated from enum value: SCS_ERROR = 4;
     */
    SCS_ERROR = 4
}
/**
 * @generated from enum livekit.SIPFeature
 */
export declare enum SIPFeature {
    /**
     * @generated from enum value: NONE = 0;
     */
    NONE = 0,
    /**
     * @generated from enum value: KRISP_ENABLED = 1;
     */
    KRISP_ENABLED = 1
}
/**
 * @generated from message livekit.CreateSIPTrunkRequest
 * @deprecated
 */
export declare class CreateSIPTrunkRequest extends Message<CreateSIPTrunkRequest> {
    /**
     * CIDR or IPs that traffic is accepted from
     * An empty list means all inbound traffic is accepted.
     *
     * @generated from field: repeated string inbound_addresses = 1;
     */
    inboundAddresses: string[];
    /**
     * IP that SIP INVITE is sent too
     *
     * @generated from field: string outbound_address = 2;
     */
    outboundAddress: string;
    /**
     * Number used to make outbound calls
     *
     * @generated from field: string outbound_number = 3;
     */
    outboundNumber: string;
    /**
     * @generated from field: repeated string inbound_numbers_regex = 4 [deprecated = true];
     * @deprecated
     */
    inboundNumbersRegex: string[];
    /**
     * Accepted `To` values. This Trunk will only accept a call made to
     * these numbers. This allows you to have distinct Trunks for different phone
     * numbers at the same provider.
     *
     * @generated from field: repeated string inbound_numbers = 9;
     */
    inboundNumbers: string[];
    /**
     * Username and password used to authenticate inbound and outbound SIP invites
     * May be empty to have no Authentication
     *
     * @generated from field: string inbound_username = 5;
     */
    inboundUsername: string;
    /**
     * @generated from field: string inbound_password = 6;
     */
    inboundPassword: string;
    /**
     * @generated from field: string outbound_username = 7;
     */
    outboundUsername: string;
    /**
     * @generated from field: string outbound_password = 8;
     */
    outboundPassword: string;
    /**
     * Optional human-readable name for the Trunk.
     *
     * @generated from field: string name = 10;
     */
    name: string;
    /**
     * Optional user-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 11;
     */
    metadata: string;
    constructor(data?: PartialMessage<CreateSIPTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateSIPTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPTrunkRequest;
    static equals(a: CreateSIPTrunkRequest | PlainMessage<CreateSIPTrunkRequest> | undefined, b: CreateSIPTrunkRequest | PlainMessage<CreateSIPTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPTrunkInfo
 * @deprecated
 */
export declare class SIPTrunkInfo extends Message<SIPTrunkInfo> {
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    /**
     * @generated from field: livekit.SIPTrunkInfo.TrunkKind kind = 14;
     */
    kind: SIPTrunkInfo_TrunkKind;
    /**
     * CIDR or IPs that traffic is accepted from
     * An empty list means all inbound traffic is accepted.
     *
     * @generated from field: repeated string inbound_addresses = 2;
     */
    inboundAddresses: string[];
    /**
     * IP that SIP INVITE is sent too
     *
     * @generated from field: string outbound_address = 3;
     */
    outboundAddress: string;
    /**
     * Number used to make outbound calls
     *
     * @generated from field: string outbound_number = 4;
     */
    outboundNumber: string;
    /**
     * Transport used for inbound and outbound calls.
     *
     * @generated from field: livekit.SIPTransport transport = 13;
     */
    transport: SIPTransport;
    /**
     * @generated from field: repeated string inbound_numbers_regex = 5 [deprecated = true];
     * @deprecated
     */
    inboundNumbersRegex: string[];
    /**
     * Accepted `To` values. This Trunk will only accept a call made to
     * these numbers. This allows you to have distinct Trunks for different phone
     * numbers at the same provider.
     *
     * @generated from field: repeated string inbound_numbers = 10;
     */
    inboundNumbers: string[];
    /**
     * Username and password used to authenticate inbound and outbound SIP invites
     * May be empty to have no Authentication
     *
     * @generated from field: string inbound_username = 6;
     */
    inboundUsername: string;
    /**
     * @generated from field: string inbound_password = 7;
     */
    inboundPassword: string;
    /**
     * @generated from field: string outbound_username = 8;
     */
    outboundUsername: string;
    /**
     * @generated from field: string outbound_password = 9;
     */
    outboundPassword: string;
    /**
     * Human-readable name for the Trunk.
     *
     * @generated from field: string name = 11;
     */
    name: string;
    /**
     * User-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 12;
     */
    metadata: string;
    constructor(data?: PartialMessage<SIPTrunkInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPTrunkInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPTrunkInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPTrunkInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPTrunkInfo;
    static equals(a: SIPTrunkInfo | PlainMessage<SIPTrunkInfo> | undefined, b: SIPTrunkInfo | PlainMessage<SIPTrunkInfo> | undefined): boolean;
}
/**
 * @generated from enum livekit.SIPTrunkInfo.TrunkKind
 */
export declare enum SIPTrunkInfo_TrunkKind {
    /**
     * @generated from enum value: TRUNK_LEGACY = 0;
     */
    TRUNK_LEGACY = 0,
    /**
     * @generated from enum value: TRUNK_INBOUND = 1;
     */
    TRUNK_INBOUND = 1,
    /**
     * @generated from enum value: TRUNK_OUTBOUND = 2;
     */
    TRUNK_OUTBOUND = 2
}
/**
 * @generated from message livekit.CreateSIPInboundTrunkRequest
 */
export declare class CreateSIPInboundTrunkRequest extends Message<CreateSIPInboundTrunkRequest> {
    /**
     * Trunk ID is ignored
     *
     * @generated from field: livekit.SIPInboundTrunkInfo trunk = 1;
     */
    trunk?: SIPInboundTrunkInfo;
    constructor(data?: PartialMessage<CreateSIPInboundTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateSIPInboundTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPInboundTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPInboundTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPInboundTrunkRequest;
    static equals(a: CreateSIPInboundTrunkRequest | PlainMessage<CreateSIPInboundTrunkRequest> | undefined, b: CreateSIPInboundTrunkRequest | PlainMessage<CreateSIPInboundTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPInboundTrunkInfo
 */
export declare class SIPInboundTrunkInfo extends Message<SIPInboundTrunkInfo> {
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    /**
     * Human-readable name for the Trunk.
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * User-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 3;
     */
    metadata: string;
    /**
     * Numbers associated with LiveKit SIP. The Trunk will only accept calls made to these numbers.
     * Creating multiple Trunks with different phone numbers allows having different rules for a single provider.
     *
     * @generated from field: repeated string numbers = 4;
     */
    numbers: string[];
    /**
     * CIDR or IPs that traffic is accepted from.
     * An empty list means all inbound traffic is accepted.
     *
     * @generated from field: repeated string allowed_addresses = 5;
     */
    allowedAddresses: string[];
    /**
     * Numbers that are allowed to make calls to this Trunk.
     * An empty list means calls from any phone number is accepted.
     *
     * @generated from field: repeated string allowed_numbers = 6;
     */
    allowedNumbers: string[];
    /**
     * Username and password used to authenticate inbound SIP invites.
     * May be empty to have no authentication.
     *
     * @generated from field: string auth_username = 7;
     */
    authUsername: string;
    /**
     * @generated from field: string auth_password = 8;
     */
    authPassword: string;
    /**
     * Include these SIP X-* headers in 200 OK responses.
     *
     * @generated from field: map<string, string> headers = 9;
     */
    headers: {
        [key: string]: string;
    };
    /**
     * Map SIP X-* headers from INVITE to SIP participant attributes.
     *
     * @generated from field: map<string, string> headers_to_attributes = 10;
     */
    headersToAttributes: {
        [key: string]: string;
    };
    /**
     * Map LiveKit attributes to SIP X-* headers when sending BYE or REFER requests.
     * Keys are the names of attributes and values are the names of X-* headers they will be mapped to.
     *
     * @generated from field: map<string, string> attributes_to_headers = 14;
     */
    attributesToHeaders: {
        [key: string]: string;
    };
    /**
     * Max time for the caller to wait for track subscription.
     *
     * @generated from field: google.protobuf.Duration ringing_timeout = 11;
     */
    ringingTimeout?: Duration;
    /**
     * Max call duration.
     *
     * @generated from field: google.protobuf.Duration max_call_duration = 12;
     */
    maxCallDuration?: Duration;
    /**
     * @generated from field: bool krisp_enabled = 13;
     */
    krispEnabled: boolean;
    constructor(data?: PartialMessage<SIPInboundTrunkInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPInboundTrunkInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPInboundTrunkInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPInboundTrunkInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPInboundTrunkInfo;
    static equals(a: SIPInboundTrunkInfo | PlainMessage<SIPInboundTrunkInfo> | undefined, b: SIPInboundTrunkInfo | PlainMessage<SIPInboundTrunkInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.CreateSIPOutboundTrunkRequest
 */
export declare class CreateSIPOutboundTrunkRequest extends Message<CreateSIPOutboundTrunkRequest> {
    /**
     * Trunk ID is ignored
     *
     * @generated from field: livekit.SIPOutboundTrunkInfo trunk = 1;
     */
    trunk?: SIPOutboundTrunkInfo;
    constructor(data?: PartialMessage<CreateSIPOutboundTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateSIPOutboundTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPOutboundTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPOutboundTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPOutboundTrunkRequest;
    static equals(a: CreateSIPOutboundTrunkRequest | PlainMessage<CreateSIPOutboundTrunkRequest> | undefined, b: CreateSIPOutboundTrunkRequest | PlainMessage<CreateSIPOutboundTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPOutboundTrunkInfo
 */
export declare class SIPOutboundTrunkInfo extends Message<SIPOutboundTrunkInfo> {
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    /**
     * Human-readable name for the Trunk.
     *
     * @generated from field: string name = 2;
     */
    name: string;
    /**
     * User-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 3;
     */
    metadata: string;
    /**
     * Hostname or IP that SIP INVITE is sent too.
     * Note that this is not a SIP URI and should not contain the 'sip:' protocol prefix.
     *
     * @generated from field: string address = 4;
     */
    address: string;
    /**
     * SIP Transport used for outbound call.
     *
     * @generated from field: livekit.SIPTransport transport = 5;
     */
    transport: SIPTransport;
    /**
     * Numbers used to make the calls. Random one from this list will be selected.
     *
     * @generated from field: repeated string numbers = 6;
     */
    numbers: string[];
    /**
     * Username and password used to authenticate with SIP server.
     * May be empty to have no authentication.
     *
     * @generated from field: string auth_username = 7;
     */
    authUsername: string;
    /**
     * @generated from field: string auth_password = 8;
     */
    authPassword: string;
    /**
     * Include these SIP X-* headers in INVITE request.
     * These headers are sent as-is and may help identify this call as coming from LiveKit for the other SIP endpoint.
     *
     * @generated from field: map<string, string> headers = 9;
     */
    headers: {
        [key: string]: string;
    };
    /**
     * Map SIP X-* headers from 200 OK to SIP participant attributes.
     * Keys are the names of X-* headers and values are the names of attributes they will be mapped to.
     *
     * @generated from field: map<string, string> headers_to_attributes = 10;
     */
    headersToAttributes: {
        [key: string]: string;
    };
    /**
     * Map LiveKit attributes to SIP X-* headers when sending BYE or REFER requests.
     * Keys are the names of attributes and values are the names of X-* headers they will be mapped to.
     *
     * @generated from field: map<string, string> attributes_to_headers = 11;
     */
    attributesToHeaders: {
        [key: string]: string;
    };
    constructor(data?: PartialMessage<SIPOutboundTrunkInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPOutboundTrunkInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPOutboundTrunkInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPOutboundTrunkInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPOutboundTrunkInfo;
    static equals(a: SIPOutboundTrunkInfo | PlainMessage<SIPOutboundTrunkInfo> | undefined, b: SIPOutboundTrunkInfo | PlainMessage<SIPOutboundTrunkInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.GetSIPInboundTrunkRequest
 */
export declare class GetSIPInboundTrunkRequest extends Message<GetSIPInboundTrunkRequest> {
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    constructor(data?: PartialMessage<GetSIPInboundTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.GetSIPInboundTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSIPInboundTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSIPInboundTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSIPInboundTrunkRequest;
    static equals(a: GetSIPInboundTrunkRequest | PlainMessage<GetSIPInboundTrunkRequest> | undefined, b: GetSIPInboundTrunkRequest | PlainMessage<GetSIPInboundTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.GetSIPInboundTrunkResponse
 */
export declare class GetSIPInboundTrunkResponse extends Message<GetSIPInboundTrunkResponse> {
    /**
     * @generated from field: livekit.SIPInboundTrunkInfo trunk = 1;
     */
    trunk?: SIPInboundTrunkInfo;
    constructor(data?: PartialMessage<GetSIPInboundTrunkResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.GetSIPInboundTrunkResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSIPInboundTrunkResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSIPInboundTrunkResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSIPInboundTrunkResponse;
    static equals(a: GetSIPInboundTrunkResponse | PlainMessage<GetSIPInboundTrunkResponse> | undefined, b: GetSIPInboundTrunkResponse | PlainMessage<GetSIPInboundTrunkResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.GetSIPOutboundTrunkRequest
 */
export declare class GetSIPOutboundTrunkRequest extends Message<GetSIPOutboundTrunkRequest> {
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    constructor(data?: PartialMessage<GetSIPOutboundTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.GetSIPOutboundTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSIPOutboundTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSIPOutboundTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSIPOutboundTrunkRequest;
    static equals(a: GetSIPOutboundTrunkRequest | PlainMessage<GetSIPOutboundTrunkRequest> | undefined, b: GetSIPOutboundTrunkRequest | PlainMessage<GetSIPOutboundTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.GetSIPOutboundTrunkResponse
 */
export declare class GetSIPOutboundTrunkResponse extends Message<GetSIPOutboundTrunkResponse> {
    /**
     * @generated from field: livekit.SIPOutboundTrunkInfo trunk = 1;
     */
    trunk?: SIPOutboundTrunkInfo;
    constructor(data?: PartialMessage<GetSIPOutboundTrunkResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.GetSIPOutboundTrunkResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSIPOutboundTrunkResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSIPOutboundTrunkResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSIPOutboundTrunkResponse;
    static equals(a: GetSIPOutboundTrunkResponse | PlainMessage<GetSIPOutboundTrunkResponse> | undefined, b: GetSIPOutboundTrunkResponse | PlainMessage<GetSIPOutboundTrunkResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPTrunkRequest
 * @deprecated
 */
export declare class ListSIPTrunkRequest extends Message<ListSIPTrunkRequest> {
    constructor(data?: PartialMessage<ListSIPTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPTrunkRequest;
    static equals(a: ListSIPTrunkRequest | PlainMessage<ListSIPTrunkRequest> | undefined, b: ListSIPTrunkRequest | PlainMessage<ListSIPTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPTrunkResponse
 * @deprecated
 */
export declare class ListSIPTrunkResponse extends Message<ListSIPTrunkResponse> {
    /**
     * @generated from field: repeated livekit.SIPTrunkInfo items = 1;
     */
    items: SIPTrunkInfo[];
    constructor(data?: PartialMessage<ListSIPTrunkResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPTrunkResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPTrunkResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPTrunkResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPTrunkResponse;
    static equals(a: ListSIPTrunkResponse | PlainMessage<ListSIPTrunkResponse> | undefined, b: ListSIPTrunkResponse | PlainMessage<ListSIPTrunkResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPInboundTrunkRequest
 */
export declare class ListSIPInboundTrunkRequest extends Message<ListSIPInboundTrunkRequest> {
    constructor(data?: PartialMessage<ListSIPInboundTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPInboundTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPInboundTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPInboundTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPInboundTrunkRequest;
    static equals(a: ListSIPInboundTrunkRequest | PlainMessage<ListSIPInboundTrunkRequest> | undefined, b: ListSIPInboundTrunkRequest | PlainMessage<ListSIPInboundTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPInboundTrunkResponse
 */
export declare class ListSIPInboundTrunkResponse extends Message<ListSIPInboundTrunkResponse> {
    /**
     * @generated from field: repeated livekit.SIPInboundTrunkInfo items = 1;
     */
    items: SIPInboundTrunkInfo[];
    constructor(data?: PartialMessage<ListSIPInboundTrunkResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPInboundTrunkResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPInboundTrunkResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPInboundTrunkResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPInboundTrunkResponse;
    static equals(a: ListSIPInboundTrunkResponse | PlainMessage<ListSIPInboundTrunkResponse> | undefined, b: ListSIPInboundTrunkResponse | PlainMessage<ListSIPInboundTrunkResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPOutboundTrunkRequest
 */
export declare class ListSIPOutboundTrunkRequest extends Message<ListSIPOutboundTrunkRequest> {
    constructor(data?: PartialMessage<ListSIPOutboundTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPOutboundTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPOutboundTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPOutboundTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPOutboundTrunkRequest;
    static equals(a: ListSIPOutboundTrunkRequest | PlainMessage<ListSIPOutboundTrunkRequest> | undefined, b: ListSIPOutboundTrunkRequest | PlainMessage<ListSIPOutboundTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPOutboundTrunkResponse
 */
export declare class ListSIPOutboundTrunkResponse extends Message<ListSIPOutboundTrunkResponse> {
    /**
     * @generated from field: repeated livekit.SIPOutboundTrunkInfo items = 1;
     */
    items: SIPOutboundTrunkInfo[];
    constructor(data?: PartialMessage<ListSIPOutboundTrunkResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPOutboundTrunkResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPOutboundTrunkResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPOutboundTrunkResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPOutboundTrunkResponse;
    static equals(a: ListSIPOutboundTrunkResponse | PlainMessage<ListSIPOutboundTrunkResponse> | undefined, b: ListSIPOutboundTrunkResponse | PlainMessage<ListSIPOutboundTrunkResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.DeleteSIPTrunkRequest
 */
export declare class DeleteSIPTrunkRequest extends Message<DeleteSIPTrunkRequest> {
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    constructor(data?: PartialMessage<DeleteSIPTrunkRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DeleteSIPTrunkRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSIPTrunkRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSIPTrunkRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSIPTrunkRequest;
    static equals(a: DeleteSIPTrunkRequest | PlainMessage<DeleteSIPTrunkRequest> | undefined, b: DeleteSIPTrunkRequest | PlainMessage<DeleteSIPTrunkRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPDispatchRuleDirect
 */
export declare class SIPDispatchRuleDirect extends Message<SIPDispatchRuleDirect> {
    /**
     * What room should call be directed into
     *
     * @generated from field: string room_name = 1;
     */
    roomName: string;
    /**
     * Optional pin required to enter room
     *
     * @generated from field: string pin = 2;
     */
    pin: string;
    constructor(data?: PartialMessage<SIPDispatchRuleDirect>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPDispatchRuleDirect";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleDirect;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleDirect;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleDirect;
    static equals(a: SIPDispatchRuleDirect | PlainMessage<SIPDispatchRuleDirect> | undefined, b: SIPDispatchRuleDirect | PlainMessage<SIPDispatchRuleDirect> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPDispatchRuleIndividual
 */
export declare class SIPDispatchRuleIndividual extends Message<SIPDispatchRuleIndividual> {
    /**
     * Prefix used on new room name
     *
     * @generated from field: string room_prefix = 1;
     */
    roomPrefix: string;
    /**
     * Optional pin required to enter room
     *
     * @generated from field: string pin = 2;
     */
    pin: string;
    constructor(data?: PartialMessage<SIPDispatchRuleIndividual>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPDispatchRuleIndividual";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleIndividual;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleIndividual;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleIndividual;
    static equals(a: SIPDispatchRuleIndividual | PlainMessage<SIPDispatchRuleIndividual> | undefined, b: SIPDispatchRuleIndividual | PlainMessage<SIPDispatchRuleIndividual> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPDispatchRuleCallee
 */
export declare class SIPDispatchRuleCallee extends Message<SIPDispatchRuleCallee> {
    /**
     * Prefix used on new room name
     *
     * @generated from field: string room_prefix = 1;
     */
    roomPrefix: string;
    /**
     * Optional pin required to enter room
     *
     * @generated from field: string pin = 2;
     */
    pin: string;
    /**
     * Optionally append random suffix
     *
     * @generated from field: bool randomize = 3;
     */
    randomize: boolean;
    constructor(data?: PartialMessage<SIPDispatchRuleCallee>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPDispatchRuleCallee";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleCallee;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleCallee;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleCallee;
    static equals(a: SIPDispatchRuleCallee | PlainMessage<SIPDispatchRuleCallee> | undefined, b: SIPDispatchRuleCallee | PlainMessage<SIPDispatchRuleCallee> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPDispatchRule
 */
export declare class SIPDispatchRule extends Message<SIPDispatchRule> {
    /**
     * @generated from oneof livekit.SIPDispatchRule.rule
     */
    rule: {
        /**
         * SIPDispatchRuleDirect is a `SIP Dispatch Rule` that puts a user directly into a room
         * This places users into an existing room. Optionally you can require a pin before a user can
         * enter the room
         *
         * @generated from field: livekit.SIPDispatchRuleDirect dispatch_rule_direct = 1;
         */
        value: SIPDispatchRuleDirect;
        case: "dispatchRuleDirect";
    } | {
        /**
         * SIPDispatchRuleIndividual is a `SIP Dispatch Rule` that creates a new room for each caller.
         *
         * @generated from field: livekit.SIPDispatchRuleIndividual dispatch_rule_individual = 2;
         */
        value: SIPDispatchRuleIndividual;
        case: "dispatchRuleIndividual";
    } | {
        /**
         * SIPDispatchRuleCallee is a `SIP Dispatch Rule` that creates a new room for each callee.
         *
         * @generated from field: livekit.SIPDispatchRuleCallee dispatch_rule_callee = 3;
         */
        value: SIPDispatchRuleCallee;
        case: "dispatchRuleCallee";
    } | {
        case: undefined;
        value?: undefined;
    };
    constructor(data?: PartialMessage<SIPDispatchRule>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPDispatchRule";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRule;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRule;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRule;
    static equals(a: SIPDispatchRule | PlainMessage<SIPDispatchRule> | undefined, b: SIPDispatchRule | PlainMessage<SIPDispatchRule> | undefined): boolean;
}
/**
 * @generated from message livekit.CreateSIPDispatchRuleRequest
 */
export declare class CreateSIPDispatchRuleRequest extends Message<CreateSIPDispatchRuleRequest> {
    /**
     * @generated from field: livekit.SIPDispatchRule rule = 1;
     */
    rule?: SIPDispatchRule;
    /**
     * What trunks are accepted for this dispatch rule
     * If empty all trunks will match this dispatch rule
     *
     * @generated from field: repeated string trunk_ids = 2;
     */
    trunkIds: string[];
    /**
     * By default the From value (Phone number) is used for participant name/identity and added to attributes.
     * If true, a random value for identity will be used and numbers will be omitted from attributes.
     *
     * @generated from field: bool hide_phone_number = 3;
     */
    hidePhoneNumber: boolean;
    /**
     * Dispatch Rule will only accept a call made to these numbers (if set).
     *
     * @generated from field: repeated string inbound_numbers = 6;
     */
    inboundNumbers: string[];
    /**
     * Optional human-readable name for the Dispatch Rule.
     *
     * @generated from field: string name = 4;
     */
    name: string;
    /**
     * User-defined metadata for the Dispatch Rule.
     * Participants created by this rule will inherit this metadata.
     *
     * @generated from field: string metadata = 5;
     */
    metadata: string;
    /**
     * User-defined attributes for the Dispatch Rule.
     * Participants created by this rule will inherit these attributes.
     *
     * @generated from field: map<string, string> attributes = 7;
     */
    attributes: {
        [key: string]: string;
    };
    constructor(data?: PartialMessage<CreateSIPDispatchRuleRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateSIPDispatchRuleRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPDispatchRuleRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPDispatchRuleRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPDispatchRuleRequest;
    static equals(a: CreateSIPDispatchRuleRequest | PlainMessage<CreateSIPDispatchRuleRequest> | undefined, b: CreateSIPDispatchRuleRequest | PlainMessage<CreateSIPDispatchRuleRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPDispatchRuleInfo
 */
export declare class SIPDispatchRuleInfo extends Message<SIPDispatchRuleInfo> {
    /**
     * @generated from field: string sip_dispatch_rule_id = 1;
     */
    sipDispatchRuleId: string;
    /**
     * @generated from field: livekit.SIPDispatchRule rule = 2;
     */
    rule?: SIPDispatchRule;
    /**
     * @generated from field: repeated string trunk_ids = 3;
     */
    trunkIds: string[];
    /**
     * @generated from field: bool hide_phone_number = 4;
     */
    hidePhoneNumber: boolean;
    /**
     * Dispatch Rule will only accept a call made to these numbers (if set).
     *
     * @generated from field: repeated string inbound_numbers = 7;
     */
    inboundNumbers: string[];
    /**
     * Human-readable name for the Dispatch Rule.
     *
     * @generated from field: string name = 5;
     */
    name: string;
    /**
     * User-defined metadata for the Dispatch Rule.
     * Participants created by this rule will inherit this metadata.
     *
     * @generated from field: string metadata = 6;
     */
    metadata: string;
    /**
     * User-defined attributes for the Dispatch Rule.
     * Participants created by this rule will inherit these attributes.
     *
     * @generated from field: map<string, string> attributes = 8;
     */
    attributes: {
        [key: string]: string;
    };
    constructor(data?: PartialMessage<SIPDispatchRuleInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPDispatchRuleInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleInfo;
    static equals(a: SIPDispatchRuleInfo | PlainMessage<SIPDispatchRuleInfo> | undefined, b: SIPDispatchRuleInfo | PlainMessage<SIPDispatchRuleInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPDispatchRuleRequest
 */
export declare class ListSIPDispatchRuleRequest extends Message<ListSIPDispatchRuleRequest> {
    constructor(data?: PartialMessage<ListSIPDispatchRuleRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPDispatchRuleRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPDispatchRuleRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleRequest;
    static equals(a: ListSIPDispatchRuleRequest | PlainMessage<ListSIPDispatchRuleRequest> | undefined, b: ListSIPDispatchRuleRequest | PlainMessage<ListSIPDispatchRuleRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.ListSIPDispatchRuleResponse
 */
export declare class ListSIPDispatchRuleResponse extends Message<ListSIPDispatchRuleResponse> {
    /**
     * @generated from field: repeated livekit.SIPDispatchRuleInfo items = 1;
     */
    items: SIPDispatchRuleInfo[];
    constructor(data?: PartialMessage<ListSIPDispatchRuleResponse>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.ListSIPDispatchRuleResponse";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPDispatchRuleResponse;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleResponse;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleResponse;
    static equals(a: ListSIPDispatchRuleResponse | PlainMessage<ListSIPDispatchRuleResponse> | undefined, b: ListSIPDispatchRuleResponse | PlainMessage<ListSIPDispatchRuleResponse> | undefined): boolean;
}
/**
 * @generated from message livekit.DeleteSIPDispatchRuleRequest
 */
export declare class DeleteSIPDispatchRuleRequest extends Message<DeleteSIPDispatchRuleRequest> {
    /**
     * @generated from field: string sip_dispatch_rule_id = 1;
     */
    sipDispatchRuleId: string;
    constructor(data?: PartialMessage<DeleteSIPDispatchRuleRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.DeleteSIPDispatchRuleRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSIPDispatchRuleRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSIPDispatchRuleRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSIPDispatchRuleRequest;
    static equals(a: DeleteSIPDispatchRuleRequest | PlainMessage<DeleteSIPDispatchRuleRequest> | undefined, b: DeleteSIPDispatchRuleRequest | PlainMessage<DeleteSIPDispatchRuleRequest> | undefined): boolean;
}
/**
 * A SIP Participant is a singular SIP session connected to a LiveKit room via
 * a SIP Trunk into a SIP DispatchRule
 *
 * @generated from message livekit.CreateSIPParticipantRequest
 */
export declare class CreateSIPParticipantRequest extends Message<CreateSIPParticipantRequest> {
    /**
     * What SIP Trunk should be used to dial the user
     *
     * @generated from field: string sip_trunk_id = 1;
     */
    sipTrunkId: string;
    /**
     * What number should be dialed via SIP
     *
     * @generated from field: string sip_call_to = 2;
     */
    sipCallTo: string;
    /**
     * Optional SIP From number to use. If empty, trunk number is used.
     *
     * @generated from field: string sip_number = 15;
     */
    sipNumber: string;
    /**
     * What LiveKit room should this participant be connected too
     *
     * @generated from field: string room_name = 3;
     */
    roomName: string;
    /**
     * Optional identity of the participant in LiveKit room
     *
     * @generated from field: string participant_identity = 4;
     */
    participantIdentity: string;
    /**
     * Optional name of the participant in LiveKit room
     *
     * @generated from field: string participant_name = 7;
     */
    participantName: string;
    /**
     * Optional user-defined metadata. Will be attached to a created Participant in the room.
     *
     * @generated from field: string participant_metadata = 8;
     */
    participantMetadata: string;
    /**
     * Optional user-defined attributes. Will be attached to a created Participant in the room.
     *
     * @generated from field: map<string, string> participant_attributes = 9;
     */
    participantAttributes: {
        [key: string]: string;
    };
    /**
     * Optionally send following DTMF digits (extension codes) when making a call.
     * Character 'w' can be used to add a 0.5 sec delay.
     *
     * @generated from field: string dtmf = 5;
     */
    dtmf: string;
    /**
     * Optionally play dialtone in the room as an audible indicator for existing participants. The `play_ringtone` option is deprectated but has the same effect.
     *
     * @generated from field: bool play_ringtone = 6 [deprecated = true];
     * @deprecated
     */
    playRingtone: boolean;
    /**
     * @generated from field: bool play_dialtone = 13;
     */
    playDialtone: boolean;
    /**
     * By default the From value (Phone number) is used for participant name/identity (if not set) and added to attributes.
     * If true, a random value for identity will be used and numbers will be omitted from attributes.
     *
     * @generated from field: bool hide_phone_number = 10;
     */
    hidePhoneNumber: boolean;
    /**
     * Max time for the callee to answer the call.
     *
     * @generated from field: google.protobuf.Duration ringing_timeout = 11;
     */
    ringingTimeout?: Duration;
    /**
     * Max call duration.
     *
     * @generated from field: google.protobuf.Duration max_call_duration = 12;
     */
    maxCallDuration?: Duration;
    /**
     * Enable voice isolation for the callee.
     *
     * NEXT ID: 16
     *
     * @generated from field: bool enable_krisp = 14;
     */
    enableKrisp: boolean;
    constructor(data?: PartialMessage<CreateSIPParticipantRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.CreateSIPParticipantRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPParticipantRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPParticipantRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPParticipantRequest;
    static equals(a: CreateSIPParticipantRequest | PlainMessage<CreateSIPParticipantRequest> | undefined, b: CreateSIPParticipantRequest | PlainMessage<CreateSIPParticipantRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPParticipantInfo
 */
export declare class SIPParticipantInfo extends Message<SIPParticipantInfo> {
    /**
     * @generated from field: string participant_id = 1;
     */
    participantId: string;
    /**
     * @generated from field: string participant_identity = 2;
     */
    participantIdentity: string;
    /**
     * @generated from field: string room_name = 3;
     */
    roomName: string;
    /**
     * @generated from field: string sip_call_id = 4;
     */
    sipCallId: string;
    constructor(data?: PartialMessage<SIPParticipantInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPParticipantInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPParticipantInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPParticipantInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPParticipantInfo;
    static equals(a: SIPParticipantInfo | PlainMessage<SIPParticipantInfo> | undefined, b: SIPParticipantInfo | PlainMessage<SIPParticipantInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.TransferSIPParticipantRequest
 */
export declare class TransferSIPParticipantRequest extends Message<TransferSIPParticipantRequest> {
    /**
     * @generated from field: string participant_identity = 1;
     */
    participantIdentity: string;
    /**
     * @generated from field: string room_name = 2;
     */
    roomName: string;
    /**
     * @generated from field: string transfer_to = 3;
     */
    transferTo: string;
    /**
     * Optionally play dialtone to the SIP participant as an audible indicator of being transferred
     *
     * @generated from field: bool play_dialtone = 4;
     */
    playDialtone: boolean;
    constructor(data?: PartialMessage<TransferSIPParticipantRequest>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.TransferSIPParticipantRequest";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TransferSIPParticipantRequest;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TransferSIPParticipantRequest;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TransferSIPParticipantRequest;
    static equals(a: TransferSIPParticipantRequest | PlainMessage<TransferSIPParticipantRequest> | undefined, b: TransferSIPParticipantRequest | PlainMessage<TransferSIPParticipantRequest> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPCallInfo
 */
export declare class SIPCallInfo extends Message<SIPCallInfo> {
    /**
     * @generated from field: string call_id = 1;
     */
    callId: string;
    /**
     * @generated from field: string trunk_id = 2;
     */
    trunkId: string;
    /**
     * @generated from field: string room_name = 3;
     */
    roomName: string;
    /**
     * ID of the current/previous room published to
     *
     * @generated from field: string room_id = 4;
     */
    roomId: string;
    /**
     * @generated from field: string participant_identity = 5;
     */
    participantIdentity: string;
    /**
     * @generated from field: livekit.SIPUri from_uri = 6;
     */
    fromUri?: SIPUri;
    /**
     * @generated from field: livekit.SIPUri to_uri = 7;
     */
    toUri?: SIPUri;
    /**
     * @generated from field: repeated livekit.SIPFeature enabled_features = 14;
     */
    enabledFeatures: SIPFeature[];
    /**
     * @generated from field: livekit.SIPCallStatus call_status = 8;
     */
    callStatus: SIPCallStatus;
    /**
     * @generated from field: int64 created_at = 9;
     */
    createdAt: bigint;
    /**
     * @generated from field: int64 started_at = 10;
     */
    startedAt: bigint;
    /**
     * @generated from field: int64 ended_at = 11;
     */
    endedAt: bigint;
    /**
     * @generated from field: livekit.DisconnectReason disconnect_reason = 12;
     */
    disconnectReason: DisconnectReason;
    /**
     * @generated from field: string error = 13;
     */
    error: string;
    constructor(data?: PartialMessage<SIPCallInfo>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPCallInfo";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPCallInfo;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPCallInfo;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPCallInfo;
    static equals(a: SIPCallInfo | PlainMessage<SIPCallInfo> | undefined, b: SIPCallInfo | PlainMessage<SIPCallInfo> | undefined): boolean;
}
/**
 * @generated from message livekit.SIPUri
 */
export declare class SIPUri extends Message<SIPUri> {
    /**
     * @generated from field: string user = 1;
     */
    user: string;
    /**
     * @generated from field: string host = 2;
     */
    host: string;
    /**
     * @generated from field: string ip = 3;
     */
    ip: string;
    /**
     * @generated from field: uint32 port = 4;
     */
    port: number;
    /**
     * @generated from field: livekit.SIPTransport transport = 5;
     */
    transport: SIPTransport;
    constructor(data?: PartialMessage<SIPUri>);
    static readonly runtime: typeof proto3;
    static readonly typeName = "livekit.SIPUri";
    static readonly fields: FieldList;
    static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPUri;
    static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPUri;
    static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPUri;
    static equals(a: SIPUri | PlainMessage<SIPUri> | undefined, b: SIPUri | PlainMessage<SIPUri> | undefined): boolean;
}
//# sourceMappingURL=livekit_sip_pb.d.ts.map