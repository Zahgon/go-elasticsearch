package bulkputrole

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

type BulkPutRole struct {
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

type NewBulkPutRole func() *BulkPutRole

func NewBulkPutRoleFunc(tp elastictransport.Interface) NewBulkPutRole {
	_ = "STUB: not implemented"
	return *new(NewBulkPutRole)
}

func New(tp elastictransport.Interface) *BulkPutRole { _ = "STUB: not implemented"; return nil }

func (r *BulkPutRole) Raw(raw io.Reader) *BulkPutRole { _ = "STUB: not implemented"; return nil }

func (r *BulkPutRole) Request(req *Request) *BulkPutRole { _ = "STUB: not implemented"; return nil }

func (r *BulkPutRole) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r BulkPutRole) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r BulkPutRole) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *BulkPutRole) Header(key, value string) *BulkPutRole { _ = "STUB: not implemented"; return nil }

func (r *BulkPutRole) Refresh(refresh refresh.Refresh) *BulkPutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkPutRole) ErrorTrace(errortrace bool) *BulkPutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkPutRole) FilterPath(filterpaths ...string) *BulkPutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkPutRole) Human(human bool) *BulkPutRole { _ = "STUB: not implemented"; return nil }

func (r *BulkPutRole) Pretty(pretty bool) *BulkPutRole { _ = "STUB: not implemented"; return nil }

func (r *BulkPutRole) Roles(roles map[string]types.RoleDescriptor) *BulkPutRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkPutRole) AddRole(key string, value types.RoleDescriptorVariant) *BulkPutRole {
	_ = "STUB: not implemented"
	return nil
}
