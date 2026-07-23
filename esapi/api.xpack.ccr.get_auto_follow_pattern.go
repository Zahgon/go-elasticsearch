package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRGetAutoFollowPatternFunc(t Transport) CCRGetAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(CCRGetAutoFollowPattern)
}

type CCRGetAutoFollowPattern func(o ...func(*CCRGetAutoFollowPatternRequest)) (*Response, error)

type CCRGetAutoFollowPatternRequest struct {
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

func (r CCRGetAutoFollowPatternRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRGetAutoFollowPattern) WithContext(v context.Context) func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithName(v string) func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithMasterTimeout(v time.Duration) func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithPretty() func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithHuman() func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithErrorTrace() func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithFilterPath(v ...string) func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithHeader(h map[string]string) func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRGetAutoFollowPattern) WithOpaqueID(s string) func(*CCRGetAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}
