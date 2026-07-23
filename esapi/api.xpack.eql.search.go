package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newEqlSearchFunc(t Transport) EqlSearch { _ = "STUB: not implemented"; return *new(EqlSearch) }

type EqlSearch func(index []string, body io.Reader, o ...func(*EqlSearchRequest)) (*Response, error)

type EqlSearchRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices              *bool
	AllowPartialSearchResults   *bool
	AllowPartialSequenceResults *bool
	CcsMinimizeRoundtrips       *bool
	ExpandWildcards             []string
	IgnoreUnavailable           *bool
	KeepAlive                   time.Duration
	KeepOnCompletion            *bool
	WaitForCompletionTimeout    time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EqlSearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EqlSearch) WithContext(v context.Context) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithAllowNoIndices(v bool) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithAllowPartialSearchResults(v bool) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithAllowPartialSequenceResults(v bool) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithCcsMinimizeRoundtrips(v bool) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithExpandWildcards(v ...string) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithIgnoreUnavailable(v bool) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithKeepAlive(v time.Duration) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithKeepOnCompletion(v bool) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithWaitForCompletionTimeout(v time.Duration) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithPretty() func(*EqlSearchRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlSearch) WithHuman() func(*EqlSearchRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlSearch) WithErrorTrace() func(*EqlSearchRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlSearch) WithFilterPath(v ...string) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithHeader(h map[string]string) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlSearch) WithOpaqueID(s string) func(*EqlSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
