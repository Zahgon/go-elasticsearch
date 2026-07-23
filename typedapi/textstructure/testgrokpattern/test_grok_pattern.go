package testgrokpattern

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type TestGrokPattern struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTestGrokPattern func() *TestGrokPattern

func NewTestGrokPatternFunc(tp elastictransport.Interface) NewTestGrokPattern {
	_ = "STUB: not implemented"
	return *new(NewTestGrokPattern)
}

func New(tp elastictransport.Interface) *TestGrokPattern { _ = "STUB: not implemented"; return nil }

func (r *TestGrokPattern) Raw(raw io.Reader) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) Request(req *Request) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TestGrokPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TestGrokPattern) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TestGrokPattern) Header(key, value string) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) EcsCompatibility(ecscompatibility string) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) ErrorTrace(errortrace bool) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) FilterPath(filterpaths ...string) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) Human(human bool) *TestGrokPattern { _ = "STUB: not implemented"; return nil }

func (r *TestGrokPattern) Pretty(pretty bool) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) GrokPattern(grokpattern string) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *TestGrokPattern) Text(texts ...string) *TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}
