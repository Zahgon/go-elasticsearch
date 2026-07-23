package esapi

import (
	"context"
	"io"
	"net/http"
)

func newRankEvalFunc(t Transport) RankEval { _ = "STUB: not implemented"; return *new(RankEval) }

type RankEval func(body io.Reader, o ...func(*RankEvalRequest)) (*Response, error)

type RankEvalRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	SearchType        string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RankEvalRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RankEval) WithContext(v context.Context) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithIndex(v ...string) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithAllowNoIndices(v bool) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithExpandWildcards(v ...string) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithIgnoreUnavailable(v bool) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithSearchType(v string) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithPretty() func(*RankEvalRequest) { _ = "STUB: not implemented"; return nil }

func (f RankEval) WithHuman() func(*RankEvalRequest) { _ = "STUB: not implemented"; return nil }

func (f RankEval) WithErrorTrace() func(*RankEvalRequest) { _ = "STUB: not implemented"; return nil }

func (f RankEval) WithFilterPath(v ...string) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithHeader(h map[string]string) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RankEval) WithOpaqueID(s string) func(*RankEvalRequest) {
	_ = "STUB: not implemented"
	return nil
}
