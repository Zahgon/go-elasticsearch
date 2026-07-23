package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSnapshotCreateRepositoryFunc(t Transport) SnapshotCreateRepository {
	_ = "STUB: not implemented"
	return *new(SnapshotCreateRepository)
}

type SnapshotCreateRepository func(repository string, body io.Reader, o ...func(*SnapshotCreateRepositoryRequest)) (*Response, error)

type SnapshotCreateRepositoryRequest struct {
	Body io.Reader

	Repository string

	MasterTimeout time.Duration
	Timeout       time.Duration
	Verify        *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotCreateRepositoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotCreateRepository) WithContext(v context.Context) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithMasterTimeout(v time.Duration) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithTimeout(v time.Duration) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithVerify(v bool) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithPretty() func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithHuman() func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithErrorTrace() func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithFilterPath(v ...string) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithHeader(h map[string]string) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCreateRepository) WithOpaqueID(s string) func(*SnapshotCreateRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}
