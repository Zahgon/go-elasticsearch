package indices

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catindicescolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Indices struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewIndices func() *Indices

func NewIndicesFunc(tp elastictransport.Interface) NewIndices {
	_ = "STUB: not implemented"
	return *new(NewIndices)
}

func New(tp elastictransport.Interface) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Indices) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Indices) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Indices) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Indices) Header(key, value string) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Index(index string) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Indices {
	_ = "STUB: not implemented"
	return nil
}

func (r *Indices) Health(health healthstatus.HealthStatus) *Indices {
	_ = "STUB: not implemented"
	return nil
}

func (r *Indices) IncludeUnloadedSegments(includeunloadedsegments bool) *Indices {
	_ = "STUB: not implemented"
	return nil
}

func (r *Indices) Pri(pri bool) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) MasterTimeout(duration string) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) H(catindicescolumns ...catindicescolumn.CatIndicesColumn) *Indices {
	_ = "STUB: not implemented"
	return nil
}

func (r *Indices) S(names ...string) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Bytes(bytes bytes.Bytes) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Format(format string) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Help(help bool) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Time(time timeunit.TimeUnit) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) V(v bool) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) ErrorTrace(errortrace bool) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) FilterPath(filterpaths ...string) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Human(human bool) *Indices { _ = "STUB: not implemented"; return nil }

func (r *Indices) Pretty(pretty bool) *Indices { _ = "STUB: not implemented"; return nil }
