package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterDeleteVotingConfigExclusionsFunc(t Transport) ClusterDeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return *new(ClusterDeleteVotingConfigExclusions)
}

type ClusterDeleteVotingConfigExclusions func(o ...func(*ClusterDeleteVotingConfigExclusionsRequest)) (*Response, error)

type ClusterDeleteVotingConfigExclusionsRequest struct {
	MasterTimeout  time.Duration
	WaitForRemoval *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterDeleteVotingConfigExclusionsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterDeleteVotingConfigExclusions) WithContext(v context.Context) func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithMasterTimeout(v time.Duration) func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithWaitForRemoval(v bool) func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithPretty() func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithHuman() func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithErrorTrace() func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithFilterPath(v ...string) func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithHeader(h map[string]string) func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteVotingConfigExclusions) WithOpaqueID(s string) func(*ClusterDeleteVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}
