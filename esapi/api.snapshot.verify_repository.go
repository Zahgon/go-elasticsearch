package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotVerifyRepositoryFunc(t Transport) SnapshotVerifyRepository {
	_ = "STUB: not implemented"
	return *new(SnapshotVerifyRepository)
}

type SnapshotVerifyRepository func(repository string, o ...func(*SnapshotVerifyRepositoryRequest)) (*Response, error)

type SnapshotVerifyRepositoryRequest struct {
	Repository string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotVerifyRepositoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotVerifyRepository) WithContext(v context.Context) func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithMasterTimeout(v time.Duration) func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithTimeout(v time.Duration) func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithPretty() func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithHuman() func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithErrorTrace() func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithFilterPath(v ...string) func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithHeader(h map[string]string) func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotVerifyRepository) WithOpaqueID(s string) func(*SnapshotVerifyRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}
