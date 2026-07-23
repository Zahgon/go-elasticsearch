package getserviceaccounts

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

type GetServiceAccounts struct {
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

type NewGetServiceAccounts func() *GetServiceAccounts

func NewGetServiceAccountsFunc(tp elastictransport.Interface) NewGetServiceAccounts {
	_ = "STUB: not implemented"
	return *new(NewGetServiceAccounts)
}

func New(tp elastictransport.Interface) *GetServiceAccounts { _ = "STUB: not implemented"; return nil }

func (r *GetServiceAccounts) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetServiceAccounts) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetServiceAccounts) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetServiceAccounts) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetServiceAccounts) Header(key, value string) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceAccounts) Namespace(namespace string) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceAccounts) Service(service string) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceAccounts) ErrorTrace(errortrace bool) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceAccounts) FilterPath(filterpaths ...string) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceAccounts) Human(human bool) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetServiceAccounts) Pretty(pretty bool) *GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}
