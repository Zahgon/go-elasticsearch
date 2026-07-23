package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndexFunc(t Transport) Index { _ = "STUB: not implemented"; return *new(Index) }

type Index func(index string, body io.Reader, o ...func(*IndexRequest)) (*Response, error)

type IndexRequest struct {
	Index      string
	DocumentID string

	Body io.Reader

	IfPrimaryTerm        *int64
	IfSeqNo              *int64
	IncludeSourceOnError *bool
	OpType               string
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

func (r IndexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Index) WithContext(v context.Context) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithDocumentID(v string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithIfPrimaryTerm(v int64) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithIfSeqNo(v int64) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithIncludeSourceOnError(v bool) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithOpType(v string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithPipeline(v string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithRefresh(v string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithRequireAlias(v bool) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithRequireDataStream(v bool) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithRouting(v ...string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithTimeout(v time.Duration) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithVersion(v int64) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithVersionType(v string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithWaitForActiveShards(v string) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithPretty() func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithHuman() func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithErrorTrace() func(*IndexRequest) { _ = "STUB: not implemented"; return nil }

func (f Index) WithFilterPath(v ...string) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithHeader(h map[string]string) func(*IndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Index) WithOpaqueID(s string) func(*IndexRequest) { _ = "STUB: not implemented"; return nil }
