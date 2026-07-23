package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTasksGetFunc(t Transport) TasksGet { _ = "STUB: not implemented"; return *new(TasksGet) }

type TasksGet func(task_id string, o ...func(*TasksGetRequest)) (*Response, error)

type TasksGetRequest struct {
	TaskID string

	Timeout           time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TasksGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TasksGet) WithContext(v context.Context) func(*TasksGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksGet) WithTimeout(v time.Duration) func(*TasksGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksGet) WithWaitForCompletion(v bool) func(*TasksGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksGet) WithPretty() func(*TasksGetRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksGet) WithHuman() func(*TasksGetRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksGet) WithErrorTrace() func(*TasksGetRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksGet) WithFilterPath(v ...string) func(*TasksGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksGet) WithHeader(h map[string]string) func(*TasksGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksGet) WithOpaqueID(s string) func(*TasksGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
