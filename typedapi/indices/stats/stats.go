package stats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/level"
)

const (
	metricMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Stats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	metric string
	index  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStats func() *Stats

func NewStatsFunc(tp elastictransport.Interface) NewStats {
	_ = "STUB: not implemented"
	return *new(NewStats)
}

func New(tp elastictransport.Interface) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Stats) Header(key, value string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Metric(metric string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Index(index string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) CompletionFields(fields ...string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Stats {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stats) FielddataFields(fields ...string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Fields(fields ...string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) ForbidClosedIndices(forbidclosedindices bool) *Stats {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stats) Groups(groups ...string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) IncludeSegmentFileSizes(includesegmentfilesizes bool) *Stats {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stats) IncludeUnloadedSegments(includeunloadedsegments bool) *Stats {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stats) Level(level level.Level) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) ErrorTrace(errortrace bool) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) FilterPath(filterpaths ...string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Human(human bool) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Pretty(pretty bool) *Stats { _ = "STUB: not implemented"; return nil }
