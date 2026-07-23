package esapi

import (
	"context"
	"net/http"
)

func newExistsSourceFunc(t Transport) ExistsSource {
	_ = "STUB: not implemented"
	return *new(ExistsSource)
}

type ExistsSource func(index string, id string, o ...func(*ExistsSourceRequest)) (*Response, error)

type ExistsSourceRequest struct {
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

func (r ExistsSourceRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ExistsSource) WithContext(v context.Context) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithPreference(v string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithRealtime(v bool) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithRefresh(v bool) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithRouting(v ...string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithSource(v ...string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithSourceExcludes(v ...string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithSourceIncludes(v ...string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithVersion(v int64) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithVersionType(v string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithPretty() func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithHuman() func(*ExistsSourceRequest) { _ = "STUB: not implemented"; return nil }

func (f ExistsSource) WithErrorTrace() func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithFilterPath(v ...string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithHeader(h map[string]string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ExistsSource) WithOpaqueID(s string) func(*ExistsSourceRequest) {
	_ = "STUB: not implemented"
	return nil
}
