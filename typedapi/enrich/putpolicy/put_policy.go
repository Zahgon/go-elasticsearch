package putpolicy

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
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutPolicy struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutPolicy func(name string) *PutPolicy

func NewPutPolicyFunc(tp elastictransport.Interface) NewPutPolicy {
	_ = "STUB: not implemented"
	return *new(NewPutPolicy)
}

func New(tp elastictransport.Interface) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) Raw(raw io.Reader) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) Request(req *Request) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutPolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutPolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutPolicy) Header(key, value string) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) _name(name string) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) MasterTimeout(duration string) *PutPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPolicy) ErrorTrace(errortrace bool) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) FilterPath(filterpaths ...string) *PutPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPolicy) Human(human bool) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) Pretty(pretty bool) *PutPolicy { _ = "STUB: not implemented"; return nil }

func (r *PutPolicy) GeoMatch(geomatch types.EnrichPolicyVariant) *PutPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPolicy) Match(match types.EnrichPolicyVariant) *PutPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPolicy) Range(range_ types.EnrichPolicyVariant) *PutPolicy {
	_ = "STUB: not implemented"
	return nil
}
