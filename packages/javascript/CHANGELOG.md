# @livekit/protocol

## 1.32.0

### Minor Changes

- Use iterators in SIP. - [#943](https://github.com/livekit/protocol/pull/943) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Rename File* related DataStream messages to Byte* - [#948](https://github.com/livekit/protocol/pull/948) ([@lukasIO](https://github.com/lukasIO))

## 1.31.0

### Minor Changes

- Add media encryption options for SIP. - [#892](https://github.com/livekit/protocol/pull/892) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Add "Enabled" flag to ingresses to allow preventing an ingress to become active without deleting it. - [#937](https://github.com/livekit/protocol/pull/937) ([@biglittlebigben](https://github.com/biglittlebigben))

- Add DataStream.Trailer for finalizing streams - [#940](https://github.com/livekit/protocol/pull/940) ([@lukasIO](https://github.com/lukasIO))

- Add SIPCallDirection to SIPCallInfo - [#938](https://github.com/livekit/protocol/pull/938) ([@biglittlebigben](https://github.com/biglittlebigben))

## 1.30.0

### Minor Changes

- Add filters to SIP list APIs. - [#913](https://github.com/livekit/protocol/pull/913) ([@dennwc](https://github.com/dennwc))

- Add headers for SIP call transfer. - [#926](https://github.com/livekit/protocol/pull/926) ([@dennwc](https://github.com/dennwc))

- Add headers for CreateSipParticipant - [#911](https://github.com/livekit/protocol/pull/911) ([@s-hamdananwar](https://github.com/s-hamdananwar))

- Allow mapping all SIP headers to attributes. - [#920](https://github.com/livekit/protocol/pull/920) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Allow SIP inbound to specify room configuration - [#923](https://github.com/livekit/protocol/pull/923) ([@davidzhao](https://github.com/davidzhao))

- Add EgressSourceType to EgressInfo - [#894](https://github.com/livekit/protocol/pull/894) ([@biglittlebigben](https://github.com/biglittlebigben))

## 1.29.4

### Patch Changes

- Add version declaration file - [#916](https://github.com/livekit/protocol/pull/916) ([@lukasIO](https://github.com/lukasIO))

## 1.29.3

### Patch Changes

- Fix treehaking possibility of individual imports - [#907](https://github.com/livekit/protocol/pull/907) ([@lukasIO](https://github.com/lukasIO))

## 1.29.2

## 1.29.1

### Patch Changes

- Actually downgrade TypeScript - [#902](https://github.com/livekit/protocol/pull/902) ([@lukasIO](https://github.com/lukasIO))

## 1.29.0

### Minor Changes

- Add DataStream support - [#871](https://github.com/livekit/protocol/pull/871) ([@lukasIO](https://github.com/lukasIO))

### Patch Changes

- Downgrade TypeScript to fix incorrect generated typings - [#901](https://github.com/livekit/protocol/pull/901) ([@lukasIO](https://github.com/lukasIO))

## 1.28.1

### Patch Changes

- fix CJS require importing .js - [#897](https://github.com/livekit/protocol/pull/897) ([@nbsp](https://github.com/nbsp))

## 1.28.0

### Minor Changes

- Add disconnect reason for SIP failures. - [#845](https://github.com/livekit/protocol/pull/845) ([@dennwc](https://github.com/dennwc))

- add native CommonJS support (#883) - [#895](https://github.com/livekit/protocol/pull/895) ([@nbsp](https://github.com/nbsp))

- Map participant attributes to SIP headers. - [#893](https://github.com/livekit/protocol/pull/893) ([@dennwc](https://github.com/dennwc))

- Allow setting number for SIP outbound. - [#891](https://github.com/livekit/protocol/pull/891) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Report call enabled features in SIPCallInfo - [#884](https://github.com/livekit/protocol/pull/884) ([@biglittlebigben](https://github.com/biglittlebigben))

- Fix port type for SIPUri. - [#882](https://github.com/livekit/protocol/pull/882) ([@dennwc](https://github.com/dennwc))

## 1.27.1

### Patch Changes

- Include room agent dispatch protobufs in JS export - [#875](https://github.com/livekit/protocol/pull/875) ([@davidzhao](https://github.com/davidzhao))

## 1.27.0

### Minor Changes

- Support for room configuration and agent dispatches - [#864](https://github.com/livekit/protocol/pull/864) ([@davidzhao](https://github.com/davidzhao))

### Patch Changes

- Allow requesting a ringtone during SIP call transfers - [#865](https://github.com/livekit/protocol/pull/865) ([@biglittlebigben](https://github.com/biglittlebigben))

## 1.26.0

## 1.25.0

### Minor Changes

- Add SIP analytics events - [#831](https://github.com/livekit/protocol/pull/831) ([@biglittlebigben](https://github.com/biglittlebigben))

- Add ringing timeout and max call duration setting for SIP. - [#844](https://github.com/livekit/protocol/pull/844) ([@dennwc](https://github.com/dennwc))

## 1.24.0

### Minor Changes

- Add disconnect reasons for SIP. - [#828](https://github.com/livekit/protocol/pull/828) ([@dennwc](https://github.com/dennwc))

### Patch Changes

- Expose metrics on javascript package - [#843](https://github.com/livekit/protocol/pull/843) ([@lukasIO](https://github.com/lukasIO))

## 1.23.0

### Minor Changes

- Add RPC types - [#814](https://github.com/livekit/protocol/pull/814) ([@bcherry](https://github.com/bcherry))

- Add TransferParticipant to sip and sip internal services - [#816](https://github.com/livekit/protocol/pull/816) ([@biglittlebigben](https://github.com/biglittlebigben))

### Patch Changes

- Add protocols for client metrics - [#821](https://github.com/livekit/protocol/pull/821) ([@davidliu](https://github.com/davidliu))

- Implement Validate on TransferSIPParticipantRequest - [#818](https://github.com/livekit/protocol/pull/818) ([@biglittlebigben](https://github.com/biglittlebigben))

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

- Add EnabledCodecs to JoinResponse - [#791](https://github.com/livekit/protocol/pull/791) ([@cnderrauber](https://github.com/cnderrauber))

- Add fast_publish option to JoinResponse - [#796](https://github.com/livekit/protocol/pull/796) ([@cnderrauber](https://github.com/cnderrauber))

## 1.20.0

### Minor Changes

- Send non-error responses also as part of RequestResponse - [#787](https://github.com/livekit/protocol/pull/787) ([@lukasIO](https://github.com/lukasIO))

### Patch Changes

- Add disconnect reason in participant info. - [#788](https://github.com/livekit/protocol/pull/788) ([@boks1971](https://github.com/boks1971))

- Add AgentDispatchInternal service. Add JobTerminate to AgentService - [#782](https://github.com/livekit/protocol/pull/782) ([@biglittlebigben](https://github.com/biglittlebigben))

## 1.19.3

### Patch Changes

- Use glob for generating JS proto definitions - [#779](https://github.com/livekit/protocol/pull/779) ([@lukasIO](https://github.com/lukasIO))

## 1.19.2

### Patch Changes

- Ability to set attributes upon job acceptance - [#769](https://github.com/livekit/protocol/pull/769) ([@davidzhao](https://github.com/davidzhao))

- Add track subscribed first time notification - [#752](https://github.com/livekit/protocol/pull/752) ([@cnderrauber](https://github.com/cnderrauber))

- Rename ErrorResponse Reason INVALID_ARGUMENT to LIMIT_EXCEEDED - [#758](https://github.com/livekit/protocol/pull/758) ([@lukasIO](https://github.com/lukasIO))

- Added RoomClosed as a DisconnectReason - [#778](https://github.com/livekit/protocol/pull/778) ([@davidzhao](https://github.com/davidzhao))

## 1.19.1

### Patch Changes

- Add ErrorResponse - [#743](https://github.com/livekit/protocol/pull/743) ([@lukasIO](https://github.com/lukasIO))

## 1.19.0

### Minor Changes

- Add SIP grants. - [#745](https://github.com/livekit/protocol/pull/745) ([@dennwc](https://github.com/dennwc))

## 1.18.0

### Minor Changes

- Split SIP Trunk configuration to inbound and outbound parts. - [#738](https://github.com/livekit/protocol/pull/738) ([@dennwc](https://github.com/dennwc))

- add Participant attributes key/val storage - [#733](https://github.com/livekit/protocol/pull/733) ([@davidzhao](https://github.com/davidzhao))

## 1.17.0

### Minor Changes

- Add support for additional SIP transports. - [#719](https://github.com/livekit/protocol/pull/719) ([@dennwc](https://github.com/dennwc))

- Add ICERestartWHIPResource to RPC - [#716](https://github.com/livekit/protocol/pull/716) ([@Sean-Der](https://github.com/Sean-Der))

### Patch Changes

- Added error code field to EgressInfo - [#714](https://github.com/livekit/protocol/pull/714) ([@frostbyte73](https://github.com/frostbyte73))

- Include node_id with analytics events - [#727](https://github.com/livekit/protocol/pull/727) ([@davidzhao](https://github.com/davidzhao))

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

## 1.12.0

### Patch Changes

- Fixed bounds check when masking short SIP numbers. - [#650](https://github.com/livekit/protocol/pull/650) ([@dennwc](https://github.com/dennwc))

- Add signal requests for local track updates - [#651](https://github.com/livekit/protocol/pull/651) ([@lukasIO](https://github.com/lukasIO))

- Allow sending DTMF when creating SIP participant. - [#658](https://github.com/livekit/protocol/pull/658) ([@dennwc](https://github.com/dennwc))

## 1.11.0

### Minor Changes

- Align package version with manually tagged golang module - [#646](https://github.com/livekit/protocol/pull/646) ([@lukasIO](https://github.com/lukasIO))

## 1.10.4

### Patch Changes

- Fix granular export paths - [#642](https://github.com/livekit/protocol/pull/642) ([@lukasIO](https://github.com/lukasIO))

## 1.10.3

### Patch Changes

- Export room stubs from package - [#640](https://github.com/livekit/protocol/pull/640) ([@lukasIO](https://github.com/lukasIO))

## 1.10.2

### Patch Changes

- Fix import paths - [#635](https://github.com/livekit/protocol/pull/635) ([@lukasIO](https://github.com/lukasIO))
