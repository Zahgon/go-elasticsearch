package createservicetoken

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

type CreateServiceToken struct {
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

type NewCreateServiceToken func(namespace, service string) *CreateServiceToken

func NewCreateServiceTokenFunc(tp elastictransport.Interface) NewCreateServiceToken {
	_ = "STUB: not implemented"
	return *new(NewCreateServiceToken)
}

func New(tp elastictransport.Interface) *CreateServiceToken { _ = "STUB: not implemented"; return nil }

func (r *CreateServiceToken) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateServiceToken) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateServiceToken) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateServiceToken) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CreateServiceToken) Header(key, value string) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) _namespace(namespace string) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) _service(service string) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) Name(name string) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) Refresh(refresh refresh.Refresh) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) ErrorTrace(errortrace bool) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) FilterPath(filterpaths ...string) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) Human(human bool) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateServiceToken) Pretty(pretty bool) *CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}
