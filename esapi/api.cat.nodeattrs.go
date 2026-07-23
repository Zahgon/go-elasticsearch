package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatNodeattrsFunc(t Transport) CatNodeattrs {
	_ = "STUB: not implemented"
	return *new(CatNodeattrs)
}

type CatNodeattrs func(o ...func(*CatNodeattrsRequest)) (*Response, error)

type CatNodeattrsRequest struct {
	Bytes         string
	Format        string
	H             []string
	Help          *bool
	Local         *bool
	MasterTimeout time.Duration
	S             []string
	Time          string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatNodeattrsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatNodeattrs) WithContext(v context.Context) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithBytes(v string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithFormat(v string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithH(v ...string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithHelp(v bool) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithLocal(v bool) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithMasterTimeout(v time.Duration) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithS(v ...string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithTime(v string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithV(v bool) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithPretty() func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithHuman() func(*CatNodeattrsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodeattrs) WithErrorTrace() func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithFilterPath(v ...string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithHeader(h map[string]string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodeattrs) WithOpaqueID(s string) func(*CatNodeattrsRequest) {
	_ = "STUB: not implemented"
	return nil
}
