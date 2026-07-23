package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMgetFunc(t Transport) Mget { _ = "STUB: not implemented"; return *new(Mget) }

type Mget func(body io.Reader, o ...func(*MgetRequest)) (*Response, error)

type MgetRequest struct {
	Index string

	Body io.Reader

	ForceSyntheticSource *bool
	Preference           string
	Realtime             *bool
	Refresh              *bool
	Routing              []string
	Source               []string
	SourceExcludes       []string
	SourceIncludes       []string
	StoredFields         []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MgetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Mget) WithContext(v context.Context) func(*MgetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mget) WithIndex(v string) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithForceSyntheticSource(v bool) func(*MgetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mget) WithPreference(v string) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithRealtime(v bool) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithRefresh(v bool) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithRouting(v ...string) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithSource(v ...string) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithSourceExcludes(v ...string) func(*MgetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mget) WithSourceIncludes(v ...string) func(*MgetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mget) WithStoredFields(v ...string) func(*MgetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mget) WithPretty() func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithHuman() func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithErrorTrace() func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithFilterPath(v ...string) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }

func (f Mget) WithHeader(h map[string]string) func(*MgetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Mget) WithOpaqueID(s string) func(*MgetRequest) { _ = "STUB: not implemented"; return nil }
