package esapi

import (
	"context"
	"net/http"
)

func newMLGetModelSnapshotUpgradeStatsFunc(t Transport) MLGetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return *new(MLGetModelSnapshotUpgradeStats)
}

type MLGetModelSnapshotUpgradeStats func(snapshot_id string, job_id string, o ...func(*MLGetModelSnapshotUpgradeStatsRequest)) (*Response, error)

type MLGetModelSnapshotUpgradeStatsRequest struct {
	JobID      string
	SnapshotID string

	AllowNoMatch *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetModelSnapshotUpgradeStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetModelSnapshotUpgradeStats) WithContext(v context.Context) func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithAllowNoMatch(v bool) func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithPretty() func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithHuman() func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithErrorTrace() func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithFilterPath(v ...string) func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithHeader(h map[string]string) func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetModelSnapshotUpgradeStats) WithOpaqueID(s string) func(*MLGetModelSnapshotUpgradeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
