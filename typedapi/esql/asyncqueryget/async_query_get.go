package asyncqueryget

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/esqlformat"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AsyncQueryGet struct {
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

type NewAsyncQueryGet func(id string) *AsyncQueryGet

func NewAsyncQueryGetFunc(tp elastictransport.Interface) NewAsyncQueryGet {
	_ = "STUB: not implemented"
	return *new(NewAsyncQueryGet)
}

func New(tp elastictransport.Interface) *AsyncQueryGet { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryGet) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryGet) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQueryGet) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r AsyncQueryGet) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *AsyncQueryGet) Header(key, value string) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) _id(id string) *AsyncQueryGet { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryGet) DropNullColumns(dropnullcolumns bool) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) Format(format esqlformat.EsqlFormat) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) KeepAlive(duration string) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) WaitForCompletionTimeout(duration string) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) ErrorTrace(errortrace bool) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) FilterPath(filterpaths ...string) *AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQueryGet) Human(human bool) *AsyncQueryGet { _ = "STUB: not implemented"; return nil }

func (r *AsyncQueryGet) Pretty(pretty bool) *AsyncQueryGet { _ = "STUB: not implemented"; return nil }
