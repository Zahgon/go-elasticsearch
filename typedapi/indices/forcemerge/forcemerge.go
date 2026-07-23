package forcemerge

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

type Forcemerge struct {
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

type NewForcemerge func() *Forcemerge

func NewForcemergeFunc(tp elastictransport.Interface) NewForcemerge {
	_ = "STUB: not implemented"
	return *new(NewForcemerge)
}

func New(tp elastictransport.Interface) *Forcemerge { _ = "STUB: not implemented"; return nil }

func (r *Forcemerge) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Forcemerge) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Forcemerge) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Forcemerge) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Forcemerge) Header(key, value string) *Forcemerge { _ = "STUB: not implemented"; return nil }

func (r *Forcemerge) Index(index string) *Forcemerge { _ = "STUB: not implemented"; return nil }

func (r *Forcemerge) AllowNoIndices(allownoindices bool) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) Flush(flush bool) *Forcemerge { _ = "STUB: not implemented"; return nil }

func (r *Forcemerge) IgnoreUnavailable(ignoreunavailable bool) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) MaxNumSegments(maxnumsegments string) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) OnlyExpungeDeletes(onlyexpungedeletes bool) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) WaitForCompletion(waitforcompletion bool) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) ErrorTrace(errortrace bool) *Forcemerge { _ = "STUB: not implemented"; return nil }

func (r *Forcemerge) FilterPath(filterpaths ...string) *Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forcemerge) Human(human bool) *Forcemerge { _ = "STUB: not implemented"; return nil }

func (r *Forcemerge) Pretty(pretty bool) *Forcemerge { _ = "STUB: not implemented"; return nil }
