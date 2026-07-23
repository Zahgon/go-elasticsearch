package followstats

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

type FollowStats struct {
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

type NewFollowStats func(index string) *FollowStats

func NewFollowStatsFunc(tp elastictransport.Interface) NewFollowStats {
	_ = "STUB: not implemented"
	return *new(NewFollowStats)
}

func New(tp elastictransport.Interface) *FollowStats { _ = "STUB: not implemented"; return nil }

func (r *FollowStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FollowStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FollowStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FollowStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *FollowStats) Header(key, value string) *FollowStats { _ = "STUB: not implemented"; return nil }

func (r *FollowStats) _index(index string) *FollowStats { _ = "STUB: not implemented"; return nil }

func (r *FollowStats) Timeout(duration string) *FollowStats { _ = "STUB: not implemented"; return nil }

func (r *FollowStats) ErrorTrace(errortrace bool) *FollowStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FollowStats) FilterPath(filterpaths ...string) *FollowStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *FollowStats) Human(human bool) *FollowStats { _ = "STUB: not implemented"; return nil }

func (r *FollowStats) Pretty(pretty bool) *FollowStats { _ = "STUB: not implemented"; return nil }
