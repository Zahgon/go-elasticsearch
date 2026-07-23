package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newNodesReloadSecureSettingsFunc(t Transport) NodesReloadSecureSettings {
	_ = "STUB: not implemented"
	return *new(NodesReloadSecureSettings)
}

type NodesReloadSecureSettings func(o ...func(*NodesReloadSecureSettingsRequest)) (*Response, error)

type NodesReloadSecureSettingsRequest struct {
	Body io.Reader

	NodeID []string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesReloadSecureSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesReloadSecureSettings) WithContext(v context.Context) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithBody(v io.Reader) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithNodeID(v ...string) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithTimeout(v time.Duration) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithPretty() func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithHuman() func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithErrorTrace() func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithFilterPath(v ...string) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithHeader(h map[string]string) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesReloadSecureSettings) WithOpaqueID(s string) func(*NodesReloadSecureSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
