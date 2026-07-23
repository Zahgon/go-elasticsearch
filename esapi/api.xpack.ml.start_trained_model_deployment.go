package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLStartTrainedModelDeploymentFunc(t Transport) MLStartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return *new(MLStartTrainedModelDeployment)
}

type MLStartTrainedModelDeployment func(model_id string, o ...func(*MLStartTrainedModelDeploymentRequest)) (*Response, error)

type MLStartTrainedModelDeploymentRequest struct {
	Body io.Reader

	ModelID string

	CacheSize            string
	DeploymentID         string
	NumberOfAllocations  *int
	Priority             string
	QueueCapacity        *int
	ThreadsPerAllocation *int
	Timeout              time.Duration
	WaitFor              string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLStartTrainedModelDeploymentRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLStartTrainedModelDeployment) WithContext(v context.Context) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithBody(v io.Reader) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithCacheSize(v string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithDeploymentID(v string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithNumberOfAllocations(v int) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithPriority(v string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithQueueCapacity(v int) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithThreadsPerAllocation(v int) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithTimeout(v time.Duration) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithWaitFor(v string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithPretty() func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithHuman() func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithErrorTrace() func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithFilterPath(v ...string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithHeader(h map[string]string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartTrainedModelDeployment) WithOpaqueID(s string) func(*MLStartTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}
