package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotDeleteFunc(t Transport) SnapshotDelete {
	_ = "STUB: not implemented"
	return *new(SnapshotDelete)
}

type SnapshotDelete func(repository string, snapshot []string, o ...func(*SnapshotDeleteRequest)) (*Response, error)

type SnapshotDeleteRequest struct {
	Repository string
	Snapshot   []string

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

func (r SnapshotDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotDelete) WithContext(v context.Context) func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithMasterTimeout(v time.Duration) func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithWaitForCompletion(v bool) func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithPretty() func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithHuman() func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithErrorTrace() func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithFilterPath(v ...string) func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithHeader(h map[string]string) func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDelete) WithOpaqueID(s string) func(*SnapshotDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
