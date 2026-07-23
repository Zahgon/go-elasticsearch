package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityInvalidateTokenFunc(t Transport) SecurityInvalidateToken {
	_ = "STUB: not implemented"
	return *new(SecurityInvalidateToken)
}

type SecurityInvalidateToken func(body io.Reader, o ...func(*SecurityInvalidateTokenRequest)) (*Response, error)

type SecurityInvalidateTokenRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityInvalidateTokenRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityInvalidateToken) WithContext(v context.Context) func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateToken) WithPretty() func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateToken) WithHuman() func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateToken) WithErrorTrace() func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateToken) WithFilterPath(v ...string) func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateToken) WithHeader(h map[string]string) func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateToken) WithOpaqueID(s string) func(*SecurityInvalidateTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}
