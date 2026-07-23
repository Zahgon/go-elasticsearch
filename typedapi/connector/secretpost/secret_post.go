package secretpost

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

type SecretPost struct {
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

type NewSecretPost func() *SecretPost

func NewSecretPostFunc(tp elastictransport.Interface) NewSecretPost {
	_ = "STUB: not implemented"
	return *new(NewSecretPost)
}

func New(tp elastictransport.Interface) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) Raw(raw io.Reader) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) Request(req *Request) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretPost) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretPost) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SecretPost) Header(key, value string) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) ErrorTrace(errortrace bool) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) FilterPath(filterpaths ...string) *SecretPost {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretPost) Human(human bool) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) Pretty(pretty bool) *SecretPost { _ = "STUB: not implemented"; return nil }

func (r *SecretPost) Value(value string) *SecretPost { _ = "STUB: not implemented"; return nil }
