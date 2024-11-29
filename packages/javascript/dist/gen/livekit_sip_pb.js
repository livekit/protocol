import { Duration, Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { DisconnectReason } from "./livekit_models_pb.js";
var SIPTransport = /* @__PURE__ */ ((SIPTransport2) => {
  SIPTransport2[SIPTransport2["SIP_TRANSPORT_AUTO"] = 0] = "SIP_TRANSPORT_AUTO";
  SIPTransport2[SIPTransport2["SIP_TRANSPORT_UDP"] = 1] = "SIP_TRANSPORT_UDP";
  SIPTransport2[SIPTransport2["SIP_TRANSPORT_TCP"] = 2] = "SIP_TRANSPORT_TCP";
  SIPTransport2[SIPTransport2["SIP_TRANSPORT_TLS"] = 3] = "SIP_TRANSPORT_TLS";
  return SIPTransport2;
})(SIPTransport || {});
proto3.util.setEnumType(SIPTransport, "livekit.SIPTransport", [
  { no: 0, name: "SIP_TRANSPORT_AUTO" },
  { no: 1, name: "SIP_TRANSPORT_UDP" },
  { no: 2, name: "SIP_TRANSPORT_TCP" },
  { no: 3, name: "SIP_TRANSPORT_TLS" }
]);
var SIPCallStatus = /* @__PURE__ */ ((SIPCallStatus2) => {
  SIPCallStatus2[SIPCallStatus2["SCS_CALL_INCOMING"] = 0] = "SCS_CALL_INCOMING";
  SIPCallStatus2[SIPCallStatus2["SCS_PARTICIPANT_JOINED"] = 1] = "SCS_PARTICIPANT_JOINED";
  SIPCallStatus2[SIPCallStatus2["SCS_ACTIVE"] = 2] = "SCS_ACTIVE";
  SIPCallStatus2[SIPCallStatus2["SCS_DISCONNECTED"] = 3] = "SCS_DISCONNECTED";
  SIPCallStatus2[SIPCallStatus2["SCS_ERROR"] = 4] = "SCS_ERROR";
  return SIPCallStatus2;
})(SIPCallStatus || {});
proto3.util.setEnumType(SIPCallStatus, "livekit.SIPCallStatus", [
  { no: 0, name: "SCS_CALL_INCOMING" },
  { no: 1, name: "SCS_PARTICIPANT_JOINED" },
  { no: 2, name: "SCS_ACTIVE" },
  { no: 3, name: "SCS_DISCONNECTED" },
  { no: 4, name: "SCS_ERROR" }
]);
var SIPFeature = /* @__PURE__ */ ((SIPFeature2) => {
  SIPFeature2[SIPFeature2["NONE"] = 0] = "NONE";
  SIPFeature2[SIPFeature2["KRISP_ENABLED"] = 1] = "KRISP_ENABLED";
  return SIPFeature2;
})(SIPFeature || {});
proto3.util.setEnumType(SIPFeature, "livekit.SIPFeature", [
  { no: 0, name: "NONE" },
  { no: 1, name: "KRISP_ENABLED" }
]);
const _CreateSIPTrunkRequest = class _CreateSIPTrunkRequest extends Message {
  constructor(data) {
    super();
    /**
     * CIDR or IPs that traffic is accepted from
     * An empty list means all inbound traffic is accepted.
     *
     * @generated from field: repeated string inbound_addresses = 1;
     */
    this.inboundAddresses = [];
    /**
     * IP that SIP INVITE is sent too
     *
     * @generated from field: string outbound_address = 2;
     */
    this.outboundAddress = "";
    /**
     * Number used to make outbound calls
     *
     * @generated from field: string outbound_number = 3;
     */
    this.outboundNumber = "";
    /**
     * @generated from field: repeated string inbound_numbers_regex = 4 [deprecated = true];
     * @deprecated
     */
    this.inboundNumbersRegex = [];
    /**
     * Accepted `To` values. This Trunk will only accept a call made to
     * these numbers. This allows you to have distinct Trunks for different phone
     * numbers at the same provider.
     *
     * @generated from field: repeated string inbound_numbers = 9;
     */
    this.inboundNumbers = [];
    /**
     * Username and password used to authenticate inbound and outbound SIP invites
     * May be empty to have no Authentication
     *
     * @generated from field: string inbound_username = 5;
     */
    this.inboundUsername = "";
    /**
     * @generated from field: string inbound_password = 6;
     */
    this.inboundPassword = "";
    /**
     * @generated from field: string outbound_username = 7;
     */
    this.outboundUsername = "";
    /**
     * @generated from field: string outbound_password = 8;
     */
    this.outboundPassword = "";
    /**
     * Optional human-readable name for the Trunk.
     *
     * @generated from field: string name = 10;
     */
    this.name = "";
    /**
     * Optional user-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 11;
     */
    this.metadata = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateSIPTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateSIPTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateSIPTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateSIPTrunkRequest, a, b);
  }
};
_CreateSIPTrunkRequest.runtime = proto3;
_CreateSIPTrunkRequest.typeName = "livekit.CreateSIPTrunkRequest";
_CreateSIPTrunkRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "inbound_addresses", kind: "scalar", T: 9, repeated: true },
  {
    no: 2,
    name: "outbound_address",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "outbound_number",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "inbound_numbers_regex", kind: "scalar", T: 9, repeated: true },
  { no: 9, name: "inbound_numbers", kind: "scalar", T: 9, repeated: true },
  {
    no: 5,
    name: "inbound_username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "inbound_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "outbound_username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "outbound_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 10,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 11,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let CreateSIPTrunkRequest = _CreateSIPTrunkRequest;
const _SIPTrunkInfo = class _SIPTrunkInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    /**
     * @generated from field: livekit.SIPTrunkInfo.TrunkKind kind = 14;
     */
    this.kind = 0 /* TRUNK_LEGACY */;
    /**
     * CIDR or IPs that traffic is accepted from
     * An empty list means all inbound traffic is accepted.
     *
     * @generated from field: repeated string inbound_addresses = 2;
     */
    this.inboundAddresses = [];
    /**
     * IP that SIP INVITE is sent too
     *
     * @generated from field: string outbound_address = 3;
     */
    this.outboundAddress = "";
    /**
     * Number used to make outbound calls
     *
     * @generated from field: string outbound_number = 4;
     */
    this.outboundNumber = "";
    /**
     * Transport used for inbound and outbound calls.
     *
     * @generated from field: livekit.SIPTransport transport = 13;
     */
    this.transport = 0 /* SIP_TRANSPORT_AUTO */;
    /**
     * @generated from field: repeated string inbound_numbers_regex = 5 [deprecated = true];
     * @deprecated
     */
    this.inboundNumbersRegex = [];
    /**
     * Accepted `To` values. This Trunk will only accept a call made to
     * these numbers. This allows you to have distinct Trunks for different phone
     * numbers at the same provider.
     *
     * @generated from field: repeated string inbound_numbers = 10;
     */
    this.inboundNumbers = [];
    /**
     * Username and password used to authenticate inbound and outbound SIP invites
     * May be empty to have no Authentication
     *
     * @generated from field: string inbound_username = 6;
     */
    this.inboundUsername = "";
    /**
     * @generated from field: string inbound_password = 7;
     */
    this.inboundPassword = "";
    /**
     * @generated from field: string outbound_username = 8;
     */
    this.outboundUsername = "";
    /**
     * @generated from field: string outbound_password = 9;
     */
    this.outboundPassword = "";
    /**
     * Human-readable name for the Trunk.
     *
     * @generated from field: string name = 11;
     */
    this.name = "";
    /**
     * User-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 12;
     */
    this.metadata = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPTrunkInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPTrunkInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPTrunkInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPTrunkInfo, a, b);
  }
};
_SIPTrunkInfo.runtime = proto3;
_SIPTrunkInfo.typeName = "livekit.SIPTrunkInfo";
_SIPTrunkInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 14, name: "kind", kind: "enum", T: proto3.getEnumType(SIPTrunkInfo_TrunkKind) },
  { no: 2, name: "inbound_addresses", kind: "scalar", T: 9, repeated: true },
  {
    no: 3,
    name: "outbound_address",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "outbound_number",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 13, name: "transport", kind: "enum", T: proto3.getEnumType(SIPTransport) },
  { no: 5, name: "inbound_numbers_regex", kind: "scalar", T: 9, repeated: true },
  { no: 10, name: "inbound_numbers", kind: "scalar", T: 9, repeated: true },
  {
    no: 6,
    name: "inbound_username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "inbound_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "outbound_username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 9,
    name: "outbound_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 11,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 12,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SIPTrunkInfo = _SIPTrunkInfo;
var SIPTrunkInfo_TrunkKind = /* @__PURE__ */ ((SIPTrunkInfo_TrunkKind2) => {
  SIPTrunkInfo_TrunkKind2[SIPTrunkInfo_TrunkKind2["TRUNK_LEGACY"] = 0] = "TRUNK_LEGACY";
  SIPTrunkInfo_TrunkKind2[SIPTrunkInfo_TrunkKind2["TRUNK_INBOUND"] = 1] = "TRUNK_INBOUND";
  SIPTrunkInfo_TrunkKind2[SIPTrunkInfo_TrunkKind2["TRUNK_OUTBOUND"] = 2] = "TRUNK_OUTBOUND";
  return SIPTrunkInfo_TrunkKind2;
})(SIPTrunkInfo_TrunkKind || {});
proto3.util.setEnumType(SIPTrunkInfo_TrunkKind, "livekit.SIPTrunkInfo.TrunkKind", [
  { no: 0, name: "TRUNK_LEGACY" },
  { no: 1, name: "TRUNK_INBOUND" },
  { no: 2, name: "TRUNK_OUTBOUND" }
]);
const _CreateSIPInboundTrunkRequest = class _CreateSIPInboundTrunkRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateSIPInboundTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateSIPInboundTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateSIPInboundTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateSIPInboundTrunkRequest, a, b);
  }
};
_CreateSIPInboundTrunkRequest.runtime = proto3;
_CreateSIPInboundTrunkRequest.typeName = "livekit.CreateSIPInboundTrunkRequest";
_CreateSIPInboundTrunkRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "trunk", kind: "message", T: SIPInboundTrunkInfo }
]);
let CreateSIPInboundTrunkRequest = _CreateSIPInboundTrunkRequest;
const _SIPInboundTrunkInfo = class _SIPInboundTrunkInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    /**
     * Human-readable name for the Trunk.
     *
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * User-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 3;
     */
    this.metadata = "";
    /**
     * Numbers associated with LiveKit SIP. The Trunk will only accept calls made to these numbers.
     * Creating multiple Trunks with different phone numbers allows having different rules for a single provider.
     *
     * @generated from field: repeated string numbers = 4;
     */
    this.numbers = [];
    /**
     * CIDR or IPs that traffic is accepted from.
     * An empty list means all inbound traffic is accepted.
     *
     * @generated from field: repeated string allowed_addresses = 5;
     */
    this.allowedAddresses = [];
    /**
     * Numbers that are allowed to make calls to this Trunk.
     * An empty list means calls from any phone number is accepted.
     *
     * @generated from field: repeated string allowed_numbers = 6;
     */
    this.allowedNumbers = [];
    /**
     * Username and password used to authenticate inbound SIP invites.
     * May be empty to have no authentication.
     *
     * @generated from field: string auth_username = 7;
     */
    this.authUsername = "";
    /**
     * @generated from field: string auth_password = 8;
     */
    this.authPassword = "";
    /**
     * Include these SIP X-* headers in 200 OK responses.
     *
     * @generated from field: map<string, string> headers = 9;
     */
    this.headers = {};
    /**
     * Map SIP X-* headers from INVITE to SIP participant attributes.
     *
     * @generated from field: map<string, string> headers_to_attributes = 10;
     */
    this.headersToAttributes = {};
    /**
     * Map LiveKit attributes to SIP X-* headers when sending BYE or REFER requests.
     * Keys are the names of attributes and values are the names of X-* headers they will be mapped to.
     *
     * @generated from field: map<string, string> attributes_to_headers = 14;
     */
    this.attributesToHeaders = {};
    /**
     * @generated from field: bool krisp_enabled = 13;
     */
    this.krispEnabled = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPInboundTrunkInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPInboundTrunkInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPInboundTrunkInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPInboundTrunkInfo, a, b);
  }
};
_SIPInboundTrunkInfo.runtime = proto3;
_SIPInboundTrunkInfo.typeName = "livekit.SIPInboundTrunkInfo";
_SIPInboundTrunkInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
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
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "numbers", kind: "scalar", T: 9, repeated: true },
  { no: 5, name: "allowed_addresses", kind: "scalar", T: 9, repeated: true },
  { no: 6, name: "allowed_numbers", kind: "scalar", T: 9, repeated: true },
  {
    no: 7,
    name: "auth_username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "auth_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 9, name: "headers", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 10, name: "headers_to_attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 14, name: "attributes_to_headers", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 11, name: "ringing_timeout", kind: "message", T: Duration },
  { no: 12, name: "max_call_duration", kind: "message", T: Duration },
  {
    no: 13,
    name: "krisp_enabled",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let SIPInboundTrunkInfo = _SIPInboundTrunkInfo;
const _CreateSIPOutboundTrunkRequest = class _CreateSIPOutboundTrunkRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateSIPOutboundTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateSIPOutboundTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateSIPOutboundTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateSIPOutboundTrunkRequest, a, b);
  }
};
_CreateSIPOutboundTrunkRequest.runtime = proto3;
_CreateSIPOutboundTrunkRequest.typeName = "livekit.CreateSIPOutboundTrunkRequest";
_CreateSIPOutboundTrunkRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "trunk", kind: "message", T: SIPOutboundTrunkInfo }
]);
let CreateSIPOutboundTrunkRequest = _CreateSIPOutboundTrunkRequest;
const _SIPOutboundTrunkInfo = class _SIPOutboundTrunkInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    /**
     * Human-readable name for the Trunk.
     *
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * User-defined metadata for the Trunk.
     *
     * @generated from field: string metadata = 3;
     */
    this.metadata = "";
    /**
     * Hostname or IP that SIP INVITE is sent too.
     * Note that this is not a SIP URI and should not contain the 'sip:' protocol prefix.
     *
     * @generated from field: string address = 4;
     */
    this.address = "";
    /**
     * SIP Transport used for outbound call.
     *
     * @generated from field: livekit.SIPTransport transport = 5;
     */
    this.transport = 0 /* SIP_TRANSPORT_AUTO */;
    /**
     * Numbers used to make the calls. Random one from this list will be selected.
     *
     * @generated from field: repeated string numbers = 6;
     */
    this.numbers = [];
    /**
     * Username and password used to authenticate with SIP server.
     * May be empty to have no authentication.
     *
     * @generated from field: string auth_username = 7;
     */
    this.authUsername = "";
    /**
     * @generated from field: string auth_password = 8;
     */
    this.authPassword = "";
    /**
     * Include these SIP X-* headers in INVITE request.
     * These headers are sent as-is and may help identify this call as coming from LiveKit for the other SIP endpoint.
     *
     * @generated from field: map<string, string> headers = 9;
     */
    this.headers = {};
    /**
     * Map SIP X-* headers from 200 OK to SIP participant attributes.
     * Keys are the names of X-* headers and values are the names of attributes they will be mapped to.
     *
     * @generated from field: map<string, string> headers_to_attributes = 10;
     */
    this.headersToAttributes = {};
    /**
     * Map LiveKit attributes to SIP X-* headers when sending BYE or REFER requests.
     * Keys are the names of attributes and values are the names of X-* headers they will be mapped to.
     *
     * @generated from field: map<string, string> attributes_to_headers = 11;
     */
    this.attributesToHeaders = {};
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPOutboundTrunkInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPOutboundTrunkInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPOutboundTrunkInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPOutboundTrunkInfo, a, b);
  }
};
_SIPOutboundTrunkInfo.runtime = proto3;
_SIPOutboundTrunkInfo.typeName = "livekit.SIPOutboundTrunkInfo";
_SIPOutboundTrunkInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
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
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "address",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "transport", kind: "enum", T: proto3.getEnumType(SIPTransport) },
  { no: 6, name: "numbers", kind: "scalar", T: 9, repeated: true },
  {
    no: 7,
    name: "auth_username",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "auth_password",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 9, name: "headers", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 10, name: "headers_to_attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  { no: 11, name: "attributes_to_headers", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } }
]);
let SIPOutboundTrunkInfo = _SIPOutboundTrunkInfo;
const _GetSIPInboundTrunkRequest = class _GetSIPInboundTrunkRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetSIPInboundTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetSIPInboundTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetSIPInboundTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_GetSIPInboundTrunkRequest, a, b);
  }
};
_GetSIPInboundTrunkRequest.runtime = proto3;
_GetSIPInboundTrunkRequest.typeName = "livekit.GetSIPInboundTrunkRequest";
_GetSIPInboundTrunkRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let GetSIPInboundTrunkRequest = _GetSIPInboundTrunkRequest;
const _GetSIPInboundTrunkResponse = class _GetSIPInboundTrunkResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetSIPInboundTrunkResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetSIPInboundTrunkResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetSIPInboundTrunkResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_GetSIPInboundTrunkResponse, a, b);
  }
};
_GetSIPInboundTrunkResponse.runtime = proto3;
_GetSIPInboundTrunkResponse.typeName = "livekit.GetSIPInboundTrunkResponse";
_GetSIPInboundTrunkResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "trunk", kind: "message", T: SIPInboundTrunkInfo }
]);
let GetSIPInboundTrunkResponse = _GetSIPInboundTrunkResponse;
const _GetSIPOutboundTrunkRequest = class _GetSIPOutboundTrunkRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetSIPOutboundTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetSIPOutboundTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetSIPOutboundTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_GetSIPOutboundTrunkRequest, a, b);
  }
};
_GetSIPOutboundTrunkRequest.runtime = proto3;
_GetSIPOutboundTrunkRequest.typeName = "livekit.GetSIPOutboundTrunkRequest";
_GetSIPOutboundTrunkRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let GetSIPOutboundTrunkRequest = _GetSIPOutboundTrunkRequest;
const _GetSIPOutboundTrunkResponse = class _GetSIPOutboundTrunkResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetSIPOutboundTrunkResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetSIPOutboundTrunkResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetSIPOutboundTrunkResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_GetSIPOutboundTrunkResponse, a, b);
  }
};
_GetSIPOutboundTrunkResponse.runtime = proto3;
_GetSIPOutboundTrunkResponse.typeName = "livekit.GetSIPOutboundTrunkResponse";
_GetSIPOutboundTrunkResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "trunk", kind: "message", T: SIPOutboundTrunkInfo }
]);
let GetSIPOutboundTrunkResponse = _GetSIPOutboundTrunkResponse;
const _ListSIPTrunkRequest = class _ListSIPTrunkRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPTrunkRequest, a, b);
  }
};
_ListSIPTrunkRequest.runtime = proto3;
_ListSIPTrunkRequest.typeName = "livekit.ListSIPTrunkRequest";
_ListSIPTrunkRequest.fields = proto3.util.newFieldList(() => []);
let ListSIPTrunkRequest = _ListSIPTrunkRequest;
const _ListSIPTrunkResponse = class _ListSIPTrunkResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.SIPTrunkInfo items = 1;
     */
    this.items = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPTrunkResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPTrunkResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPTrunkResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPTrunkResponse, a, b);
  }
};
_ListSIPTrunkResponse.runtime = proto3;
_ListSIPTrunkResponse.typeName = "livekit.ListSIPTrunkResponse";
_ListSIPTrunkResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "items", kind: "message", T: SIPTrunkInfo, repeated: true }
]);
let ListSIPTrunkResponse = _ListSIPTrunkResponse;
const _ListSIPInboundTrunkRequest = class _ListSIPInboundTrunkRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPInboundTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPInboundTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPInboundTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPInboundTrunkRequest, a, b);
  }
};
_ListSIPInboundTrunkRequest.runtime = proto3;
_ListSIPInboundTrunkRequest.typeName = "livekit.ListSIPInboundTrunkRequest";
_ListSIPInboundTrunkRequest.fields = proto3.util.newFieldList(() => []);
let ListSIPInboundTrunkRequest = _ListSIPInboundTrunkRequest;
const _ListSIPInboundTrunkResponse = class _ListSIPInboundTrunkResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.SIPInboundTrunkInfo items = 1;
     */
    this.items = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPInboundTrunkResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPInboundTrunkResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPInboundTrunkResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPInboundTrunkResponse, a, b);
  }
};
_ListSIPInboundTrunkResponse.runtime = proto3;
_ListSIPInboundTrunkResponse.typeName = "livekit.ListSIPInboundTrunkResponse";
_ListSIPInboundTrunkResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "items", kind: "message", T: SIPInboundTrunkInfo, repeated: true }
]);
let ListSIPInboundTrunkResponse = _ListSIPInboundTrunkResponse;
const _ListSIPOutboundTrunkRequest = class _ListSIPOutboundTrunkRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPOutboundTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPOutboundTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPOutboundTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPOutboundTrunkRequest, a, b);
  }
};
_ListSIPOutboundTrunkRequest.runtime = proto3;
_ListSIPOutboundTrunkRequest.typeName = "livekit.ListSIPOutboundTrunkRequest";
_ListSIPOutboundTrunkRequest.fields = proto3.util.newFieldList(() => []);
let ListSIPOutboundTrunkRequest = _ListSIPOutboundTrunkRequest;
const _ListSIPOutboundTrunkResponse = class _ListSIPOutboundTrunkResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.SIPOutboundTrunkInfo items = 1;
     */
    this.items = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPOutboundTrunkResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPOutboundTrunkResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPOutboundTrunkResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPOutboundTrunkResponse, a, b);
  }
};
_ListSIPOutboundTrunkResponse.runtime = proto3;
_ListSIPOutboundTrunkResponse.typeName = "livekit.ListSIPOutboundTrunkResponse";
_ListSIPOutboundTrunkResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "items", kind: "message", T: SIPOutboundTrunkInfo, repeated: true }
]);
let ListSIPOutboundTrunkResponse = _ListSIPOutboundTrunkResponse;
const _DeleteSIPTrunkRequest = class _DeleteSIPTrunkRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteSIPTrunkRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteSIPTrunkRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteSIPTrunkRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_DeleteSIPTrunkRequest, a, b);
  }
};
_DeleteSIPTrunkRequest.runtime = proto3;
_DeleteSIPTrunkRequest.typeName = "livekit.DeleteSIPTrunkRequest";
_DeleteSIPTrunkRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let DeleteSIPTrunkRequest = _DeleteSIPTrunkRequest;
const _SIPDispatchRuleDirect = class _SIPDispatchRuleDirect extends Message {
  constructor(data) {
    super();
    /**
     * What room should call be directed into
     *
     * @generated from field: string room_name = 1;
     */
    this.roomName = "";
    /**
     * Optional pin required to enter room
     *
     * @generated from field: string pin = 2;
     */
    this.pin = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPDispatchRuleDirect().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPDispatchRuleDirect().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPDispatchRuleDirect().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPDispatchRuleDirect, a, b);
  }
};
_SIPDispatchRuleDirect.runtime = proto3;
_SIPDispatchRuleDirect.typeName = "livekit.SIPDispatchRuleDirect";
_SIPDispatchRuleDirect.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "pin",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SIPDispatchRuleDirect = _SIPDispatchRuleDirect;
const _SIPDispatchRuleIndividual = class _SIPDispatchRuleIndividual extends Message {
  constructor(data) {
    super();
    /**
     * Prefix used on new room name
     *
     * @generated from field: string room_prefix = 1;
     */
    this.roomPrefix = "";
    /**
     * Optional pin required to enter room
     *
     * @generated from field: string pin = 2;
     */
    this.pin = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPDispatchRuleIndividual().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPDispatchRuleIndividual().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPDispatchRuleIndividual().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPDispatchRuleIndividual, a, b);
  }
};
_SIPDispatchRuleIndividual.runtime = proto3;
_SIPDispatchRuleIndividual.typeName = "livekit.SIPDispatchRuleIndividual";
_SIPDispatchRuleIndividual.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_prefix",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "pin",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SIPDispatchRuleIndividual = _SIPDispatchRuleIndividual;
const _SIPDispatchRuleCallee = class _SIPDispatchRuleCallee extends Message {
  constructor(data) {
    super();
    /**
     * Prefix used on new room name
     *
     * @generated from field: string room_prefix = 1;
     */
    this.roomPrefix = "";
    /**
     * Optional pin required to enter room
     *
     * @generated from field: string pin = 2;
     */
    this.pin = "";
    /**
     * Optionally append random suffix
     *
     * @generated from field: bool randomize = 3;
     */
    this.randomize = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPDispatchRuleCallee().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPDispatchRuleCallee().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPDispatchRuleCallee().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPDispatchRuleCallee, a, b);
  }
};
_SIPDispatchRuleCallee.runtime = proto3;
_SIPDispatchRuleCallee.typeName = "livekit.SIPDispatchRuleCallee";
_SIPDispatchRuleCallee.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "room_prefix",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "pin",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "randomize",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let SIPDispatchRuleCallee = _SIPDispatchRuleCallee;
const _SIPDispatchRule = class _SIPDispatchRule extends Message {
  constructor(data) {
    super();
    /**
     * @generated from oneof livekit.SIPDispatchRule.rule
     */
    this.rule = { case: void 0 };
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPDispatchRule().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPDispatchRule().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPDispatchRule().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPDispatchRule, a, b);
  }
};
_SIPDispatchRule.runtime = proto3;
_SIPDispatchRule.typeName = "livekit.SIPDispatchRule";
_SIPDispatchRule.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "dispatch_rule_direct", kind: "message", T: SIPDispatchRuleDirect, oneof: "rule" },
  { no: 2, name: "dispatch_rule_individual", kind: "message", T: SIPDispatchRuleIndividual, oneof: "rule" },
  { no: 3, name: "dispatch_rule_callee", kind: "message", T: SIPDispatchRuleCallee, oneof: "rule" }
]);
let SIPDispatchRule = _SIPDispatchRule;
const _CreateSIPDispatchRuleRequest = class _CreateSIPDispatchRuleRequest extends Message {
  constructor(data) {
    super();
    /**
     * What trunks are accepted for this dispatch rule
     * If empty all trunks will match this dispatch rule
     *
     * @generated from field: repeated string trunk_ids = 2;
     */
    this.trunkIds = [];
    /**
     * By default the From value (Phone number) is used for participant name/identity and added to attributes.
     * If true, a random value for identity will be used and numbers will be omitted from attributes.
     *
     * @generated from field: bool hide_phone_number = 3;
     */
    this.hidePhoneNumber = false;
    /**
     * Dispatch Rule will only accept a call made to these numbers (if set).
     *
     * @generated from field: repeated string inbound_numbers = 6;
     */
    this.inboundNumbers = [];
    /**
     * Optional human-readable name for the Dispatch Rule.
     *
     * @generated from field: string name = 4;
     */
    this.name = "";
    /**
     * User-defined metadata for the Dispatch Rule.
     * Participants created by this rule will inherit this metadata.
     *
     * @generated from field: string metadata = 5;
     */
    this.metadata = "";
    /**
     * User-defined attributes for the Dispatch Rule.
     * Participants created by this rule will inherit these attributes.
     *
     * @generated from field: map<string, string> attributes = 7;
     */
    this.attributes = {};
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateSIPDispatchRuleRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateSIPDispatchRuleRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateSIPDispatchRuleRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateSIPDispatchRuleRequest, a, b);
  }
};
_CreateSIPDispatchRuleRequest.runtime = proto3;
_CreateSIPDispatchRuleRequest.typeName = "livekit.CreateSIPDispatchRuleRequest";
_CreateSIPDispatchRuleRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "rule", kind: "message", T: SIPDispatchRule },
  { no: 2, name: "trunk_ids", kind: "scalar", T: 9, repeated: true },
  {
    no: 3,
    name: "hide_phone_number",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 6, name: "inbound_numbers", kind: "scalar", T: 9, repeated: true },
  {
    no: 4,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 7, name: "attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } }
]);
let CreateSIPDispatchRuleRequest = _CreateSIPDispatchRuleRequest;
const _SIPDispatchRuleInfo = class _SIPDispatchRuleInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_dispatch_rule_id = 1;
     */
    this.sipDispatchRuleId = "";
    /**
     * @generated from field: repeated string trunk_ids = 3;
     */
    this.trunkIds = [];
    /**
     * @generated from field: bool hide_phone_number = 4;
     */
    this.hidePhoneNumber = false;
    /**
     * Dispatch Rule will only accept a call made to these numbers (if set).
     *
     * @generated from field: repeated string inbound_numbers = 7;
     */
    this.inboundNumbers = [];
    /**
     * Human-readable name for the Dispatch Rule.
     *
     * @generated from field: string name = 5;
     */
    this.name = "";
    /**
     * User-defined metadata for the Dispatch Rule.
     * Participants created by this rule will inherit this metadata.
     *
     * @generated from field: string metadata = 6;
     */
    this.metadata = "";
    /**
     * User-defined attributes for the Dispatch Rule.
     * Participants created by this rule will inherit these attributes.
     *
     * @generated from field: map<string, string> attributes = 8;
     */
    this.attributes = {};
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPDispatchRuleInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPDispatchRuleInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPDispatchRuleInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPDispatchRuleInfo, a, b);
  }
};
_SIPDispatchRuleInfo.runtime = proto3;
_SIPDispatchRuleInfo.typeName = "livekit.SIPDispatchRuleInfo";
_SIPDispatchRuleInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_dispatch_rule_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "rule", kind: "message", T: SIPDispatchRule },
  { no: 3, name: "trunk_ids", kind: "scalar", T: 9, repeated: true },
  {
    no: 4,
    name: "hide_phone_number",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 7, name: "inbound_numbers", kind: "scalar", T: 9, repeated: true },
  {
    no: 5,
    name: "name",
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
  { no: 8, name: "attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } }
]);
let SIPDispatchRuleInfo = _SIPDispatchRuleInfo;
const _ListSIPDispatchRuleRequest = class _ListSIPDispatchRuleRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPDispatchRuleRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPDispatchRuleRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPDispatchRuleRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPDispatchRuleRequest, a, b);
  }
};
_ListSIPDispatchRuleRequest.runtime = proto3;
_ListSIPDispatchRuleRequest.typeName = "livekit.ListSIPDispatchRuleRequest";
_ListSIPDispatchRuleRequest.fields = proto3.util.newFieldList(() => []);
let ListSIPDispatchRuleRequest = _ListSIPDispatchRuleRequest;
const _ListSIPDispatchRuleResponse = class _ListSIPDispatchRuleResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated livekit.SIPDispatchRuleInfo items = 1;
     */
    this.items = [];
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListSIPDispatchRuleResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListSIPDispatchRuleResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListSIPDispatchRuleResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_ListSIPDispatchRuleResponse, a, b);
  }
};
_ListSIPDispatchRuleResponse.runtime = proto3;
_ListSIPDispatchRuleResponse.typeName = "livekit.ListSIPDispatchRuleResponse";
_ListSIPDispatchRuleResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "items", kind: "message", T: SIPDispatchRuleInfo, repeated: true }
]);
let ListSIPDispatchRuleResponse = _ListSIPDispatchRuleResponse;
const _DeleteSIPDispatchRuleRequest = class _DeleteSIPDispatchRuleRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string sip_dispatch_rule_id = 1;
     */
    this.sipDispatchRuleId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteSIPDispatchRuleRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteSIPDispatchRuleRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteSIPDispatchRuleRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_DeleteSIPDispatchRuleRequest, a, b);
  }
};
_DeleteSIPDispatchRuleRequest.runtime = proto3;
_DeleteSIPDispatchRuleRequest.typeName = "livekit.DeleteSIPDispatchRuleRequest";
_DeleteSIPDispatchRuleRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_dispatch_rule_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let DeleteSIPDispatchRuleRequest = _DeleteSIPDispatchRuleRequest;
const _CreateSIPParticipantRequest = class _CreateSIPParticipantRequest extends Message {
  constructor(data) {
    super();
    /**
     * What SIP Trunk should be used to dial the user
     *
     * @generated from field: string sip_trunk_id = 1;
     */
    this.sipTrunkId = "";
    /**
     * What number should be dialed via SIP
     *
     * @generated from field: string sip_call_to = 2;
     */
    this.sipCallTo = "";
    /**
     * Optional SIP From number to use. If empty, trunk number is used.
     *
     * @generated from field: string sip_number = 15;
     */
    this.sipNumber = "";
    /**
     * What LiveKit room should this participant be connected too
     *
     * @generated from field: string room_name = 3;
     */
    this.roomName = "";
    /**
     * Optional identity of the participant in LiveKit room
     *
     * @generated from field: string participant_identity = 4;
     */
    this.participantIdentity = "";
    /**
     * Optional name of the participant in LiveKit room
     *
     * @generated from field: string participant_name = 7;
     */
    this.participantName = "";
    /**
     * Optional user-defined metadata. Will be attached to a created Participant in the room.
     *
     * @generated from field: string participant_metadata = 8;
     */
    this.participantMetadata = "";
    /**
     * Optional user-defined attributes. Will be attached to a created Participant in the room.
     *
     * @generated from field: map<string, string> participant_attributes = 9;
     */
    this.participantAttributes = {};
    /**
     * Optionally send following DTMF digits (extension codes) when making a call.
     * Character 'w' can be used to add a 0.5 sec delay.
     *
     * @generated from field: string dtmf = 5;
     */
    this.dtmf = "";
    /**
     * Optionally play dialtone in the room as an audible indicator for existing participants. The `play_ringtone` option is deprectated but has the same effect.
     *
     * @generated from field: bool play_ringtone = 6 [deprecated = true];
     * @deprecated
     */
    this.playRingtone = false;
    /**
     * @generated from field: bool play_dialtone = 13;
     */
    this.playDialtone = false;
    /**
     * By default the From value (Phone number) is used for participant name/identity (if not set) and added to attributes.
     * If true, a random value for identity will be used and numbers will be omitted from attributes.
     *
     * @generated from field: bool hide_phone_number = 10;
     */
    this.hidePhoneNumber = false;
    /**
     * Enable voice isolation for the callee.
     *
     * NEXT ID: 16
     *
     * @generated from field: bool enable_krisp = 14;
     */
    this.enableKrisp = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateSIPParticipantRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateSIPParticipantRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateSIPParticipantRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_CreateSIPParticipantRequest, a, b);
  }
};
_CreateSIPParticipantRequest.runtime = proto3;
_CreateSIPParticipantRequest.typeName = "livekit.CreateSIPParticipantRequest";
_CreateSIPParticipantRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "sip_trunk_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "sip_call_to",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 15,
    name: "sip_number",
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
    no: 7,
    name: "participant_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "participant_metadata",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 9, name: "participant_attributes", kind: "map", K: 9, V: {
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  } },
  {
    no: 5,
    name: "dtmf",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "play_ringtone",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 13,
    name: "play_dialtone",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 10,
    name: "hide_phone_number",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 11, name: "ringing_timeout", kind: "message", T: Duration },
  { no: 12, name: "max_call_duration", kind: "message", T: Duration },
  {
    no: 14,
    name: "enable_krisp",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let CreateSIPParticipantRequest = _CreateSIPParticipantRequest;
const _SIPParticipantInfo = class _SIPParticipantInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string participant_id = 1;
     */
    this.participantId = "";
    /**
     * @generated from field: string participant_identity = 2;
     */
    this.participantIdentity = "";
    /**
     * @generated from field: string room_name = 3;
     */
    this.roomName = "";
    /**
     * @generated from field: string sip_call_id = 4;
     */
    this.sipCallId = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPParticipantInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPParticipantInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPParticipantInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPParticipantInfo, a, b);
  }
};
_SIPParticipantInfo.runtime = proto3;
_SIPParticipantInfo.typeName = "livekit.SIPParticipantInfo";
_SIPParticipantInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "participant_identity",
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
    name: "sip_call_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SIPParticipantInfo = _SIPParticipantInfo;
const _TransferSIPParticipantRequest = class _TransferSIPParticipantRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string participant_identity = 1;
     */
    this.participantIdentity = "";
    /**
     * @generated from field: string room_name = 2;
     */
    this.roomName = "";
    /**
     * @generated from field: string transfer_to = 3;
     */
    this.transferTo = "";
    /**
     * Optionally play dialtone to the SIP participant as an audible indicator of being transferred
     *
     * @generated from field: bool play_dialtone = 4;
     */
    this.playDialtone = false;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _TransferSIPParticipantRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _TransferSIPParticipantRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _TransferSIPParticipantRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_TransferSIPParticipantRequest, a, b);
  }
};
_TransferSIPParticipantRequest.runtime = proto3;
_TransferSIPParticipantRequest.typeName = "livekit.TransferSIPParticipantRequest";
_TransferSIPParticipantRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "participant_identity",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "room_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "transfer_to",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "play_dialtone",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
let TransferSIPParticipantRequest = _TransferSIPParticipantRequest;
const _SIPCallInfo = class _SIPCallInfo extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string call_id = 1;
     */
    this.callId = "";
    /**
     * @generated from field: string trunk_id = 2;
     */
    this.trunkId = "";
    /**
     * @generated from field: string room_name = 3;
     */
    this.roomName = "";
    /**
     * ID of the current/previous room published to
     *
     * @generated from field: string room_id = 4;
     */
    this.roomId = "";
    /**
     * @generated from field: string participant_identity = 5;
     */
    this.participantIdentity = "";
    /**
     * @generated from field: repeated livekit.SIPFeature enabled_features = 14;
     */
    this.enabledFeatures = [];
    /**
     * @generated from field: livekit.SIPCallStatus call_status = 8;
     */
    this.callStatus = 0 /* SCS_CALL_INCOMING */;
    /**
     * @generated from field: int64 created_at = 9;
     */
    this.createdAt = protoInt64.zero;
    /**
     * @generated from field: int64 started_at = 10;
     */
    this.startedAt = protoInt64.zero;
    /**
     * @generated from field: int64 ended_at = 11;
     */
    this.endedAt = protoInt64.zero;
    /**
     * @generated from field: livekit.DisconnectReason disconnect_reason = 12;
     */
    this.disconnectReason = DisconnectReason.UNKNOWN_REASON;
    /**
     * @generated from field: string error = 13;
     */
    this.error = "";
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPCallInfo().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPCallInfo().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPCallInfo().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPCallInfo, a, b);
  }
};
_SIPCallInfo.runtime = proto3;
_SIPCallInfo.typeName = "livekit.SIPCallInfo";
_SIPCallInfo.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "call_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "trunk_id",
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
    name: "room_id",
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
  { no: 6, name: "from_uri", kind: "message", T: SIPUri },
  { no: 7, name: "to_uri", kind: "message", T: SIPUri },
  { no: 14, name: "enabled_features", kind: "enum", T: proto3.getEnumType(SIPFeature), repeated: true },
  { no: 8, name: "call_status", kind: "enum", T: proto3.getEnumType(SIPCallStatus) },
  {
    no: 9,
    name: "created_at",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
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
  { no: 12, name: "disconnect_reason", kind: "enum", T: proto3.getEnumType(DisconnectReason) },
  {
    no: 13,
    name: "error",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
let SIPCallInfo = _SIPCallInfo;
const _SIPUri = class _SIPUri extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string user = 1;
     */
    this.user = "";
    /**
     * @generated from field: string host = 2;
     */
    this.host = "";
    /**
     * @generated from field: string ip = 3;
     */
    this.ip = "";
    /**
     * @generated from field: uint32 port = 4;
     */
    this.port = 0;
    /**
     * @generated from field: livekit.SIPTransport transport = 5;
     */
    this.transport = 0 /* SIP_TRANSPORT_AUTO */;
    proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SIPUri().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SIPUri().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SIPUri().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_SIPUri, a, b);
  }
};
_SIPUri.runtime = proto3;
_SIPUri.typeName = "livekit.SIPUri";
_SIPUri.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "user",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "host",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "ip",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "port",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 5, name: "transport", kind: "enum", T: proto3.getEnumType(SIPTransport) }
]);
let SIPUri = _SIPUri;
export {
  CreateSIPDispatchRuleRequest,
  CreateSIPInboundTrunkRequest,
  CreateSIPOutboundTrunkRequest,
  CreateSIPParticipantRequest,
  CreateSIPTrunkRequest,
  DeleteSIPDispatchRuleRequest,
  DeleteSIPTrunkRequest,
  GetSIPInboundTrunkRequest,
  GetSIPInboundTrunkResponse,
  GetSIPOutboundTrunkRequest,
  GetSIPOutboundTrunkResponse,
  ListSIPDispatchRuleRequest,
  ListSIPDispatchRuleResponse,
  ListSIPInboundTrunkRequest,
  ListSIPInboundTrunkResponse,
  ListSIPOutboundTrunkRequest,
  ListSIPOutboundTrunkResponse,
  ListSIPTrunkRequest,
  ListSIPTrunkResponse,
  SIPCallInfo,
  SIPCallStatus,
  SIPDispatchRule,
  SIPDispatchRuleCallee,
  SIPDispatchRuleDirect,
  SIPDispatchRuleIndividual,
  SIPDispatchRuleInfo,
  SIPFeature,
  SIPInboundTrunkInfo,
  SIPOutboundTrunkInfo,
  SIPParticipantInfo,
  SIPTransport,
  SIPTrunkInfo,
  SIPTrunkInfo_TrunkKind,
  SIPUri,
  TransferSIPParticipantRequest
};
//# sourceMappingURL=livekit_sip_pb.js.map