package esapi

import (
	"context"
	"net/http"
)

func newCatFielddataFunc(t Transport) CatFielddata {
	_ = "STUB: not implemented"
	return *new(CatFielddata)
}

type CatFielddata func(o ...func(*CatFielddataRequest)) (*Response, error)

type CatFielddataRequest struct {
	Fields []string

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

func (r CatFielddataRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatFielddata) WithContext(v context.Context) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithFields(v ...string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithBytes(v string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithFormat(v string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithH(v ...string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithHelp(v bool) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithS(v ...string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithTime(v string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithV(v bool) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithPretty() func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithHuman() func(*CatFielddataRequest) { _ = "STUB: not implemented"; return nil }

func (f CatFielddata) WithErrorTrace() func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithFilterPath(v ...string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithHeader(h map[string]string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatFielddata) WithOpaqueID(s string) func(*CatFielddataRequest) {
	_ = "STUB: not implemented"
	return nil
}
