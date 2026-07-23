package queryuser

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

type QueryUser struct {
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

type NewQueryUser func() *QueryUser

func NewQueryUserFunc(tp elastictransport.Interface) NewQueryUser {
	_ = "STUB: not implemented"
	return *new(NewQueryUser)
}

func New(tp elastictransport.Interface) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) Raw(raw io.Reader) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) Request(req *Request) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryUser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryUser) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *QueryUser) Header(key, value string) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) WithProfileUid(withprofileuid bool) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryUser) ErrorTrace(errortrace bool) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) FilterPath(filterpaths ...string) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryUser) Human(human bool) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) Pretty(pretty bool) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) From(from int) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) Query(query types.UserQueryContainerVariant) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryUser) SearchAfter(sortresults ...types.FieldValueVariant) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryUser) SearchAfterValues(sortresultsvalues []types.FieldValue) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryUser) Size(size int) *QueryUser { _ = "STUB: not implemented"; return nil }

func (r *QueryUser) Sort(sorts ...types.SortCombinationsVariant) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryUser) SortValues(sortvalues []types.SortCombinations) *QueryUser {
	_ = "STUB: not implemented"
	return nil
}
