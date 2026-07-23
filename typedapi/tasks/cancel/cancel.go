package cancel

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

type Cancel struct {
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

type NewCancel func() *Cancel

func NewCancelFunc(tp elastictransport.Interface) NewCancel {
	_ = "STUB: not implemented"
	return *new(NewCancel)
}

func New(tp elastictransport.Interface) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Cancel) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Cancel) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Cancel) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Cancel) Header(key, value string) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) TaskId(taskid string) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) Actions(actions ...string) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) Nodes(nodes ...string) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) ParentTaskId(parenttaskid string) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) WaitForCompletion(waitforcompletion bool) *Cancel {
	_ = "STUB: not implemented"
	return nil
}

func (r *Cancel) ErrorTrace(errortrace bool) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) FilterPath(filterpaths ...string) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) Human(human bool) *Cancel { _ = "STUB: not implemented"; return nil }

func (r *Cancel) Pretty(pretty bool) *Cancel { _ = "STUB: not implemented"; return nil }
