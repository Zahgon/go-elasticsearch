package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutTemplateFunc(t Transport) IndicesPutTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesPutTemplate)
}

type IndicesPutTemplate func(name string, body io.Reader, o ...func(*IndicesPutTemplateRequest)) (*Response, error)

type IndicesPutTemplateRequest struct {
	Body io.Reader

	Name string

	Cause         string
	Create        *bool
	MasterTimeout time.Duration
	Order         *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutTemplate) WithContext(v context.Context) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithCause(v string) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithCreate(v bool) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithMasterTimeout(v time.Duration) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithOrder(v int) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithPretty() func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithHuman() func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithErrorTrace() func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithFilterPath(v ...string) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithHeader(h map[string]string) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutTemplate) WithOpaqueID(s string) func(*IndicesPutTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
