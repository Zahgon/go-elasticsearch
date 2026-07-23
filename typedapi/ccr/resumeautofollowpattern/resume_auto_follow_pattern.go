package resumeautofollowpattern

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

type ResumeAutoFollowPattern struct {
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

type NewResumeAutoFollowPattern func(name string) *ResumeAutoFollowPattern

func NewResumeAutoFollowPatternFunc(tp elastictransport.Interface) NewResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(NewResumeAutoFollowPattern)
}

func New(tp elastictransport.Interface) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResumeAutoFollowPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResumeAutoFollowPattern) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResumeAutoFollowPattern) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ResumeAutoFollowPattern) Header(key, value string) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) _name(name string) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) MasterTimeout(duration string) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) ErrorTrace(errortrace bool) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) FilterPath(filterpaths ...string) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) Human(human bool) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeAutoFollowPattern) Pretty(pretty bool) *ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}
