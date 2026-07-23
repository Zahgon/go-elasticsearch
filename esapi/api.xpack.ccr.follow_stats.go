package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRFollowStatsFunc(t Transport) CCRFollowStats {
	_ = "STUB: not implemented"
	return *new(CCRFollowStats)
}

type CCRFollowStats func(index []string, o ...func(*CCRFollowStatsRequest)) (*Response, error)

type CCRFollowStatsRequest struct {
	Index []string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRFollowStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRFollowStats) WithContext(v context.Context) func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithTimeout(v time.Duration) func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithPretty() func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithHuman() func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithErrorTrace() func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithFilterPath(v ...string) func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithHeader(h map[string]string) func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRFollowStats) WithOpaqueID(s string) func(*CCRFollowStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
