package esapi

import (
	"context"
	"net/http"
)

func newLicenseGetTrialStatusFunc(t Transport) LicenseGetTrialStatus {
	_ = "STUB: not implemented"
	return *new(LicenseGetTrialStatus)
}

type LicenseGetTrialStatus func(o ...func(*LicenseGetTrialStatusRequest)) (*Response, error)

type LicenseGetTrialStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LicenseGetTrialStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicenseGetTrialStatus) WithContext(v context.Context) func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetTrialStatus) WithPretty() func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetTrialStatus) WithHuman() func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetTrialStatus) WithErrorTrace() func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetTrialStatus) WithFilterPath(v ...string) func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetTrialStatus) WithHeader(h map[string]string) func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGetTrialStatus) WithOpaqueID(s string) func(*LicenseGetTrialStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
