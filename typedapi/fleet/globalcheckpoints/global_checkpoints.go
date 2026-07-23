package globalcheckpoints

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GlobalCheckpoints struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGlobalCheckpoints func(index string) *GlobalCheckpoints

func NewGlobalCheckpointsFunc(tp elastictransport.Interface) NewGlobalCheckpoints {
	_ = "STUB: not implemented"
	return *new(NewGlobalCheckpoints)
}

func New(tp elastictransport.Interface) *GlobalCheckpoints { _ = "STUB: not implemented"; return nil }

func (r *GlobalCheckpoints) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GlobalCheckpoints) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GlobalCheckpoints) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GlobalCheckpoints) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GlobalCheckpoints) Header(key, value string) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) _index(index string) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) WaitForAdvance(waitforadvance bool) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) WaitForIndex(waitforindex bool) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) Checkpoints(checkpoints ...int64) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) Timeout(duration string) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) ErrorTrace(errortrace bool) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) FilterPath(filterpaths ...string) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) Human(human bool) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (r *GlobalCheckpoints) Pretty(pretty bool) *GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}
