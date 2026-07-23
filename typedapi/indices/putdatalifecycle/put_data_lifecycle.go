package putdatalifecycle

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataLifecycle struct {
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

type NewPutDataLifecycle func(name string) *PutDataLifecycle

func NewPutDataLifecycleFunc(tp elastictransport.Interface) NewPutDataLifecycle {
	_ = "STUB: not implemented"
	return *new(NewPutDataLifecycle)
}

func New(tp elastictransport.Interface) *PutDataLifecycle { _ = "STUB: not implemented"; return nil }

func (r *PutDataLifecycle) Raw(raw io.Reader) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) Request(req *Request) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDataLifecycle) Header(key, value string) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) _name(name string) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) MasterTimeout(duration string) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) Timeout(duration string) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) ErrorTrace(errortrace bool) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) FilterPath(filterpaths ...string) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) Human(human bool) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) Pretty(pretty bool) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) DataRetention(duration types.DurationVariant) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) Downsampling(downsamplings ...types.DownsamplingRoundVariant) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) DownsamplingValues(downsamplingvalues []types.DownsamplingRound) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) DownsamplingMethod(downsamplingmethod samplingmethod.SamplingMethod) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataLifecycle) Enabled(enabled bool) *PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}
