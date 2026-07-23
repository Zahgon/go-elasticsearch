package explainlifecycle

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

type ExplainLifecycle struct {
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

type NewExplainLifecycle func(index string) *ExplainLifecycle

func NewExplainLifecycleFunc(tp elastictransport.Interface) NewExplainLifecycle {
	_ = "STUB: not implemented"
	return *new(NewExplainLifecycle)
}

func New(tp elastictransport.Interface) *ExplainLifecycle { _ = "STUB: not implemented"; return nil }

func (r *ExplainLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExplainLifecycle) Header(key, value string) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) _index(index string) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) OnlyErrors(onlyerrors bool) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) OnlyManaged(onlymanaged bool) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) MasterTimeout(duration string) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) ErrorTrace(errortrace bool) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) FilterPath(filterpaths ...string) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) Human(human bool) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainLifecycle) Pretty(pretty bool) *ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}
