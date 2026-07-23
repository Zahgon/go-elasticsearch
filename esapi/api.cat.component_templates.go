package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatComponentTemplatesFunc(t Transport) CatComponentTemplates {
	_ = "STUB: not implemented"
	return *new(CatComponentTemplates)
}

type CatComponentTemplates func(o ...func(*CatComponentTemplatesRequest)) (*Response, error)

type CatComponentTemplatesRequest struct {
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

func (r CatComponentTemplatesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatComponentTemplates) WithContext(v context.Context) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithName(v string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithBytes(v string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithFormat(v string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithH(v ...string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithHelp(v bool) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithLocal(v bool) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithMasterTimeout(v time.Duration) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithS(v ...string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithTime(v string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithV(v bool) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithPretty() func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithHuman() func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithErrorTrace() func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithFilterPath(v ...string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithHeader(h map[string]string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatComponentTemplates) WithOpaqueID(s string) func(*CatComponentTemplatesRequest) {
	_ = "STUB: not implemented"
	return nil
}
