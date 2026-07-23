package retry

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Retry struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRetry func(index string) *Retry

func NewRetryFunc(tp elastictransport.Interface) NewRetry {
	_ = "STUB: not implemented"
	return *new(NewRetry)
}

func New(tp elastictransport.Interface) *Retry { _ = "STUB: not implemented"; return nil }

func (r *Retry) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Retry) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Retry) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Retry) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Retry) Header(key, value string) *Retry { _ = "STUB: not implemented"; return nil }

func (r *Retry) _index(index string) *Retry { _ = "STUB: not implemented"; return nil }

func (r *Retry) ErrorTrace(errortrace bool) *Retry { _ = "STUB: not implemented"; return nil }

func (r *Retry) FilterPath(filterpaths ...string) *Retry { _ = "STUB: not implemented"; return nil }

func (r *Retry) Human(human bool) *Retry { _ = "STUB: not implemented"; return nil }

func (r *Retry) Pretty(pretty bool) *Retry { _ = "STUB: not implemented"; return nil }
