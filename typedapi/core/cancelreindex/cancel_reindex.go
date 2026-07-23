package cancelreindex

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

type CancelReindex struct {
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

type NewCancelReindex func(taskid string) *CancelReindex

func NewCancelReindexFunc(tp elastictransport.Interface) NewCancelReindex {
	_ = "STUB: not implemented"
	return *new(NewCancelReindex)
}

func New(tp elastictransport.Interface) *CancelReindex { _ = "STUB: not implemented"; return nil }

func (r *CancelReindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CancelReindex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CancelReindex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CancelReindex) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CancelReindex) Header(key, value string) *CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelReindex) _taskid(taskid string) *CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelReindex) WaitForCompletion(waitforcompletion bool) *CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelReindex) ErrorTrace(errortrace bool) *CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelReindex) FilterPath(filterpaths ...string) *CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelReindex) Human(human bool) *CancelReindex { _ = "STUB: not implemented"; return nil }

func (r *CancelReindex) Pretty(pretty bool) *CancelReindex { _ = "STUB: not implemented"; return nil }
