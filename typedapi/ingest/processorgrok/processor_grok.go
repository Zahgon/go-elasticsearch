package processorgrok

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ProcessorGrok struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewProcessorGrok func() *ProcessorGrok

func NewProcessorGrokFunc(tp elastictransport.Interface) NewProcessorGrok {
	_ = "STUB: not implemented"
	return *new(NewProcessorGrok)
}

func New(tp elastictransport.Interface) *ProcessorGrok { _ = "STUB: not implemented"; return nil }

func (r *ProcessorGrok) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ProcessorGrok) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ProcessorGrok) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ProcessorGrok) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ProcessorGrok) Header(key, value string) *ProcessorGrok {
	_ = "STUB: not implemented"
	return nil
}

func (r *ProcessorGrok) ErrorTrace(errortrace bool) *ProcessorGrok {
	_ = "STUB: not implemented"
	return nil
}

func (r *ProcessorGrok) FilterPath(filterpaths ...string) *ProcessorGrok {
	_ = "STUB: not implemented"
	return nil
}

func (r *ProcessorGrok) Human(human bool) *ProcessorGrok { _ = "STUB: not implemented"; return nil }

func (r *ProcessorGrok) Pretty(pretty bool) *ProcessorGrok { _ = "STUB: not implemented"; return nil }
