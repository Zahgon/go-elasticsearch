package esapi

import (
	"context"
	"io"
	"net/http"
)

func newTermvectorsFunc(t Transport) Termvectors {
	_ = "STUB: not implemented"
	return *new(Termvectors)
}

type Termvectors func(index string, o ...func(*TermvectorsRequest)) (*Response, error)

type TermvectorsRequest struct {
	Index      string
	DocumentID string

	Body io.Reader

	Fields          []string
	FieldStatistics *bool
	Offsets         *bool
	Payloads        *bool
	Positions       *bool
	Preference      string
	Realtime        *bool
	Routing         []string
	TermStatistics  *bool
	Version         *int64
	VersionType     string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TermvectorsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Termvectors) WithContext(v context.Context) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithBody(v io.Reader) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithDocumentID(v string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithFields(v ...string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithFieldStatistics(v bool) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithOffsets(v bool) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithPayloads(v bool) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithPositions(v bool) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithPreference(v string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithRealtime(v bool) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithRouting(v ...string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithTermStatistics(v bool) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithVersion(v int64) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithVersionType(v string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithPretty() func(*TermvectorsRequest) { _ = "STUB: not implemented"; return nil }

func (f Termvectors) WithHuman() func(*TermvectorsRequest) { _ = "STUB: not implemented"; return nil }

func (f Termvectors) WithErrorTrace() func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithFilterPath(v ...string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithHeader(h map[string]string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Termvectors) WithOpaqueID(s string) func(*TermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}
