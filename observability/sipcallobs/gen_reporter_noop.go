// Code generated; DO NOT EDIT.

package sipcallobs

import (
	"time"
)

var (
	_ Reporter        = (*noopReporter)(nil)
	_ ProjectReporter = (*noopProjectReporter)(nil)
	_ ProjectTx       = (*noopProjectTx)(nil)
	_ CallReporter    = (*noopCallReporter)(nil)
	_ CallTx          = (*noopCallTx)(nil)
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
func (r *noopProjectReporter) WithCall(id string) CallReporter {
	return &noopCallReporter{}
}
func (r *noopProjectReporter) WithDeferredCall() (CallReporter, KeyResolver) {
	return &noopCallReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopCallReporter struct{}

func NewNoopCallReporter() CallReporter {
	return &noopCallReporter{}
}

func (r *noopCallReporter) RegisterFunc(f func(ts time.Time, tx CallTx) bool)           {}
func (r *noopCallReporter) Tx(f func(CallTx))                                           {}
func (r *noopCallReporter) TxAt(ts time.Time, f func(CallTx))                           {}
func (r *noopCallReporter) ReportStartTime(v time.Time)                                 {}
func (r *noopCallReporter) ReportEndTime(v time.Time)                                   {}
func (r *noopCallReporter) ReportDuration(v uint64)                                     {}
func (r *noopCallReporter) ReportDurationSeconds(v uint64)                              {}
func (r *noopCallReporter) ReportDurationMinutes(v uint16)                              {}
func (r *noopCallReporter) ReportTrunkID(v string)                                      {}
func (r *noopCallReporter) ReportTrunkType(v CallTrunkType)                             {}
func (r *noopCallReporter) ReportDispatchID(v string)                                   {}
func (r *noopCallReporter) ReportToNumber(v string)                                     {}
func (r *noopCallReporter) ReportToHost(v string)                                       {}
func (r *noopCallReporter) ReportFromNumber(v string)                                   {}
func (r *noopCallReporter) ReportFromHost(v string)                                     {}
func (r *noopCallReporter) ReportNumberType(v CallNumberType)                           {}
func (r *noopCallReporter) ReportCountryCode(v string)                                  {}
func (r *noopCallReporter) ReportDirection(v CallDirection)                             {}
func (r *noopCallReporter) ReportTransport(v CallTransport)                             {}
func (r *noopCallReporter) ReportProviderCallID(v string)                               {}
func (r *noopCallReporter) ReportProviderName(v string)                                 {}
func (r *noopCallReporter) ReportSIPCallID(v string)                                    {}
func (r *noopCallReporter) ReportRoomID(v string)                                       {}
func (r *noopCallReporter) ReportRoomName(v string)                                     {}
func (r *noopCallReporter) ReportParticipantIdentity(v string)                          {}
func (r *noopCallReporter) ReportError(v string)                                        {}
func (r *noopCallReporter) ReportStatus(v CallStatus)                                   {}
func (r *noopCallReporter) ReportResponseCode(v uint16)                                 {}
func (r *noopCallReporter) ReportDisconnectReason(v string)                             {}
func (r *noopCallReporter) ReportTransferID(v string)                                   {}
func (r *noopCallReporter) ReportTransferTo(v string)                                   {}
func (r *noopCallReporter) ReportTransferDuration(v uint32)                             {}
func (r *noopCallReporter) ReportTransferStatus(v CallTransferStatus)                   {}
func (r *noopCallReporter) ReportTransferStatusCode(v uint16)                           {}
func (r *noopCallReporter) ReportTransferError(v string)                                {}
func (r *noopCallReporter) ReportCodec(v string)                                        {}
func (r *noopCallReporter) ReportRegion(v string)                                       {}
func (r *noopCallReporter) ReportPcapLink(v string)                                     {}
func (r *noopCallReporter) ReportAttributes(v string)                                   {}
func (r *noopCallReporter) ReportFeatures(v uint16)                                     {}
func (r *noopCallReporter) ReportMediaEncryptionSettings(v CallMediaEncryptionSettings) {}
func (r *noopCallReporter) ReportMediaEncryption(v string)                              {}

type noopCallTx struct{}

func (t *noopCallTx) Project() ProjectTx {
	return &noopProjectTx{}
}

func (t *noopCallTx) ReportStartTime(v time.Time)                                 {}
func (t *noopCallTx) ReportEndTime(v time.Time)                                   {}
func (t *noopCallTx) ReportDuration(v uint64)                                     {}
func (t *noopCallTx) ReportDurationSeconds(v uint64)                              {}
func (t *noopCallTx) ReportDurationMinutes(v uint16)                              {}
func (t *noopCallTx) ReportTrunkID(v string)                                      {}
func (t *noopCallTx) ReportTrunkType(v CallTrunkType)                             {}
func (t *noopCallTx) ReportDispatchID(v string)                                   {}
func (t *noopCallTx) ReportToNumber(v string)                                     {}
func (t *noopCallTx) ReportToHost(v string)                                       {}
func (t *noopCallTx) ReportFromNumber(v string)                                   {}
func (t *noopCallTx) ReportFromHost(v string)                                     {}
func (t *noopCallTx) ReportNumberType(v CallNumberType)                           {}
func (t *noopCallTx) ReportCountryCode(v string)                                  {}
func (t *noopCallTx) ReportDirection(v CallDirection)                             {}
func (t *noopCallTx) ReportTransport(v CallTransport)                             {}
func (t *noopCallTx) ReportProviderCallID(v string)                               {}
func (t *noopCallTx) ReportProviderName(v string)                                 {}
func (t *noopCallTx) ReportSIPCallID(v string)                                    {}
func (t *noopCallTx) ReportRoomID(v string)                                       {}
func (t *noopCallTx) ReportRoomName(v string)                                     {}
func (t *noopCallTx) ReportParticipantIdentity(v string)                          {}
func (t *noopCallTx) ReportError(v string)                                        {}
func (t *noopCallTx) ReportStatus(v CallStatus)                                   {}
func (t *noopCallTx) ReportResponseCode(v uint16)                                 {}
func (t *noopCallTx) ReportDisconnectReason(v string)                             {}
func (t *noopCallTx) ReportTransferID(v string)                                   {}
func (t *noopCallTx) ReportTransferTo(v string)                                   {}
func (t *noopCallTx) ReportTransferDuration(v uint32)                             {}
func (t *noopCallTx) ReportTransferStatus(v CallTransferStatus)                   {}
func (t *noopCallTx) ReportTransferStatusCode(v uint16)                           {}
func (t *noopCallTx) ReportTransferError(v string)                                {}
func (t *noopCallTx) ReportCodec(v string)                                        {}
func (t *noopCallTx) ReportRegion(v string)                                       {}
func (t *noopCallTx) ReportPcapLink(v string)                                     {}
func (t *noopCallTx) ReportAttributes(v string)                                   {}
func (t *noopCallTx) ReportFeatures(v uint16)                                     {}
func (t *noopCallTx) ReportMediaEncryptionSettings(v CallMediaEncryptionSettings) {}
func (t *noopCallTx) ReportMediaEncryption(v string)                              {}
