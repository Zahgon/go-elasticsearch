package esapi

import (
	"context"
	"io"
	"net/http"
)

func newCountFunc(t Transport) Count { _ = "STUB: not implemented"; return *new(Count) }

type Count func(o ...func(*CountRequest)) (*Response, error)

type CountRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices    *bool
	Analyzer          string
	AnalyzeWildcard   *bool
	DefaultOperator   string
	Df                string
	ExpandWildcards   []string
	IgnoreThrottled   *bool
	IgnoreUnavailable *bool
	Lenient           *bool
	MinScore          interface{}
	Preference        string
	Query             string
	Routing           []string
	TerminateAfter    *int64

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CountRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Count) WithContext(v context.Context) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithBody(v io.Reader) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithIndex(v ...string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithAllowNoIndices(v bool) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithAnalyzer(v string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithAnalyzeWildcard(v bool) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithDefaultOperator(v string) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithDf(v string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithExpandWildcards(v ...string) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithIgnoreThrottled(v bool) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithIgnoreUnavailable(v bool) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithLenient(v bool) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithMinScore(v interface{}) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithPreference(v string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithQuery(v string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithRouting(v ...string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithTerminateAfter(v int64) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithPretty() func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithHuman() func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithErrorTrace() func(*CountRequest) { _ = "STUB: not implemented"; return nil }

func (f Count) WithFilterPath(v ...string) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithHeader(h map[string]string) func(*CountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Count) WithOpaqueID(s string) func(*CountRequest) { _ = "STUB: not implemented"; return nil }
