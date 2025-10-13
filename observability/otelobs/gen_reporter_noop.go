// Code generated; DO NOT EDIT.

package otelobs

import (
	"time"
)

var (
	_ Reporter        = (*noopReporter)(nil)
	_ ProjectReporter = (*noopProjectReporter)(nil)
	_ SourceReporter  = (*noopSourceReporter)(nil)
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
func (r *noopProjectReporter) WithSource(name string) SourceReporter {
	return &noopSourceReporter{}
}
func (r *noopProjectReporter) WithDeferredSource() (SourceReporter, KeyResolver) {
	return &noopSourceReporter{}, noopKeyResolver{}
}

type noopSourceReporter struct{}

func NewNoopSourceReporter() SourceReporter {
	return &noopSourceReporter{}
}

func (r *noopSourceReporter) RegisterFunc(f func(ts time.Time, tx SourceTx) bool) {}
func (r *noopSourceReporter) Tx(f func(SourceTx))                                 {}
func (r *noopSourceReporter) TxAt(ts time.Time, f func(SourceTx))                 {}
func (r *noopSourceReporter) ReportDataType(v DataType)                           {}
func (r *noopSourceReporter) ReportBytes(v uint64)                                {}
func (r *noopSourceReporter) ReportDuration(v uint64)                             {}
func (r *noopSourceReporter) ReportAllowedForTraining(v bool)                     {}
