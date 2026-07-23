package termvectors

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

	idMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Termvectors struct {
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
	id    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTermvectors func(index string) *Termvectors

func NewTermvectorsFunc(tp elastictransport.Interface) NewTermvectors {
	_ = "STUB: not implemented"
	return *new(NewTermvectors)
}

func New(tp elastictransport.Interface) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Raw(raw io.Reader) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Request(req *Request) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Termvectors) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Termvectors) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Termvectors) Header(key, value string) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) _index(index string) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Id(id string) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Preference(preference string) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) Realtime(realtime bool) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) ErrorTrace(errortrace bool) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) FilterPath(filterpaths ...string) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) Human(human bool) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Pretty(pretty bool) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Doc(doc any) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) FieldStatistics(fieldstatistics bool) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) Fields(fields ...string) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Filter(filter types.TermVectorsFilterVariant) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) Offsets(offsets bool) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Payloads(payloads bool) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) PerFieldAnalyzer(perfieldanalyzer map[string]string) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) AddPerFieldAnalyzer(key string, value string) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) Positions(positions bool) *Termvectors { _ = "STUB: not implemented"; return nil }

func (r *Termvectors) Routing(routings ...string) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) TermStatistics(termstatistics bool) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) Version(versionnumber int64) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (r *Termvectors) VersionType(versiontype versiontype.VersionType) *Termvectors {
	_ = "STUB: not implemented"
	return nil
}
