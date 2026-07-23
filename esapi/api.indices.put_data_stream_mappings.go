package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutDataStreamMappingsFunc(t Transport) IndicesPutDataStreamMappings {
	_ = "STUB: not implemented"
	return *new(IndicesPutDataStreamMappings)
}

type IndicesPutDataStreamMappings func(name []string, body io.Reader, o ...func(*IndicesPutDataStreamMappingsRequest)) (*Response, error)

type IndicesPutDataStreamMappingsRequest struct {
	Body io.Reader

	Name []string

	DryRun        *bool
	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutDataStreamMappingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutDataStreamMappings) WithContext(v context.Context) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithDryRun(v bool) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithMasterTimeout(v time.Duration) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithTimeout(v time.Duration) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithPretty() func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithHuman() func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithErrorTrace() func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithFilterPath(v ...string) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithHeader(h map[string]string) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamMappings) WithOpaqueID(s string) func(*IndicesPutDataStreamMappingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
