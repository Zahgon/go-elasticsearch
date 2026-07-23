package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMonitoringBulkFunc(t Transport) MonitoringBulk {
	_ = "STUB: not implemented"
	return *new(MonitoringBulk)
}

type MonitoringBulk func(body io.Reader, interval time.Duration, system_api_version string, system_id string, o ...func(*MonitoringBulkRequest)) (*Response, error)

type MonitoringBulkRequest struct {
	Body io.Reader

	Interval         time.Duration
	SystemAPIVersion string
	SystemID         string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MonitoringBulkRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MonitoringBulk) WithContext(v context.Context) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithInterval(v time.Duration) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithSystemAPIVersion(v string) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithSystemID(v string) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithPretty() func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithHuman() func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithErrorTrace() func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithFilterPath(v ...string) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithHeader(h map[string]string) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MonitoringBulk) WithOpaqueID(s string) func(*MonitoringBulkRequest) {
	_ = "STUB: not implemented"
	return nil
}
