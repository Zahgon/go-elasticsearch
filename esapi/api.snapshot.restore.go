package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSnapshotRestoreFunc(t Transport) SnapshotRestore {
	_ = "STUB: not implemented"
	return *new(SnapshotRestore)
}

type SnapshotRestore func(repository string, snapshot string, o ...func(*SnapshotRestoreRequest)) (*Response, error)

type SnapshotRestoreRequest struct {
	Body io.Reader

	Repository string
	Snapshot   string

	MasterTimeout     time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotRestoreRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotRestore) WithContext(v context.Context) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithBody(v io.Reader) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithMasterTimeout(v time.Duration) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithWaitForCompletion(v bool) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithPretty() func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithHuman() func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithErrorTrace() func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithFilterPath(v ...string) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithHeader(h map[string]string) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRestore) WithOpaqueID(s string) func(*SnapshotRestoreRequest) {
	_ = "STUB: not implemented"
	return nil
}
