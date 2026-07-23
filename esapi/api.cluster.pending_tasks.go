package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterPendingTasksFunc(t Transport) ClusterPendingTasks {
	_ = "STUB: not implemented"
	return *new(ClusterPendingTasks)
}

type ClusterPendingTasks func(o ...func(*ClusterPendingTasksRequest)) (*Response, error)

type ClusterPendingTasksRequest struct {
	Local         *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterPendingTasksRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterPendingTasks) WithContext(v context.Context) func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithLocal(v bool) func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithMasterTimeout(v time.Duration) func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithPretty() func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithHuman() func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithErrorTrace() func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithFilterPath(v ...string) func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithHeader(h map[string]string) func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPendingTasks) WithOpaqueID(s string) func(*ClusterPendingTasksRequest) {
	_ = "STUB: not implemented"
	return nil
}
