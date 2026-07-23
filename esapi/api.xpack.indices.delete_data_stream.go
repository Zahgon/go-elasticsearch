package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteDataStreamFunc(t Transport) IndicesDeleteDataStream {
	_ = "STUB: not implemented"
	return *new(IndicesDeleteDataStream)
}

type IndicesDeleteDataStream func(name []string, o ...func(*IndicesDeleteDataStreamRequest)) (*Response, error)

type IndicesDeleteDataStreamRequest struct {
	Name []string

	ExpandWildcards []string
	MasterTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesDeleteDataStreamRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDeleteDataStream) WithContext(v context.Context) func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithExpandWildcards(v ...string) func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithMasterTimeout(v time.Duration) func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithPretty() func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithHuman() func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithErrorTrace() func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithFilterPath(v ...string) func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithHeader(h map[string]string) func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStream) WithOpaqueID(s string) func(*IndicesDeleteDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}
