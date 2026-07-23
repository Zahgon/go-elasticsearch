package split

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	indexMask = iota + 1

	targetMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Split struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index  string
	target string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSplit func(index, target string) *Split

func NewSplitFunc(tp elastictransport.Interface) NewSplit {
	_ = "STUB: not implemented"
	return *new(NewSplit)
}

func New(tp elastictransport.Interface) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) Raw(raw io.Reader) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) Request(req *Request) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Split) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Split) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Split) Header(key, value string) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) _index(index string) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) _target(target string) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) MasterTimeout(duration string) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) Timeout(duration string) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) WaitForActiveShards(waitforactiveshards string) *Split {
	_ = "STUB: not implemented"
	return nil
}

func (r *Split) ErrorTrace(errortrace bool) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) FilterPath(filterpaths ...string) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) Human(human bool) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) Pretty(pretty bool) *Split { _ = "STUB: not implemented"; return nil }

func (r *Split) Aliases(aliases map[string]types.Alias) *Split {
	_ = "STUB: not implemented"
	return nil
}

func (r *Split) AddAlias(key string, value types.AliasVariant) *Split {
	_ = "STUB: not implemented"
	return nil
}

func (r *Split) Settings(settings map[string]json.RawMessage) *Split {
	_ = "STUB: not implemented"
	return nil
}

func (r *Split) AddSetting(key string, value json.RawMessage) *Split {
	_ = "STUB: not implemented"
	return nil
}
