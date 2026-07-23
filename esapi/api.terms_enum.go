package esapi

import (
	"context"
	"io"
	"net/http"
)

func newTermsEnumFunc(t Transport) TermsEnum { _ = "STUB: not implemented"; return *new(TermsEnum) }

type TermsEnum func(index []string, body io.Reader, o ...func(*TermsEnumRequest)) (*Response, error)

type TermsEnumRequest struct {
	Index []string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TermsEnumRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TermsEnum) WithContext(v context.Context) func(*TermsEnumRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TermsEnum) WithPretty() func(*TermsEnumRequest) { _ = "STUB: not implemented"; return nil }

func (f TermsEnum) WithHuman() func(*TermsEnumRequest) { _ = "STUB: not implemented"; return nil }

func (f TermsEnum) WithErrorTrace() func(*TermsEnumRequest) { _ = "STUB: not implemented"; return nil }

func (f TermsEnum) WithFilterPath(v ...string) func(*TermsEnumRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TermsEnum) WithHeader(h map[string]string) func(*TermsEnumRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TermsEnum) WithOpaqueID(s string) func(*TermsEnumRequest) {
	_ = "STUB: not implemented"
	return nil
}
