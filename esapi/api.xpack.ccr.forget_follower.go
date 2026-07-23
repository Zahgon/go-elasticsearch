package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newCCRForgetFollowerFunc(t Transport) CCRForgetFollower {
	_ = "STUB: not implemented"
	return *new(CCRForgetFollower)
}

type CCRForgetFollower func(index string, body io.Reader, o ...func(*CCRForgetFollowerRequest)) (*Response, error)

type CCRForgetFollowerRequest struct {
	Index string

	Body io.Reader

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRForgetFollowerRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRForgetFollower) WithContext(v context.Context) func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithTimeout(v time.Duration) func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithPretty() func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithHuman() func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithErrorTrace() func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithFilterPath(v ...string) func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithHeader(h map[string]string) func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRForgetFollower) WithOpaqueID(s string) func(*CCRForgetFollowerRequest) {
	_ = "STUB: not implemented"
	return nil
}
