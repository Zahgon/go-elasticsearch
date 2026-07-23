package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatSnapshotsFunc(t Transport) CatSnapshots {
	_ = "STUB: not implemented"
	return *new(CatSnapshots)
}

type CatSnapshots func(o ...func(*CatSnapshotsRequest)) (*Response, error)

type CatSnapshotsRequest struct {
	Repository []string

	Bytes             string
	Format            string
	H                 []string
	Help              *bool
	IgnoreUnavailable *bool
	MasterTimeout     time.Duration
	S                 []string
	Time              string
	V                 *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatSnapshotsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatSnapshots) WithContext(v context.Context) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithRepository(v ...string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithBytes(v string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithFormat(v string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithH(v ...string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithHelp(v bool) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithIgnoreUnavailable(v bool) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithMasterTimeout(v time.Duration) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithS(v ...string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithTime(v string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithV(v bool) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithPretty() func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithHuman() func(*CatSnapshotsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatSnapshots) WithErrorTrace() func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithFilterPath(v ...string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithHeader(h map[string]string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatSnapshots) WithOpaqueID(s string) func(*CatSnapshotsRequest) {
	_ = "STUB: not implemented"
	return nil
}
