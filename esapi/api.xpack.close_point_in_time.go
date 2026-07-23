package esapi

import (
	"context"
	"io"
	"net/http"
)

func newClosePointInTimeFunc(t Transport) ClosePointInTime {
	_ = "STUB: not implemented"
	return *new(ClosePointInTime)
}

type ClosePointInTime func(body io.Reader, o ...func(*ClosePointInTimeRequest)) (*Response, error)

type ClosePointInTimeRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClosePointInTimeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClosePointInTime) WithContext(v context.Context) func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClosePointInTime) WithPretty() func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClosePointInTime) WithHuman() func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClosePointInTime) WithErrorTrace() func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClosePointInTime) WithFilterPath(v ...string) func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClosePointInTime) WithHeader(h map[string]string) func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClosePointInTime) WithOpaqueID(s string) func(*ClosePointInTimeRequest) {
	_ = "STUB: not implemented"
	return nil
}
