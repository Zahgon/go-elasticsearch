package secretput

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SecretPut struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSecretPut func(id string) *SecretPut

func NewSecretPutFunc(tp elastictransport.Interface) NewSecretPut {
	_ = "STUB: not implemented"
	return *new(NewSecretPut)
}

func New(tp elastictransport.Interface) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) Raw(raw io.Reader) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) Request(req *Request) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretPut) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretPut) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SecretPut) Header(key, value string) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) _id(id string) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) ErrorTrace(errortrace bool) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) FilterPath(filterpaths ...string) *SecretPut {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretPut) Human(human bool) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) Pretty(pretty bool) *SecretPut { _ = "STUB: not implemented"; return nil }

func (r *SecretPut) Value(value string) *SecretPut { _ = "STUB: not implemented"; return nil }
