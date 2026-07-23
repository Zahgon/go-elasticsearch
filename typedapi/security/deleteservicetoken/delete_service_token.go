package deleteservicetoken

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	namespaceMask = iota + 1

	serviceMask

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteServiceToken struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	namespace string
	service   string
	name      string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteServiceToken func(namespace, service, name string) *DeleteServiceToken

func NewDeleteServiceTokenFunc(tp elastictransport.Interface) NewDeleteServiceToken {
	_ = "STUB: not implemented"
	return *new(NewDeleteServiceToken)
}

func New(tp elastictransport.Interface) *DeleteServiceToken { _ = "STUB: not implemented"; return nil }

func (r *DeleteServiceToken) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteServiceToken) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteServiceToken) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteServiceToken) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteServiceToken) Header(key, value string) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) _namespace(namespace string) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) _service(service string) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) _name(name string) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) Refresh(refresh refresh.Refresh) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) ErrorTrace(errortrace bool) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) FilterPath(filterpaths ...string) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) Human(human bool) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteServiceToken) Pretty(pretty bool) *DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}
