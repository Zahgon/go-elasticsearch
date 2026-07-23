package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newLicensePostFunc(t Transport) LicensePost {
	_ = "STUB: not implemented"
	return *new(LicensePost)
}

type LicensePost func(o ...func(*LicensePostRequest)) (*Response, error)

type LicensePostRequest struct {
	Body io.Reader

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

func (r LicensePostRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LicensePost) WithContext(v context.Context) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithBody(v io.Reader) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithAcknowledge(v bool) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithMasterTimeout(v time.Duration) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithTimeout(v time.Duration) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithPretty() func(*LicensePostRequest) { _ = "STUB: not implemented"; return nil }

func (f LicensePost) WithHuman() func(*LicensePostRequest) { _ = "STUB: not implemented"; return nil }

func (f LicensePost) WithErrorTrace() func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithFilterPath(v ...string) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithHeader(h map[string]string) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LicensePost) WithOpaqueID(s string) func(*LicensePostRequest) {
	_ = "STUB: not implemented"
	return nil
}
