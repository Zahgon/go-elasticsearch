package rollover

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
	aliasMask = iota + 1

	newindexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Rollover struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	alias    string
	newindex string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRollover func(alias string) *Rollover

func NewRolloverFunc(tp elastictransport.Interface) NewRollover {
	_ = "STUB: not implemented"
	return *new(NewRollover)
}

func New(tp elastictransport.Interface) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) Raw(raw io.Reader) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) Request(req *Request) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Rollover) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Rollover) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Rollover) Header(key, value string) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) _alias(alias string) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) NewIndex(newindex string) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) DryRun(dryrun bool) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) MasterTimeout(duration string) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) Timeout(duration string) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) WaitForActiveShards(waitforactiveshards string) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) Lazy(lazy bool) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) ErrorTrace(errortrace bool) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) FilterPath(filterpaths ...string) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) Human(human bool) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) Pretty(pretty bool) *Rollover { _ = "STUB: not implemented"; return nil }

func (r *Rollover) Aliases(aliases map[string]types.Alias) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) AddAlias(key string, value types.AliasVariant) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) Conditions(conditions types.RolloverConditionsVariant) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) Mappings(mappings types.TypeMappingVariant) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) Settings(settings map[string]json.RawMessage) *Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rollover) AddSetting(key string, value json.RawMessage) *Rollover {
	_ = "STUB: not implemented"
	return nil
}
