package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatAliasesFunc(t Transport) CatAliases { _ = "STUB: not implemented"; return *new(CatAliases) }

type CatAliases func(o ...func(*CatAliasesRequest)) (*Response, error)

type CatAliasesRequest struct {
	Name []string

	Bytes           string
	ExpandWildcards []string
	Format          string
	H               []string
	Help            *bool
	MasterTimeout   time.Duration
	S               []string
	Time            string
	V               *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatAliasesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatAliases) WithContext(v context.Context) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithName(v ...string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithBytes(v string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithExpandWildcards(v ...string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithFormat(v string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithH(v ...string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithHelp(v bool) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithMasterTimeout(v time.Duration) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithS(v ...string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithTime(v string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithV(v bool) func(*CatAliasesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatAliases) WithPretty() func(*CatAliasesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatAliases) WithHuman() func(*CatAliasesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatAliases) WithErrorTrace() func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithFilterPath(v ...string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithHeader(h map[string]string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAliases) WithOpaqueID(s string) func(*CatAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}
