package esapi

import (
	"context"
	"net/http"
)

func newLicenseGetFunc(t Transport) LicenseGet { _ = "STUB: not implemented"; return *new(LicenseGet) }

type LicenseGet func(o ...func(*LicenseGetRequest)) (*Response, error)

type LicenseGetRequest struct {
	AcceptEnterprise *bool
	Local            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LicenseGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicenseGet) WithContext(v context.Context) func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGet) WithAcceptEnterprise(v bool) func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGet) WithLocal(v bool) func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGet) WithPretty() func(*LicenseGetRequest) { _ = "STUB: not implemented"; return nil }

func (f LicenseGet) WithHuman() func(*LicenseGetRequest) { _ = "STUB: not implemented"; return nil }

func (f LicenseGet) WithErrorTrace() func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGet) WithFilterPath(v ...string) func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGet) WithHeader(h map[string]string) func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicenseGet) WithOpaqueID(s string) func(*LicenseGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
