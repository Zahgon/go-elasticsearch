package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSearchTemplateFunc(t Transport) SearchTemplate {
	_ = "STUB: not implemented"
	return *new(SearchTemplate)
}

type SearchTemplate func(body io.Reader, o ...func(*SearchTemplateRequest)) (*Response, error)

type SearchTemplateRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices        *bool
	CcsMinimizeRoundtrips *bool
	ExpandWildcards       []string
	Explain               *bool
	IgnoreThrottled       *bool
	IgnoreUnavailable     *bool
	Preference            string
	Profile               *bool
	RestTotalHitsAsInt    *bool
	Routing               []string
	Scroll                time.Duration
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

func (r SearchTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchTemplate) WithContext(v context.Context) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithIndex(v ...string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithAllowNoIndices(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithCcsMinimizeRoundtrips(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithExpandWildcards(v ...string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithExplain(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithIgnoreThrottled(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithIgnoreUnavailable(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithPreference(v string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithProfile(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithRestTotalHitsAsInt(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithRouting(v ...string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithScroll(v time.Duration) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithSearchType(v string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithTypedKeys(v bool) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithPretty() func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithHuman() func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithErrorTrace() func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithFilterPath(v ...string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithHeader(h map[string]string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchTemplate) WithOpaqueID(s string) func(*SearchTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
