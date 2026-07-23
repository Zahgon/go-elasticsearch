package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutIndexTemplateFunc(t Transport) IndicesPutIndexTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesPutIndexTemplate)
}

type IndicesPutIndexTemplate func(name string, body io.Reader, o ...func(*IndicesPutIndexTemplateRequest)) (*Response, error)

type IndicesPutIndexTemplateRequest struct {
	Body io.Reader

	Name string

	Cause         string
	Create        *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutIndexTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutIndexTemplate) WithContext(v context.Context) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithCause(v string) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithCreate(v bool) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithMasterTimeout(v time.Duration) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithPretty() func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithHuman() func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithErrorTrace() func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithFilterPath(v ...string) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithHeader(h map[string]string) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutIndexTemplate) WithOpaqueID(s string) func(*IndicesPutIndexTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
