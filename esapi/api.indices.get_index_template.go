package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetIndexTemplateFunc(t Transport) IndicesGetIndexTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesGetIndexTemplate)
}

type IndicesGetIndexTemplate func(o ...func(*IndicesGetIndexTemplateRequest)) (*Response, error)

type IndicesGetIndexTemplateRequest struct {
	Name string

	FlatSettings    *bool
	IncludeDefaults *bool
	Local           *bool
	MasterTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetIndexTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetIndexTemplate) WithContext(v context.Context) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithName(v string) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithFlatSettings(v bool) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithIncludeDefaults(v bool) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithLocal(v bool) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithMasterTimeout(v time.Duration) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithPretty() func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithHuman() func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithErrorTrace() func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithFilterPath(v ...string) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithHeader(h map[string]string) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetIndexTemplate) WithOpaqueID(s string) func(*IndicesGetIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
