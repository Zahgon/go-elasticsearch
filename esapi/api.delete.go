package esapi

import (
	"context"
	"net/http"
	"time"
)

func newDeleteFunc(t Transport) Delete { _ = "STUB: not implemented"; return *new(Delete) }

type Delete func(index string, id string, o ...func(*DeleteRequest)) (*Response, error)

type DeleteRequest struct {
	Index      string
	DocumentID string

	IfPrimaryTerm       *int64
	IfSeqNo             *int64
	Refresh             string
	Routing             []string
	Timeout             time.Duration
	Version             *int64
	VersionType         string
	WaitForActiveShards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r DeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Delete) WithContext(v context.Context) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithIfPrimaryTerm(v int64) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithIfSeqNo(v int64) func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f Delete) WithRefresh(v string) func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f Delete) WithRouting(v ...string) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithTimeout(v time.Duration) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithVersion(v int64) func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f Delete) WithVersionType(v string) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithWaitForActiveShards(v string) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithPretty() func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f Delete) WithHuman() func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f Delete) WithErrorTrace() func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f Delete) WithFilterPath(v ...string) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithHeader(h map[string]string) func(*DeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Delete) WithOpaqueID(s string) func(*DeleteRequest) { _ = "STUB: not implemented"; return nil }
