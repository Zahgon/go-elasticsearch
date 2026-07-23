package esapi

import (
	"context"
	"net/http"
)

func newSecurityEnrollKibanaFunc(t Transport) SecurityEnrollKibana {
	_ = "STUB: not implemented"
	return *new(SecurityEnrollKibana)
}

type SecurityEnrollKibana func(o ...func(*SecurityEnrollKibanaRequest)) (*Response, error)

type SecurityEnrollKibanaRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityEnrollKibanaRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityEnrollKibana) WithContext(v context.Context) func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollKibana) WithPretty() func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollKibana) WithHuman() func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollKibana) WithErrorTrace() func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollKibana) WithFilterPath(v ...string) func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollKibana) WithHeader(h map[string]string) func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollKibana) WithOpaqueID(s string) func(*SecurityEnrollKibanaRequest) {
	_ = "STUB: not implemented"
	return nil
}
