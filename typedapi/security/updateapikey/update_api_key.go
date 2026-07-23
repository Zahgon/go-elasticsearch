package updateapikey

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateApiKey struct {
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

type NewUpdateApiKey func(id string) *UpdateApiKey

func NewUpdateApiKeyFunc(tp elastictransport.Interface) NewUpdateApiKey {
	_ = "STUB: not implemented"
	return *new(NewUpdateApiKey)
}

func New(tp elastictransport.Interface) *UpdateApiKey { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKey) Raw(raw io.Reader) *UpdateApiKey { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKey) Request(req *Request) *UpdateApiKey { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateApiKey) Header(key, value string) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKey) _id(id string) *UpdateApiKey { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKey) ErrorTrace(errortrace bool) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKey) FilterPath(filterpaths ...string) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKey) Human(human bool) *UpdateApiKey { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKey) Pretty(pretty bool) *UpdateApiKey { _ = "STUB: not implemented"; return nil }

func (r *UpdateApiKey) Expiration(duration types.DurationVariant) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKey) Metadata(metadata types.MetadataVariant) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKey) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateApiKey) AddRoleDescriptor(key string, value types.RoleDescriptorVariant) *UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}
