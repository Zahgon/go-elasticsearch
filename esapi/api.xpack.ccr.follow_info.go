package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRFollowInfoFunc(t Transport) CCRFollowInfo {
	_ = "STUB: not implemented"
	return *new(CCRFollowInfo)
}

type CCRFollowInfo func(index []string, o ...func(*CCRFollowInfoRequest)) (*Response, error)

type CCRFollowInfoRequest struct {
	Index []string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRFollowInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRFollowInfo) WithContext(v context.Context) func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithMasterTimeout(v time.Duration) func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithPretty() func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithHuman() func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithErrorTrace() func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithFilterPath(v ...string) func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithHeader(h map[string]string) func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowInfo) WithOpaqueID(s string) func(*CCRFollowInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}
