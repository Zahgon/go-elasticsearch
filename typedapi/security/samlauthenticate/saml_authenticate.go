package samlauthenticate

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

type SamlAuthenticate struct {
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

type NewSamlAuthenticate func() *SamlAuthenticate

func NewSamlAuthenticateFunc(tp elastictransport.Interface) NewSamlAuthenticate {
	_ = "STUB: not implemented"
	return *new(NewSamlAuthenticate)
}

func New(tp elastictransport.Interface) *SamlAuthenticate { _ = "STUB: not implemented"; return nil }

func (r *SamlAuthenticate) Raw(raw io.Reader) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) Request(req *Request) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlAuthenticate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlAuthenticate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SamlAuthenticate) Header(key, value string) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) ErrorTrace(errortrace bool) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) FilterPath(filterpaths ...string) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) Human(human bool) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) Pretty(pretty bool) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) Content(content string) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) Ids(ids ...string) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlAuthenticate) Realm(realm string) *SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}
