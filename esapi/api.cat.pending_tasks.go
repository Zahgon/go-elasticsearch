package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatPendingTasksFunc(t Transport) CatPendingTasks {
	_ = "STUB: not implemented"
	return *new(CatPendingTasks)
}

type CatPendingTasks func(o ...func(*CatPendingTasksRequest)) (*Response, error)

type CatPendingTasksRequest struct {
	Bytes         string
	Format        string
	H             []string
	Help          *bool
	Local         *bool
	MasterTimeout time.Duration
	S             []string
	Time          string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatPendingTasksRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatPendingTasks) WithContext(v context.Context) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithBytes(v string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithFormat(v string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithH(v ...string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithHelp(v bool) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithLocal(v bool) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithMasterTimeout(v time.Duration) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithS(v ...string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithTime(v string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithV(v bool) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithPretty() func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithHuman() func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithErrorTrace() func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithFilterPath(v ...string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithHeader(h map[string]string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPendingTasks) WithOpaqueID(s string) func(*CatPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}
