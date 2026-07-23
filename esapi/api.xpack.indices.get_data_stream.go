package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetDataStreamFunc(t Transport) IndicesGetDataStream {
	_ = "STUB: not implemented"
	return *new(IndicesGetDataStream)
}

type IndicesGetDataStream func(o ...func(*IndicesGetDataStreamRequest)) (*Response, error)

type IndicesGetDataStreamRequest struct {
	Name []string

	ExpandWildcards []string
	IncludeDefaults *bool
	MasterTimeout   time.Duration
	Verbose         *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetDataStreamRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetDataStream) WithContext(v context.Context) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithName(v ...string) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithExpandWildcards(v ...string) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithIncludeDefaults(v bool) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithMasterTimeout(v time.Duration) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithVerbose(v bool) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithPretty() func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithHuman() func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithErrorTrace() func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithFilterPath(v ...string) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithHeader(h map[string]string) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStream) WithOpaqueID(s string) func(*IndicesGetDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}
