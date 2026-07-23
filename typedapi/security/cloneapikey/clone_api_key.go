package cloneapikey

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

type CloneApiKey struct {
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

type NewCloneApiKey func() *CloneApiKey

func NewCloneApiKeyFunc(tp elastictransport.Interface) NewCloneApiKey {
	_ = "STUB: not implemented"
	return *new(NewCloneApiKey)
}

func New(tp elastictransport.Interface) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) Raw(raw io.Reader) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) Request(req *Request) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CloneApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CloneApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CloneApiKey) Header(key, value string) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) Refresh(refresh refresh.Refresh) *CloneApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CloneApiKey) ErrorTrace(errortrace bool) *CloneApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CloneApiKey) FilterPath(filterpaths ...string) *CloneApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CloneApiKey) Human(human bool) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) Pretty(pretty bool) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) ApiKey(apikey string) *CloneApiKey { _ = "STUB: not implemented"; return nil }

func (r *CloneApiKey) Expiration(duration types.DurationVariant) *CloneApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CloneApiKey) Metadata(metadata types.MetadataVariant) *CloneApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *CloneApiKey) Name(name string) *CloneApiKey { _ = "STUB: not implemented"; return nil }
