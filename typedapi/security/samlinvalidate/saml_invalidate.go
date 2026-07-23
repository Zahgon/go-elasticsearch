package samlinvalidate

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

type SamlInvalidate struct {
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

type NewSamlInvalidate func() *SamlInvalidate

func NewSamlInvalidateFunc(tp elastictransport.Interface) NewSamlInvalidate {
	_ = "STUB: not implemented"
	return *new(NewSamlInvalidate)
}

func New(tp elastictransport.Interface) *SamlInvalidate { _ = "STUB: not implemented"; return nil }

func (r *SamlInvalidate) Raw(raw io.Reader) *SamlInvalidate { _ = "STUB: not implemented"; return nil }

func (r *SamlInvalidate) Request(req *Request) *SamlInvalidate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlInvalidate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlInvalidate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlInvalidate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SamlInvalidate) Header(key, value string) *SamlInvalidate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlInvalidate) ErrorTrace(errortrace bool) *SamlInvalidate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlInvalidate) FilterPath(filterpaths ...string) *SamlInvalidate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlInvalidate) Human(human bool) *SamlInvalidate { _ = "STUB: not implemented"; return nil }

func (r *SamlInvalidate) Pretty(pretty bool) *SamlInvalidate { _ = "STUB: not implemented"; return nil }

func (r *SamlInvalidate) Acs(acs string) *SamlInvalidate { _ = "STUB: not implemented"; return nil }

func (r *SamlInvalidate) QueryString(querystring string) *SamlInvalidate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlInvalidate) Realm(realm string) *SamlInvalidate { _ = "STUB: not implemented"; return nil }
