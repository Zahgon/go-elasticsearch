package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteTemplateFunc(t Transport) IndicesDeleteTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesDeleteTemplate)
}

type IndicesDeleteTemplate func(name string, o ...func(*IndicesDeleteTemplateRequest)) (*Response, error)

type IndicesDeleteTemplateRequest struct {
	Name string

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

func (r IndicesDeleteTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDeleteTemplate) WithContext(v context.Context) func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithMasterTimeout(v time.Duration) func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithTimeout(v time.Duration) func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithPretty() func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithHuman() func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithErrorTrace() func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithFilterPath(v ...string) func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithHeader(h map[string]string) func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteTemplate) WithOpaqueID(s string) func(*IndicesDeleteTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
