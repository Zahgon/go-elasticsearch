package esapi

import (
	"context"
	"net/http"
)

func newGetSourceFunc(t Transport) GetSource { _ = "STUB: not implemented"; return *new(GetSource) }

type GetSource func(index string, id string, o ...func(*GetSourceRequest)) (*Response, error)

type GetSourceRequest struct {
	Index      string
	DocumentID string

	Preference     string
	Realtime       *bool
	Refresh        *bool
	Routing        []string
	Source         []string
	SourceExcludes []string
	SourceIncludes []string
	Version        *int64
	VersionType    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r GetSourceRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f GetSource) WithContext(v context.Context) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithPreference(v string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithRealtime(v bool) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithRefresh(v bool) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithRouting(v ...string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithSource(v ...string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithSourceExcludes(v ...string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithSourceIncludes(v ...string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithVersion(v int64) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithVersionType(v string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithPretty() func(*GetSourceRequest) { _ = "STUB: not implemented"; return nil }

func (f GetSource) WithHuman() func(*GetSourceRequest) { _ = "STUB: not implemented"; return nil }

func (f GetSource) WithErrorTrace() func(*GetSourceRequest) { _ = "STUB: not implemented"; return nil }

func (f GetSource) WithFilterPath(v ...string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithHeader(h map[string]string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GetSource) WithOpaqueID(s string) func(*GetSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}
