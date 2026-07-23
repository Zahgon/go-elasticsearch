package updatebyqueryrethrottle

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	taskidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateByQueryRethrottle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	taskid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateByQueryRethrottle func(taskid string) *UpdateByQueryRethrottle

func NewUpdateByQueryRethrottleFunc(tp elastictransport.Interface) NewUpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return *new(NewUpdateByQueryRethrottle)
}

func New(tp elastictransport.Interface) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateByQueryRethrottle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateByQueryRethrottle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateByQueryRethrottle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *UpdateByQueryRethrottle) Header(key, value string) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) _taskid(taskid string) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) RequestsPerSecond(requestspersecond string) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) ErrorTrace(errortrace bool) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) FilterPath(filterpaths ...string) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) Human(human bool) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQueryRethrottle) Pretty(pretty bool) *UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}
