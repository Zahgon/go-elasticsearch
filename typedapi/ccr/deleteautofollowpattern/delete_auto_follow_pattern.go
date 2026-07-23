package deleteautofollowpattern

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

type DeleteAutoFollowPattern struct {
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

type NewDeleteAutoFollowPattern func(name string) *DeleteAutoFollowPattern

func NewDeleteAutoFollowPatternFunc(tp elastictransport.Interface) NewDeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(NewDeleteAutoFollowPattern)
}

func New(tp elastictransport.Interface) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAutoFollowPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAutoFollowPattern) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAutoFollowPattern) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteAutoFollowPattern) Header(key, value string) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) _name(name string) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) MasterTimeout(duration string) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) ErrorTrace(errortrace bool) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) FilterPath(filterpaths ...string) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) Human(human bool) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoFollowPattern) Pretty(pretty bool) *DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}
