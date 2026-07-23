package reindexrethrottle

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/groupby"
)

const (
	taskidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ReindexRethrottle struct {
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

type NewReindexRethrottle func(taskid string) *ReindexRethrottle

func NewReindexRethrottleFunc(tp elastictransport.Interface) NewReindexRethrottle {
	_ = "STUB: not implemented"
	return *new(NewReindexRethrottle)
}

func New(tp elastictransport.Interface) *ReindexRethrottle { _ = "STUB: not implemented"; return nil }

func (r *ReindexRethrottle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReindexRethrottle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReindexRethrottle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReindexRethrottle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ReindexRethrottle) Header(key, value string) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) _taskid(taskid string) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) RequestsPerSecond(requestspersecond string) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) GroupBy(groupby groupby.GroupBy) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) ErrorTrace(errortrace bool) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) FilterPath(filterpaths ...string) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) Human(human bool) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReindexRethrottle) Pretty(pretty bool) *ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}
