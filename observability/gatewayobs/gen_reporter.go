// Code generated; DO NOT EDIT.

package gatewayobs

import (
	"time"
)

const Version_NI039G8 = true

type KeyResolver interface {
	Resolve(string)
	Reset()
}

type Reporter interface {
	WithProject(id string) ProjectReporter
	WithDeferredProject() (ProjectReporter, KeyResolver)
}

type projectReporter interface {
}

type ProjectTx interface {
	projectReporter
}

type ProjectReporter interface {
	RegisterFunc(func(ts time.Time, tx ProjectTx) bool)
	Tx(func(tx ProjectTx))
	TxAt(time.Time, func(tx ProjectTx))
	WithRequestedPriority(priority string) RequestedPriorityReporter
	WithDeferredRequestedPriority() (RequestedPriorityReporter, KeyResolver)
	projectReporter
}

type requestedPriorityReporter interface {
}

type RequestedPriorityTx interface {
	Project() ProjectTx
	requestedPriorityReporter
}

type RequestedPriorityReporter interface {
	RegisterFunc(func(ts time.Time, tx RequestedPriorityTx) bool)
	Tx(func(tx RequestedPriorityTx))
	TxAt(time.Time, func(tx RequestedPriorityTx))
	WithGrantedPriority(priority string) GrantedPriorityReporter
	WithDeferredGrantedPriority() (GrantedPriorityReporter, KeyResolver)
	requestedPriorityReporter
}

type grantedPriorityReporter interface {
}

type GrantedPriorityTx interface {
	RequestedPriority() RequestedPriorityTx
	grantedPriorityReporter
}

type GrantedPriorityReporter interface {
	RegisterFunc(func(ts time.Time, tx GrantedPriorityTx) bool)
	Tx(func(tx GrantedPriorityTx))
	TxAt(time.Time, func(tx GrantedPriorityTx))
	WithBillablePriority(priority string) BillablePriorityReporter
	WithDeferredBillablePriority() (BillablePriorityReporter, KeyResolver)
	grantedPriorityReporter
}

type billablePriorityReporter interface {
}

type BillablePriorityTx interface {
	GrantedPriority() GrantedPriorityTx
	billablePriorityReporter
}

type BillablePriorityReporter interface {
	RegisterFunc(func(ts time.Time, tx BillablePriorityTx) bool)
	Tx(func(tx BillablePriorityTx))
	TxAt(time.Time, func(tx BillablePriorityTx))
	WithProvider(name string) ProviderReporter
	WithDeferredProvider() (ProviderReporter, KeyResolver)
	billablePriorityReporter
}

type providerReporter interface {
}

type ProviderTx interface {
	BillablePriority() BillablePriorityTx
	providerReporter
}

type ProviderReporter interface {
	RegisterFunc(func(ts time.Time, tx ProviderTx) bool)
	Tx(func(tx ProviderTx))
	TxAt(time.Time, func(tx ProviderTx))
	WithModel(name string) ModelReporter
	WithDeferredModel() (ModelReporter, KeyResolver)
	providerReporter
}

type modelReporter interface {
	ReportInferencePromptTokens(v uint64)
	ReportInferencePromptCacheTokens(v uint64)
	ReportInferenceCompletionTokens(v uint64)
	ReportInferenceTotalTokens(v uint64)
	ReportInferenceCacheCreateTokens(v uint64)
	ReportInferenceCacheReadTokens(v uint64)
	ReportSttDuration(v uint32)
	ReportTtsChars(v uint32)
	ReportBargeInRequests(v uint64)
	ReportBargeInRequestTypes(v ModelBargeInRequestTypes)
	ReportVoiceCloneRequests(v uint64)
}

type ModelTx interface {
	Provider() ProviderTx
	modelReporter
}

type ModelReporter interface {
	RegisterFunc(func(ts time.Time, tx ModelTx) bool)
	Tx(func(tx ModelTx))
	TxAt(time.Time, func(tx ModelTx))
	modelReporter
}
