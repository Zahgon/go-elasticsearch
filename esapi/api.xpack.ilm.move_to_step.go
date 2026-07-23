package esapi

import (
	"context"
	"io"
	"net/http"
)

func newILMMoveToStepFunc(t Transport) ILMMoveToStep {
	_ = "STUB: not implemented"
	return *new(ILMMoveToStep)
}

type ILMMoveToStep func(index string, body io.Reader, o ...func(*ILMMoveToStepRequest)) (*Response, error)

type ILMMoveToStepRequest struct {
	Index string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ILMMoveToStepRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMMoveToStep) WithContext(v context.Context) func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMoveToStep) WithPretty() func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMoveToStep) WithHuman() func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMoveToStep) WithErrorTrace() func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMoveToStep) WithFilterPath(v ...string) func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMoveToStep) WithHeader(h map[string]string) func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMoveToStep) WithOpaqueID(s string) func(*ILMMoveToStepRequest) {
	_ = "STUB: not implemented"
	return nil
}
