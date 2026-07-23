package count

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catcountcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Count struct {
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

type NewCount func() *Count

func NewCountFunc(tp elastictransport.Interface) NewCount {
	_ = "STUB: not implemented"
	return *new(NewCount)
}

func New(tp elastictransport.Interface) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Raw(raw io.Reader) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Request(req *Request) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Count) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Count) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *Count) Header(key, value string) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Index(index string) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) H(catcountcolumns ...catcountcolumn.CatCountColumn) *Count {
	_ = "STUB: not implemented"
	return nil
}

func (r *Count) S(names ...string) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Bytes(bytes bytes.Bytes) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Format(format string) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Help(help bool) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Time(time timeunit.TimeUnit) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) V(v bool) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) ErrorTrace(errortrace bool) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) FilterPath(filterpaths ...string) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Human(human bool) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) Pretty(pretty bool) *Count { _ = "STUB: not implemented"; return nil }

func (r *Count) ProjectRouting(projectrouting string) *Count { _ = "STUB: not implemented"; return nil }
