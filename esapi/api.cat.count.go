package esapi

import (
	"context"
	"net/http"
)

func newCatCountFunc(t Transport) CatCount { _ = "STUB: not implemented"; return *new(CatCount) }

type CatCount func(o ...func(*CatCountRequest)) (*Response, error)

type CatCountRequest struct {
	Index []string

	Bytes  string
	Format string
	H      []string
	Help   *bool
	S      []string
	Time   string
	V      *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatCountRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatCount) WithContext(v context.Context) func(*CatCountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCount) WithIndex(v ...string) func(*CatCountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCount) WithBytes(v string) func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithFormat(v string) func(*CatCountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCount) WithH(v ...string) func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithHelp(v bool) func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithS(v ...string) func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithTime(v string) func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithV(v bool) func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithPretty() func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithHuman() func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithErrorTrace() func(*CatCountRequest) { _ = "STUB: not implemented"; return nil }

func (f CatCount) WithFilterPath(v ...string) func(*CatCountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCount) WithHeader(h map[string]string) func(*CatCountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCount) WithOpaqueID(s string) func(*CatCountRequest) {
	_ = "STUB: not implemented"
	return nil
}
