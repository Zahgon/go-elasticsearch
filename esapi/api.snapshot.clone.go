package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSnapshotCloneFunc(t Transport) SnapshotClone {
	_ = "STUB: not implemented"
	return *new(SnapshotClone)
}

type SnapshotClone func(repository string, snapshot string, body io.Reader, target_snapshot string, o ...func(*SnapshotCloneRequest)) (*Response, error)

type SnapshotCloneRequest struct {
	Body io.Reader

	Repository     string
	Snapshot       string
	TargetSnapshot string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotCloneRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotClone) WithContext(v context.Context) func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithMasterTimeout(v time.Duration) func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithPretty() func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithHuman() func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithErrorTrace() func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithFilterPath(v ...string) func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithHeader(h map[string]string) func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotClone) WithOpaqueID(s string) func(*SnapshotCloneRequest) {
	_ = "STUB: not implemented"
	return nil
}
