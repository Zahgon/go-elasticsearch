package esapi

import (
	"context"
	"net/http"
)

func newNodesGetRepositoriesMeteringInfoFunc(t Transport) NodesGetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return *new(NodesGetRepositoriesMeteringInfo)
}

type NodesGetRepositoriesMeteringInfo func(node_id []string, o ...func(*NodesGetRepositoriesMeteringInfoRequest)) (*Response, error)

type NodesGetRepositoriesMeteringInfoRequest struct {
	NodeID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesGetRepositoriesMeteringInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesGetRepositoriesMeteringInfo) WithContext(v context.Context) func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesGetRepositoriesMeteringInfo) WithPretty() func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesGetRepositoriesMeteringInfo) WithHuman() func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesGetRepositoriesMeteringInfo) WithErrorTrace() func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesGetRepositoriesMeteringInfo) WithFilterPath(v ...string) func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesGetRepositoriesMeteringInfo) WithHeader(h map[string]string) func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesGetRepositoriesMeteringInfo) WithOpaqueID(s string) func(*NodesGetRepositoriesMeteringInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}
