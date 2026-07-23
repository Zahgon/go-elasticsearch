package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newOpenPointInTimeFunc(t Transport) OpenPointInTime {
	_ = "STUB: not implemented"
	return *new(OpenPointInTime)
}

type OpenPointInTime func(index []string, keep_alive time.Duration, o ...func(*OpenPointInTimeRequest)) (*Response, error)

type OpenPointInTimeRequest struct {
	Index []string

	Body io.Reader

	AllowPartialSearchResults  *bool
	ExpandWildcards            []string
	IgnoreUnavailable          *bool
	KeepAlive                  time.Duration
	MaxConcurrentShardRequests *int
	Preference                 string
	Routing                    []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r OpenPointInTimeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f OpenPointInTime) WithContext(v context.Context) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithBody(v io.Reader) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithAllowPartialSearchResults(v bool) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithExpandWildcards(v ...string) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithIgnoreUnavailable(v bool) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithKeepAlive(v time.Duration) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithMaxConcurrentShardRequests(v int) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithPreference(v string) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithRouting(v ...string) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithPretty() func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithHuman() func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithErrorTrace() func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithFilterPath(v ...string) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithHeader(h map[string]string) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f OpenPointInTime) WithOpaqueID(s string) func(*OpenPointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}
