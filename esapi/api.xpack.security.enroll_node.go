package esapi

import (
	"context"
	"net/http"
)

func newSecurityEnrollNodeFunc(t Transport) SecurityEnrollNode {
	_ = "STUB: not implemented"
	return *new(SecurityEnrollNode)
}

type SecurityEnrollNode func(o ...func(*SecurityEnrollNodeRequest)) (*Response, error)

type SecurityEnrollNodeRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityEnrollNodeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityEnrollNode) WithContext(v context.Context) func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollNode) WithPretty() func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollNode) WithHuman() func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollNode) WithErrorTrace() func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollNode) WithFilterPath(v ...string) func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollNode) WithHeader(h map[string]string) func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnrollNode) WithOpaqueID(s string) func(*SecurityEnrollNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}
