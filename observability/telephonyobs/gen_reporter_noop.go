// Code generated; DO NOT EDIT.

package telephonyobs

import (
	"time"
)

var (
	_ Reporter        = (*noopReporter)(nil)
	_ ProjectReporter = (*noopProjectReporter)(nil)
	_ ProjectTx       = (*noopProjectTx)(nil)
	_ CarrierReporter = (*noopCarrierReporter)(nil)
	_ CarrierTx       = (*noopCarrierTx)(nil)
	_ CountryReporter = (*noopCountryReporter)(nil)
	_ CountryTx       = (*noopCountryTx)(nil)
	_ PhoneReporter   = (*noopPhoneReporter)(nil)
	_ PhoneTx         = (*noopPhoneTx)(nil)
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
func (r *noopProjectReporter) WithCarrier(id string) CarrierReporter {
	return &noopCarrierReporter{}
}
func (r *noopProjectReporter) WithDeferredCarrier() (CarrierReporter, KeyResolver) {
	return &noopCarrierReporter{}, noopKeyResolver{}
}

type noopProjectTx struct{}

type noopCarrierReporter struct{}

func NewNoopCarrierReporter() CarrierReporter {
	return &noopCarrierReporter{}
}

func (r *noopCarrierReporter) RegisterFunc(f func(ts time.Time, tx CarrierTx) bool) {}
func (r *noopCarrierReporter) Tx(f func(CarrierTx))                                 {}
func (r *noopCarrierReporter) TxAt(ts time.Time, f func(CarrierTx))                 {}
func (r *noopCarrierReporter) WithCountry(code string) CountryReporter {
	return &noopCountryReporter{}
}
func (r *noopCarrierReporter) WithDeferredCountry() (CountryReporter, KeyResolver) {
	return &noopCountryReporter{}, noopKeyResolver{}
}

type noopCarrierTx struct{}

func (t *noopCarrierTx) Project() ProjectTx {
	return &noopProjectTx{}
}

type noopCountryReporter struct{}

func NewNoopCountryReporter() CountryReporter {
	return &noopCountryReporter{}
}

func (r *noopCountryReporter) RegisterFunc(f func(ts time.Time, tx CountryTx) bool) {}
func (r *noopCountryReporter) Tx(f func(CountryTx))                                 {}
func (r *noopCountryReporter) TxAt(ts time.Time, f func(CountryTx))                 {}
func (r *noopCountryReporter) WithPhone(number string) PhoneReporter {
	return &noopPhoneReporter{}
}
func (r *noopCountryReporter) WithDeferredPhone() (PhoneReporter, KeyResolver) {
	return &noopPhoneReporter{}, noopKeyResolver{}
}

type noopCountryTx struct{}

func (t *noopCountryTx) Carrier() CarrierTx {
	return &noopCarrierTx{}
}

type noopPhoneReporter struct{}

func NewNoopPhoneReporter() PhoneReporter {
	return &noopPhoneReporter{}
}

func (r *noopPhoneReporter) RegisterFunc(f func(ts time.Time, tx PhoneTx) bool) {}
func (r *noopPhoneReporter) Tx(f func(PhoneTx))                                 {}
func (r *noopPhoneReporter) TxAt(ts time.Time, f func(PhoneTx))                 {}
func (r *noopPhoneReporter) WithCall(id string) CallReporter {
	return &noopCallReporter{}
}
func (r *noopPhoneReporter) WithDeferredCall() (CallReporter, KeyResolver) {
	return &noopCallReporter{}, noopKeyResolver{}
}

type noopPhoneTx struct{}

func (t *noopPhoneTx) Country() CountryTx {
	return &noopCountryTx{}
}

type noopCallReporter struct{}

func NewNoopCallReporter() CallReporter {
	return &noopCallReporter{}
}

func (r *noopCallReporter) RegisterFunc(f func(ts time.Time, tx CallTx) bool) {}
func (r *noopCallReporter) Tx(f func(CallTx))                                 {}
func (r *noopCallReporter) TxAt(ts time.Time, f func(CallTx))                 {}
func (r *noopCallReporter) ReportDirection(v DirectionType)                   {}
func (r *noopCallReporter) ReportNumberType(v NumberType)                     {}
func (r *noopCallReporter) ReportStatus(v CallStatus)                         {}
func (r *noopCallReporter) ReportTrunkType(v TrunkType)                       {}
func (r *noopCallReporter) ReportCountryCode(v string)                        {}
func (r *noopCallReporter) ReportPhoneNumber(v string)                        {}
func (r *noopCallReporter) ReportDuration(v uint32)                           {}
func (r *noopCallReporter) ReportDurationMinutes(v uint16)                    {}
func (r *noopCallReporter) ReportStartTime(v time.Time)                       {}
func (r *noopCallReporter) ReportEndTime(v time.Time)                         {}

type noopCallTx struct{}

func (t *noopCallTx) Phone() PhoneTx {
	return &noopPhoneTx{}
}

func (t *noopCallTx) ReportDirection(v DirectionType) {}
func (t *noopCallTx) ReportNumberType(v NumberType)   {}
func (t *noopCallTx) ReportStatus(v CallStatus)       {}
func (t *noopCallTx) ReportTrunkType(v TrunkType)     {}
func (t *noopCallTx) ReportCountryCode(v string)      {}
func (t *noopCallTx) ReportPhoneNumber(v string)      {}
func (t *noopCallTx) ReportDuration(v uint32)         {}
func (t *noopCallTx) ReportDurationMinutes(v uint16)  {}
func (t *noopCallTx) ReportStartTime(v time.Time)     {}
func (t *noopCallTx) ReportEndTime(v time.Time)       {}
