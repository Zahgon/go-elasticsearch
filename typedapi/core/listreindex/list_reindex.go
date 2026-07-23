package listreindex

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ListReindex struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewListReindex func() *ListReindex

func NewListReindexFunc(tp elastictransport.Interface) NewListReindex {
	_ = "STUB: not implemented"
	return *new(NewListReindex)
}

func New(tp elastictransport.Interface) *ListReindex { _ = "STUB: not implemented"; return nil }

func (r *ListReindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListReindex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListReindex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListReindex) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ListReindex) Header(key, value string) *ListReindex { _ = "STUB: not implemented"; return nil }

func (r *ListReindex) Detailed(detailed bool) *ListReindex { _ = "STUB: not implemented"; return nil }

func (r *ListReindex) ErrorTrace(errortrace bool) *ListReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListReindex) FilterPath(filterpaths ...string) *ListReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListReindex) Human(human bool) *ListReindex { _ = "STUB: not implemented"; return nil }

func (r *ListReindex) Pretty(pretty bool) *ListReindex { _ = "STUB: not implemented"; return nil }
