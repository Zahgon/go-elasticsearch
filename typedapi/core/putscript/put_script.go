package putscript

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

	contextMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutScript struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id      string
	context string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutScript func(id string) *PutScript

func NewPutScriptFunc(tp elastictransport.Interface) NewPutScript {
	_ = "STUB: not implemented"
	return *new(NewPutScript)
}

func New(tp elastictransport.Interface) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) Raw(raw io.Reader) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) Request(req *Request) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutScript) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutScript) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutScript) Header(key, value string) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) _id(id string) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) Context(context string) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) MasterTimeout(duration string) *PutScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutScript) Timeout(duration string) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) ErrorTrace(errortrace bool) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) FilterPath(filterpaths ...string) *PutScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutScript) Human(human bool) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) Pretty(pretty bool) *PutScript { _ = "STUB: not implemented"; return nil }

func (r *PutScript) Script(script types.StoredScriptVariant) *PutScript {
	_ = "STUB: not implemented"
	return nil
}
