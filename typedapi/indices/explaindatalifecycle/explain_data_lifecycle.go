package explaindatalifecycle

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

type ExplainDataLifecycle struct {
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

type NewExplainDataLifecycle func(index string) *ExplainDataLifecycle

func NewExplainDataLifecycleFunc(tp elastictransport.Interface) NewExplainDataLifecycle {
	_ = "STUB: not implemented"
	return *new(NewExplainDataLifecycle)
}

func New(tp elastictransport.Interface) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainDataLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainDataLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainDataLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExplainDataLifecycle) Header(key, value string) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) _index(index string) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) IncludeDefaults(includedefaults bool) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) MasterTimeout(duration string) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) ErrorTrace(errortrace bool) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) FilterPath(filterpaths ...string) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) Human(human bool) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataLifecycle) Pretty(pretty bool) *ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}
