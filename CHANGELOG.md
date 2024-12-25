# github.com/livekit/protocol

## 1.30.0

### Minor Changes

- Add filters to SIP list APIs. - [#913](https://github.com/livekit/protocol/pull/913) ([@dennwc](https://github.com/dennwc))

- Add headers for SIP call transfer. - [#926](https://github.com/livekit/protocol/pull/926) ([@dennwc](https://github.com/dennwc))

- Add headers for CreateSipParticipant - [#911](https://github.com/livekit/protocol/pull/911) ([@s-hamdananwar](https://github.com/s-hamdananwar))

- Allow mapping all SIP headers to attributes. - [#920](https://github.com/livekit/protocol/pull/920) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Allow SIP inbound to specify room configuration - [#923](https://github.com/livekit/protocol/pull/923) ([@davidzhao](https://github.com/davidzhao))

- Relax SIP header validation. - [#925](https://github.com/livekit/protocol/pull/925) ([@dennwc](https://github.com/dennwc))

## 1.29.4

## 1.29.3

## 1.29.2

### Patch Changes

- Revert back to Go 1.22 - [#905](https://github.com/livekit/protocol/pull/905) ([@davidzhao](https://github.com/davidzhao))

## 1.29.1

## 1.29.0

### Minor Changes

- Add DataStream support - [#871](https://github.com/livekit/protocol/pull/871) ([@lukasIO](https://github.com/lukasIO))

## 1.28.1

## 1.28.0

### Minor Changes

- Add disconnect reason for SIP failures. - [#845](https://github.com/livekit/protocol/pull/845) ([@dennwc](https://github.com/dennwc))

- Map participant attributes to SIP headers. - [#893](https://github.com/livekit/protocol/pull/893) ([@dennwc](https://github.com/dennwc))

- Allow setting number for SIP outbound. - [#891](https://github.com/livekit/protocol/pull/891) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Annotate SIP errors with Twirp codes. - [#879](https://github.com/livekit/protocol/pull/879) ([@dennwc](https://github.com/dennwc))

- Fix port type for SIPUri. - [#882](https://github.com/livekit/protocol/pull/882) ([@dennwc](https://github.com/dennwc))

- Require number, CIDR or auth for SIP inbound. - [#890](https://github.com/livekit/protocol/pull/890) ([@dennwc](https://github.com/dennwc))

## 1.27.1

### Patch Changes

- Add extra fields to SIP analytics events - [#872](https://github.com/livekit/protocol/pull/872) ([@biglittlebigben](https://github.com/biglittlebigben))

## 1.27.0

### Minor Changes

- Support for room configuration and agent dispatches - [#864](https://github.com/livekit/protocol/pull/864) ([@davidzhao](https://github.com/davidzhao))

### Patch Changes

- added manifest details to egress info - [#862](https://github.com/livekit/protocol/pull/862) ([@frostbyte73](https://github.com/frostbyte73))

- Type safe IP checks for SIP Trunks. - [#857](https://github.com/livekit/protocol/pull/857) ([@dennwc](https://github.com/dennwc))

- enable krisp SIP setting - [#866](https://github.com/livekit/protocol/pull/866) ([@frostbyte73](https://github.com/frostbyte73))

## 1.26.0

### Minor Changes

- allow Agents to pass through initial attributes - [#852](https://github.com/livekit/protocol/pull/852) ([@davidzhao](https://github.com/davidzhao))

## 1.25.0

### Minor Changes

- Add ringing timeout and max call duration setting for SIP. - [#844](https://github.com/livekit/protocol/pull/844) ([@dennwc](https://github.com/dennwc))

## 1.24.0

### Minor Changes

- Add disconnect reasons for SIP. - [#828](https://github.com/livekit/protocol/pull/828) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Add missing SIP status attribute. - [#833](https://github.com/livekit/protocol/pull/833) ([@dennwc](https://github.com/dennwc))

- Validate SipCallTo in CreateSIPParticipantRequest. - [#823](https://github.com/livekit/protocol/pull/823) ([@dennwc](https://github.com/dennwc))

## 1.23.0

### Patch Changes

- Add protocols for client metrics - [#821](https://github.com/livekit/protocol/pull/821) ([@davidliu](https://github.com/davidliu))

- Add other_sdks field to ClientInfo - [#804](https://github.com/livekit/protocol/pull/804) ([@bcherry](https://github.com/bcherry))

## 1.22.0

### Minor Changes

- Update SIP protocol. Pass headers, project ID and hostname. - [#805](https://github.com/livekit/protocol/pull/805) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Add generated flag to ChatMessage - [#813](https://github.com/livekit/protocol/pull/813) ([@lukasIO](https://github.com/lukasIO))

- Add dedicated chat message definition - [#785](https://github.com/livekit/protocol/pull/785) ([@lukasIO](https://github.com/lukasIO))

- Update @bufbuild/protobuf dependency - [#800](https://github.com/livekit/protocol/pull/800) ([@lukasIO](https://github.com/lukasIO))

- Metrics server side timestamp - [#806](https://github.com/livekit/protocol/pull/806) ([@boks1971](https://github.com/boks1971))

## 1.21.0

### Minor Changes

- Add Callee dispatch rule type for SIP. - [#798](https://github.com/livekit/protocol/pull/798) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Add SDK values for Unity-Web and NodeJS. - [#797](https://github.com/livekit/protocol/pull/797) ([@bcherry](https://github.com/bcherry))

## 1.20.1

### Patch Changes

- Add fast_publish option to JoinResponse - [#796](https://github.com/livekit/protocol/pull/796) ([@cnderrauber](https://github.com/cnderrauber))

## 1.20.0

### Minor Changes

- Send non-error responses also as part of RequestResponse - [#787](https://github.com/livekit/protocol/pull/787) ([@lukasIO](https://github.com/lukasIO))

### Patch Changes

- Add disconnect reason in participant info. - [#788](https://github.com/livekit/protocol/pull/788) ([@boks1971](https://github.com/boks1971))

- Use multiple webhook workers for each URL to improve parallelism - [#784](https://github.com/livekit/protocol/pull/784) ([@davidzhao](https://github.com/davidzhao))

## 1.19.3

## 1.19.2

### Patch Changes

- Add FilenamePrefix to ImagesInfo - [#751](https://github.com/livekit/protocol/pull/751) ([@frostbyte73](https://github.com/frostbyte73))

- Add experimental replay api - [#755](https://github.com/livekit/protocol/pull/755) ([@frostbyte73](https://github.com/frostbyte73))

- Implement a custom YAML unmarshaller for RoomConfiguration - [#765](https://github.com/livekit/protocol/pull/765) ([@biglittlebigben](https://github.com/biglittlebigben))

- Add AgentDispatchPrefix to guid - [#757](https://github.com/livekit/protocol/pull/757) ([@biglittlebigben](https://github.com/biglittlebigben))

- Ability to set attributes upon job acceptance - [#769](https://github.com/livekit/protocol/pull/769) ([@davidzhao](https://github.com/davidzhao))

- Remove unused fields from RegisterWorkerRequest - [#761](https://github.com/livekit/protocol/pull/761) ([@biglittlebigben](https://github.com/biglittlebigben))

- Integrate feedback on the agents protocol: - [#750](https://github.com/livekit/protocol/pull/750) ([@biglittlebigben](https://github.com/biglittlebigben))

  - Rename JobDescription to AgentDispatch
  - Remove participant_identity entry in the dispatch
  - Deprecate namespace
  - Add an agent_name field to specify what agent workers a job should be dispatched to

  Also allow setting a room configuration in the token.

- add srt output for egress - [#766](https://github.com/livekit/protocol/pull/766) ([@frostbyte73](https://github.com/frostbyte73))

- Internal agent protocol improvements - [#764](https://github.com/livekit/protocol/pull/764) ([@biglittlebigben](https://github.com/biglittlebigben))

- Implement MarshalYAML on RoomConfiguration and related types - [#771](https://github.com/livekit/protocol/pull/771) ([@biglittlebigben](https://github.com/biglittlebigben))

- Use consistent json field name for room configuration grant - [#762](https://github.com/livekit/protocol/pull/762) ([@biglittlebigben](https://github.com/biglittlebigben))

- Added RoomClosed as a DisconnectReason - [#778](https://github.com/livekit/protocol/pull/778) ([@davidzhao](https://github.com/davidzhao))

## 1.19.1

## 1.19.0

### Minor Changes

- Add SIP grants. - [#745](https://github.com/livekit/protocol/pull/745) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Fix typo in SIP Trunk conversion. - [#744](https://github.com/livekit/protocol/pull/744) ([@dennwc](https://github.com/dennwc))

## 1.18.0

### Minor Changes

- Split SIP Trunk configuration to inbound and outbound parts. - [#738](https://github.com/livekit/protocol/pull/738) ([@dennwc](https://github.com/dennwc))

- add Participant attributes key/val storage - [#733](https://github.com/livekit/protocol/pull/733) ([@davidzhao](https://github.com/davidzhao))

### Patch Changes

- Simplify number matching rules for SIP. - [#737](https://github.com/livekit/protocol/pull/737) ([@dennwc](https://github.com/dennwc))

- Added session token option for s3 uploads - [#734](https://github.com/livekit/protocol/pull/734) ([@frostbyte73](https://github.com/frostbyte73))

- Include analytics event ids - [#739](https://github.com/livekit/protocol/pull/739) ([@davidzhao](https://github.com/davidzhao))

## 1.17.0

### Minor Changes

- Add support for additional SIP transports. - [#719](https://github.com/livekit/protocol/pull/719) ([@dennwc](https://github.com/dennwc))

- Add ICERestartWHIPResource to RPC - [#716](https://github.com/livekit/protocol/pull/716) ([@Sean-Der](https://github.com/Sean-Der))

### Patch Changes

- Added error code field to EgressInfo - [#714](https://github.com/livekit/protocol/pull/714) ([@frostbyte73](https://github.com/frostbyte73))

- Include node_id with analytics events - [#727](https://github.com/livekit/protocol/pull/727) ([@davidzhao](https://github.com/davidzhao))

- Match SIP Trunks by source IP or mask. - [#724](https://github.com/livekit/protocol/pull/724) ([@dennwc](https://github.com/dennwc))

- Add notice to use participant identity from DataPacket for transcriptions. - [#706](https://github.com/livekit/protocol/pull/706) ([@dennwc](https://github.com/dennwc))

## 1.16.0

### Minor Changes

- - Add RedactAutoEncodedOutput to support redacting auto egress types - [#712](https://github.com/livekit/protocol/pull/712) ([@biglittlebigben](https://github.com/biglittlebigben))
  - Redact image outputs
  - Support AliOSS uploads for auto egress

- Allow inbound number filtering on SIP Dispatch Rules - [#707](https://github.com/livekit/protocol/pull/707) ([@dennwc](https://github.com/dennwc))

- Move egress request redacting routines to protocol - [#711](https://github.com/livekit/protocol/pull/711) ([@biglittlebigben](https://github.com/biglittlebigben))

## 1.15.0

### Minor Changes

- Expose Logger constructor for Zap. - [#701](https://github.com/livekit/protocol/pull/701) ([@dennwc](https://github.com/dennwc))

- Add metadata to SIP trunks, dispatch rules and participants. Change SIP participant identity prefix to `sip_`. - [#696](https://github.com/livekit/protocol/pull/696) ([@dennwc](https://github.com/dennwc))

- Move language into TranscriptionSegment #703 - [#704](https://github.com/livekit/protocol/pull/704) ([@lukasIO](https://github.com/lukasIO))

## 1.14.0

### Minor Changes

- Added real-time Transcription protocol - [#686](https://github.com/livekit/protocol/pull/686) ([@davidzhao](https://github.com/davidzhao))

- WHIP protocol change - [#680](https://github.com/livekit/protocol/pull/680) ([@biglittlebigben](https://github.com/biglittlebigben))

  Deprecate the bypass_transcoding property in all ingress APIs and introduce the optional enable_transcoding property. This property will default to false for WHIP and to true for all other ingress types.

### Patch Changes

- Add SIP participant name. - [#687](https://github.com/livekit/protocol/pull/687) ([@dennwc](https://github.com/dennwc))

## 1.13.0

### Minor Changes

- Add and option to play ringtone for SIP outbound calls. - [#671](https://github.com/livekit/protocol/pull/671) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Add initial support for slog. - [#668](https://github.com/livekit/protocol/pull/668) ([@dennwc](https://github.com/dennwc))

## 1.12.0

### Minor Changes

- Moved CPU stats to a separate hwstats package, removing cgo dependency. - [#660](https://github.com/livekit/protocol/pull/660) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Fixed bounds check when masking short SIP numbers. - [#650](https://github.com/livekit/protocol/pull/650) ([@dennwc](https://github.com/dennwc))

- Add signal requests for local track updates - [#651](https://github.com/livekit/protocol/pull/651) ([@lukasIO](https://github.com/lukasIO))

- Allow sending DTMF when creating SIP participant. - [#658](https://github.com/livekit/protocol/pull/658) ([@dennwc](https://github.com/dennwc))

## 1.11.0

### Minor Changes

- Align package version with manually tagged golang module - [#646](https://github.com/livekit/protocol/pull/646) ([@lukasIO](https://github.com/lukasIO))

## 1.10.4

## 1.10.3

## 1.10.2
