package esapi

import (
	"context"
	"net/http"
)

func newCatHealthFunc(t Transport) CatHealth { _ = "STUB: not implemented"; return *new(CatHealth) }

type CatHealth func(o ...func(*CatHealthRequest)) (*Response, error)

type CatHealthRequest struct {
	Bytes  string
	Format string
	H      []string
	Help   *bool
	S      []string
	Time   string
	Ts     *bool
	V      *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatHealthRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatHealth) WithContext(v context.Context) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithBytes(v string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithFormat(v string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithH(v ...string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithHelp(v bool) func(*CatHealthRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHealth) WithS(v ...string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithTime(v string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithTs(v bool) func(*CatHealthRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHealth) WithV(v bool) func(*CatHealthRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHealth) WithPretty() func(*CatHealthRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHealth) WithHuman() func(*CatHealthRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHealth) WithErrorTrace() func(*CatHealthRequest) { _ = "STUB: not implemented"; return nil }

func (f CatHealth) WithFilterPath(v ...string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithHeader(h map[string]string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatHealth) WithOpaqueID(s string) func(*CatHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}
