// Code generated; DO NOT EDIT.

package otelobs

import (
	"time"
)

const Version_2AT6MB8 = true

type KeyResolver interface {
	Resolve(string)
	Reset()
}

type Reporter interface {
	WithProject(id string) ProjectReporter
	WithDeferredProject() (ProjectReporter, KeyResolver)
}

type ProjectTx interface{}

type ProjectReporter interface {
	RegisterFunc(func(ts time.Time, tx ProjectTx) bool)
	Tx(func(tx ProjectTx))
	TxAt(time.Time, func(tx ProjectTx))
	WithSource(name string) SourceReporter
	WithDeferredSource() (SourceReporter, KeyResolver)
	ProjectTx
}

type SourceTx interface {
	ReportDataType(v DataType)
	ReportBytes(v uint64)
	ReportDuration(v uint64)
	ReportAllowedForTraining(v bool)
}

type SourceReporter interface {
	RegisterFunc(func(ts time.Time, tx SourceTx) bool)
	Tx(func(tx SourceTx))
	TxAt(time.Time, func(tx SourceTx))
	SourceTx
}
