package esapi

import (
	"context"
	"net/http"
	"time"
)

func newMLUpgradeJobSnapshotFunc(t Transport) MLUpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return *new(MLUpgradeJobSnapshot)
}

type MLUpgradeJobSnapshot func(snapshot_id string, job_id string, o ...func(*MLUpgradeJobSnapshotRequest)) (*Response, error)

type MLUpgradeJobSnapshotRequest struct {
	JobID      string
	SnapshotID string

	Timeout           time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLUpgradeJobSnapshotRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpgradeJobSnapshot) WithContext(v context.Context) func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithTimeout(v time.Duration) func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithWaitForCompletion(v bool) func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithPretty() func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithHuman() func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithErrorTrace() func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithFilterPath(v ...string) func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithHeader(h map[string]string) func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpgradeJobSnapshot) WithOpaqueID(s string) func(*MLUpgradeJobSnapshotRequest) {
	_ = "STUB: not implemented"
	return nil
}
