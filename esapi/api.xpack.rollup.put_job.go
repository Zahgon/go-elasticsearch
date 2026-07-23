package esapi

import (
	"context"
	"io"
	"net/http"
)

func newRollupPutJobFunc(t Transport) RollupPutJob {
	_ = "STUB: not implemented"
	return *new(RollupPutJob)
}

type RollupPutJob func(id string, body io.Reader, o ...func(*RollupPutJobRequest)) (*Response, error)

type RollupPutJobRequest struct {
	JobID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupPutJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupPutJob) WithContext(v context.Context) func(*RollupPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupPutJob) WithPretty() func(*RollupPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupPutJob) WithHuman() func(*RollupPutJobRequest) { _ = "STUB: not implemented"; return nil }

func (f RollupPutJob) WithErrorTrace() func(*RollupPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupPutJob) WithFilterPath(v ...string) func(*RollupPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupPutJob) WithHeader(h map[string]string) func(*RollupPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupPutJob) WithOpaqueID(s string) func(*RollupPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
