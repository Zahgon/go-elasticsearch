package esapi

import (
	"context"
	"net/http"
)

func newGetScriptLanguagesFunc(t Transport) GetScriptLanguages {
	_ = "STUB: not implemented"
	return *new(GetScriptLanguages)
}

type GetScriptLanguages func(o ...func(*GetScriptLanguagesRequest)) (*Response, error)

type GetScriptLanguagesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r GetScriptLanguagesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f GetScriptLanguages) WithContext(v context.Context) func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptLanguages) WithPretty() func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptLanguages) WithHuman() func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptLanguages) WithErrorTrace() func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptLanguages) WithFilterPath(v ...string) func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptLanguages) WithHeader(h map[string]string) func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScriptLanguages) WithOpaqueID(s string) func(*GetScriptLanguagesRequest) {
	_ = "STUB: not implemented"
	return nil
}
