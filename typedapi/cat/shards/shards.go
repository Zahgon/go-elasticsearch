package shards

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catshardcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Shards struct {
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

type NewShards func() *Shards

func NewShardsFunc(tp elastictransport.Interface) NewShards {
	_ = "STUB: not implemented"
	return *new(NewShards)
}

func New(tp elastictransport.Interface) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Shards) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Shards) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Shards) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Shards) Header(key, value string) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Index(index string) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) H(catshardcolumns ...catshardcolumn.CatShardColumn) *Shards {
	_ = "STUB: not implemented"
	return nil
}

func (r *Shards) S(names ...string) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) MasterTimeout(duration string) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Bytes(bytes bytes.Bytes) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Format(format string) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Help(help bool) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Time(time timeunit.TimeUnit) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) V(v bool) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) ErrorTrace(errortrace bool) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) FilterPath(filterpaths ...string) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Human(human bool) *Shards { _ = "STUB: not implemented"; return nil }

func (r *Shards) Pretty(pretty bool) *Shards { _ = "STUB: not implemented"; return nil }
