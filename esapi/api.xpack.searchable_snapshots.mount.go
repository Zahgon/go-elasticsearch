package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSearchableSnapshotsMountFunc(t Transport) SearchableSnapshotsMount {
	_ = "STUB: not implemented"
	return *new(SearchableSnapshotsMount)
}

type SearchableSnapshotsMount func(repository string, snapshot string, body io.Reader, o ...func(*SearchableSnapshotsMountRequest)) (*Response, error)

type SearchableSnapshotsMountRequest struct {
	Body io.Reader

	Repository string
	Snapshot   string

	MasterTimeout     time.Duration
	Storage           string
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchableSnapshotsMountRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchableSnapshotsMount) WithContext(v context.Context) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithMasterTimeout(v time.Duration) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithStorage(v string) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithWaitForCompletion(v bool) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithPretty() func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithHuman() func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithErrorTrace() func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithFilterPath(v ...string) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithHeader(h map[string]string) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsMount) WithOpaqueID(s string) func(*SearchableSnapshotsMountRequest) {
	_ = "STUB: not implemented"
	return nil
}
