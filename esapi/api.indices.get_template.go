package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetTemplateFunc(t Transport) IndicesGetTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesGetTemplate)
}

type IndicesGetTemplate func(o ...func(*IndicesGetTemplateRequest)) (*Response, error)

type IndicesGetTemplateRequest struct {
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

func (r IndicesGetTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetTemplate) WithContext(v context.Context) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithName(v ...string) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithFlatSettings(v bool) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithLocal(v bool) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithMasterTimeout(v time.Duration) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithPretty() func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithHuman() func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithErrorTrace() func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithFilterPath(v ...string) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithHeader(h map[string]string) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetTemplate) WithOpaqueID(s string) func(*IndicesGetTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
