package bulkdeleterole

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type BulkDeleteRole struct {
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

type NewBulkDeleteRole func() *BulkDeleteRole

func NewBulkDeleteRoleFunc(tp elastictransport.Interface) NewBulkDeleteRole {
	_ = "STUB: not implemented"
	return *new(NewBulkDeleteRole)
}

func New(tp elastictransport.Interface) *BulkDeleteRole { _ = "STUB: not implemented"; return nil }

func (r *BulkDeleteRole) Raw(raw io.Reader) *BulkDeleteRole { _ = "STUB: not implemented"; return nil }

func (r *BulkDeleteRole) Request(req *Request) *BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkDeleteRole) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r BulkDeleteRole) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r BulkDeleteRole) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *BulkDeleteRole) Header(key, value string) *BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkDeleteRole) Refresh(refresh refresh.Refresh) *BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkDeleteRole) ErrorTrace(errortrace bool) *BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkDeleteRole) FilterPath(filterpaths ...string) *BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *BulkDeleteRole) Human(human bool) *BulkDeleteRole { _ = "STUB: not implemented"; return nil }

func (r *BulkDeleteRole) Pretty(pretty bool) *BulkDeleteRole { _ = "STUB: not implemented"; return nil }

func (r *BulkDeleteRole) Names(names ...string) *BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}
