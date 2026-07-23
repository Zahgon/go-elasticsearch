package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteModelSnapshotFunc(t Transport) MLDeleteModelSnapshot {
	_ = "STUB: not implemented"
	return *new(MLDeleteModelSnapshot)
}

type MLDeleteModelSnapshot func(snapshot_id string, job_id string, o ...func(*MLDeleteModelSnapshotRequest)) (*Response, error)

type MLDeleteModelSnapshotRequest struct {
	JobID      string
	SnapshotID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteModelSnapshotRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteModelSnapshot) WithContext(v context.Context) func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteModelSnapshot) WithPretty() func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteModelSnapshot) WithHuman() func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteModelSnapshot) WithErrorTrace() func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteModelSnapshot) WithFilterPath(v ...string) func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteModelSnapshot) WithHeader(h map[string]string) func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteModelSnapshot) WithOpaqueID(s string) func(*MLDeleteModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}
