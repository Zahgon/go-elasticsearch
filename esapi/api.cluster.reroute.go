package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newClusterRerouteFunc(t Transport) ClusterReroute {
	_ = "STUB: not implemented"
	return *new(ClusterReroute)
}

type ClusterReroute func(o ...func(*ClusterRerouteRequest)) (*Response, error)

type ClusterRerouteRequest struct {
	Body io.Reader

	DryRun        *bool
	Explain       *bool
	MasterTimeout time.Duration
	Metric        []string
	RetryFailed   *bool
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterRerouteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterReroute) WithContext(v context.Context) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithBody(v io.Reader) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithDryRun(v bool) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithExplain(v bool) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithMasterTimeout(v time.Duration) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithMetric(v ...string) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithRetryFailed(v bool) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithTimeout(v time.Duration) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithPretty() func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithHuman() func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithErrorTrace() func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithFilterPath(v ...string) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithHeader(h map[string]string) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterReroute) WithOpaqueID(s string) func(*ClusterRerouteRequest) {
	_ = "STUB: not implemented"
	return nil
}
