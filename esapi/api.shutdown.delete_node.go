package esapi

import (
	"context"
	"net/http"
	"time"
)

func newShutdownDeleteNodeFunc(t Transport) ShutdownDeleteNode {
	_ = "STUB: not implemented"
	return *new(ShutdownDeleteNode)
}

type ShutdownDeleteNode func(node_id string, o ...func(*ShutdownDeleteNodeRequest)) (*Response, error)

type ShutdownDeleteNodeRequest struct {
	NodeID string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ShutdownDeleteNodeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ShutdownDeleteNode) WithContext(v context.Context) func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithMasterTimeout(v time.Duration) func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithTimeout(v time.Duration) func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithPretty() func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithHuman() func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithErrorTrace() func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithFilterPath(v ...string) func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithHeader(h map[string]string) func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownDeleteNode) WithOpaqueID(s string) func(*ShutdownDeleteNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}
