package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newDeleteByQueryFunc(t Transport) DeleteByQuery {
	_ = "STUB: not implemented"
	return *new(DeleteByQuery)
}

type DeleteByQuery func(index []string, body io.Reader, o ...func(*DeleteByQueryRequest)) (*Response, error)

type DeleteByQueryRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices      *bool
	Analyzer            string
	AnalyzeWildcard     *bool
	Conflicts           string
	DefaultOperator     string
	Df                  string
	ExpandWildcards     []string
	From                *int64
	IgnoreUnavailable   *bool
	Lenient             *bool
	MaxDocs             *int64
	Preference          string
	Query               string
	Refresh             *bool
	RequestCache        *bool
	RequestsPerSecond   *int
	Routing             []string
	Scroll              time.Duration
	ScrollSize          *int64
	SearchTimeout       time.Duration
	SearchType          string
	Slices              interface{}
	Sort                []string
	Stats               []string
	TerminateAfter      *int64
	Timeout             time.Duration
	Version             *bool
	WaitForActiveShards string
	WaitForCompletion   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r DeleteByQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f DeleteByQuery) WithContext(v context.Context) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithAllowNoIndices(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithAnalyzer(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithAnalyzeWildcard(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithConflicts(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithDefaultOperator(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithDf(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithExpandWildcards(v ...string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithFrom(v int64) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithIgnoreUnavailable(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithLenient(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithMaxDocs(v int64) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithPreference(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithQuery(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithRefresh(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithRequestCache(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithRequestsPerSecond(v int) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithRouting(v ...string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithScroll(v time.Duration) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithScrollSize(v int64) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithSearchTimeout(v time.Duration) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithSearchType(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithSlices(v interface{}) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithSort(v ...string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithStats(v ...string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithTerminateAfter(v int64) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithTimeout(v time.Duration) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithVersion(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithWaitForActiveShards(v string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithWaitForCompletion(v bool) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithPretty() func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithHuman() func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithErrorTrace() func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithFilterPath(v ...string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithHeader(h map[string]string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQuery) WithOpaqueID(s string) func(*DeleteByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
