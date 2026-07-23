package getdatalifecycle

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDataLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetDataLifecycle func(name string) *GetDataLifecycle

func NewGetDataLifecycleFunc(tp elastictransport.Interface) NewGetDataLifecycle {
	_ = "STUB: not implemented"
	return *new(NewGetDataLifecycle)
}

func New(tp elastictransport.Interface) *GetDataLifecycle { _ = "STUB: not implemented"; return nil }

func (r *GetDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataLifecycle) Header(key, value string) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) _name(name string) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) IncludeDefaults(includedefaults bool) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) MasterTimeout(duration string) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) ErrorTrace(errortrace bool) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) FilterPath(filterpaths ...string) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) Human(human bool) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycle) Pretty(pretty bool) *GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}
