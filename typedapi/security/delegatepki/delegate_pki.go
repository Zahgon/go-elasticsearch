package delegatepki

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

type DelegatePki struct {
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

type NewDelegatePki func() *DelegatePki

func NewDelegatePkiFunc(tp elastictransport.Interface) NewDelegatePki {
	_ = "STUB: not implemented"
	return *new(NewDelegatePki)
}

func New(tp elastictransport.Interface) *DelegatePki { _ = "STUB: not implemented"; return nil }

func (r *DelegatePki) Raw(raw io.Reader) *DelegatePki { _ = "STUB: not implemented"; return nil }

func (r *DelegatePki) Request(req *Request) *DelegatePki { _ = "STUB: not implemented"; return nil }

func (r *DelegatePki) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DelegatePki) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DelegatePki) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *DelegatePki) Header(key, value string) *DelegatePki { _ = "STUB: not implemented"; return nil }

func (r *DelegatePki) ErrorTrace(errortrace bool) *DelegatePki {
	_ = "STUB: not implemented"
	return nil
}

func (r *DelegatePki) FilterPath(filterpaths ...string) *DelegatePki {
	_ = "STUB: not implemented"
	return nil
}

func (r *DelegatePki) Human(human bool) *DelegatePki { _ = "STUB: not implemented"; return nil }

func (r *DelegatePki) Pretty(pretty bool) *DelegatePki { _ = "STUB: not implemented"; return nil }

func (r *DelegatePki) X509CertificateChain(x509certificatechains ...string) *DelegatePki {
	_ = "STUB: not implemented"
	return nil
}
