package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatIndicesFunc(t Transport) CatIndices { _ = "STUB: not implemented"; return *new(CatIndices) }

type CatIndices func(o ...func(*CatIndicesRequest)) (*Response, error)

type CatIndicesRequest struct {
	Index []string

	Bytes                   string
	ExpandWildcards         []string
	Format                  string
	H                       []string
	Health                  string
	Help                    *bool
	IncludeUnloadedSegments *bool
	MasterTimeout           time.Duration
	Pri                     *bool
	S                       []string
	Time                    string
	V                       *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatIndicesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatIndices) WithContext(v context.Context) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithIndex(v ...string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithBytes(v string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithExpandWildcards(v ...string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithFormat(v string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithH(v ...string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithHealth(v string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithHelp(v bool) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithIncludeUnloadedSegments(v bool) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithMasterTimeout(v time.Duration) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithPri(v bool) func(*CatIndicesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatIndices) WithS(v ...string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithTime(v string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithV(v bool) func(*CatIndicesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatIndices) WithPretty() func(*CatIndicesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatIndices) WithHuman() func(*CatIndicesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatIndices) WithErrorTrace() func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithFilterPath(v ...string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithHeader(h map[string]string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatIndices) WithOpaqueID(s string) func(*CatIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}
