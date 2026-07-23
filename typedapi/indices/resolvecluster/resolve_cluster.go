package resolvecluster

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
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResolveCluster struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewResolveCluster func() *ResolveCluster

func NewResolveClusterFunc(tp elastictransport.Interface) NewResolveCluster {
	_ = "STUB: not implemented"
	return *new(NewResolveCluster)
}

func New(tp elastictransport.Interface) *ResolveCluster { _ = "STUB: not implemented"; return nil }

func (r *ResolveCluster) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResolveCluster) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResolveCluster) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r ResolveCluster) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ResolveCluster) Header(key, value string) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) Name(name string) *ResolveCluster { _ = "STUB: not implemented"; return nil }

func (r *ResolveCluster) AllowNoIndices(allownoindices bool) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) IgnoreThrottled(ignorethrottled bool) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) IgnoreUnavailable(ignoreunavailable bool) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) Timeout(duration string) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) ErrorTrace(errortrace bool) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) FilterPath(filterpaths ...string) *ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResolveCluster) Human(human bool) *ResolveCluster { _ = "STUB: not implemented"; return nil }

func (r *ResolveCluster) Pretty(pretty bool) *ResolveCluster { _ = "STUB: not implemented"; return nil }
