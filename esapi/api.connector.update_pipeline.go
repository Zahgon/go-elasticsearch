package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdatePipelineFunc(t Transport) ConnectorUpdatePipeline {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdatePipeline)
}

type ConnectorUpdatePipeline func(body io.Reader, connector_id string, o ...func(*ConnectorUpdatePipelineRequest)) (*Response, error)

type ConnectorUpdatePipelineRequest struct {
	Body io.Reader

	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorUpdatePipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdatePipeline) WithContext(v context.Context) func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdatePipeline) WithPretty() func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdatePipeline) WithHuman() func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdatePipeline) WithErrorTrace() func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdatePipeline) WithFilterPath(v ...string) func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdatePipeline) WithHeader(h map[string]string) func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdatePipeline) WithOpaqueID(s string) func(*ConnectorUpdatePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
