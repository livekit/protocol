// Code generated; DO NOT EDIT.

package telephonyobs

import (
	"time"
)

const Version_6TP95J0 = true

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
	WithCarrier(id string) CarrierReporter
	WithDeferredCarrier() (CarrierReporter, KeyResolver)
	ProjectTx
}

type CarrierTx interface{}

type CarrierReporter interface {
	RegisterFunc(func(ts time.Time, tx CarrierTx) bool)
	Tx(func(tx CarrierTx))
	TxAt(time.Time, func(tx CarrierTx))
	WithCountry(callCountryCode string) CountryReporter
	WithDeferredCountry() (CountryReporter, KeyResolver)
	CarrierTx
}

type CountryTx interface{}

type CountryReporter interface {
	RegisterFunc(func(ts time.Time, tx CountryTx) bool)
	Tx(func(tx CountryTx))
	TxAt(time.Time, func(tx CountryTx))
	WithCall(id string) CallReporter
	WithDeferredCall() (CallReporter, KeyResolver)
	CountryTx
}

type CallTx interface {
	ReportDirection(v DirectionType)
	ReportNumberType(v NumberType)
	ReportStatus(v CallStatus)
	ReportTrunkType(v TrunkType)
	ReportCountryCode(v string)
	ReportPhoneNumber(v string)
	ReportDuration(v uint32)
	ReportDurationMinutes(v uint16)
	ReportStartTime(v time.Time)
	ReportEndTime(v time.Time)
}

type CallReporter interface {
	RegisterFunc(func(ts time.Time, tx CallTx) bool)
	Tx(func(tx CallTx))
	TxAt(time.Time, func(tx CallTx))
	CallTx
}
