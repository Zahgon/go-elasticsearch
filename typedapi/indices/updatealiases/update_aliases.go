package updatealiases

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

type UpdateAliases struct {
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

type NewUpdateAliases func() *UpdateAliases

func NewUpdateAliasesFunc(tp elastictransport.Interface) NewUpdateAliases {
	_ = "STUB: not implemented"
	return *new(NewUpdateAliases)
}

func New(tp elastictransport.Interface) *UpdateAliases { _ = "STUB: not implemented"; return nil }

func (r *UpdateAliases) Raw(raw io.Reader) *UpdateAliases { _ = "STUB: not implemented"; return nil }

func (r *UpdateAliases) Request(req *Request) *UpdateAliases { _ = "STUB: not implemented"; return nil }

func (r *UpdateAliases) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateAliases) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateAliases) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateAliases) Header(key, value string) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateAliases) MasterTimeout(duration string) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateAliases) Timeout(duration string) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateAliases) ErrorTrace(errortrace bool) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateAliases) FilterPath(filterpaths ...string) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateAliases) Human(human bool) *UpdateAliases { _ = "STUB: not implemented"; return nil }

func (r *UpdateAliases) Pretty(pretty bool) *UpdateAliases { _ = "STUB: not implemented"; return nil }

func (r *UpdateAliases) Actions(actions ...types.IndicesActionVariant) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateAliases) ActionsValues(actionsvalues []types.IndicesAction) *UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}
