package putdatastreamoptions

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataStreamOptions struct {
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

type NewPutDataStreamOptions func(name string) *PutDataStreamOptions

func NewPutDataStreamOptionsFunc(tp elastictransport.Interface) NewPutDataStreamOptions {
	_ = "STUB: not implemented"
	return *new(NewPutDataStreamOptions)
}

func New(tp elastictransport.Interface) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) Raw(raw io.Reader) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) Request(req *Request) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataStreamOptions) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataStreamOptions) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDataStreamOptions) Header(key, value string) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) _name(name string) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) MasterTimeout(duration string) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) Timeout(duration string) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) ErrorTrace(errortrace bool) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) FilterPath(filterpaths ...string) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) Human(human bool) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) Pretty(pretty bool) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamOptions) FailureStore(failurestore types.DataStreamFailureStoreVariant) *PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}
