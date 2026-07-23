package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newAsyncSearchSubmitFunc(t Transport) AsyncSearchSubmit {
	_ = "STUB: not implemented"
	return *new(AsyncSearchSubmit)
}

type AsyncSearchSubmit func(o ...func(*AsyncSearchSubmitRequest)) (*Response, error)

type AsyncSearchSubmitRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices             *bool
	AllowPartialSearchResults  *bool
	Analyzer                   string
	AnalyzeWildcard            *bool
	BatchedReduceSize          *int64
	CcsMinimizeRoundtrips      *bool
	DefaultOperator            string
	Df                         string
	DocvalueFields             []string
	ExpandWildcards            []string
	Explain                    *bool
	From                       *int
	IgnoreThrottled            *bool
	IgnoreUnavailable          *bool
	KeepAlive                  time.Duration
	KeepOnCompletion           *bool
	Lenient                    *bool
	MaxConcurrentShardRequests *int
	Preference                 string
	Query                      string
	RequestCache               *bool
	RestTotalHitsAsInt         *bool
	Routing                    []string
	SearchType                 string
	SeqNoPrimaryTerm           *bool
	Size                       *int
	Sort                       []string
	Source                     []string
	SourceExcludes             []string
	SourceIncludes             []string
	Stats                      []string
	StoredFields               []string
	SuggestField               string
	SuggestMode                string
	SuggestSize                *int64
	SuggestText                string
	TerminateAfter             *int64
	Timeout                    time.Duration
	TrackScores                *bool
	TrackTotalHits             interface{}
	TypedKeys                  *bool
	Version                    *bool
	WaitForCompletionTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r AsyncSearchSubmitRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AsyncSearchSubmit) WithContext(v context.Context) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithBody(v io.Reader) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithIndex(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithAllowNoIndices(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithAllowPartialSearchResults(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithAnalyzer(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithAnalyzeWildcard(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithBatchedReduceSize(v int64) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithCcsMinimizeRoundtrips(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithDefaultOperator(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithDf(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithDocvalueFields(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithExpandWildcards(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithExplain(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithFrom(v int) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithIgnoreThrottled(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithIgnoreUnavailable(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithKeepAlive(v time.Duration) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithKeepOnCompletion(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithLenient(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithMaxConcurrentShardRequests(v int) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithPreference(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithQuery(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithRequestCache(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithRestTotalHitsAsInt(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithRouting(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSearchType(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSeqNoPrimaryTerm(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSize(v int) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSort(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSource(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSourceExcludes(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSourceIncludes(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithStats(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithStoredFields(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSuggestField(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSuggestMode(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSuggestSize(v int64) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithSuggestText(v string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithTerminateAfter(v int64) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithTimeout(v time.Duration) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithTrackScores(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithTrackTotalHits(v interface{}) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithTypedKeys(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithVersion(v bool) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithWaitForCompletionTimeout(v time.Duration) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithPretty() func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithHuman() func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithErrorTrace() func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithFilterPath(v ...string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithHeader(h map[string]string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchSubmit) WithOpaqueID(s string) func(*AsyncSearchSubmitRequest) {
	_ = "STUB: not implemented"
	return nil
}
