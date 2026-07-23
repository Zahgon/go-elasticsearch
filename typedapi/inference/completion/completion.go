package completion

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

type Completion struct {
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

type NewCompletion func(inferenceid string) *Completion

func NewCompletionFunc(tp elastictransport.Interface) NewCompletion {
	_ = "STUB: not implemented"
	return *new(NewCompletion)
}

func New(tp elastictransport.Interface) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) Raw(raw io.Reader) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) Request(req *Request) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Completion) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Completion) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *Completion) Header(key, value string) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) _inferenceid(inferenceid string) *Completion {
	_ = "STUB: not implemented"
	return nil
}

func (r *Completion) Timeout(duration string) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) ErrorTrace(errortrace bool) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) FilterPath(filterpaths ...string) *Completion {
	_ = "STUB: not implemented"
	return nil
}

func (r *Completion) Human(human bool) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) Pretty(pretty bool) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) Input(inputs ...string) *Completion { _ = "STUB: not implemented"; return nil }

func (r *Completion) TaskSettings(tasksettings json.RawMessage) *Completion {
	_ = "STUB: not implemented"
	return nil
}
