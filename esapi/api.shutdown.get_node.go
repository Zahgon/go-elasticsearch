package esapi

import (
	"context"
	"net/http"
	"time"
)

func newShutdownGetNodeFunc(t Transport) ShutdownGetNode {
	_ = "STUB: not implemented"
	return *new(ShutdownGetNode)
}

type ShutdownGetNode func(o ...func(*ShutdownGetNodeRequest)) (*Response, error)

type ShutdownGetNodeRequest struct {
	NodeID []string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ShutdownGetNodeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ShutdownGetNode) WithContext(v context.Context) func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithNodeID(v ...string) func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithMasterTimeout(v time.Duration) func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithPretty() func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithHuman() func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithErrorTrace() func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithFilterPath(v ...string) func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithHeader(h map[string]string) func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ShutdownGetNode) WithOpaqueID(s string) func(*ShutdownGetNodeRequest) {
	_ = "STUB: not implemented"
	return nil
}
