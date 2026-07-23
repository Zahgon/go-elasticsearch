package esapi

import (
	"context"
	"net/http"
)

func newLicenseGetBasicStatusFunc(t Transport) LicenseGetBasicStatus {
	_ = "STUB: not implemented"
	return *new(LicenseGetBasicStatus)
}

type LicenseGetBasicStatus func(o ...func(*LicenseGetBasicStatusRequest)) (*Response, error)

type LicenseGetBasicStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LicenseGetBasicStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicenseGetBasicStatus) WithContext(v context.Context) func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetBasicStatus) WithPretty() func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetBasicStatus) WithHuman() func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetBasicStatus) WithErrorTrace() func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetBasicStatus) WithFilterPath(v ...string) func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetBasicStatus) WithHeader(h map[string]string) func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetBasicStatus) WithOpaqueID(s string) func(*LicenseGetBasicStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
