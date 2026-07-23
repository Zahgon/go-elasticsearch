package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityDelegatePkiFunc(t Transport) SecurityDelegatePki {
	_ = "STUB: not implemented"
	return *new(SecurityDelegatePki)
}

type SecurityDelegatePki func(body io.Reader, o ...func(*SecurityDelegatePkiRequest)) (*Response, error)

type SecurityDelegatePkiRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityDelegatePkiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDelegatePki) WithContext(v context.Context) func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDelegatePki) WithPretty() func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDelegatePki) WithHuman() func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDelegatePki) WithErrorTrace() func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDelegatePki) WithFilterPath(v ...string) func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDelegatePki) WithHeader(h map[string]string) func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDelegatePki) WithOpaqueID(s string) func(*SecurityDelegatePkiRequest) {
	_ = "STUB: not implemented"
	return nil
}
