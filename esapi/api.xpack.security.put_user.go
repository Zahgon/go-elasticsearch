package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityPutUserFunc(t Transport) SecurityPutUser {
	_ = "STUB: not implemented"
	return *new(SecurityPutUser)
}

type SecurityPutUser func(username string, body io.Reader, o ...func(*SecurityPutUserRequest)) (*Response, error)

type SecurityPutUserRequest struct {
	Body io.Reader

	Username string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityPutUserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityPutUser) WithContext(v context.Context) func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithRefresh(v string) func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithPretty() func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithHuman() func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithErrorTrace() func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithFilterPath(v ...string) func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithHeader(h map[string]string) func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutUser) WithOpaqueID(s string) func(*SecurityPutUserRequest) {
	_ = "STUB: not implemented"
	return nil
}
