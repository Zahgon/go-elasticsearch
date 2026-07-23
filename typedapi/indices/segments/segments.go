package segments

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Segments struct {
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

type NewSegments func() *Segments

func NewSegmentsFunc(tp elastictransport.Interface) NewSegments {
	_ = "STUB: not implemented"
	return *new(NewSegments)
}

func New(tp elastictransport.Interface) *Segments { _ = "STUB: not implemented"; return nil }

func (r *Segments) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Segments) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Segments) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Segments) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Segments) Header(key, value string) *Segments { _ = "STUB: not implemented"; return nil }

func (r *Segments) Index(index string) *Segments { _ = "STUB: not implemented"; return nil }

func (r *Segments) AllowNoIndices(allownoindices bool) *Segments {
	_ = "STUB: not implemented"
	return nil
}

func (r *Segments) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Segments {
	_ = "STUB: not implemented"
	return nil
}

func (r *Segments) IgnoreUnavailable(ignoreunavailable bool) *Segments {
	_ = "STUB: not implemented"
	return nil
}

func (r *Segments) ErrorTrace(errortrace bool) *Segments { _ = "STUB: not implemented"; return nil }

func (r *Segments) FilterPath(filterpaths ...string) *Segments {
	_ = "STUB: not implemented"
	return nil
}

func (r *Segments) Human(human bool) *Segments { _ = "STUB: not implemented"; return nil }

func (r *Segments) Pretty(pretty bool) *Segments { _ = "STUB: not implemented"; return nil }
