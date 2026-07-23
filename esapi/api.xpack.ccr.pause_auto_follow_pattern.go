package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRPauseAutoFollowPatternFunc(t Transport) CCRPauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(CCRPauseAutoFollowPattern)
}

type CCRPauseAutoFollowPattern func(name string, o ...func(*CCRPauseAutoFollowPatternRequest)) (*Response, error)

type CCRPauseAutoFollowPatternRequest struct {
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

func (r CCRPauseAutoFollowPatternRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRPauseAutoFollowPattern) WithContext(v context.Context) func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithMasterTimeout(v time.Duration) func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithPretty() func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithHuman() func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithErrorTrace() func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithFilterPath(v ...string) func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithHeader(h map[string]string) func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseAutoFollowPattern) WithOpaqueID(s string) func(*CCRPauseAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}
