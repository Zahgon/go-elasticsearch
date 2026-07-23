package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotRepositoryAnalyzeFunc(t Transport) SnapshotRepositoryAnalyze {
	_ = "STUB: not implemented"
	return *new(SnapshotRepositoryAnalyze)
}

type SnapshotRepositoryAnalyze func(repository string, o ...func(*SnapshotRepositoryAnalyzeRequest)) (*Response, error)

type SnapshotRepositoryAnalyzeRequest struct {
	Repository string

	BlobCount              *int
	Concurrency            *int
	Detailed               *bool
	EarlyReadNodeCount     *int
	MaxBlobSize            string
	MaxTotalDataSize       string
	RareActionProbability  interface{}
	RarelyAbortWrites      *bool
	ReadNodeCount          *int
	RegisterOperationCount *int
	Seed                   *int
	Timeout                time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotRepositoryAnalyzeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotRepositoryAnalyze) WithContext(v context.Context) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithBlobCount(v int) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithConcurrency(v int) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithDetailed(v bool) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithEarlyReadNodeCount(v int) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithMaxBlobSize(v string) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithMaxTotalDataSize(v string) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithRareActionProbability(v interface{}) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithRarelyAbortWrites(v bool) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithReadNodeCount(v int) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithRegisterOperationCount(v int) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithSeed(v int) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithTimeout(v time.Duration) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithPretty() func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithHuman() func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithErrorTrace() func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithFilterPath(v ...string) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithHeader(h map[string]string) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryAnalyze) WithOpaqueID(s string) func(*SnapshotRepositoryAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}
