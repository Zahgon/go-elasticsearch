package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRDeleteAutoFollowPatternFunc(t Transport) CCRDeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(CCRDeleteAutoFollowPattern)
}

type CCRDeleteAutoFollowPattern func(name string, o ...func(*CCRDeleteAutoFollowPatternRequest)) (*Response, error)

type CCRDeleteAutoFollowPatternRequest struct {
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

func (r CCRDeleteAutoFollowPatternRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRDeleteAutoFollowPattern) WithContext(v context.Context) func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithMasterTimeout(v time.Duration) func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithPretty() func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithHuman() func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithErrorTrace() func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithFilterPath(v ...string) func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithHeader(h map[string]string) func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRDeleteAutoFollowPattern) WithOpaqueID(s string) func(*CCRDeleteAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}
