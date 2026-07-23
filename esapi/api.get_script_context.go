package esapi

import (
	"context"
	"net/http"
)

func newGetScriptContextFunc(t Transport) GetScriptContext {
	_ = "STUB: not implemented"
	return *new(GetScriptContext)
}

type GetScriptContext func(o ...func(*GetScriptContextRequest)) (*Response, error)

type GetScriptContextRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r GetScriptContextRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f GetScriptContext) WithContext(v context.Context) func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptContext) WithPretty() func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptContext) WithHuman() func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptContext) WithErrorTrace() func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptContext) WithFilterPath(v ...string) func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptContext) WithHeader(h map[string]string) func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptContext) WithOpaqueID(s string) func(*GetScriptContextRequest) {
	_ = "STUB: not implemented"
	return nil
}
