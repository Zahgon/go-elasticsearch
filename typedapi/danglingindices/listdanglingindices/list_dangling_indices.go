package listdanglingindices

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ListDanglingIndices struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewListDanglingIndices func() *ListDanglingIndices

func NewListDanglingIndicesFunc(tp elastictransport.Interface) NewListDanglingIndices {
	_ = "STUB: not implemented"
	return *new(NewListDanglingIndices)
}

func New(tp elastictransport.Interface) *ListDanglingIndices { _ = "STUB: not implemented"; return nil }

func (r *ListDanglingIndices) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListDanglingIndices) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListDanglingIndices) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListDanglingIndices) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ListDanglingIndices) Header(key, value string) *ListDanglingIndices {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListDanglingIndices) ErrorTrace(errortrace bool) *ListDanglingIndices {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListDanglingIndices) FilterPath(filterpaths ...string) *ListDanglingIndices {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListDanglingIndices) Human(human bool) *ListDanglingIndices {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListDanglingIndices) Pretty(pretty bool) *ListDanglingIndices {
	_ = "STUB: not implemented"
	return nil
}
