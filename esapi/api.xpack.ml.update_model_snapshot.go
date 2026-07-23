package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLUpdateModelSnapshotFunc(t Transport) MLUpdateModelSnapshot {
	_ = "STUB: not implemented"
	return *new(MLUpdateModelSnapshot)
}

type MLUpdateModelSnapshot func(snapshot_id string, job_id string, body io.Reader, o ...func(*MLUpdateModelSnapshotRequest)) (*Response, error)

type MLUpdateModelSnapshotRequest struct {
	Body io.Reader

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

func (r MLUpdateModelSnapshotRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpdateModelSnapshot) WithContext(v context.Context) func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateModelSnapshot) WithPretty() func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateModelSnapshot) WithHuman() func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateModelSnapshot) WithErrorTrace() func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateModelSnapshot) WithFilterPath(v ...string) func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateModelSnapshot) WithHeader(h map[string]string) func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateModelSnapshot) WithOpaqueID(s string) func(*MLUpdateModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}
