package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSnapshotCreateFunc(t Transport) SnapshotCreate {
	_ = "STUB: not implemented"
	return *new(SnapshotCreate)
}

type SnapshotCreate func(repository string, snapshot string, o ...func(*SnapshotCreateRequest)) (*Response, error)

type SnapshotCreateRequest struct {
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

func (r SnapshotCreateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotCreate) WithContext(v context.Context) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithBody(v io.Reader) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithMasterTimeout(v time.Duration) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithWaitForCompletion(v bool) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithPretty() func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithHuman() func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithErrorTrace() func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithFilterPath(v ...string) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithHeader(h map[string]string) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreate) WithOpaqueID(s string) func(*SnapshotCreateRequest) {
	_ = "STUB: not implemented"
	return nil
}
