package esapi

import (
	"context"
	"net/http"
)

func newRollupGetRollupIndexCapsFunc(t Transport) RollupGetRollupIndexCaps {
	_ = "STUB: not implemented"
	return *new(RollupGetRollupIndexCaps)
}

type RollupGetRollupIndexCaps func(index []string, o ...func(*RollupGetRollupIndexCapsRequest)) (*Response, error)

type RollupGetRollupIndexCapsRequest struct {
	Index []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupGetRollupIndexCapsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupGetRollupIndexCaps) WithContext(v context.Context) func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupIndexCaps) WithPretty() func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupIndexCaps) WithHuman() func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupIndexCaps) WithErrorTrace() func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupIndexCaps) WithFilterPath(v ...string) func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupIndexCaps) WithHeader(h map[string]string) func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetRollupIndexCaps) WithOpaqueID(s string) func(*RollupGetRollupIndexCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}
