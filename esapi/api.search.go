package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSearchFunc(t Transport) Search { _ = "STUB: not implemented"; return *new(Search) }

type Search func(o ...func(*SearchRequest)) (*Response, error)

type SearchRequest struct {
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
	ForceSyntheticSource       *bool
	From                       *int
	IgnoreThrottled            *bool
	IgnoreUnavailable          *bool
	IncludeNamedQueriesScore   *bool
	Lenient                    *bool
	MaxConcurrentShardRequests *int
	Preference                 string
	PreFilterShardSize         *int64
	Query                      string
	RequestCache               *bool
	RestTotalHitsAsInt         *bool
	Routing                    []string
	Scroll                     time.Duration
	SearchType                 string
	SeqNoPrimaryTerm           *bool
	Size                       *int
	Sort                       []string
	Source                     []string
	SourceExcludeVectors       *bool
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

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Search) WithContext(v context.Context) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithBody(v io.Reader) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithIndex(v ...string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithAllowNoIndices(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithAllowPartialSearchResults(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithAnalyzer(v string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithAnalyzeWildcard(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithBatchedReduceSize(v int64) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithCcsMinimizeRoundtrips(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithDefaultOperator(v string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithDf(v string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithDocvalueFields(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithExpandWildcards(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithExplain(v bool) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithForceSyntheticSource(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithFrom(v int) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithIgnoreThrottled(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithIgnoreUnavailable(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithIncludeNamedQueriesScore(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithLenient(v bool) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithMaxConcurrentShardRequests(v int) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithPreference(v string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithPreFilterShardSize(v int64) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithQuery(v string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithRequestCache(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithRestTotalHitsAsInt(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithRouting(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithScroll(v time.Duration) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSearchType(v string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSeqNoPrimaryTerm(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSize(v int) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithSort(v ...string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithSource(v ...string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithSourceExcludeVectors(v bool) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSourceExcludes(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSourceIncludes(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithStats(v ...string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithStoredFields(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSuggestField(v string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSuggestMode(v string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSuggestSize(v int64) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithSuggestText(v string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithTerminateAfter(v int64) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithTimeout(v time.Duration) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithTrackScores(v bool) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithTrackTotalHits(v interface{}) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithTypedKeys(v bool) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithVersion(v bool) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithPretty() func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithHuman() func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithErrorTrace() func(*SearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Search) WithFilterPath(v ...string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithHeader(h map[string]string) func(*SearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Search) WithOpaqueID(s string) func(*SearchRequest) { _ = "STUB: not implemented"; return nil }
