package getservicecredentials

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	namespaceMask = iota + 1

	serviceMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetServiceCredentials struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	namespace string
	service   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetServiceCredentials func(namespace, service string) *GetServiceCredentials

func NewGetServiceCredentialsFunc(tp elastictransport.Interface) NewGetServiceCredentials {
	_ = "STUB: not implemented"
	return *new(NewGetServiceCredentials)
}

func New(tp elastictransport.Interface) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetServiceCredentials) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetServiceCredentials) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetServiceCredentials) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetServiceCredentials) Header(key, value string) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) _namespace(namespace string) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) _service(service string) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) ErrorTrace(errortrace bool) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) FilterPath(filterpaths ...string) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) Human(human bool) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceCredentials) Pretty(pretty bool) *GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}
