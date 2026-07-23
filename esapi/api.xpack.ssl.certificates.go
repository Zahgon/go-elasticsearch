package esapi

import (
	"context"
	"net/http"
)

func newSSLCertificatesFunc(t Transport) SSLCertificates {
	_ = "STUB: not implemented"
	return *new(SSLCertificates)
}

type SSLCertificates func(o ...func(*SSLCertificatesRequest)) (*Response, error)

type SSLCertificatesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SSLCertificatesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SSLCertificates) WithContext(v context.Context) func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SSLCertificates) WithPretty() func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SSLCertificates) WithHuman() func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SSLCertificates) WithErrorTrace() func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SSLCertificates) WithFilterPath(v ...string) func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SSLCertificates) WithHeader(h map[string]string) func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SSLCertificates) WithOpaqueID(s string) func(*SSLCertificatesRequest) {
	_ = "STUB: not implemented"
	return nil
}
