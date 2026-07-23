package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLDeleteExpiredDataFunc(t Transport) MLDeleteExpiredData {
	_ = "STUB: not implemented"
	return *new(MLDeleteExpiredData)
}

type MLDeleteExpiredData func(o ...func(*MLDeleteExpiredDataRequest)) (*Response, error)

type MLDeleteExpiredDataRequest struct {
	Body io.Reader

	JobID string

	RequestsPerSecond *int
	Timeout           time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteExpiredDataRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteExpiredData) WithContext(v context.Context) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithBody(v io.Reader) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithJobID(v string) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithRequestsPerSecond(v int) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithTimeout(v time.Duration) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithPretty() func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithHuman() func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithErrorTrace() func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithFilterPath(v ...string) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithHeader(h map[string]string) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteExpiredData) WithOpaqueID(s string) func(*MLDeleteExpiredDataRequest) {
	_ = "STUB: not implemented"
	return nil
}
