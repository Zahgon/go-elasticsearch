package queryrole

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

type QueryRole struct {
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

type NewQueryRole func() *QueryRole

func NewQueryRoleFunc(tp elastictransport.Interface) NewQueryRole {
	_ = "STUB: not implemented"
	return *new(NewQueryRole)
}

func New(tp elastictransport.Interface) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) Raw(raw io.Reader) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) Request(req *Request) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryRole) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryRole) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *QueryRole) Header(key, value string) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) ErrorTrace(errortrace bool) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) FilterPath(filterpaths ...string) *QueryRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryRole) Human(human bool) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) Pretty(pretty bool) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) From(from int) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) Query(query types.RoleQueryContainerVariant) *QueryRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryRole) SearchAfter(sortresults ...types.FieldValueVariant) *QueryRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryRole) SearchAfterValues(sortresultsvalues []types.FieldValue) *QueryRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryRole) Size(size int) *QueryRole { _ = "STUB: not implemented"; return nil }

func (r *QueryRole) Sort(sorts ...types.SortCombinationsVariant) *QueryRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryRole) SortValues(sortvalues []types.SortCombinations) *QueryRole {
	_ = "STUB: not implemented"
	return nil
}
