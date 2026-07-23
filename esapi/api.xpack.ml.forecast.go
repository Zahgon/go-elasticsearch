package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLForecastFunc(t Transport) MLForecast { _ = "STUB: not implemented"; return *new(MLForecast) }

type MLForecast func(job_id string, o ...func(*MLForecastRequest)) (*Response, error)

type MLForecastRequest struct {
	Body io.Reader

	JobID string

	Duration       time.Duration
	ExpiresIn      time.Duration
	MaxModelMemory string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLForecastRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLForecast) WithContext(v context.Context) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithBody(v io.Reader) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithDuration(v time.Duration) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithExpiresIn(v time.Duration) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithMaxModelMemory(v string) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithPretty() func(*MLForecastRequest) { _ = "STUB: not implemented"; return nil }

func (f MLForecast) WithHuman() func(*MLForecastRequest) { _ = "STUB: not implemented"; return nil }

func (f MLForecast) WithErrorTrace() func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithFilterPath(v ...string) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithHeader(h map[string]string) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLForecast) WithOpaqueID(s string) func(*MLForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}
