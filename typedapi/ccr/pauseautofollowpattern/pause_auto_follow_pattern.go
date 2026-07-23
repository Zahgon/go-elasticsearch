package pauseautofollowpattern

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

type PauseAutoFollowPattern struct {
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

type NewPauseAutoFollowPattern func(name string) *PauseAutoFollowPattern

func NewPauseAutoFollowPatternFunc(tp elastictransport.Interface) NewPauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(NewPauseAutoFollowPattern)
}

func New(tp elastictransport.Interface) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PauseAutoFollowPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PauseAutoFollowPattern) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PauseAutoFollowPattern) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PauseAutoFollowPattern) Header(key, value string) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) _name(name string) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) MasterTimeout(duration string) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) ErrorTrace(errortrace bool) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) FilterPath(filterpaths ...string) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) Human(human bool) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseAutoFollowPattern) Pretty(pretty bool) *PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}
