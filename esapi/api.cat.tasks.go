package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatTasksFunc(t Transport) CatTasks { _ = "STUB: not implemented"; return *new(CatTasks) }

type CatTasks func(o ...func(*CatTasksRequest)) (*Response, error)

type CatTasksRequest struct {
	Actions           []string
	Bytes             string
	Detailed          *bool
	Format            string
	H                 []string
	Help              *bool
	Nodes             []string
	ParentTaskID      string
	S                 []string
	Time              string
	Timeout           time.Duration
	V                 *bool
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatTasksRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatTasks) WithContext(v context.Context) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithActions(v ...string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithBytes(v string) func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithDetailed(v bool) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithFormat(v string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithH(v ...string) func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithHelp(v bool) func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithNodes(v ...string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithParentTaskID(v string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithS(v ...string) func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithTime(v string) func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithTimeout(v time.Duration) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithV(v bool) func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithWaitForCompletion(v bool) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithPretty() func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithHuman() func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithErrorTrace() func(*CatTasksRequest) { _ = "STUB: not implemented"; return nil }

func (f CatTasks) WithFilterPath(v ...string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithHeader(h map[string]string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTasks) WithOpaqueID(s string) func(*CatTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}
