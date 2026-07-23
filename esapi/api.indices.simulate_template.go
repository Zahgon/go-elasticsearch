package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesSimulateTemplateFunc(t Transport) IndicesSimulateTemplate {
	_ = "STUB: not implemented"
	return *new(IndicesSimulateTemplate)
}

type IndicesSimulateTemplate func(o ...func(*IndicesSimulateTemplateRequest)) (*Response, error)

type IndicesSimulateTemplateRequest struct {
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

func (r IndicesSimulateTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesSimulateTemplate) WithContext(v context.Context) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithBody(v io.Reader) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithName(v string) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithCause(v string) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithCreate(v bool) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithIncludeDefaults(v bool) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithMasterTimeout(v time.Duration) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithPretty() func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithHuman() func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithErrorTrace() func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithFilterPath(v ...string) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithHeader(h map[string]string) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSimulateTemplate) WithOpaqueID(s string) func(*IndicesSimulateTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
