package create

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
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCreate func(index string) *Create

func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	_ = "STUB: not implemented"
	return *new(NewCreate)
}

func New(tp elastictransport.Interface) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Raw(raw io.Reader) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Request(req *Request) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Create) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Create) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Create) Header(key, value string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) _index(index string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) MasterTimeout(duration string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Timeout(duration string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) WaitForActiveShards(waitforactiveshards string) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) ErrorTrace(errortrace bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) FilterPath(filterpaths ...string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Human(human bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Pretty(pretty bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Aliases(aliases map[string]types.Alias) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) AddAlias(key string, value types.AliasVariant) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) Mappings(mappings types.TypeMappingVariant) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) Settings(settings types.IndexSettingsVariant) *Create {
	_ = "STUB: not implemented"
	return nil
}
