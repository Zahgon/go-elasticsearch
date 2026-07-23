package getquery

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

type GetQuery struct {
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

type NewGetQuery func(id string) *GetQuery

func NewGetQueryFunc(tp elastictransport.Interface) NewGetQuery {
	_ = "STUB: not implemented"
	return *new(NewGetQuery)
}

func New(tp elastictransport.Interface) *GetQuery { _ = "STUB: not implemented"; return nil }

func (r *GetQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetQuery) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetQuery) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetQuery) Header(key, value string) *GetQuery { _ = "STUB: not implemented"; return nil }

func (r *GetQuery) _id(id string) *GetQuery { _ = "STUB: not implemented"; return nil }

func (r *GetQuery) ErrorTrace(errortrace bool) *GetQuery { _ = "STUB: not implemented"; return nil }

func (r *GetQuery) FilterPath(filterpaths ...string) *GetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetQuery) Human(human bool) *GetQuery { _ = "STUB: not implemented"; return nil }

func (r *GetQuery) Pretty(pretty bool) *GetQuery { _ = "STUB: not implemented"; return nil }
