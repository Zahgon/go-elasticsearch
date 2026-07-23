package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCapabilitiesFunc(t Transport) Capabilities {
	_ = "STUB: not implemented"
	return *new(Capabilities)
}

type Capabilities func(o ...func(*CapabilitiesRequest)) (*Response, error)

type CapabilitiesRequest struct {
	Capabilities []string
	LocalOnly    *bool
	Method       string
	Parameters   []string
	Path         string
	Timeout      time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CapabilitiesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Capabilities) WithContext(v context.Context) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithCapabilities(v ...string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithLocalOnly(v bool) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithMethod(v string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithParameters(v ...string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithPath(v string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithTimeout(v time.Duration) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithPretty() func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithHuman() func(*CapabilitiesRequest) { _ = "STUB: not implemented"; return nil }

func (f Capabilities) WithErrorTrace() func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithFilterPath(v ...string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithHeader(h map[string]string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Capabilities) WithOpaqueID(s string) func(*CapabilitiesRequest) {
	_ = "STUB: not implemented"
	return nil
}
