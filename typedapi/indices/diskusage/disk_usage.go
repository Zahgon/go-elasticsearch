package diskusage

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

type DiskUsage struct {
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

type NewDiskUsage func(index string) *DiskUsage

func NewDiskUsageFunc(tp elastictransport.Interface) NewDiskUsage {
	_ = "STUB: not implemented"
	return *new(NewDiskUsage)
}

func New(tp elastictransport.Interface) *DiskUsage { _ = "STUB: not implemented"; return nil }

func (r *DiskUsage) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DiskUsage) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DiskUsage) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r DiskUsage) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DiskUsage) Header(key, value string) *DiskUsage { _ = "STUB: not implemented"; return nil }

func (r *DiskUsage) _index(index string) *DiskUsage { _ = "STUB: not implemented"; return nil }

func (r *DiskUsage) AllowNoIndices(allownoindices bool) *DiskUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *DiskUsage) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DiskUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *DiskUsage) Flush(flush bool) *DiskUsage { _ = "STUB: not implemented"; return nil }

func (r *DiskUsage) IgnoreUnavailable(ignoreunavailable bool) *DiskUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *DiskUsage) RunExpensiveTasks(runexpensivetasks bool) *DiskUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *DiskUsage) ErrorTrace(errortrace bool) *DiskUsage { _ = "STUB: not implemented"; return nil }

func (r *DiskUsage) FilterPath(filterpaths ...string) *DiskUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *DiskUsage) Human(human bool) *DiskUsage { _ = "STUB: not implemented"; return nil }

func (r *DiskUsage) Pretty(pretty bool) *DiskUsage { _ = "STUB: not implemented"; return nil }
