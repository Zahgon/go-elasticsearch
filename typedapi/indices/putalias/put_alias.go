package putalias

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

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAlias struct {
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
	name  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAlias func(index, name string) *PutAlias

func NewPutAliasFunc(tp elastictransport.Interface) NewPutAlias {
	_ = "STUB: not implemented"
	return *new(NewPutAlias)
}

func New(tp elastictransport.Interface) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) Raw(raw io.Reader) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) Request(req *Request) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAlias) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAlias) Header(key, value string) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) _index(index string) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) _name(name string) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) MasterTimeout(duration string) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) Timeout(duration string) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) ErrorTrace(errortrace bool) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) FilterPath(filterpaths ...string) *PutAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlias) Human(human bool) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) Pretty(pretty bool) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) Filter(filter types.QueryVariant) *PutAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlias) IndexRouting(indexrouting string) *PutAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlias) IsWriteIndex(iswriteindex bool) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) Routing(routing string) *PutAlias { _ = "STUB: not implemented"; return nil }

func (r *PutAlias) SearchRouting(searchrouting string) *PutAlias {
	_ = "STUB: not implemented"
	return nil
}
