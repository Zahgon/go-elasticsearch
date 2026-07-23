package esapi

import (
	"context"
	"net/http"
)

func newExistsFunc(t Transport) Exists { _ = "STUB: not implemented"; return *new(Exists) }

type Exists func(index string, id string, o ...func(*ExistsRequest)) (*Response, error)

type ExistsRequest struct {
	Index      string
	DocumentID string

	Preference     string
	Realtime       *bool
	Refresh        *bool
	Routing        []string
	Source         []string
	SourceExcludes []string
	SourceIncludes []string
	StoredFields   []string
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

func (r ExistsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Exists) WithContext(v context.Context) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithPreference(v string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithRealtime(v bool) func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithRefresh(v bool) func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithRouting(v ...string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithSource(v ...string) func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithSourceExcludes(v ...string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithSourceIncludes(v ...string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithStoredFields(v ...string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithVersion(v int64) func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithVersionType(v string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithPretty() func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithHuman() func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithErrorTrace() func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }

func (f Exists) WithFilterPath(v ...string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithHeader(h map[string]string) func(*ExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Exists) WithOpaqueID(s string) func(*ExistsRequest) { _ = "STUB: not implemented"; return nil }
