package mtermvectors

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mtermvectors struct {
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

type NewMtermvectors func() *Mtermvectors

func NewMtermvectorsFunc(tp elastictransport.Interface) NewMtermvectors {
	_ = "STUB: not implemented"
	return *new(NewMtermvectors)
}

func New(tp elastictransport.Interface) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Raw(raw io.Reader) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Request(req *Request) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Mtermvectors) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Mtermvectors) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Mtermvectors) Header(key, value string) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Index(index string) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Fields(fields ...string) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) FieldStatistics(fieldstatistics bool) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Offsets(offsets bool) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Payloads(payloads bool) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Positions(positions bool) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Preference(preference string) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Realtime(realtime bool) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Routing(routings ...string) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) TermStatistics(termstatistics bool) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Version(versionnumber string) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) VersionType(versiontype versiontype.VersionType) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) ErrorTrace(errortrace bool) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) FilterPath(filterpaths ...string) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Human(human bool) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Pretty(pretty bool) *Mtermvectors { _ = "STUB: not implemented"; return nil }

func (r *Mtermvectors) Docs(docs ...types.MTermVectorsOperationVariant) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) DocsValues(docsvalues []types.MTermVectorsOperation) *Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mtermvectors) Ids(ids ...string) *Mtermvectors { _ = "STUB: not implemented"; return nil }
