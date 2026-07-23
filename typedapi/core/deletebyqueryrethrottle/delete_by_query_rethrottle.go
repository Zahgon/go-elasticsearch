package deletebyqueryrethrottle

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

type DeleteByQueryRethrottle struct {
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

type NewDeleteByQueryRethrottle func(taskid string) *DeleteByQueryRethrottle

func NewDeleteByQueryRethrottleFunc(tp elastictransport.Interface) NewDeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return *new(NewDeleteByQueryRethrottle)
}

func New(tp elastictransport.Interface) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteByQueryRethrottle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteByQueryRethrottle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteByQueryRethrottle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteByQueryRethrottle) Header(key, value string) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) _taskid(taskid string) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) RequestsPerSecond(requestspersecond string) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) ErrorTrace(errortrace bool) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) FilterPath(filterpaths ...string) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) Human(human bool) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQueryRethrottle) Pretty(pretty bool) *DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}
