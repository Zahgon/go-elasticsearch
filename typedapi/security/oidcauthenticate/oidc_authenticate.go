package oidcauthenticate

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

type OidcAuthenticate struct {
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

type NewOidcAuthenticate func() *OidcAuthenticate

func NewOidcAuthenticateFunc(tp elastictransport.Interface) NewOidcAuthenticate {
	_ = "STUB: not implemented"
	return *new(NewOidcAuthenticate)
}

func New(tp elastictransport.Interface) *OidcAuthenticate { _ = "STUB: not implemented"; return nil }

func (r *OidcAuthenticate) Raw(raw io.Reader) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) Request(req *Request) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OidcAuthenticate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OidcAuthenticate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OidcAuthenticate) Header(key, value string) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) ErrorTrace(errortrace bool) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) FilterPath(filterpaths ...string) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) Human(human bool) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) Pretty(pretty bool) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) Nonce(nonce string) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) Realm(realm string) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) RedirectUri(redirecturi string) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcAuthenticate) State(state string) *OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}
