package asyncquerydelete

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AsyncQueryDelete struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewAsyncQueryDelete func(id string) *AsyncQueryDelete

func NewAsyncQueryDeleteFunc(tp elastictransport.Interface) NewAsyncQueryDelete {
	_ = "STUB: not implemented"
	return *new(NewAsyncQueryDelete)
}

func New(tp elastictransport.Interface) *AsyncQueryDelete { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryDelete) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryDelete) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryDelete) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryDelete) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *AsyncQueryDelete) Header(key, value string) *AsyncQueryDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryDelete) _id(id string) *AsyncQueryDelete { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryDelete) ErrorTrace(errortrace bool) *AsyncQueryDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryDelete) FilterPath(filterpaths ...string) *AsyncQueryDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryDelete) Human(human bool) *AsyncQueryDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryDelete) Pretty(pretty bool) *AsyncQueryDelete {
	_ = "STUB: not implemented"
	return nil
}
