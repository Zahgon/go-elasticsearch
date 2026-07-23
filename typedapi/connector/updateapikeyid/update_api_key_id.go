package updateapikeyid

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
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateApiKeyId struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateApiKeyId func(connectorid string) *UpdateApiKeyId

func NewUpdateApiKeyIdFunc(tp elastictransport.Interface) NewUpdateApiKeyId {
	_ = "STUB: not implemented"
	return *new(NewUpdateApiKeyId)
}

func New(tp elastictransport.Interface) *UpdateApiKeyId { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKeyId) Raw(raw io.Reader) *UpdateApiKeyId { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKeyId) Request(req *Request) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKeyId) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateApiKeyId) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateApiKeyId) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateApiKeyId) Header(key, value string) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKeyId) _connectorid(connectorid string) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKeyId) ErrorTrace(errortrace bool) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKeyId) FilterPath(filterpaths ...string) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKeyId) Human(human bool) *UpdateApiKeyId { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKeyId) Pretty(pretty bool) *UpdateApiKeyId { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKeyId) ApiKeyId(apikeyid string) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKeyId) ApiKeySecretId(apikeysecretid string) *UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}
