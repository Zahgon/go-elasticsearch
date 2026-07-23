package getreindex

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

type GetReindex struct {
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

type NewGetReindex func(taskid string) *GetReindex

func NewGetReindexFunc(tp elastictransport.Interface) NewGetReindex {
	_ = "STUB: not implemented"
	return *new(NewGetReindex)
}

func New(tp elastictransport.Interface) *GetReindex { _ = "STUB: not implemented"; return nil }

func (r *GetReindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetReindex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetReindex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetReindex) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetReindex) Header(key, value string) *GetReindex { _ = "STUB: not implemented"; return nil }

func (r *GetReindex) _taskid(taskid string) *GetReindex { _ = "STUB: not implemented"; return nil }

func (r *GetReindex) WaitForCompletion(waitforcompletion bool) *GetReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetReindex) Timeout(duration string) *GetReindex { _ = "STUB: not implemented"; return nil }

func (r *GetReindex) ErrorTrace(errortrace bool) *GetReindex { _ = "STUB: not implemented"; return nil }

func (r *GetReindex) FilterPath(filterpaths ...string) *GetReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetReindex) Human(human bool) *GetReindex { _ = "STUB: not implemented"; return nil }

func (r *GetReindex) Pretty(pretty bool) *GetReindex { _ = "STUB: not implemented"; return nil }
