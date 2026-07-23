package deleterouting

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteRouting struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteRouting func(name string) *DeleteRouting

func NewDeleteRoutingFunc(tp elastictransport.Interface) NewDeleteRouting {
	_ = "STUB: not implemented"
	return *new(NewDeleteRouting)
}

func New(tp elastictransport.Interface) *DeleteRouting { _ = "STUB: not implemented"; return nil }

func (r *DeleteRouting) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRouting) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRouting) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRouting) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteRouting) Header(key, value string) *DeleteRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRouting) _name(name string) *DeleteRouting { _ = "STUB: not implemented"; return nil }

func (r *DeleteRouting) ErrorTrace(errortrace bool) *DeleteRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRouting) FilterPath(filterpaths ...string) *DeleteRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRouting) Human(human bool) *DeleteRouting { _ = "STUB: not implemented"; return nil }

func (r *DeleteRouting) Pretty(pretty bool) *DeleteRouting { _ = "STUB: not implemented"; return nil }
