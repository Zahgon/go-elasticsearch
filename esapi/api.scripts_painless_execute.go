package esapi

import (
	"context"
	"io"
	"net/http"
)

func newScriptsPainlessExecuteFunc(t Transport) ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return *new(ScriptsPainlessExecute)
}

type ScriptsPainlessExecute func(body io.Reader, o ...func(*ScriptsPainlessExecuteRequest)) (*Response, error)

type ScriptsPainlessExecuteRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ScriptsPainlessExecuteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ScriptsPainlessExecute) WithContext(v context.Context) func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ScriptsPainlessExecute) WithPretty() func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ScriptsPainlessExecute) WithHuman() func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ScriptsPainlessExecute) WithErrorTrace() func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ScriptsPainlessExecute) WithFilterPath(v ...string) func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ScriptsPainlessExecute) WithHeader(h map[string]string) func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ScriptsPainlessExecute) WithOpaqueID(s string) func(*ScriptsPainlessExecuteRequest) {
	_ = "STUB: not implemented"
	return nil
}
