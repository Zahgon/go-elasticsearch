package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesCreateFunc(t Transport) IndicesCreate {
	_ = "STUB: not implemented"
	return *new(IndicesCreate)
}

type IndicesCreate func(index string, o ...func(*IndicesCreateRequest)) (*Response, error)

type IndicesCreateRequest struct {
	Index string

	Body io.Reader

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

func (r IndicesCreateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesCreate) WithContext(v context.Context) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithBody(v io.Reader) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithMasterTimeout(v time.Duration) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithTimeout(v time.Duration) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithWaitForActiveShards(v string) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithPretty() func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithHuman() func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithErrorTrace() func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithFilterPath(v ...string) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithHeader(h map[string]string) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreate) WithOpaqueID(s string) func(*IndicesCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}
