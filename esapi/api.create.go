package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newCreateFunc(t Transport) Create { _ = "STUB: not implemented"; return *new(Create) }

type Create func(index string, id string, body io.Reader, o ...func(*CreateRequest)) (*Response, error)

type CreateRequest struct {
	Index      string
	DocumentID string

	Body io.Reader

	IncludeSourceOnError *bool
	Pipeline             string
	Refresh              string
	RequireAlias         *bool
	RequireDataStream    *bool
	Routing              []string
	Timeout              time.Duration
	Version              *int64
	VersionType          string
	WaitForActiveShards  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CreateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Create) WithContext(v context.Context) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithIncludeSourceOnError(v bool) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithPipeline(v string) func(*CreateRequest) { _ = "STUB: not implemented"; return nil }

func (f Create) WithRefresh(v string) func(*CreateRequest) { _ = "STUB: not implemented"; return nil }

func (f Create) WithRequireAlias(v bool) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithRequireDataStream(v bool) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithRouting(v ...string) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithTimeout(v time.Duration) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithVersion(v int64) func(*CreateRequest) { _ = "STUB: not implemented"; return nil }

func (f Create) WithVersionType(v string) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithWaitForActiveShards(v string) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithPretty() func(*CreateRequest) { _ = "STUB: not implemented"; return nil }

func (f Create) WithHuman() func(*CreateRequest) { _ = "STUB: not implemented"; return nil }

func (f Create) WithErrorTrace() func(*CreateRequest) { _ = "STUB: not implemented"; return nil }

func (f Create) WithFilterPath(v ...string) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithHeader(h map[string]string) func(*CreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Create) WithOpaqueID(s string) func(*CreateRequest) { _ = "STUB: not implemented"; return nil }
