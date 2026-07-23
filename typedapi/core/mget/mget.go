package mget

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

type Mget struct {
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

type NewMget func() *Mget

func NewMgetFunc(tp elastictransport.Interface) NewMget {
	_ = "STUB: not implemented"
	return *new(NewMget)
}

func New(tp elastictransport.Interface) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Raw(raw io.Reader) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Request(req *Request) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Mget) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Mget) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Mget) Header(key, value string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Index(index string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) ForceSyntheticSource(forcesyntheticsource bool) *Mget {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mget) Preference(preference string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Realtime(realtime bool) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Refresh(refresh bool) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Routing(routings ...string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Source_(sourceconfigparam string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) SourceExcludes_(fields ...string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) SourceIncludes_(fields ...string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) StoredFields(fields ...string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) ErrorTrace(errortrace bool) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) FilterPath(filterpaths ...string) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Human(human bool) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Pretty(pretty bool) *Mget { _ = "STUB: not implemented"; return nil }

func (r *Mget) Docs(docs ...types.MgetOperationVariant) *Mget {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mget) DocsValues(docsvalues []types.MgetOperation) *Mget {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mget) Ids(ids ...string) *Mget { _ = "STUB: not implemented"; return nil }
