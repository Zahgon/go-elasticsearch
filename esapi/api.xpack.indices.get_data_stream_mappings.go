package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetDataStreamMappingsFunc(t Transport) IndicesGetDataStreamMappings {
	_ = "STUB: not implemented"
	return *new(IndicesGetDataStreamMappings)
}

type IndicesGetDataStreamMappings func(name []string, o ...func(*IndicesGetDataStreamMappingsRequest)) (*Response, error)

type IndicesGetDataStreamMappingsRequest struct {
	Name []string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetDataStreamMappingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetDataStreamMappings) WithContext(v context.Context) func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithMasterTimeout(v time.Duration) func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithPretty() func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithHuman() func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithErrorTrace() func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithFilterPath(v ...string) func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithHeader(h map[string]string) func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamMappings) WithOpaqueID(s string) func(*IndicesGetDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
