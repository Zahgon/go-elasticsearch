package downsample

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

const (
	indexMask = iota + 1

	targetindexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Downsample struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index       string
	targetindex string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDownsample func(index, targetindex string) *Downsample

func NewDownsampleFunc(tp elastictransport.Interface) NewDownsample {
	_ = "STUB: not implemented"
	return *new(NewDownsample)
}

func New(tp elastictransport.Interface) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) Raw(raw io.Reader) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) Request(req *Request) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Downsample) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Downsample) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *Downsample) Header(key, value string) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) _index(index string) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) _targetindex(targetindex string) *Downsample {
	_ = "STUB: not implemented"
	return nil
}

func (r *Downsample) ErrorTrace(errortrace bool) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) FilterPath(filterpaths ...string) *Downsample {
	_ = "STUB: not implemented"
	return nil
}

func (r *Downsample) Human(human bool) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) Pretty(pretty bool) *Downsample { _ = "STUB: not implemented"; return nil }

func (r *Downsample) FixedInterval(durationlarge string) *Downsample {
	_ = "STUB: not implemented"
	return nil
}

func (r *Downsample) SamplingMethod(samplingmethod samplingmethod.SamplingMethod) *Downsample {
	_ = "STUB: not implemented"
	return nil
}
