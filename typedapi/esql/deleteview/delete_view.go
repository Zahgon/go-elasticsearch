package deleteview

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

type DeleteView struct {
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

type NewDeleteView func(name string) *DeleteView

func NewDeleteViewFunc(tp elastictransport.Interface) NewDeleteView {
	_ = "STUB: not implemented"
	return *new(NewDeleteView)
}

func New(tp elastictransport.Interface) *DeleteView { _ = "STUB: not implemented"; return nil }

func (r *DeleteView) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteView) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteView) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteView) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteView) Header(key, value string) *DeleteView { _ = "STUB: not implemented"; return nil }

func (r *DeleteView) _name(name string) *DeleteView { _ = "STUB: not implemented"; return nil }

func (r *DeleteView) ErrorTrace(errortrace bool) *DeleteView { _ = "STUB: not implemented"; return nil }

func (r *DeleteView) FilterPath(filterpaths ...string) *DeleteView {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteView) Human(human bool) *DeleteView { _ = "STUB: not implemented"; return nil }

func (r *DeleteView) Pretty(pretty bool) *DeleteView { _ = "STUB: not implemented"; return nil }
