package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatNodesFunc(t Transport) CatNodes { _ = "STUB: not implemented"; return *new(CatNodes) }

type CatNodes func(o ...func(*CatNodesRequest)) (*Response, error)

type CatNodesRequest struct {
	Bytes                   string
	Format                  string
	FullID                  *bool
	H                       []string
	Help                    *bool
	IncludeUnloadedSegments *bool
	MasterTimeout           time.Duration
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

func (r CatNodesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatNodes) WithContext(v context.Context) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodes) WithBytes(v string) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithFormat(v string) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodes) WithFullID(v bool) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithH(v ...string) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithHelp(v bool) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithIncludeUnloadedSegments(v bool) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodes) WithMasterTimeout(v time.Duration) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodes) WithS(v ...string) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithTime(v string) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithV(v bool) func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithPretty() func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithHuman() func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithErrorTrace() func(*CatNodesRequest) { _ = "STUB: not implemented"; return nil }

func (f CatNodes) WithFilterPath(v ...string) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodes) WithHeader(h map[string]string) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatNodes) WithOpaqueID(s string) func(*CatNodesRequest) {
	_ = "STUB: not implemented"
	return nil
}
