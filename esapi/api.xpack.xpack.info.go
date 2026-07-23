package esapi

import (
	"context"
	"net/http"
)

func newXPackInfoFunc(t Transport) XPackInfo { _ = "STUB: not implemented"; return *new(XPackInfo) }

type XPackInfo func(o ...func(*XPackInfoRequest)) (*Response, error)

type XPackInfoRequest struct {
	AcceptEnterprise *bool
	Categories       []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r XPackInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f XPackInfo) WithContext(v context.Context) func(*XPackInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackInfo) WithAcceptEnterprise(v bool) func(*XPackInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackInfo) WithCategories(v ...string) func(*XPackInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackInfo) WithPretty() func(*XPackInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f XPackInfo) WithHuman() func(*XPackInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f XPackInfo) WithErrorTrace() func(*XPackInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f XPackInfo) WithFilterPath(v ...string) func(*XPackInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackInfo) WithHeader(h map[string]string) func(*XPackInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackInfo) WithOpaqueID(s string) func(*XPackInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}
