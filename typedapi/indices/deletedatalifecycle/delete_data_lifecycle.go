package deletedatalifecycle

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

type DeleteDataLifecycle struct {
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

type NewDeleteDataLifecycle func(name string) *DeleteDataLifecycle

func NewDeleteDataLifecycleFunc(tp elastictransport.Interface) NewDeleteDataLifecycle {
	_ = "STUB: not implemented"
	return *new(NewDeleteDataLifecycle)
}

func New(tp elastictransport.Interface) *DeleteDataLifecycle { _ = "STUB: not implemented"; return nil }

func (r *DeleteDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteDataLifecycle) Header(key, value string) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) _name(name string) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) MasterTimeout(duration string) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) Timeout(duration string) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) ErrorTrace(errortrace bool) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) FilterPath(filterpaths ...string) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) Human(human bool) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataLifecycle) Pretty(pretty bool) *DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}
