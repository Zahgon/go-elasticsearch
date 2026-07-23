package getfilters

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	filteridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetFilters struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	filterid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetFilters func() *GetFilters

func NewGetFiltersFunc(tp elastictransport.Interface) NewGetFilters {
	_ = "STUB: not implemented"
	return *new(NewGetFilters)
}

func New(tp elastictransport.Interface) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFilters) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFilters) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFilters) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetFilters) Header(key, value string) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) FilterId(filterid string) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) From(from int) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) Size(size int) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) ErrorTrace(errortrace bool) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) FilterPath(filterpaths ...string) *GetFilters {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFilters) Human(human bool) *GetFilters { _ = "STUB: not implemented"; return nil }

func (r *GetFilters) Pretty(pretty bool) *GetFilters { _ = "STUB: not implemented"; return nil }
