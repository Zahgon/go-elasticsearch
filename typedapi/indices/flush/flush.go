package flush

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

type Flush struct {
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

type NewFlush func() *Flush

func NewFlushFunc(tp elastictransport.Interface) NewFlush {
	_ = "STUB: not implemented"
	return *new(NewFlush)
}

func New(tp elastictransport.Interface) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Flush) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Flush) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Flush) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Flush) Header(key, value string) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) Index(index string) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) AllowNoIndices(allownoindices bool) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Flush {
	_ = "STUB: not implemented"
	return nil
}

func (r *Flush) Force(force bool) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) IgnoreUnavailable(ignoreunavailable bool) *Flush {
	_ = "STUB: not implemented"
	return nil
}

func (r *Flush) WaitIfOngoing(waitifongoing bool) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) ErrorTrace(errortrace bool) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) FilterPath(filterpaths ...string) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) Human(human bool) *Flush { _ = "STUB: not implemented"; return nil }

func (r *Flush) Pretty(pretty bool) *Flush { _ = "STUB: not implemented"; return nil }
