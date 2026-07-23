package esapi

import (
	"context"
	"net/http"
	"time"
)

func newDeleteScriptFunc(t Transport) DeleteScript {
	_ = "STUB: not implemented"
	return *new(DeleteScript)
}

type DeleteScript func(id string, o ...func(*DeleteScriptRequest)) (*Response, error)

type DeleteScriptRequest struct {
	ScriptID string

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

func (r DeleteScriptRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f DeleteScript) WithContext(v context.Context) func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithMasterTimeout(v time.Duration) func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithTimeout(v time.Duration) func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithPretty() func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithHuman() func(*DeleteScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f DeleteScript) WithErrorTrace() func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithFilterPath(v ...string) func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithHeader(h map[string]string) func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteScript) WithOpaqueID(s string) func(*DeleteScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}
