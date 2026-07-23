package samlcompletelogout

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

type SamlCompleteLogout struct {
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

type NewSamlCompleteLogout func() *SamlCompleteLogout

func NewSamlCompleteLogoutFunc(tp elastictransport.Interface) NewSamlCompleteLogout {
	_ = "STUB: not implemented"
	return *new(NewSamlCompleteLogout)
}

func New(tp elastictransport.Interface) *SamlCompleteLogout { _ = "STUB: not implemented"; return nil }

func (r *SamlCompleteLogout) Raw(raw io.Reader) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) Request(req *Request) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlCompleteLogout) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SamlCompleteLogout) Header(key, value string) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) ErrorTrace(errortrace bool) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) FilterPath(filterpaths ...string) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) Human(human bool) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) Pretty(pretty bool) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) Content(content string) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) Ids(ids ...string) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) QueryString(querystring string) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlCompleteLogout) Realm(realm string) *SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}
