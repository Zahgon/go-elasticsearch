package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newUpdateByQueryFunc(t Transport) UpdateByQuery {
	_ = "STUB: not implemented"
	return *new(UpdateByQuery)
}

type UpdateByQuery func(index []string, o ...func(*UpdateByQueryRequest)) (*Response, error)

type UpdateByQueryRequest struct {
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
	Pipeline            string
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
	VersionType         *bool
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

func (r UpdateByQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f UpdateByQuery) WithContext(v context.Context) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithBody(v io.Reader) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithAllowNoIndices(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithAnalyzer(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithAnalyzeWildcard(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithConflicts(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithDefaultOperator(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithDf(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithExpandWildcards(v ...string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithFrom(v int64) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithIgnoreUnavailable(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithLenient(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithMaxDocs(v int64) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithPipeline(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithPreference(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithQuery(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithRefresh(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithRequestCache(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithRequestsPerSecond(v int) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithRouting(v ...string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithScroll(v time.Duration) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithScrollSize(v int64) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithSearchTimeout(v time.Duration) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithSearchType(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithSlices(v interface{}) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithSort(v ...string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithStats(v ...string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithTerminateAfter(v int64) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithTimeout(v time.Duration) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithVersion(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithVersionType(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithWaitForActiveShards(v string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithWaitForCompletion(v bool) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithPretty() func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithHuman() func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithErrorTrace() func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithFilterPath(v ...string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithHeader(h map[string]string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQuery) WithOpaqueID(s string) func(*UpdateByQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
