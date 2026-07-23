package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLGetInfluencersFunc(t Transport) MLGetInfluencers {
	_ = "STUB: not implemented"
	return *new(MLGetInfluencers)
}

type MLGetInfluencers func(job_id string, o ...func(*MLGetInfluencersRequest)) (*Response, error)

type MLGetInfluencersRequest struct {
	Body io.Reader

	JobID string

	Desc            *bool
	End             string
	ExcludeInterim  *bool
	From            *int
	InfluencerScore interface{}
	Size            *int
	Sort            string
	Start           string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetInfluencersRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetInfluencers) WithContext(v context.Context) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithBody(v io.Reader) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithDesc(v bool) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithEnd(v string) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithExcludeInterim(v bool) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithFrom(v int) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithInfluencerScore(v interface{}) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithSize(v int) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithSort(v string) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithStart(v string) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithPretty() func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithHuman() func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithErrorTrace() func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithFilterPath(v ...string) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithHeader(h map[string]string) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetInfluencers) WithOpaqueID(s string) func(*MLGetInfluencersRequest) {
	_ = "STUB: not implemented"
	return nil
}
