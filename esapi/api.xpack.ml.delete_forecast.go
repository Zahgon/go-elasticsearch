package esapi

import (
	"context"
	"net/http"
	"time"
)

func newMLDeleteForecastFunc(t Transport) MLDeleteForecast {
	_ = "STUB: not implemented"
	return *new(MLDeleteForecast)
}

type MLDeleteForecast func(job_id string, o ...func(*MLDeleteForecastRequest)) (*Response, error)

type MLDeleteForecastRequest struct {
	ForecastID string
	JobID      string

	AllowNoForecasts *bool
	Timeout          time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteForecastRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteForecast) WithContext(v context.Context) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithForecastID(v string) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithAllowNoForecasts(v bool) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithTimeout(v time.Duration) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithPretty() func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithHuman() func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithErrorTrace() func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithFilterPath(v ...string) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithHeader(h map[string]string) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteForecast) WithOpaqueID(s string) func(*MLDeleteForecastRequest) {
	_ = "STUB: not implemented"
	return nil
}
