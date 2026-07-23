package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLGetModelSnapshotsFunc(t Transport) MLGetModelSnapshots {
	_ = "STUB: not implemented"
	return *new(MLGetModelSnapshots)
}

type MLGetModelSnapshots func(job_id string, o ...func(*MLGetModelSnapshotsRequest)) (*Response, error)

type MLGetModelSnapshotsRequest struct {
	Body io.Reader

	JobID      string
	SnapshotID string

	Desc  *bool
	End   string
	From  *int
	Size  *int
	Sort  string
	Start string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetModelSnapshotsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetModelSnapshots) WithContext(v context.Context) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithBody(v io.Reader) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithSnapshotID(v string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithDesc(v bool) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithEnd(v string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithFrom(v int) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithSize(v int) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithSort(v string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithStart(v string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithPretty() func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithHuman() func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithErrorTrace() func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithFilterPath(v ...string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithHeader(h map[string]string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshots) WithOpaqueID(s string) func(*MLGetModelSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}
