package esapi

import (
	"context"
	"net/http"
	"time"
)

func newLicensePostStartBasicFunc(t Transport) LicensePostStartBasic {
	_ = "STUB: not implemented"
	return *new(LicensePostStartBasic)
}

type LicensePostStartBasic func(o ...func(*LicensePostStartBasicRequest)) (*Response, error)

type LicensePostStartBasicRequest struct {
	Acknowledge   *bool
	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LicensePostStartBasicRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicensePostStartBasic) WithContext(v context.Context) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithAcknowledge(v bool) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithMasterTimeout(v time.Duration) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithTimeout(v time.Duration) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithPretty() func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithHuman() func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithErrorTrace() func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithFilterPath(v ...string) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithHeader(h map[string]string) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePostStartBasic) WithOpaqueID(s string) func(*LicensePostStartBasicRequest) {
	_ = "STUB: not implemented"
	return nil
}
