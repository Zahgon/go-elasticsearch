package updateactivefiltering

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateActiveFiltering struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateActiveFiltering func(connectorid string) *UpdateActiveFiltering

func NewUpdateActiveFilteringFunc(tp elastictransport.Interface) NewUpdateActiveFiltering {
	_ = "STUB: not implemented"
	return *new(NewUpdateActiveFiltering)
}

func New(tp elastictransport.Interface) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateActiveFiltering) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateActiveFiltering) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateActiveFiltering) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateActiveFiltering) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *UpdateActiveFiltering) Header(key, value string) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateActiveFiltering) _connectorid(connectorid string) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateActiveFiltering) ErrorTrace(errortrace bool) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateActiveFiltering) FilterPath(filterpaths ...string) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateActiveFiltering) Human(human bool) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateActiveFiltering) Pretty(pretty bool) *UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}
