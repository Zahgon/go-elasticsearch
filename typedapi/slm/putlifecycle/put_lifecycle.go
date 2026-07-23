package putlifecycle

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
	policyidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	policyid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutLifecycle func(policyid string) *PutLifecycle

func NewPutLifecycleFunc(tp elastictransport.Interface) NewPutLifecycle {
	_ = "STUB: not implemented"
	return *new(NewPutLifecycle)
}

func New(tp elastictransport.Interface) *PutLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutLifecycle) Raw(raw io.Reader) *PutLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutLifecycle) Request(req *Request) *PutLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutLifecycle) Header(key, value string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) _policyid(policyid string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) MasterTimeout(duration string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) Timeout(duration string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) ErrorTrace(errortrace bool) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) FilterPath(filterpaths ...string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) Human(human bool) *PutLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutLifecycle) Pretty(pretty bool) *PutLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutLifecycle) Config(config types.ConfigurationVariant) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) Name(name string) *PutLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutLifecycle) Repository(repository string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) Retention(retention types.RetentionVariant) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLifecycle) Schedule(cronexpression string) *PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}
