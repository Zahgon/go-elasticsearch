package esapi

import (
	"context"
	"net/http"
	"time"
)

func newFleetGlobalCheckpointsFunc(t Transport) FleetGlobalCheckpoints {
	_ = "STUB: not implemented"
	return *new(FleetGlobalCheckpoints)
}

type FleetGlobalCheckpoints func(index string, o ...func(*FleetGlobalCheckpointsRequest)) (*Response, error)

type FleetGlobalCheckpointsRequest struct {
	Index string

	Checkpoints    []string
	Timeout        time.Duration
	WaitForAdvance *bool
	WaitForIndex   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FleetGlobalCheckpointsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FleetGlobalCheckpoints) WithContext(v context.Context) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithCheckpoints(v ...string) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithTimeout(v time.Duration) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithWaitForAdvance(v bool) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithWaitForIndex(v bool) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithPretty() func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithHuman() func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithErrorTrace() func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithFilterPath(v ...string) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithHeader(h map[string]string) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGlobalCheckpoints) WithOpaqueID(s string) func(*FleetGlobalCheckpointsRequest) {
	_ = "STUB: not implemented"
	return nil
}
