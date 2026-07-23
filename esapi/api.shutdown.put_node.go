package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newShutdownPutNodeFunc(t Transport) ShutdownPutNode {
	_ = "STUB: not implemented"
	return *new(ShutdownPutNode)
}

type ShutdownPutNode func(body io.Reader, node_id string, o ...func(*ShutdownPutNodeRequest)) (*Response, error)

type ShutdownPutNodeRequest struct {
	Body io.Reader

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

func (r ShutdownPutNodeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ShutdownPutNode) WithContext(v context.Context) func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithMasterTimeout(v time.Duration) func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithTimeout(v time.Duration) func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithPretty() func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithHuman() func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithErrorTrace() func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithFilterPath(v ...string) func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithHeader(h map[string]string) func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownPutNode) WithOpaqueID(s string) func(*ShutdownPutNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}
