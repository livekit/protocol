// Code generated; DO NOT EDIT.

package storageobs

import (
	"time"
)

var (
	_ Reporter        = (*noopReporter)(nil)
	_ ProjectReporter = (*noopProjectReporter)(nil)
	_ ProjectTx       = (*noopProjectTx)(nil)
	_ EventReporter   = (*noopEventReporter)(nil)
	_ EventTx         = (*noopEventTx)(nil)
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
func (r *noopProjectReporter) WithEvent(id string) EventReporter {
	return &noopEventReporter{}
}
func (r *noopProjectReporter) WithDeferredEvent() (EventReporter, KeyResolver) {
	return &noopEventReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopEventReporter struct{}

func NewNoopEventReporter() EventReporter {
	return &noopEventReporter{}
}

func (r *noopEventReporter) RegisterFunc(f func(ts time.Time, tx EventTx) bool) {}
func (r *noopEventReporter) Tx(f func(EventTx))                                 {}
func (r *noopEventReporter) TxAt(ts time.Time, f func(EventTx))                 {}
func (r *noopEventReporter) ReportService(v EventService)                       {}
func (r *noopEventReporter) ReportServiceID(v string)                           {}
func (r *noopEventReporter) ReportOperation(v EventOperation)                   {}
func (r *noopEventReporter) ReportPath(v string)                                {}
func (r *noopEventReporter) ReportSize(v uint64)                                {}
func (r *noopEventReporter) ReportLifetime(v uint64)                            {}

type noopEventTx struct{}

func (t *noopEventTx) Project() ProjectTx {
	return &noopProjectTx{}
}

func (t *noopEventTx) ReportService(v EventService)     {}
func (t *noopEventTx) ReportServiceID(v string)         {}
func (t *noopEventTx) ReportOperation(v EventOperation) {}
func (t *noopEventTx) ReportPath(v string)              {}
func (t *noopEventTx) ReportSize(v uint64)              {}
func (t *noopEventTx) ReportLifetime(v uint64)          {}
