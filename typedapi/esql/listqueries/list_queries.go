package listqueries

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ListQueries struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewListQueries func() *ListQueries

func NewListQueriesFunc(tp elastictransport.Interface) NewListQueries {
	_ = "STUB: not implemented"
	return *new(NewListQueries)
}

func New(tp elastictransport.Interface) *ListQueries { _ = "STUB: not implemented"; return nil }

func (r *ListQueries) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListQueries) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListQueries) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListQueries) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ListQueries) Header(key, value string) *ListQueries { _ = "STUB: not implemented"; return nil }

func (r *ListQueries) ErrorTrace(errortrace bool) *ListQueries {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListQueries) FilterPath(filterpaths ...string) *ListQueries {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListQueries) Human(human bool) *ListQueries { _ = "STUB: not implemented"; return nil }

func (r *ListQueries) Pretty(pretty bool) *ListQueries { _ = "STUB: not implemented"; return nil }
