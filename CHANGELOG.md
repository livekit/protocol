# github.com/livekit/protocol

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
