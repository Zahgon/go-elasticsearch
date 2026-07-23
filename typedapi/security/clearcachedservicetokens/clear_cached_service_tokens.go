package clearcachedservicetokens

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

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearCachedServiceTokens struct {
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

type NewClearCachedServiceTokens func(namespace, service, name string) *ClearCachedServiceTokens

func NewClearCachedServiceTokensFunc(tp elastictransport.Interface) NewClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return *new(NewClearCachedServiceTokens)
}

func New(tp elastictransport.Interface) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedServiceTokens) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedServiceTokens) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedServiceTokens) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearCachedServiceTokens) Header(key, value string) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) _namespace(namespace string) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) _service(service string) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) _name(name string) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) ErrorTrace(errortrace bool) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) FilterPath(filterpaths ...string) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) Human(human bool) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedServiceTokens) Pretty(pretty bool) *ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}
