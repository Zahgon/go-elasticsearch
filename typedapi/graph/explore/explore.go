package explore

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

type Explore struct {
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

type NewExplore func(index string) *Explore

func NewExploreFunc(tp elastictransport.Interface) NewExplore {
	_ = "STUB: not implemented"
	return *new(NewExplore)
}

func New(tp elastictransport.Interface) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Raw(raw io.Reader) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Request(req *Request) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Explore) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Explore) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Explore) Header(key, value string) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) _index(index string) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Routing(routings ...string) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Timeout(duration string) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) ErrorTrace(errortrace bool) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) FilterPath(filterpaths ...string) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Human(human bool) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Pretty(pretty bool) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Connections(connections types.HopVariant) *Explore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Explore) Controls(controls types.ExploreControlsVariant) *Explore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Explore) Query(query types.QueryVariant) *Explore { _ = "STUB: not implemented"; return nil }

func (r *Explore) Vertices(vertices ...types.VertexDefinitionVariant) *Explore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Explore) VerticesValues(verticesvalues []types.VertexDefinition) *Explore {
	_ = "STUB: not implemented"
	return nil
}
