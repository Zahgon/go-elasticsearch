package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRResumeAutoFollowPatternFunc(t Transport) CCRResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(CCRResumeAutoFollowPattern)
}

type CCRResumeAutoFollowPattern func(name string, o ...func(*CCRResumeAutoFollowPatternRequest)) (*Response, error)

type CCRResumeAutoFollowPatternRequest struct {
	Name string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRResumeAutoFollowPatternRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRResumeAutoFollowPattern) WithContext(v context.Context) func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithMasterTimeout(v time.Duration) func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithPretty() func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithHuman() func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithErrorTrace() func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithFilterPath(v ...string) func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithHeader(h map[string]string) func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRResumeAutoFollowPattern) WithOpaqueID(s string) func(*CCRResumeAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}
