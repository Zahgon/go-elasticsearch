package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotStatusFunc(t Transport) SnapshotStatus {
	_ = "STUB: not implemented"
	return *new(SnapshotStatus)
}

type SnapshotStatus func(o ...func(*SnapshotStatusRequest)) (*Response, error)

type SnapshotStatusRequest struct {
	Repository string
	Snapshot   []string

	IgnoreUnavailable *bool
	MasterTimeout     time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotStatus) WithContext(v context.Context) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithRepository(v string) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithSnapshot(v ...string) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithIgnoreUnavailable(v bool) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithMasterTimeout(v time.Duration) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithPretty() func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithHuman() func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithErrorTrace() func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithFilterPath(v ...string) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithHeader(h map[string]string) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotStatus) WithOpaqueID(s string) func(*SnapshotStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
