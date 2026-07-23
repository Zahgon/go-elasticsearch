package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteIndexTemplateFunc(t Transport) IndicesDeleteIndexTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesDeleteIndexTemplate)
}

type IndicesDeleteIndexTemplate func(name []string, o ...func(*IndicesDeleteIndexTemplateRequest)) (*Response, error)

type IndicesDeleteIndexTemplateRequest struct {
	Name []string

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

func (r IndicesDeleteIndexTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDeleteIndexTemplate) WithContext(v context.Context) func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithMasterTimeout(v time.Duration) func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithTimeout(v time.Duration) func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithPretty() func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithHuman() func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithErrorTrace() func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithFilterPath(v ...string) func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithHeader(h map[string]string) func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteIndexTemplate) WithOpaqueID(s string) func(*IndicesDeleteIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
