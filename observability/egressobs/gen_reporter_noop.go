// Code generated; DO NOT EDIT.

package egressobs

import (
	"time"
)

var (
	_ Reporter        = (*noopReporter)(nil)
	_ ProjectReporter = (*noopProjectReporter)(nil)
	_ ProjectTx       = (*noopProjectTx)(nil)
	_ EgressReporter  = (*noopEgressReporter)(nil)
	_ EgressTx        = (*noopEgressTx)(nil)
	_ SessionReporter = (*noopSessionReporter)(nil)
	_ SessionTx       = (*noopSessionTx)(nil)
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
func (r *noopProjectReporter) WithEgress(id string) EgressReporter {
	return &noopEgressReporter{}
}
func (r *noopProjectReporter) WithDeferredEgress() (EgressReporter, KeyResolver) {
	return &noopEgressReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopEgressReporter struct{}

func NewNoopEgressReporter() EgressReporter {
	return &noopEgressReporter{}
}

func (r *noopEgressReporter) RegisterFunc(f func(ts time.Time, tx EgressTx) bool) {}
func (r *noopEgressReporter) Tx(f func(EgressTx))                                 {}
func (r *noopEgressReporter) TxAt(ts time.Time, f func(EgressTx))                 {}
func (r *noopEgressReporter) ReportRequestType(v EgressRequestType)               {}
func (r *noopEgressReporter) ReportRoomName(v string)                             {}
func (r *noopEgressReporter) ReportRequest(v string)                              {}
func (r *noopEgressReporter) ReportAudioOnly(v bool)                              {}
func (r *noopEgressReporter) ReportStartTime(v time.Time)                         {}
func (r *noopEgressReporter) ReportEndTime(v time.Time)                           {}
func (r *noopEgressReporter) ReportUpdateTime(v time.Time)                        {}
func (r *noopEgressReporter) ReportStatus(v string)                               {}
func (r *noopEgressReporter) ReportDetails(v string)                              {}
func (r *noopEgressReporter) ReportError(v string)                                {}
func (r *noopEgressReporter) ReportErrorCode(v int32)                             {}
func (r *noopEgressReporter) ReportResult(v string)                               {}
func (r *noopEgressReporter) ReportManifestLocation(v string)                     {}
func (r *noopEgressReporter) WithSession(id string) SessionReporter {
	return &noopSessionReporter{}
}
func (r *noopEgressReporter) WithDeferredSession() (SessionReporter, KeyResolver) {
	return &noopSessionReporter{}, noopKeyResolver{}
}

type noopEgressTx struct{}

func (t *noopEgressTx) Project() ProjectTx {
	return &noopProjectTx{}
}

func (t *noopEgressTx) ReportRequestType(v EgressRequestType) {}
func (t *noopEgressTx) ReportRoomName(v string)               {}
func (t *noopEgressTx) ReportRequest(v string)                {}
func (t *noopEgressTx) ReportAudioOnly(v bool)                {}
func (t *noopEgressTx) ReportStartTime(v time.Time)           {}
func (t *noopEgressTx) ReportEndTime(v time.Time)             {}
func (t *noopEgressTx) ReportUpdateTime(v time.Time)          {}
func (t *noopEgressTx) ReportStatus(v string)                 {}
func (t *noopEgressTx) ReportDetails(v string)                {}
func (t *noopEgressTx) ReportError(v string)                  {}
func (t *noopEgressTx) ReportErrorCode(v int32)               {}
func (t *noopEgressTx) ReportResult(v string)                 {}
func (t *noopEgressTx) ReportManifestLocation(v string)       {}

type noopSessionReporter struct{}

func NewNoopSessionReporter() SessionReporter {
	return &noopSessionReporter{}
}

func (r *noopSessionReporter) RegisterFunc(f func(ts time.Time, tx SessionTx) bool) {}
func (r *noopSessionReporter) Tx(f func(SessionTx))                                 {}
func (r *noopSessionReporter) TxAt(ts time.Time, f func(SessionTx))                 {}
func (r *noopSessionReporter) ReportStartTime(v time.Time)                          {}
func (r *noopSessionReporter) ReportEndTime(v time.Time)                            {}
func (r *noopSessionReporter) ReportUpdateTime(v time.Time)                         {}
func (r *noopSessionReporter) ReportDuration(v uint64)                              {}
func (r *noopSessionReporter) ReportRetryCount(v uint32)                            {}
func (r *noopSessionReporter) ReportSourceType(v SessionSourceType)                 {}
func (r *noopSessionReporter) ReportRegion(v string)                                {}
func (r *noopSessionReporter) ReportRoomID(v string)                                {}
func (r *noopSessionReporter) ReportStatus(v string)                                {}
func (r *noopSessionReporter) ReportDetails(v string)                               {}
func (r *noopSessionReporter) ReportError(v string)                                 {}
func (r *noopSessionReporter) ReportErrorCode(v int32)                              {}
func (r *noopSessionReporter) ReportManifestLocation(v string)                      {}
func (r *noopSessionReporter) ReportBackupStorageUsed(v bool)                       {}
func (r *noopSessionReporter) ReportResult(v string)                                {}

type noopSessionTx struct{}

func (t *noopSessionTx) Egress() EgressTx {
	return &noopEgressTx{}
}

func (t *noopSessionTx) ReportStartTime(v time.Time)          {}
func (t *noopSessionTx) ReportEndTime(v time.Time)            {}
func (t *noopSessionTx) ReportUpdateTime(v time.Time)         {}
func (t *noopSessionTx) ReportDuration(v uint64)              {}
func (t *noopSessionTx) ReportRetryCount(v uint32)            {}
func (t *noopSessionTx) ReportSourceType(v SessionSourceType) {}
func (t *noopSessionTx) ReportRegion(v string)                {}
func (t *noopSessionTx) ReportRoomID(v string)                {}
func (t *noopSessionTx) ReportStatus(v string)                {}
func (t *noopSessionTx) ReportDetails(v string)               {}
func (t *noopSessionTx) ReportError(v string)                 {}
func (t *noopSessionTx) ReportErrorCode(v int32)              {}
func (t *noopSessionTx) ReportManifestLocation(v string)      {}
func (t *noopSessionTx) ReportBackupStorageUsed(v bool)       {}
func (t *noopSessionTx) ReportResult(v string)                {}
