package deprecations

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Deprecations struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeprecations func() *Deprecations

func NewDeprecationsFunc(tp elastictransport.Interface) NewDeprecations {
	_ = "STUB: not implemented"
	return *new(NewDeprecations)
}

func New(tp elastictransport.Interface) *Deprecations { _ = "STUB: not implemented"; return nil }

func (r *Deprecations) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Deprecations) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Deprecations) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Deprecations) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Deprecations) Header(key, value string) *Deprecations {
	_ = "STUB: not implemented"
	return nil
}

func (r *Deprecations) Index(index string) *Deprecations { _ = "STUB: not implemented"; return nil }

func (r *Deprecations) ErrorTrace(errortrace bool) *Deprecations {
	_ = "STUB: not implemented"
	return nil
}

func (r *Deprecations) FilterPath(filterpaths ...string) *Deprecations {
	_ = "STUB: not implemented"
	return nil
}

func (r *Deprecations) Human(human bool) *Deprecations { _ = "STUB: not implemented"; return nil }

func (r *Deprecations) Pretty(pretty bool) *Deprecations { _ = "STUB: not implemented"; return nil }
