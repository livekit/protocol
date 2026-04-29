// Code generated; DO NOT EDIT.

package ingressobs

import (
	"time"
)

var (
	_ Reporter        = (*noopReporter)(nil)
	_ ProjectReporter = (*noopProjectReporter)(nil)
	_ ProjectTx       = (*noopProjectTx)(nil)
	_ IngressReporter = (*noopIngressReporter)(nil)
	_ IngressTx       = (*noopIngressTx)(nil)
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
func (r *noopProjectReporter) WithIngress(id string) IngressReporter {
	return &noopIngressReporter{}
}
func (r *noopProjectReporter) WithDeferredIngress() (IngressReporter, KeyResolver) {
	return &noopIngressReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopIngressReporter struct{}

func NewNoopIngressReporter() IngressReporter {
	return &noopIngressReporter{}
}

func (r *noopIngressReporter) RegisterFunc(f func(ts time.Time, tx IngressTx) bool) {}
func (r *noopIngressReporter) Tx(f func(IngressTx))                                 {}
func (r *noopIngressReporter) TxAt(ts time.Time, f func(IngressTx))                 {}
func (r *noopIngressReporter) WithSession(id string) SessionReporter {
	return &noopSessionReporter{}
}
func (r *noopIngressReporter) WithDeferredSession() (SessionReporter, KeyResolver) {
	return &noopSessionReporter{}, noopKeyResolver{}
}

type noopIngressTx struct{}

func (t *noopIngressTx) Project() ProjectTx {
	return &noopProjectTx{}
}

type noopSessionReporter struct{}

func NewNoopSessionReporter() SessionReporter {
	return &noopSessionReporter{}
}

func (r *noopSessionReporter) RegisterFunc(f func(ts time.Time, tx SessionTx) bool) {}
func (r *noopSessionReporter) Tx(f func(SessionTx))                                 {}
func (r *noopSessionReporter) TxAt(ts time.Time, f func(SessionTx))                 {}
func (r *noopSessionReporter) ReportStartTime(v time.Time)                          {}
func (r *noopSessionReporter) ReportEndTime(v time.Time)                            {}
func (r *noopSessionReporter) ReportDuration(v uint64)                              {}
func (r *noopSessionReporter) ReportInputType(v SessionInputType)                   {}
func (r *noopSessionReporter) ReportRegion(v string)                                {}
func (r *noopSessionReporter) ReportRoomName(v string)                              {}
func (r *noopSessionReporter) ReportRoomID(v string)                                {}
func (r *noopSessionReporter) ReportError(v string)                                 {}
func (r *noopSessionReporter) ReportStatus(v SessionStatus)                         {}
func (r *noopSessionReporter) ReportAudioOnly(v bool)                               {}
func (r *noopSessionReporter) ReportTranscoded(v bool)                              {}
func (r *noopSessionReporter) ReportReusable(v bool)                                {}

type noopSessionTx struct{}

func (t *noopSessionTx) Ingress() IngressTx {
	return &noopIngressTx{}
}

func (t *noopSessionTx) ReportStartTime(v time.Time)        {}
func (t *noopSessionTx) ReportEndTime(v time.Time)          {}
func (t *noopSessionTx) ReportDuration(v uint64)            {}
func (t *noopSessionTx) ReportInputType(v SessionInputType) {}
func (t *noopSessionTx) ReportRegion(v string)              {}
func (t *noopSessionTx) ReportRoomName(v string)            {}
func (t *noopSessionTx) ReportRoomID(v string)              {}
func (t *noopSessionTx) ReportError(v string)               {}
func (t *noopSessionTx) ReportStatus(v SessionStatus)       {}
func (t *noopSessionTx) ReportAudioOnly(v bool)             {}
func (t *noopSessionTx) ReportTranscoded(v bool)            {}
func (t *noopSessionTx) ReportReusable(v bool)              {}
