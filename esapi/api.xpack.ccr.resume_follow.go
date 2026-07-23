package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newCCRResumeFollowFunc(t Transport) CCRResumeFollow {
	_ = "STUB: not implemented"
	return *new(CCRResumeFollow)
}

type CCRResumeFollow func(index string, o ...func(*CCRResumeFollowRequest)) (*Response, error)

type CCRResumeFollowRequest struct {
	Index string

	Body io.Reader

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRResumeFollowRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRResumeFollow) WithContext(v context.Context) func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithBody(v io.Reader) func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithMasterTimeout(v time.Duration) func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithPretty() func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithHuman() func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithErrorTrace() func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithFilterPath(v ...string) func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithHeader(h map[string]string) func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeFollow) WithOpaqueID(s string) func(*CCRResumeFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}
