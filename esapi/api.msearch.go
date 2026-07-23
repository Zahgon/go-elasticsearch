package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMsearchFunc(t Transport) Msearch { _ = "STUB: not implemented"; return *new(Msearch) }

type Msearch func(body io.Reader, o ...func(*MsearchRequest)) (*Response, error)

type MsearchRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices             *bool
	CcsMinimizeRoundtrips      *bool
	ExpandWildcards            []string
	IgnoreThrottled            *bool
	IgnoreUnavailable          *bool
	IncludeNamedQueriesScore   *bool
	MaxConcurrentSearches      *int
	MaxConcurrentShardRequests *int
	PreFilterShardSize         *int64
	ProjectRouting             string
	RestTotalHitsAsInt         *bool
	Routing                    []string
	SearchType                 string
	TypedKeys                  *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MsearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Msearch) WithContext(v context.Context) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithIndex(v ...string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithAllowNoIndices(v bool) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithCcsMinimizeRoundtrips(v bool) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithExpandWildcards(v ...string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithIgnoreThrottled(v bool) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithIgnoreUnavailable(v bool) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithIncludeNamedQueriesScore(v bool) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithMaxConcurrentSearches(v int) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithMaxConcurrentShardRequests(v int) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithPreFilterShardSize(v int64) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithProjectRouting(v string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithRestTotalHitsAsInt(v bool) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithRouting(v ...string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithSearchType(v string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithTypedKeys(v bool) func(*MsearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Msearch) WithPretty() func(*MsearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Msearch) WithHuman() func(*MsearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Msearch) WithErrorTrace() func(*MsearchRequest) { _ = "STUB: not implemented"; return nil }

func (f Msearch) WithFilterPath(v ...string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithHeader(h map[string]string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Msearch) WithOpaqueID(s string) func(*MsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
