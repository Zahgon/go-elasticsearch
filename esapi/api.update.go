package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newUpdateFunc(t Transport) Update { _ = "STUB: not implemented"; return *new(Update) }

type Update func(index string, id string, body io.Reader, o ...func(*UpdateRequest)) (*Response, error)

type UpdateRequest struct {
	Index      string
	DocumentID string

	Body io.Reader

	IfPrimaryTerm        *int64
	IfSeqNo              *int64
	IncludeSourceOnError *bool
	Lang                 string
	Refresh              string
	RequireAlias         *bool
	RetryOnConflict      *int
	Routing              []string
	Source               []string
	SourceExcludes       []string
	SourceIncludes       []string
	Timeout              time.Duration
	WaitForActiveShards  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r UpdateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Update) WithContext(v context.Context) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithIfPrimaryTerm(v int64) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithIfSeqNo(v int64) func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithIncludeSourceOnError(v bool) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithLang(v string) func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithRefresh(v string) func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithRequireAlias(v bool) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithRetryOnConflict(v int) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithRouting(v ...string) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithSource(v ...string) func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithSourceExcludes(v ...string) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithSourceIncludes(v ...string) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithTimeout(v time.Duration) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithWaitForActiveShards(v string) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithPretty() func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithHuman() func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithErrorTrace() func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }

func (f Update) WithFilterPath(v ...string) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithHeader(h map[string]string) func(*UpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Update) WithOpaqueID(s string) func(*UpdateRequest) { _ = "STUB: not implemented"; return nil }
