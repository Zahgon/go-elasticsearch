package oidcprepareauthentication

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type OidcPrepareAuthentication struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewOidcPrepareAuthentication func() *OidcPrepareAuthentication

func NewOidcPrepareAuthenticationFunc(tp elastictransport.Interface) NewOidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return *new(NewOidcPrepareAuthentication)
}

func New(tp elastictransport.Interface) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Raw(raw io.Reader) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Request(req *Request) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OidcPrepareAuthentication) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OidcPrepareAuthentication) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OidcPrepareAuthentication) Header(key, value string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) ErrorTrace(errortrace bool) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) FilterPath(filterpaths ...string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Human(human bool) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Pretty(pretty bool) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Iss(iss string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) LoginHint(loginhint string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Nonce(nonce string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) Realm(realm string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcPrepareAuthentication) State(state string) *OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}
