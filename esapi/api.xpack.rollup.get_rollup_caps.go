package esapi

import (
	"context"
	"net/http"
)

func newRollupGetRollupCapsFunc(t Transport) RollupGetRollupCaps {
	_ = "STUB: not implemented"
	return *new(RollupGetRollupCaps)
}

type RollupGetRollupCaps func(o ...func(*RollupGetRollupCapsRequest)) (*Response, error)

type RollupGetRollupCapsRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupGetRollupCapsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupGetRollupCaps) WithContext(v context.Context) func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithIndex(v string) func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithPretty() func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithHuman() func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithErrorTrace() func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithFilterPath(v ...string) func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithHeader(h map[string]string) func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupCaps) WithOpaqueID(s string) func(*RollupGetRollupCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}
