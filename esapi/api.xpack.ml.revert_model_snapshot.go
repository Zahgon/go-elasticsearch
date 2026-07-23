package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLRevertModelSnapshotFunc(t Transport) MLRevertModelSnapshot {
	_ = "STUB: not implemented"
	return *new(MLRevertModelSnapshot)
}

type MLRevertModelSnapshot func(snapshot_id string, job_id string, o ...func(*MLRevertModelSnapshotRequest)) (*Response, error)

type MLRevertModelSnapshotRequest struct {
	Body io.Reader

	JobID      string
	SnapshotID string

	DeleteInterveningResults *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLRevertModelSnapshotRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLRevertModelSnapshot) WithContext(v context.Context) func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithBody(v io.Reader) func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithDeleteInterveningResults(v bool) func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithPretty() func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithHuman() func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithErrorTrace() func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithFilterPath(v ...string) func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithHeader(h map[string]string) func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLRevertModelSnapshot) WithOpaqueID(s string) func(*MLRevertModelSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}
