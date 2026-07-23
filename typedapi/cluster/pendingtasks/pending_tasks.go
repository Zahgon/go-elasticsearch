package pendingtasks

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PendingTasks struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPendingTasks func() *PendingTasks

func NewPendingTasksFunc(tp elastictransport.Interface) NewPendingTasks {
	_ = "STUB: not implemented"
	return *new(NewPendingTasks)
}

func New(tp elastictransport.Interface) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PendingTasks) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PendingTasks) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PendingTasks) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PendingTasks) Header(key, value string) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) Local(local bool) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) MasterTimeout(duration string) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) ErrorTrace(errortrace bool) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) FilterPath(filterpaths ...string) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) Human(human bool) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) Pretty(pretty bool) *PendingTasks { _ = "STUB: not implemented"; return nil }
