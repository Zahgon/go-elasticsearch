package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotGetFunc(t Transport) SnapshotGet {
	_ = "STUB: not implemented"
	return *new(SnapshotGet)
}

type SnapshotGet func(repository string, snapshot []string, o ...func(*SnapshotGetRequest)) (*Response, error)

type SnapshotGetRequest struct {
	Repository string
	Snapshot   []string

	After             string
	FromSortValue     string
	IgnoreUnavailable *bool
	IncludeRepository *bool
	IndexDetails      *bool
	IndexNames        *bool
	MasterTimeout     time.Duration
	Offset            *int
	Order             string
	Size              *int
	SlmPolicyFilter   string
	Sort              string
	State             []string
	Verbose           *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotGet) WithContext(v context.Context) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithAfter(v string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithFromSortValue(v string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithIgnoreUnavailable(v bool) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithIncludeRepository(v bool) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithIndexDetails(v bool) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithIndexNames(v bool) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithMasterTimeout(v time.Duration) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithOffset(v int) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithOrder(v string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithSize(v int) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithSlmPolicyFilter(v string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithSort(v string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithState(v ...string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithVerbose(v bool) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithPretty() func(*SnapshotGetRequest) { _ = "STUB: not implemented"; return nil }

func (f SnapshotGet) WithHuman() func(*SnapshotGetRequest) { _ = "STUB: not implemented"; return nil }

func (f SnapshotGet) WithErrorTrace() func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithFilterPath(v ...string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithHeader(h map[string]string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGet) WithOpaqueID(s string) func(*SnapshotGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
