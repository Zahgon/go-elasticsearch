package shrink

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

type Shrink struct {
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

type NewShrink func(index, target string) *Shrink

func NewShrinkFunc(tp elastictransport.Interface) NewShrink {
	_ = "STUB: not implemented"
	return *new(NewShrink)
}

func New(tp elastictransport.Interface) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) Raw(raw io.Reader) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) Request(req *Request) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Shrink) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Shrink) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Shrink) Header(key, value string) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) _index(index string) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) _target(target string) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) MasterTimeout(duration string) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) Timeout(duration string) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) WaitForActiveShards(waitforactiveshards string) *Shrink {
	_ = "STUB: not implemented"
	return nil
}

func (r *Shrink) ErrorTrace(errortrace bool) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) FilterPath(filterpaths ...string) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) Human(human bool) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) Pretty(pretty bool) *Shrink { _ = "STUB: not implemented"; return nil }

func (r *Shrink) Aliases(aliases map[string]types.Alias) *Shrink {
	_ = "STUB: not implemented"
	return nil
}

func (r *Shrink) AddAlias(key string, value types.AliasVariant) *Shrink {
	_ = "STUB: not implemented"
	return nil
}

func (r *Shrink) Settings(settings map[string]json.RawMessage) *Shrink {
	_ = "STUB: not implemented"
	return nil
}

func (r *Shrink) AddSetting(key string, value json.RawMessage) *Shrink {
	_ = "STUB: not implemented"
	return nil
}
