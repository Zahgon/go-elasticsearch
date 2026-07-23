package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesSimulateIndexTemplateFunc(t Transport) IndicesSimulateIndexTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesSimulateIndexTemplate)
}

type IndicesSimulateIndexTemplate func(name string, o ...func(*IndicesSimulateIndexTemplateRequest)) (*Response, error)

type IndicesSimulateIndexTemplateRequest struct {
	Body io.Reader

	Name string

	Cause           string
	Create          *bool
	IncludeDefaults *bool
	MasterTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesSimulateIndexTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesSimulateIndexTemplate) WithContext(v context.Context) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithBody(v io.Reader) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithCause(v string) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithCreate(v bool) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithIncludeDefaults(v bool) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithMasterTimeout(v time.Duration) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithPretty() func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithHuman() func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithErrorTrace() func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithFilterPath(v ...string) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithHeader(h map[string]string) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateIndexTemplate) WithOpaqueID(s string) func(*IndicesSimulateIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
