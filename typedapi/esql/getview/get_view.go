package getview

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

type GetView struct {
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

type NewGetView func() *GetView

func NewGetViewFunc(tp elastictransport.Interface) NewGetView {
	_ = "STUB: not implemented"
	return *new(NewGetView)
}

func New(tp elastictransport.Interface) *GetView { _ = "STUB: not implemented"; return nil }

func (r *GetView) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetView) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetView) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetView) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetView) Header(key, value string) *GetView { _ = "STUB: not implemented"; return nil }

func (r *GetView) Name(name string) *GetView { _ = "STUB: not implemented"; return nil }

func (r *GetView) ErrorTrace(errortrace bool) *GetView { _ = "STUB: not implemented"; return nil }

func (r *GetView) FilterPath(filterpaths ...string) *GetView { _ = "STUB: not implemented"; return nil }

func (r *GetView) Human(human bool) *GetView { _ = "STUB: not implemented"; return nil }

func (r *GetView) Pretty(pretty bool) *GetView { _ = "STUB: not implemented"; return nil }
