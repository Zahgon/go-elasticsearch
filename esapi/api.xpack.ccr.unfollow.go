package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRUnfollowFunc(t Transport) CCRUnfollow {
	_ = "STUB: not implemented"
	return *new(CCRUnfollow)
}

type CCRUnfollow func(index string, o ...func(*CCRUnfollowRequest)) (*Response, error)

type CCRUnfollowRequest struct {
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

func (r CCRUnfollowRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRUnfollow) WithContext(v context.Context) func(*CCRUnfollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRUnfollow) WithMasterTimeout(v time.Duration) func(*CCRUnfollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRUnfollow) WithPretty() func(*CCRUnfollowRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRUnfollow) WithHuman() func(*CCRUnfollowRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRUnfollow) WithErrorTrace() func(*CCRUnfollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRUnfollow) WithFilterPath(v ...string) func(*CCRUnfollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRUnfollow) WithHeader(h map[string]string) func(*CCRUnfollowRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRUnfollow) WithOpaqueID(s string) func(*CCRUnfollowRequest) {
	_ = "STUB: not implemented"
	return nil
}
