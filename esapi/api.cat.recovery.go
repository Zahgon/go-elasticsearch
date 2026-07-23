package esapi

import (
	"context"
	"net/http"
)

func newCatRecoveryFunc(t Transport) CatRecovery {
	_ = "STUB: not implemented"
	return *new(CatRecovery)
}

type CatRecovery func(o ...func(*CatRecoveryRequest)) (*Response, error)

type CatRecoveryRequest struct {
	Index []string

	ActiveOnly *bool
	Bytes      string
	Detailed   *bool
	Format     string
	H          []string
	Help       *bool
	S          []string
	Time       string
	V          *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatRecoveryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatRecovery) WithContext(v context.Context) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithIndex(v ...string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithActiveOnly(v bool) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithBytes(v string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithDetailed(v bool) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithFormat(v string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithH(v ...string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithHelp(v bool) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithS(v ...string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithTime(v string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithV(v bool) func(*CatRecoveryRequest) { _ = "STUB: not implemented"; return nil }

func (f CatRecovery) WithPretty() func(*CatRecoveryRequest) { _ = "STUB: not implemented"; return nil }

func (f CatRecovery) WithHuman() func(*CatRecoveryRequest) { _ = "STUB: not implemented"; return nil }

func (f CatRecovery) WithErrorTrace() func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithFilterPath(v ...string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithHeader(h map[string]string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRecovery) WithOpaqueID(s string) func(*CatRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}
