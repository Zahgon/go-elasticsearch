package inference

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
	tasktypeMask = iota + 1

	inferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Inference struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype    string
	inferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewInference func(inferenceid string) *Inference

func NewInferenceFunc(tp elastictransport.Interface) NewInference {
	_ = "STUB: not implemented"
	return *new(NewInference)
}

func New(tp elastictransport.Interface) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) Raw(raw io.Reader) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) Request(req *Request) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Inference) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Inference) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Inference) Header(key, value string) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) TaskType(tasktype string) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) _inferenceid(inferenceid string) *Inference {
	_ = "STUB: not implemented"
	return nil
}

func (r *Inference) Timeout(duration string) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) ErrorTrace(errortrace bool) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) FilterPath(filterpaths ...string) *Inference {
	_ = "STUB: not implemented"
	return nil
}

func (r *Inference) Human(human bool) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) Pretty(pretty bool) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) Input(inputs ...string) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) InputType(inputtype string) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) Query(query string) *Inference { _ = "STUB: not implemented"; return nil }

func (r *Inference) TaskSettings(tasksettings json.RawMessage) *Inference {
	_ = "STUB: not implemented"
	return nil
}
