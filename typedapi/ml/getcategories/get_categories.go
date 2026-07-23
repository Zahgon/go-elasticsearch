package getcategories

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
	jobidMask = iota + 1

	categoryidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetCategories struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid      string
	categoryid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetCategories func(jobid string) *GetCategories

func NewGetCategoriesFunc(tp elastictransport.Interface) NewGetCategories {
	_ = "STUB: not implemented"
	return *new(NewGetCategories)
}

func New(tp elastictransport.Interface) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) Raw(raw io.Reader) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) Request(req *Request) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCategories) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCategories) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetCategories) Header(key, value string) *GetCategories {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCategories) _jobid(jobid string) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) CategoryId(categoryid string) *GetCategories {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCategories) From(from int) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) PartitionFieldValue(partitionfieldvalue string) *GetCategories {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCategories) Size(size int) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) ErrorTrace(errortrace bool) *GetCategories {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCategories) FilterPath(filterpaths ...string) *GetCategories {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCategories) Human(human bool) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) Pretty(pretty bool) *GetCategories { _ = "STUB: not implemented"; return nil }

func (r *GetCategories) Page(page types.PageVariant) *GetCategories {
	_ = "STUB: not implemented"
	return nil
}
