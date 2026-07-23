package getautofollowpattern

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

type GetAutoFollowPattern struct {
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

type NewGetAutoFollowPattern func() *GetAutoFollowPattern

func NewGetAutoFollowPatternFunc(tp elastictransport.Interface) NewGetAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(NewGetAutoFollowPattern)
}

func New(tp elastictransport.Interface) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoFollowPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoFollowPattern) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoFollowPattern) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetAutoFollowPattern) Header(key, value string) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) Name(name string) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) MasterTimeout(duration string) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) ErrorTrace(errortrace bool) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) FilterPath(filterpaths ...string) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) Human(human bool) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoFollowPattern) Pretty(pretty bool) *GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}
