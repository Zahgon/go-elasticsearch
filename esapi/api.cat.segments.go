package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatSegmentsFunc(t Transport) CatSegments {
	_ = "STUB: not implemented"
	return *new(CatSegments)
}

type CatSegments func(o ...func(*CatSegmentsRequest)) (*Response, error)

type CatSegmentsRequest struct {
	Index []string

	AllowClosed       *bool
	AllowNoIndices    *bool
	Bytes             string
	ExpandWildcards   []string
	Format            string
	H                 []string
	Help              *bool
	IgnoreThrottled   *bool
	IgnoreUnavailable *bool
	Local             *bool
	MasterTimeout     time.Duration
	S                 []string
	Time              string
	V                 *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatSegmentsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatSegments) WithContext(v context.Context) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithIndex(v ...string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithAllowClosed(v bool) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithAllowNoIndices(v bool) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithBytes(v string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithExpandWildcards(v ...string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithFormat(v string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithH(v ...string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithHelp(v bool) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithIgnoreThrottled(v bool) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithIgnoreUnavailable(v bool) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithLocal(v bool) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithMasterTimeout(v time.Duration) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithS(v ...string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithTime(v string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithV(v bool) func(*CatSegmentsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatSegments) WithPretty() func(*CatSegmentsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatSegments) WithHuman() func(*CatSegmentsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatSegments) WithErrorTrace() func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithFilterPath(v ...string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithHeader(h map[string]string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSegments) WithOpaqueID(s string) func(*CatSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}
