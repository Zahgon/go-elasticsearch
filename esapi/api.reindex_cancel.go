package esapi

import (
	"context"
	"net/http"
)

func newReindexCancelFunc(t Transport) ReindexCancel {
	_ = "STUB: not implemented"
	return *new(ReindexCancel)
}

type ReindexCancel func(task_id string, o ...func(*ReindexCancelRequest)) (*Response, error)

type ReindexCancelRequest struct {
	TaskID string

	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ReindexCancelRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ReindexCancel) WithContext(v context.Context) func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithWaitForCompletion(v bool) func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithPretty() func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithHuman() func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithErrorTrace() func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithFilterPath(v ...string) func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithHeader(h map[string]string) func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexCancel) WithOpaqueID(s string) func(*ReindexCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}
