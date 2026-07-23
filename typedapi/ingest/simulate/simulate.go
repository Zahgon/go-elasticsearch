package simulate

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
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Simulate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSimulate func() *Simulate

func NewSimulateFunc(tp elastictransport.Interface) NewSimulate {
	_ = "STUB: not implemented"
	return *new(NewSimulate)
}

func New(tp elastictransport.Interface) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) Raw(raw io.Reader) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) Request(req *Request) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Simulate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Simulate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Simulate) Header(key, value string) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) Id(id string) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) Verbose(verbose bool) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) ErrorTrace(errortrace bool) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) FilterPath(filterpaths ...string) *Simulate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Simulate) Human(human bool) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) Pretty(pretty bool) *Simulate { _ = "STUB: not implemented"; return nil }

func (r *Simulate) Docs(docs ...types.DocumentVariant) *Simulate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Simulate) DocsValues(docsvalues []types.Document) *Simulate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Simulate) Pipeline(pipeline types.IngestPipelineVariant) *Simulate {
	_ = "STUB: not implemented"
	return nil
}
