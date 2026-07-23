package samllogout

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

type SamlLogout struct {
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

type NewSamlLogout func() *SamlLogout

func NewSamlLogoutFunc(tp elastictransport.Interface) NewSamlLogout {
	_ = "STUB: not implemented"
	return *new(NewSamlLogout)
}

func New(tp elastictransport.Interface) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) Raw(raw io.Reader) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) Request(req *Request) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlLogout) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlLogout) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SamlLogout) Header(key, value string) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) ErrorTrace(errortrace bool) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) FilterPath(filterpaths ...string) *SamlLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlLogout) Human(human bool) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) Pretty(pretty bool) *SamlLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlLogout) RefreshToken(refreshtoken string) *SamlLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlLogout) Token(token string) *SamlLogout { _ = "STUB: not implemented"; return nil }
