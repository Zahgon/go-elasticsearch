package esapi

import (
	"context"
	"io"
	"net/http"
)

func newRenderSearchTemplateFunc(t Transport) RenderSearchTemplate {
	_ = "STUB: not implemented"
	return *new(RenderSearchTemplate)
}

type RenderSearchTemplate func(body io.Reader, o ...func(*RenderSearchTemplateRequest)) (*Response, error)

type RenderSearchTemplateRequest struct {
	TemplateID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RenderSearchTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RenderSearchTemplate) WithContext(v context.Context) func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithTemplateID(v string) func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithPretty() func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithHuman() func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithErrorTrace() func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithFilterPath(v ...string) func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithHeader(h map[string]string) func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RenderSearchTemplate) WithOpaqueID(s string) func(*RenderSearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
