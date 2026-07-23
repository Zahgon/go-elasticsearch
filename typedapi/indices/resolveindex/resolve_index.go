package resolveindex

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexmode"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResolveIndex struct {
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

type NewResolveIndex func(name string) *ResolveIndex

func NewResolveIndexFunc(tp elastictransport.Interface) NewResolveIndex {
	_ = "STUB: not implemented"
	return *new(NewResolveIndex)
}

func New(tp elastictransport.Interface) *ResolveIndex { _ = "STUB: not implemented"; return nil }

func (r *ResolveIndex) Raw(raw io.Reader) *ResolveIndex { _ = "STUB: not implemented"; return nil }

func (r *ResolveIndex) Request(req *Request) *ResolveIndex { _ = "STUB: not implemented"; return nil }

func (r *ResolveIndex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResolveIndex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResolveIndex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ResolveIndex) Header(key, value string) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) _name(name string) *ResolveIndex { _ = "STUB: not implemented"; return nil }

func (r *ResolveIndex) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) IgnoreUnavailable(ignoreunavailable bool) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) AllowNoIndices(allownoindices bool) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) Mode(modes ...indexmode.IndexMode) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) ErrorTrace(errortrace bool) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) FilterPath(filterpaths ...string) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveIndex) Human(human bool) *ResolveIndex { _ = "STUB: not implemented"; return nil }

func (r *ResolveIndex) Pretty(pretty bool) *ResolveIndex { _ = "STUB: not implemented"; return nil }

func (r *ResolveIndex) ProjectRouting(projectrouting string) *ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}
