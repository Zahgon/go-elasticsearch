package esapi

import (
	"context"
	"io"
	"net/http"
)

func newClearScrollFunc(t Transport) ClearScroll {
	_ = "STUB: not implemented"
	return *new(ClearScroll)
}

type ClearScroll func(o ...func(*ClearScrollRequest)) (*Response, error)

type ClearScrollRequest struct {
	Body io.Reader

	ScrollID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClearScrollRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClearScroll) WithContext(v context.Context) func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClearScroll) WithBody(v io.Reader) func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClearScroll) WithScrollID(v ...string) func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClearScroll) WithPretty() func(*ClearScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f ClearScroll) WithHuman() func(*ClearScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f ClearScroll) WithErrorTrace() func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClearScroll) WithFilterPath(v ...string) func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClearScroll) WithHeader(h map[string]string) func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClearScroll) WithOpaqueID(s string) func(*ClearScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}
