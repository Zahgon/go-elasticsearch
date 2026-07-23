package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRPauseFollowFunc(t Transport) CCRPauseFollow {
	_ = "STUB: not implemented"
	return *new(CCRPauseFollow)
}

type CCRPauseFollow func(index string, o ...func(*CCRPauseFollowRequest)) (*Response, error)

type CCRPauseFollowRequest struct {
	Index string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRPauseFollowRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRPauseFollow) WithContext(v context.Context) func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithMasterTimeout(v time.Duration) func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithPretty() func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithHuman() func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithErrorTrace() func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithFilterPath(v ...string) func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithHeader(h map[string]string) func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPauseFollow) WithOpaqueID(s string) func(*CCRPauseFollowRequest) {
	_ = "STUB: not implemented"
	return nil
}
