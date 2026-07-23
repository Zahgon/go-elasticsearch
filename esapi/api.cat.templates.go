package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatTemplatesFunc(t Transport) CatTemplates {
	_ = "STUB: not implemented"
	return *new(CatTemplates)
}

type CatTemplates func(o ...func(*CatTemplatesRequest)) (*Response, error)

type CatTemplatesRequest struct {
	Name string

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

func (r CatTemplatesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatTemplates) WithContext(v context.Context) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithName(v string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithBytes(v string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithFormat(v string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithH(v ...string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithHelp(v bool) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithLocal(v bool) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithMasterTimeout(v time.Duration) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithS(v ...string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithTime(v string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithV(v bool) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithPretty() func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithHuman() func(*CatTemplatesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTemplates) WithErrorTrace() func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithFilterPath(v ...string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithHeader(h map[string]string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTemplates) WithOpaqueID(s string) func(*CatTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}
