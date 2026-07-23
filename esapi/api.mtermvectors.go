package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMtermvectorsFunc(t Transport) Mtermvectors {
	_ = "STUB: not implemented"
	return *new(Mtermvectors)
}

type Mtermvectors func(o ...func(*MtermvectorsRequest)) (*Response, error)

type MtermvectorsRequest struct {
	Index string

	Body io.Reader

	Fields          []string
	FieldStatistics *bool
	Ids             []string
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

func (r MtermvectorsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Mtermvectors) WithContext(v context.Context) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithBody(v io.Reader) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithIndex(v string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithFields(v ...string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithFieldStatistics(v bool) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithIds(v ...string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithOffsets(v bool) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithPayloads(v bool) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithPositions(v bool) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithPreference(v string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithRealtime(v bool) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithRouting(v ...string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithTermStatistics(v bool) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithVersion(v int64) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithVersionType(v string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithPretty() func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithHuman() func(*MtermvectorsRequest) { _ = "STUB: not implemented"; return nil }

func (f Mtermvectors) WithErrorTrace() func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithFilterPath(v ...string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithHeader(h map[string]string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mtermvectors) WithOpaqueID(s string) func(*MtermvectorsRequest) {
	_ = "STUB: not implemented"
	return nil
}
