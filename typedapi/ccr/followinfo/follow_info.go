package followinfo

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

type FollowInfo struct {
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

type NewFollowInfo func(index string) *FollowInfo

func NewFollowInfoFunc(tp elastictransport.Interface) NewFollowInfo {
	_ = "STUB: not implemented"
	return *new(NewFollowInfo)
}

func New(tp elastictransport.Interface) *FollowInfo { _ = "STUB: not implemented"; return nil }

func (r *FollowInfo) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FollowInfo) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FollowInfo) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FollowInfo) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *FollowInfo) Header(key, value string) *FollowInfo { _ = "STUB: not implemented"; return nil }

func (r *FollowInfo) _index(index string) *FollowInfo { _ = "STUB: not implemented"; return nil }

func (r *FollowInfo) MasterTimeout(duration string) *FollowInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *FollowInfo) ErrorTrace(errortrace bool) *FollowInfo { _ = "STUB: not implemented"; return nil }

func (r *FollowInfo) FilterPath(filterpaths ...string) *FollowInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *FollowInfo) Human(human bool) *FollowInfo { _ = "STUB: not implemented"; return nil }

func (r *FollowInfo) Pretty(pretty bool) *FollowInfo { _ = "STUB: not implemented"; return nil }
