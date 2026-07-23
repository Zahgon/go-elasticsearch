package esapi

import (
	"context"
	"net/http"
)

func newTasksCancelFunc(t Transport) TasksCancel {
	_ = "STUB: not implemented"
	return *new(TasksCancel)
}

type TasksCancel func(o ...func(*TasksCancelRequest)) (*Response, error)

type TasksCancelRequest struct {
	TaskID string

	Actions           []string
	Nodes             []string
	ParentTaskID      string
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TasksCancelRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TasksCancel) WithContext(v context.Context) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithTaskID(v string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithActions(v ...string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithNodes(v ...string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithParentTaskID(v string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithWaitForCompletion(v bool) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithPretty() func(*TasksCancelRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksCancel) WithHuman() func(*TasksCancelRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksCancel) WithErrorTrace() func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithFilterPath(v ...string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithHeader(h map[string]string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksCancel) WithOpaqueID(s string) func(*TasksCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}
