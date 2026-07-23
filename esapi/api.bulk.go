package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newBulkFunc(t Transport) Bulk { _ = "STUB: not implemented"; return *new(Bulk) }

type Bulk func(body io.Reader, o ...func(*BulkRequest)) (*Response, error)

type BulkRequest struct {
	Index string

	Body io.Reader

	IncludeSourceOnError  *bool
	ListExecutedPipelines *bool
	Pipeline              string
	Refresh               string
	RequireAlias          *bool
	RequireDataStream     *bool
	Routing               []string
	Source                []string
	SourceExcludes        []string
	SourceIncludes        []string
	Timeout               time.Duration
	WaitForActiveShards   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r BulkRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Bulk) WithContext(v context.Context) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithIndex(v string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithIncludeSourceOnError(v bool) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithListExecutedPipelines(v bool) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithPipeline(v string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithRefresh(v string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithRequireAlias(v bool) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithRequireDataStream(v bool) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithRouting(v ...string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithSource(v ...string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithSourceExcludes(v ...string) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithSourceIncludes(v ...string) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithTimeout(v time.Duration) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithWaitForActiveShards(v string) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithPretty() func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithHuman() func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithErrorTrace() func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithFilterPath(v ...string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }

func (f Bulk) WithHeader(h map[string]string) func(*BulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Bulk) WithOpaqueID(s string) func(*BulkRequest) { _ = "STUB: not implemented"; return nil }
