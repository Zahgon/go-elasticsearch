package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTasksListFunc(t Transport) TasksList { _ = "STUB: not implemented"; return *new(TasksList) }

type TasksList func(o ...func(*TasksListRequest)) (*Response, error)

type TasksListRequest struct {
	Actions           []string
	Detailed          *bool
	GroupBy           string
	Nodes             []string
	ParentTaskID      string
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

func (r TasksListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TasksList) WithContext(v context.Context) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithActions(v ...string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithDetailed(v bool) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithGroupBy(v string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithNodes(v ...string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithParentTaskID(v string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithTimeout(v time.Duration) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithWaitForCompletion(v bool) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithPretty() func(*TasksListRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksList) WithHuman() func(*TasksListRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksList) WithErrorTrace() func(*TasksListRequest) { _ = "STUB: not implemented"; return nil }

func (f TasksList) WithFilterPath(v ...string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithHeader(h map[string]string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TasksList) WithOpaqueID(s string) func(*TasksListRequest) {
	_ = "STUB: not implemented"
	return nil
}
