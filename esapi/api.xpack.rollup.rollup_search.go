package esapi

import (
	"context"
	"io"
	"net/http"
)

func newRollupRollupSearchFunc(t Transport) RollupRollupSearch {
	_ = "STUB: not implemented"
	return *new(RollupRollupSearch)
}

type RollupRollupSearch func(index []string, body io.Reader, o ...func(*RollupRollupSearchRequest)) (*Response, error)

type RollupRollupSearchRequest struct {
	Index []string

	Body io.Reader

	RestTotalHitsAsInt *bool
	TypedKeys          *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupRollupSearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupRollupSearch) WithContext(v context.Context) func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithRestTotalHitsAsInt(v bool) func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithTypedKeys(v bool) func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithPretty() func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithHuman() func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithErrorTrace() func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithFilterPath(v ...string) func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithHeader(h map[string]string) func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupRollupSearch) WithOpaqueID(s string) func(*RollupRollupSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
