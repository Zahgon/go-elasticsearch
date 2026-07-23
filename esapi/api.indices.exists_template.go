package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesExistsTemplateFunc(t Transport) IndicesExistsTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesExistsTemplate)
}

type IndicesExistsTemplate func(name []string, o ...func(*IndicesExistsTemplateRequest)) (*Response, error)

type IndicesExistsTemplateRequest struct {
	Name []string

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

func (r IndicesExistsTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesExistsTemplate) WithContext(v context.Context) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithFlatSettings(v bool) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithLocal(v bool) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithMasterTimeout(v time.Duration) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithPretty() func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithHuman() func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithErrorTrace() func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithFilterPath(v ...string) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithHeader(h map[string]string) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsTemplate) WithOpaqueID(s string) func(*IndicesExistsTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
