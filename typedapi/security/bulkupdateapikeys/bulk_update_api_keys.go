package bulkupdateapikeys

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

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type BulkUpdateApiKeys struct {
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

type NewBulkUpdateApiKeys func() *BulkUpdateApiKeys

func NewBulkUpdateApiKeysFunc(tp elastictransport.Interface) NewBulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return *new(NewBulkUpdateApiKeys)
}

func New(tp elastictransport.Interface) *BulkUpdateApiKeys { _ = "STUB: not implemented"; return nil }

func (r *BulkUpdateApiKeys) Raw(raw io.Reader) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) Request(req *Request) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r BulkUpdateApiKeys) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r BulkUpdateApiKeys) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *BulkUpdateApiKeys) Header(key, value string) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) ErrorTrace(errortrace bool) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) FilterPath(filterpaths ...string) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) Human(human bool) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) Pretty(pretty bool) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) Expiration(duration types.DurationVariant) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) Ids(ids ...string) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) Metadata(metadata types.MetadataVariant) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkUpdateApiKeys) AddRoleDescriptor(key string, value types.RoleDescriptorVariant) *BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}
