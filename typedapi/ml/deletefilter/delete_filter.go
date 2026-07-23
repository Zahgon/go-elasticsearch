package deletefilter

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

type DeleteFilter struct {
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

type NewDeleteFilter func(filterid string) *DeleteFilter

func NewDeleteFilterFunc(tp elastictransport.Interface) NewDeleteFilter {
	_ = "STUB: not implemented"
	return *new(NewDeleteFilter)
}

func New(tp elastictransport.Interface) *DeleteFilter { _ = "STUB: not implemented"; return nil }

func (r *DeleteFilter) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteFilter) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteFilter) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteFilter) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteFilter) Header(key, value string) *DeleteFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteFilter) _filterid(filterid string) *DeleteFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteFilter) ErrorTrace(errortrace bool) *DeleteFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteFilter) FilterPath(filterpaths ...string) *DeleteFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteFilter) Human(human bool) *DeleteFilter { _ = "STUB: not implemented"; return nil }

func (r *DeleteFilter) Pretty(pretty bool) *DeleteFilter { _ = "STUB: not implemented"; return nil }
