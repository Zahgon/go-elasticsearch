package close

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

type Close struct {
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

type NewClose func(index string) *Close

func NewCloseFunc(tp elastictransport.Interface) NewClose {
	_ = "STUB: not implemented"
	return *new(NewClose)
}

func New(tp elastictransport.Interface) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Close) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Close) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Close) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Close) Header(key, value string) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) _index(index string) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) AllowNoIndices(allownoindices bool) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Close {
	_ = "STUB: not implemented"
	return nil
}

func (r *Close) IgnoreUnavailable(ignoreunavailable bool) *Close {
	_ = "STUB: not implemented"
	return nil
}

func (r *Close) MasterTimeout(duration string) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) Timeout(duration string) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) WaitForActiveShards(waitforactiveshards string) *Close {
	_ = "STUB: not implemented"
	return nil
}

func (r *Close) ErrorTrace(errortrace bool) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) FilterPath(filterpaths ...string) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) Human(human bool) *Close { _ = "STUB: not implemented"; return nil }

func (r *Close) Pretty(pretty bool) *Close { _ = "STUB: not implemented"; return nil }
