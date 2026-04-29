// Code generated; DO NOT EDIT.

package corecallobs

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

func (r *noopCallReporter) RegisterFunc(f func(ts time.Time, tx CallTx) bool) {}
func (r *noopCallReporter) Tx(f func(CallTx))                                 {}
func (r *noopCallReporter) TxAt(ts time.Time, f func(CallTx))                 {}
func (r *noopCallReporter) ReportStartTime(v time.Time)                       {}
func (r *noopCallReporter) ReportEndTime(v time.Time)                         {}
func (r *noopCallReporter) ReportDuration(v uint64)                           {}
func (r *noopCallReporter) ReportDurationMinutes(v uint16)                    {}
func (r *noopCallReporter) ReportDirection(v CallDirection)                   {}
func (r *noopCallReporter) ReportCallType(v CallCallType)                     {}
func (r *noopCallReporter) ReportFrom(v string)                               {}
func (r *noopCallReporter) ReportTo(v string)                                 {}
func (r *noopCallReporter) ReportRegion(v string)                             {}
func (r *noopCallReporter) ReportRoomID(v string)                             {}
func (r *noopCallReporter) ReportRoomName(v string)                           {}
func (r *noopCallReporter) ReportError(v string)                              {}
func (r *noopCallReporter) ReportStatus(v CallStatus)                         {}

type noopCallTx struct{}

func (t *noopCallTx) Project() ProjectTx {
	return &noopProjectTx{}
}

func (t *noopCallTx) ReportStartTime(v time.Time)     {}
func (t *noopCallTx) ReportEndTime(v time.Time)       {}
func (t *noopCallTx) ReportDuration(v uint64)         {}
func (t *noopCallTx) ReportDurationMinutes(v uint16)  {}
func (t *noopCallTx) ReportDirection(v CallDirection) {}
func (t *noopCallTx) ReportCallType(v CallCallType)   {}
func (t *noopCallTx) ReportFrom(v string)             {}
func (t *noopCallTx) ReportTo(v string)               {}
func (t *noopCallTx) ReportRegion(v string)           {}
func (t *noopCallTx) ReportRoomID(v string)           {}
func (t *noopCallTx) ReportRoomName(v string)         {}
func (t *noopCallTx) ReportError(v string)            {}
func (t *noopCallTx) ReportStatus(v CallStatus)       {}
