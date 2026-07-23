package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatMasterFunc(t Transport) CatMaster { _ = "STUB: not implemented"; return *new(CatMaster) }

type CatMaster func(o ...func(*CatMasterRequest)) (*Response, error)

type CatMasterRequest struct {
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

func (r CatMasterRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatMaster) WithContext(v context.Context) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithBytes(v string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithFormat(v string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithH(v ...string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithHelp(v bool) func(*CatMasterRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMaster) WithLocal(v bool) func(*CatMasterRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMaster) WithMasterTimeout(v time.Duration) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithS(v ...string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithTime(v string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithV(v bool) func(*CatMasterRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMaster) WithPretty() func(*CatMasterRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMaster) WithHuman() func(*CatMasterRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMaster) WithErrorTrace() func(*CatMasterRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMaster) WithFilterPath(v ...string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithHeader(h map[string]string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMaster) WithOpaqueID(s string) func(*CatMasterRequest) {
	_ = "STUB: not implemented"
	return nil
}
