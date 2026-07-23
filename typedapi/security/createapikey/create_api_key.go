package createapikey

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateApiKey struct {
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

type NewCreateApiKey func() *CreateApiKey

func NewCreateApiKeyFunc(tp elastictransport.Interface) NewCreateApiKey {
	_ = "STUB: not implemented"
	return *new(NewCreateApiKey)
}

func New(tp elastictransport.Interface) *CreateApiKey { _ = "STUB: not implemented"; return nil }

func (r *CreateApiKey) Raw(raw io.Reader) *CreateApiKey { _ = "STUB: not implemented"; return nil }

func (r *CreateApiKey) Request(req *Request) *CreateApiKey { _ = "STUB: not implemented"; return nil }

func (r *CreateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CreateApiKey) Header(key, value string) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) Refresh(refresh refresh.Refresh) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) ErrorTrace(errortrace bool) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) FilterPath(filterpaths ...string) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) Human(human bool) *CreateApiKey { _ = "STUB: not implemented"; return nil }

func (r *CreateApiKey) Pretty(pretty bool) *CreateApiKey { _ = "STUB: not implemented"; return nil }

func (r *CreateApiKey) Expiration(duration types.DurationVariant) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) Metadata(metadata types.MetadataVariant) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) Name(name string) *CreateApiKey { _ = "STUB: not implemented"; return nil }

func (r *CreateApiKey) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateApiKey) AddRoleDescriptor(key string, value types.RoleDescriptorVariant) *CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}
