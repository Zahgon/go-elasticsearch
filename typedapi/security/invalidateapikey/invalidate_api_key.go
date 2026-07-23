package invalidateapikey

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

type InvalidateApiKey struct {
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

type NewInvalidateApiKey func() *InvalidateApiKey

func NewInvalidateApiKeyFunc(tp elastictransport.Interface) NewInvalidateApiKey {
	_ = "STUB: not implemented"
	return *new(NewInvalidateApiKey)
}

func New(tp elastictransport.Interface) *InvalidateApiKey { _ = "STUB: not implemented"; return nil }

func (r *InvalidateApiKey) Raw(raw io.Reader) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Request(req *Request) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r InvalidateApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r InvalidateApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *InvalidateApiKey) Header(key, value string) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) ErrorTrace(errortrace bool) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) FilterPath(filterpaths ...string) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Human(human bool) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Pretty(pretty bool) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Id(id string) *InvalidateApiKey { _ = "STUB: not implemented"; return nil }

func (r *InvalidateApiKey) Ids(ids ...string) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Name(name string) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Owner(owner bool) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) RealmName(realmname string) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateApiKey) Username(username string) *InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}
