package esapi

import (
	"context"
	"net/http"
	"time"
)

func newGetScriptFunc(t Transport) GetScript { _ = "STUB: not implemented"; return *new(GetScript) }

type GetScript func(id string, o ...func(*GetScriptRequest)) (*Response, error)

type GetScriptRequest struct {
	ScriptID string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r GetScriptRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f GetScript) WithContext(v context.Context) func(*GetScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScript) WithMasterTimeout(v time.Duration) func(*GetScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScript) WithPretty() func(*GetScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f GetScript) WithHuman() func(*GetScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f GetScript) WithErrorTrace() func(*GetScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f GetScript) WithFilterPath(v ...string) func(*GetScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScript) WithHeader(h map[string]string) func(*GetScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetScript) WithOpaqueID(s string) func(*GetScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}
