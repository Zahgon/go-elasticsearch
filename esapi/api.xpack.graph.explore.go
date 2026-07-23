package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newGraphExploreFunc(t Transport) GraphExplore {
	_ = "STUB: not implemented"
	return *new(GraphExplore)
}

type GraphExplore func(index []string, body io.Reader, o ...func(*GraphExploreRequest)) (*Response, error)

type GraphExploreRequest struct {
	Index []string

	Body io.Reader

	Routing []string
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r GraphExploreRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f GraphExplore) WithContext(v context.Context) func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithRouting(v ...string) func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithTimeout(v time.Duration) func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithPretty() func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithHuman() func(*GraphExploreRequest) { _ = "STUB: not implemented"; return nil }

func (f GraphExplore) WithErrorTrace() func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithFilterPath(v ...string) func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithHeader(h map[string]string) func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f GraphExplore) WithOpaqueID(s string) func(*GraphExploreRequest) {
	_ = "STUB: not implemented"
	return nil
}
