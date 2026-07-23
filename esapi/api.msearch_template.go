package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMsearchTemplateFunc(t Transport) MsearchTemplate {
	_ = "STUB: not implemented"
	return *new(MsearchTemplate)
}

type MsearchTemplate func(body io.Reader, o ...func(*MsearchTemplateRequest)) (*Response, error)

type MsearchTemplateRequest struct {
	Index []string

	Body io.Reader

	CcsMinimizeRoundtrips *bool
	MaxConcurrentSearches *int64
	ProjectRouting        string
	RestTotalHitsAsInt    *bool
	SearchType            string
	TypedKeys             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MsearchTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MsearchTemplate) WithContext(v context.Context) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithIndex(v ...string) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithCcsMinimizeRoundtrips(v bool) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithMaxConcurrentSearches(v int64) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithProjectRouting(v string) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithRestTotalHitsAsInt(v bool) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithSearchType(v string) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithTypedKeys(v bool) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithPretty() func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithHuman() func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithErrorTrace() func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithFilterPath(v ...string) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithHeader(h map[string]string) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MsearchTemplate) WithOpaqueID(s string) func(*MsearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
