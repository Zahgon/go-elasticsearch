package fieldusagestats

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

type FieldUsageStats struct {
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

type NewFieldUsageStats func(index string) *FieldUsageStats

func NewFieldUsageStatsFunc(tp elastictransport.Interface) NewFieldUsageStats {
	_ = "STUB: not implemented"
	return *new(NewFieldUsageStats)
}

func New(tp elastictransport.Interface) *FieldUsageStats { _ = "STUB: not implemented"; return nil }

func (r *FieldUsageStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FieldUsageStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FieldUsageStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FieldUsageStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *FieldUsageStats) Header(key, value string) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) _index(index string) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) AllowNoIndices(allownoindices bool) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) IgnoreUnavailable(ignoreunavailable bool) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) Fields(fields ...string) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) ErrorTrace(errortrace bool) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) FilterPath(filterpaths ...string) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldUsageStats) Human(human bool) *FieldUsageStats { _ = "STUB: not implemented"; return nil }

func (r *FieldUsageStats) Pretty(pretty bool) *FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}
