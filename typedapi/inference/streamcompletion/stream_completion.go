package streamcompletion

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	inferenceidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StreamCompletion struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	inferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStreamCompletion func(inferenceid string) *StreamCompletion

func NewStreamCompletionFunc(tp elastictransport.Interface) NewStreamCompletion {
	_ = "STUB: not implemented"
	return *new(NewStreamCompletion)
}

func New(tp elastictransport.Interface) *StreamCompletion { _ = "STUB: not implemented"; return nil }

func (r *StreamCompletion) Raw(raw io.Reader) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) Request(req *Request) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StreamCompletion) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StreamCompletion) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *StreamCompletion) Header(key, value string) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) _inferenceid(inferenceid string) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) Timeout(duration string) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) ErrorTrace(errortrace bool) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) FilterPath(filterpaths ...string) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) Human(human bool) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) Pretty(pretty bool) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) Input(inputs ...string) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (r *StreamCompletion) TaskSettings(tasksettings json.RawMessage) *StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}
