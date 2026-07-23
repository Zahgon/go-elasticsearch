package asyncquerystop

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

type AsyncQueryStop struct {
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

type NewAsyncQueryStop func(id string) *AsyncQueryStop

func NewAsyncQueryStopFunc(tp elastictransport.Interface) NewAsyncQueryStop {
	_ = "STUB: not implemented"
	return *new(NewAsyncQueryStop)
}

func New(tp elastictransport.Interface) *AsyncQueryStop { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryStop) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryStop) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryStop) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r AsyncQueryStop) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *AsyncQueryStop) Header(key, value string) *AsyncQueryStop {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryStop) _id(id string) *AsyncQueryStop { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryStop) DropNullColumns(dropnullcolumns bool) *AsyncQueryStop {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryStop) ErrorTrace(errortrace bool) *AsyncQueryStop {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryStop) FilterPath(filterpaths ...string) *AsyncQueryStop {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryStop) Human(human bool) *AsyncQueryStop { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryStop) Pretty(pretty bool) *AsyncQueryStop { _ = "STUB: not implemented"; return nil }
