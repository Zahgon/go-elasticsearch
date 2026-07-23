package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesExistsIndexTemplateFunc(t Transport) IndicesExistsIndexTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesExistsIndexTemplate)
}

type IndicesExistsIndexTemplate func(name string, o ...func(*IndicesExistsIndexTemplateRequest)) (*Response, error)

type IndicesExistsIndexTemplateRequest struct {
	Name string

	FlatSettings  *bool
	Local         *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesExistsIndexTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesExistsIndexTemplate) WithContext(v context.Context) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithFlatSettings(v bool) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithLocal(v bool) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithMasterTimeout(v time.Duration) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithPretty() func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithHuman() func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithErrorTrace() func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithFilterPath(v ...string) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithHeader(h map[string]string) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsIndexTemplate) WithOpaqueID(s string) func(*IndicesExistsIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
