package esapi

import (
	"context"
	"io"
	"net/http"
)

func newExplainFunc(t Transport) Explain { _ = "STUB: not implemented"; return *new(Explain) }

type Explain func(index string, id string, o ...func(*ExplainRequest)) (*Response, error)

type ExplainRequest struct {
	Index      string
	DocumentID string

	Body io.Reader

	Analyzer        string
	AnalyzeWildcard *bool
	DefaultOperator string
	Df              string
	Lenient         *bool
	Preference      string
	Query           string
	Routing         []string
	Source          []string
	SourceExcludes  []string
	SourceIncludes  []string
	StoredFields    []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ExplainRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Explain) WithContext(v context.Context) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithBody(v io.Reader) func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithAnalyzer(v string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithAnalyzeWildcard(v bool) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithDefaultOperator(v string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithDf(v string) func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithLenient(v bool) func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithPreference(v string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithQuery(v string) func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithRouting(v ...string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithSource(v ...string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithSourceExcludes(v ...string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithSourceIncludes(v ...string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithStoredFields(v ...string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithPretty() func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithHuman() func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithErrorTrace() func(*ExplainRequest) { _ = "STUB: not implemented"; return nil }

func (f Explain) WithFilterPath(v ...string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithHeader(h map[string]string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Explain) WithOpaqueID(s string) func(*ExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}
