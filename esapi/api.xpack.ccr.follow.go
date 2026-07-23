package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newCCRFollowFunc(t Transport) CCRFollow { _ = "STUB: not implemented"; return *new(CCRFollow) }

type CCRFollow func(index string, body io.Reader, o ...func(*CCRFollowRequest)) (*Response, error)

type CCRFollowRequest struct {
	Index string

	Body io.Reader

	MasterTimeout       time.Duration
	WaitForActiveShards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRFollowRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRFollow) WithContext(v context.Context) func(*CCRFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollow) WithMasterTimeout(v time.Duration) func(*CCRFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollow) WithWaitForActiveShards(v string) func(*CCRFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollow) WithPretty() func(*CCRFollowRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRFollow) WithHuman() func(*CCRFollowRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRFollow) WithErrorTrace() func(*CCRFollowRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRFollow) WithFilterPath(v ...string) func(*CCRFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollow) WithHeader(h map[string]string) func(*CCRFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollow) WithOpaqueID(s string) func(*CCRFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}
