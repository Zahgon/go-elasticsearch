package esapi

import (
	"context"
	"net/http"
	"time"
)

func newLicenseDeleteFunc(t Transport) LicenseDelete {
	_ = "STUB: not implemented"
	return *new(LicenseDelete)
}

type LicenseDelete func(o ...func(*LicenseDeleteRequest)) (*Response, error)

type LicenseDeleteRequest struct {
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

func (r LicenseDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicenseDelete) WithContext(v context.Context) func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithMasterTimeout(v time.Duration) func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithTimeout(v time.Duration) func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithPretty() func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithHuman() func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithErrorTrace() func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithFilterPath(v ...string) func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithHeader(h map[string]string) func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseDelete) WithOpaqueID(s string) func(*LicenseDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
