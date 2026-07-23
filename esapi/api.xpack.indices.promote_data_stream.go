package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesPromoteDataStreamFunc(t Transport) IndicesPromoteDataStream {
	_ = "STUB: not implemented"
	return *new(IndicesPromoteDataStream)
}

type IndicesPromoteDataStream func(name string, o ...func(*IndicesPromoteDataStreamRequest)) (*Response, error)

type IndicesPromoteDataStreamRequest struct {
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

func (r IndicesPromoteDataStreamRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPromoteDataStream) WithContext(v context.Context) func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithMasterTimeout(v time.Duration) func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithPretty() func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithHuman() func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithErrorTrace() func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithFilterPath(v ...string) func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithHeader(h map[string]string) func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPromoteDataStream) WithOpaqueID(s string) func(*IndicesPromoteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}
