package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newPutScriptFunc(t Transport) PutScript { _ = "STUB: not implemented"; return *new(PutScript) }

type PutScript func(id string, body io.Reader, o ...func(*PutScriptRequest)) (*Response, error)

type PutScriptRequest struct {
	ScriptID string

	Body io.Reader

	ScriptContext string

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

func (r PutScriptRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f PutScript) WithContext(v context.Context) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f PutScript) WithScriptContext(v string) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f PutScript) WithMasterTimeout(v time.Duration) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f PutScript) WithTimeout(v time.Duration) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f PutScript) WithPretty() func(*PutScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f PutScript) WithHuman() func(*PutScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f PutScript) WithErrorTrace() func(*PutScriptRequest) { _ = "STUB: not implemented"; return nil }

func (f PutScript) WithFilterPath(v ...string) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f PutScript) WithHeader(h map[string]string) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f PutScript) WithOpaqueID(s string) func(*PutScriptRequest) {
	_ = "STUB: not implemented"
	return nil
}
