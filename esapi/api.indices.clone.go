package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesCloneFunc(t Transport) IndicesClone {
	_ = "STUB: not implemented"
	return *new(IndicesClone)
}

type IndicesClone func(index string, target string, o ...func(*IndicesCloneRequest)) (*Response, error)

type IndicesCloneRequest struct {
	Index string

	Body io.Reader

	Target string

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

func (r IndicesCloneRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesClone) WithContext(v context.Context) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithBody(v io.Reader) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithMasterTimeout(v time.Duration) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithTimeout(v time.Duration) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithWaitForActiveShards(v string) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithPretty() func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithHuman() func(*IndicesCloneRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesClone) WithErrorTrace() func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithFilterPath(v ...string) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithHeader(h map[string]string) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClone) WithOpaqueID(s string) func(*IndicesCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}
