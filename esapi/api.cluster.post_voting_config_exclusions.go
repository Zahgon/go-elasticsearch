package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterPostVotingConfigExclusionsFunc(t Transport) ClusterPostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return *new(ClusterPostVotingConfigExclusions)
}

type ClusterPostVotingConfigExclusions func(o ...func(*ClusterPostVotingConfigExclusionsRequest)) (*Response, error)

type ClusterPostVotingConfigExclusionsRequest struct {
	MasterTimeout time.Duration
	NodeIds       []string
	NodeNames     []string
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterPostVotingConfigExclusionsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterPostVotingConfigExclusions) WithContext(v context.Context) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithMasterTimeout(v time.Duration) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithNodeIds(v ...string) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithNodeNames(v ...string) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithTimeout(v time.Duration) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithPretty() func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithHuman() func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithErrorTrace() func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithFilterPath(v ...string) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithHeader(h map[string]string) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPostVotingConfigExclusions) WithOpaqueID(s string) func(*ClusterPostVotingConfigExclusionsRequest) {
	_ = "STUB: not implemented"
	return nil
}
