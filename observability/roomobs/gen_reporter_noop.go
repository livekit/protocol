// Code generated; DO NOT EDIT.

package roomobs

import (
	"time"
)

var (
	_ Reporter                   = (*noopReporter)(nil)
	_ ProjectReporter            = (*noopProjectReporter)(nil)
	_ ProjectTx                  = (*noopProjectTx)(nil)
	_ RoomReporter               = (*noopRoomReporter)(nil)
	_ RoomTx                     = (*noopRoomTx)(nil)
	_ RoomSessionReporter        = (*noopRoomSessionReporter)(nil)
	_ RoomSessionTx              = (*noopRoomSessionTx)(nil)
	_ ParticipantReporter        = (*noopParticipantReporter)(nil)
	_ ParticipantTx              = (*noopParticipantTx)(nil)
	_ ParticipantSessionReporter = (*noopParticipantSessionReporter)(nil)
	_ ParticipantSessionTx       = (*noopParticipantSessionTx)(nil)
	_ TrackReporter              = (*noopTrackReporter)(nil)
	_ TrackTx                    = (*noopTrackTx)(nil)
)

type noopKeyResolver struct{}

func (noopKeyResolver) Resolve(string) {}
func (noopKeyResolver) Reset()         {}

type noopReporter struct{}

func NewNoopReporter() Reporter {
	return &noopReporter{}
}

func (r *noopReporter) WithProject(id string) ProjectReporter {
	return &noopProjectReporter{}
}

func (r *noopReporter) WithDeferredProject() (ProjectReporter, KeyResolver) {
	return &noopProjectReporter{}, noopKeyResolver{}
}

type noopProjectReporter struct{}

func NewNoopProjectReporter() ProjectReporter {
	return &noopProjectReporter{}
}

func (r *noopProjectReporter) RegisterFunc(f func(ts time.Time, tx ProjectTx) bool) {}
func (r *noopProjectReporter) Tx(f func(ProjectTx))                                 {}
func (r *noopProjectReporter) TxAt(ts time.Time, f func(ProjectTx))                 {}
func (r *noopProjectReporter) WithRoom(name string) RoomReporter {
	return &noopRoomReporter{}
}
func (r *noopProjectReporter) WithDeferredRoom() (RoomReporter, KeyResolver) {
	return &noopRoomReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopRoomReporter struct{}

func NewNoopRoomReporter() RoomReporter {
	return &noopRoomReporter{}
}

func (r *noopRoomReporter) RegisterFunc(f func(ts time.Time, tx RoomTx) bool) {}
func (r *noopRoomReporter) Tx(f func(RoomTx))                                 {}
func (r *noopRoomReporter) TxAt(ts time.Time, f func(RoomTx))                 {}
func (r *noopRoomReporter) WithRoomSession(id string) RoomSessionReporter {
	return &noopRoomSessionReporter{}
}
func (r *noopRoomReporter) WithDeferredRoomSession() (RoomSessionReporter, KeyResolver) {
	return &noopRoomSessionReporter{}, noopKeyResolver{}
}

type noopRoomTx struct{}

func (t *noopRoomTx) Project() ProjectTx {
	return &noopProjectTx{}
}

type noopRoomSessionReporter struct{}

func NewNoopRoomSessionReporter() RoomSessionReporter {
	return &noopRoomSessionReporter{}
}

func (r *noopRoomSessionReporter) RegisterFunc(f func(ts time.Time, tx RoomSessionTx) bool) {}
func (r *noopRoomSessionReporter) Tx(f func(RoomSessionTx))                                 {}
func (r *noopRoomSessionReporter) TxAt(ts time.Time, f func(RoomSessionTx))                 {}
func (r *noopRoomSessionReporter) ReportStartTime(v time.Time)                              {}
func (r *noopRoomSessionReporter) ReportEndTime(v time.Time)                                {}
func (r *noopRoomSessionReporter) ReportFeatures(v uint16)                                  {}
func (r *noopRoomSessionReporter) ReportRoomDuration(v uint32)                              {}
func (r *noopRoomSessionReporter) ReportTags(v []string)                                    {}
func (r *noopRoomSessionReporter) ReportClosed(v bool)                                      {}
func (r *noopRoomSessionReporter) WithParticipant(identity string) ParticipantReporter {
	return &noopParticipantReporter{}
}
func (r *noopRoomSessionReporter) WithDeferredParticipant() (ParticipantReporter, KeyResolver) {
	return &noopParticipantReporter{}, noopKeyResolver{}
}

type noopRoomSessionTx struct{}

func (t *noopRoomSessionTx) Room() RoomTx {
	return &noopRoomTx{}
}

func (t *noopRoomSessionTx) ReportStartTime(v time.Time) {}
func (t *noopRoomSessionTx) ReportEndTime(v time.Time)   {}
func (t *noopRoomSessionTx) ReportFeatures(v uint16)     {}
func (t *noopRoomSessionTx) ReportRoomDuration(v uint32) {}
func (t *noopRoomSessionTx) ReportTags(v []string)       {}
func (t *noopRoomSessionTx) ReportClosed(v bool)         {}

type noopParticipantReporter struct{}

func NewNoopParticipantReporter() ParticipantReporter {
	return &noopParticipantReporter{}
}

func (r *noopParticipantReporter) RegisterFunc(f func(ts time.Time, tx ParticipantTx) bool) {}
func (r *noopParticipantReporter) Tx(f func(ParticipantTx))                                 {}
func (r *noopParticipantReporter) TxAt(ts time.Time, f func(ParticipantTx))                 {}
func (r *noopParticipantReporter) WithParticipantSession(id string) ParticipantSessionReporter {
	return &noopParticipantSessionReporter{}
}
func (r *noopParticipantReporter) WithDeferredParticipantSession() (ParticipantSessionReporter, KeyResolver) {
	return &noopParticipantSessionReporter{}, noopKeyResolver{}
}

type noopParticipantTx struct{}

func (t *noopParticipantTx) RoomSession() RoomSessionTx {
	return &noopRoomSessionTx{}
}

type noopParticipantSessionReporter struct{}

func NewNoopParticipantSessionReporter() ParticipantSessionReporter {
	return &noopParticipantSessionReporter{}
}

func (r *noopParticipantSessionReporter) RegisterFunc(f func(ts time.Time, tx ParticipantSessionTx) bool) {
}
func (r *noopParticipantSessionReporter) Tx(f func(ParticipantSessionTx))                 {}
func (r *noopParticipantSessionReporter) TxAt(ts time.Time, f func(ParticipantSessionTx)) {}
func (r *noopParticipantSessionReporter) ReportRegion(v string)                           {}
func (r *noopParticipantSessionReporter) ReportClientConnectTime(v uint16)                {}
func (r *noopParticipantSessionReporter) ReportConnectResult(v ConnectionResult)          {}
func (r *noopParticipantSessionReporter) ReportConnectionType(v ConnectionType)           {}
func (r *noopParticipantSessionReporter) ReportOs(v ClientOS)                             {}
func (r *noopParticipantSessionReporter) ReportDeviceModel(v string)                      {}
func (r *noopParticipantSessionReporter) ReportBrowser(v string)                          {}
func (r *noopParticipantSessionReporter) ReportSdkVersion(v string)                       {}
func (r *noopParticipantSessionReporter) ReportCountry(v uint16)                          {}
func (r *noopParticipantSessionReporter) ReportIspAsn(v uint32)                           {}
func (r *noopParticipantSessionReporter) ReportStartTime(v time.Time)                     {}
func (r *noopParticipantSessionReporter) ReportEndTime(v time.Time)                       {}
func (r *noopParticipantSessionReporter) ReportDuration(v uint16)                         {}
func (r *noopParticipantSessionReporter) ReportDurationSeconds(v uint16)                  {}
func (r *noopParticipantSessionReporter) ReportDurationMinutes(v uint8)                   {}
func (r *noopParticipantSessionReporter) ReportKind(v string)                             {}
func (r *noopParticipantSessionReporter) ReportName(v string)                             {}
func (r *noopParticipantSessionReporter) ReportFeatures(v uint16)                         {}
func (r *noopParticipantSessionReporter) WithTrack(id string) TrackReporter {
	return &noopTrackReporter{}
}
func (r *noopParticipantSessionReporter) WithDeferredTrack() (TrackReporter, KeyResolver) {
	return &noopTrackReporter{}, noopKeyResolver{}
}

type noopParticipantSessionTx struct{}

func (t *noopParticipantSessionTx) Participant() ParticipantTx {
	return &noopParticipantTx{}
}

func (t *noopParticipantSessionTx) ReportRegion(v string)                  {}
func (t *noopParticipantSessionTx) ReportClientConnectTime(v uint16)       {}
func (t *noopParticipantSessionTx) ReportConnectResult(v ConnectionResult) {}
func (t *noopParticipantSessionTx) ReportConnectionType(v ConnectionType)  {}
func (t *noopParticipantSessionTx) ReportOs(v ClientOS)                    {}
func (t *noopParticipantSessionTx) ReportDeviceModel(v string)             {}
func (t *noopParticipantSessionTx) ReportBrowser(v string)                 {}
func (t *noopParticipantSessionTx) ReportSdkVersion(v string)              {}
func (t *noopParticipantSessionTx) ReportCountry(v uint16)                 {}
func (t *noopParticipantSessionTx) ReportIspAsn(v uint32)                  {}
func (t *noopParticipantSessionTx) ReportStartTime(v time.Time)            {}
func (t *noopParticipantSessionTx) ReportEndTime(v time.Time)              {}
func (t *noopParticipantSessionTx) ReportDuration(v uint16)                {}
func (t *noopParticipantSessionTx) ReportDurationSeconds(v uint16)         {}
func (t *noopParticipantSessionTx) ReportDurationMinutes(v uint8)          {}
func (t *noopParticipantSessionTx) ReportKind(v string)                    {}
func (t *noopParticipantSessionTx) ReportName(v string)                    {}
func (t *noopParticipantSessionTx) ReportFeatures(v uint16)                {}

type noopTrackReporter struct{}

func NewNoopTrackReporter() TrackReporter {
	return &noopTrackReporter{}
}

func (r *noopTrackReporter) RegisterFunc(f func(ts time.Time, tx TrackTx) bool) {}
func (r *noopTrackReporter) Tx(f func(TrackTx))                                 {}
func (r *noopTrackReporter) TxAt(ts time.Time, f func(TrackTx))                 {}
func (r *noopTrackReporter) ReportName(v string)                                {}
func (r *noopTrackReporter) ReportKind(v TrackKind)                             {}
func (r *noopTrackReporter) ReportType(v TrackType)                             {}
func (r *noopTrackReporter) ReportSource(v TrackSource)                         {}
func (r *noopTrackReporter) ReportMime(v MimeType)                              {}
func (r *noopTrackReporter) ReportLayer(v uint32)                               {}
func (r *noopTrackReporter) ReportDuration(v uint16)                            {}
func (r *noopTrackReporter) ReportFrames(v uint16)                              {}
func (r *noopTrackReporter) ReportSendBytes(v uint32)                           {}
func (r *noopTrackReporter) ReportRecvBytes(v uint32)                           {}
func (r *noopTrackReporter) ReportSendPackets(v uint32)                         {}
func (r *noopTrackReporter) ReportRecvPackets(v uint32)                         {}
func (r *noopTrackReporter) ReportPacketsLost(v uint32)                         {}
func (r *noopTrackReporter) ReportScore(v float32)                              {}

type noopTrackTx struct{}

func (t *noopTrackTx) ParticipantSession() ParticipantSessionTx {
	return &noopParticipantSessionTx{}
}

func (t *noopTrackTx) ReportName(v string)        {}
func (t *noopTrackTx) ReportKind(v TrackKind)     {}
func (t *noopTrackTx) ReportType(v TrackType)     {}
func (t *noopTrackTx) ReportSource(v TrackSource) {}
func (t *noopTrackTx) ReportMime(v MimeType)      {}
func (t *noopTrackTx) ReportLayer(v uint32)       {}
func (t *noopTrackTx) ReportDuration(v uint16)    {}
func (t *noopTrackTx) ReportFrames(v uint16)      {}
func (t *noopTrackTx) ReportSendBytes(v uint32)   {}
func (t *noopTrackTx) ReportRecvBytes(v uint32)   {}
func (t *noopTrackTx) ReportSendPackets(v uint32) {}
func (t *noopTrackTx) ReportRecvPackets(v uint32) {}
func (t *noopTrackTx) ReportPacketsLost(v uint32) {}
func (t *noopTrackTx) ReportScore(v float32)      {}
