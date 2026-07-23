package esapi

import (
	"context"
	"net/http"
)

func newGetFunc(t Transport) Get { _ = "STUB: not implemented"; return *new(Get) }

type Get func(index string, id string, o ...func(*GetRequest)) (*Response, error)

type GetRequest struct {
	Index      string
	DocumentID string

	ForceSyntheticSource *bool
	Preference           string
	Realtime             *bool
	Refresh              *bool
	Routing              []string
	Source               []string
	SourceExcludeVectors *bool
	SourceExcludes       []string
	SourceIncludes       []string
	StoredFields         []string
	Version              *int64
	VersionType          string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r GetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Get) WithContext(v context.Context) func(*GetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Get) WithForceSyntheticSource(v bool) func(*GetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Get) WithPreference(v string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithRealtime(v bool) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithRefresh(v bool) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithRouting(v ...string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithSource(v ...string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithSourceExcludeVectors(v bool) func(*GetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Get) WithSourceExcludes(v ...string) func(*GetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Get) WithSourceIncludes(v ...string) func(*GetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Get) WithStoredFields(v ...string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithVersion(v int64) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithVersionType(v string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithPretty() func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithHuman() func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithErrorTrace() func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithFilterPath(v ...string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }

func (f Get) WithHeader(h map[string]string) func(*GetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Get) WithOpaqueID(s string) func(*GetRequest) { _ = "STUB: not implemented"; return nil }
