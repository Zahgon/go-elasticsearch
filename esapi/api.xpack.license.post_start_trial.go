package esapi

import (
	"context"
	"net/http"
	"time"
)

func newLicensePostStartTrialFunc(t Transport) LicensePostStartTrial {
	_ = "STUB: not implemented"
	return *new(LicensePostStartTrial)
}

type LicensePostStartTrial func(o ...func(*LicensePostStartTrialRequest)) (*Response, error)

type LicensePostStartTrialRequest struct {
	Acknowledge   *bool
	MasterTimeout time.Duration
	DocumentType  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LicensePostStartTrialRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicensePostStartTrial) WithContext(v context.Context) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithAcknowledge(v bool) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithMasterTimeout(v time.Duration) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithDocumentType(v string) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithPretty() func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithHuman() func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithErrorTrace() func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithFilterPath(v ...string) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithHeader(h map[string]string) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartTrial) WithOpaqueID(s string) func(*LicensePostStartTrialRequest) {
	_ = "STUB: not implemented"
	return nil
}
