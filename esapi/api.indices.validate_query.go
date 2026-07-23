package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIndicesValidateQueryFunc(t Transport) IndicesValidateQuery {
	_ = "STUB: not implemented"
	return *new(IndicesValidateQuery)
}

type IndicesValidateQuery func(o ...func(*IndicesValidateQueryRequest)) (*Response, error)

type IndicesValidateQueryRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices    *bool
	AllShards         *bool
	Analyzer          string
	AnalyzeWildcard   *bool
	DefaultOperator   string
	Df                string
	ExpandWildcards   []string
	Explain           *bool
	IgnoreUnavailable *bool
	Lenient           *bool
	Query             string
	Rewrite           *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesValidateQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesValidateQuery) WithContext(v context.Context) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithBody(v io.Reader) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithIndex(v ...string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithAllowNoIndices(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithAllShards(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithAnalyzer(v string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithAnalyzeWildcard(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithDefaultOperator(v string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithDf(v string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithExpandWildcards(v ...string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithExplain(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithIgnoreUnavailable(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithLenient(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithQuery(v string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithRewrite(v bool) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithPretty() func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithHuman() func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithErrorTrace() func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithFilterPath(v ...string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithHeader(h map[string]string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesValidateQuery) WithOpaqueID(s string) func(*IndicesValidateQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
