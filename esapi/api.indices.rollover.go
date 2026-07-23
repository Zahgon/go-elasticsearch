package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesRolloverFunc(t Transport) IndicesRollover {
	_ = "STUB: not implemented"
	return *new(IndicesRollover)
}

type IndicesRollover func(alias string, o ...func(*IndicesRolloverRequest)) (*Response, error)

type IndicesRolloverRequest struct {
	Body io.Reader

	Alias    string
	NewIndex string

	DryRun              *bool
	Lazy                *bool
	MasterTimeout       time.Duration
	Timeout             time.Duration
	WaitForActiveShards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesRolloverRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesRollover) WithContext(v context.Context) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithBody(v io.Reader) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithNewIndex(v string) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithDryRun(v bool) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithLazy(v bool) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithMasterTimeout(v time.Duration) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithTimeout(v time.Duration) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithWaitForActiveShards(v string) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithPretty() func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithHuman() func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithErrorTrace() func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithFilterPath(v ...string) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithHeader(h map[string]string) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRollover) WithOpaqueID(s string) func(*IndicesRolloverRequest) {
	_ = "STUB: not implemented"
	return nil
}
