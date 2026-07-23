package list

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type List struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewList func() *List

func NewListFunc(tp elastictransport.Interface) NewList {
	_ = "STUB: not implemented"
	return *new(NewList)
}

func New(tp elastictransport.Interface) *List { _ = "STUB: not implemented"; return nil }

func (r *List) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r List) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r List) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r List) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *List) Header(key, value string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) From(from int) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Size(size int) *List { _ = "STUB: not implemented"; return nil }

func (r *List) IndexName(indices ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) ConnectorName(names ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) ServiceType(names ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) IncludeDeleted(includedeleted bool) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Query(query string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) ErrorTrace(errortrace bool) *List { _ = "STUB: not implemented"; return nil }

func (r *List) FilterPath(filterpaths ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Human(human bool) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Pretty(pretty bool) *List { _ = "STUB: not implemented"; return nil }
