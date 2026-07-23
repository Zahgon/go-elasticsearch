package esapi

import (
	"context"
	"net/http"
)

func newCatHelpFunc(t Transport) CatHelp { _ = "STUB: not implemented"; return *new(CatHelp) }

type CatHelp func(o ...func(*CatHelpRequest)) (*Response, error)

type CatHelpRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatHelpRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatHelp) WithContext(v context.Context) func(*CatHelpRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHelp) WithPretty() func(*CatHelpRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHelp) WithHuman() func(*CatHelpRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHelp) WithErrorTrace() func(*CatHelpRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHelp) WithFilterPath(v ...string) func(*CatHelpRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHelp) WithHeader(h map[string]string) func(*CatHelpRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHelp) WithOpaqueID(s string) func(*CatHelpRequest) {
	_ = "STUB: not implemented"
	return nil
}
