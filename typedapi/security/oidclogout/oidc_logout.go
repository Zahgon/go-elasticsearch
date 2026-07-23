package oidclogout

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

type OidcLogout struct {
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

type NewOidcLogout func() *OidcLogout

func NewOidcLogoutFunc(tp elastictransport.Interface) NewOidcLogout {
	_ = "STUB: not implemented"
	return *new(NewOidcLogout)
}

func New(tp elastictransport.Interface) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) Raw(raw io.Reader) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) Request(req *Request) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OidcLogout) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OidcLogout) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OidcLogout) Header(key, value string) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) ErrorTrace(errortrace bool) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) FilterPath(filterpaths ...string) *OidcLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcLogout) Human(human bool) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) Pretty(pretty bool) *OidcLogout { _ = "STUB: not implemented"; return nil }

func (r *OidcLogout) RefreshToken(refreshtoken string) *OidcLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *OidcLogout) Token(token string) *OidcLogout { _ = "STUB: not implemented"; return nil }
