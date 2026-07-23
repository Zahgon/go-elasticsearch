package forgetfollower

import (
	gobytes "bytes"
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

type ForgetFollower struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewForgetFollower func(index string) *ForgetFollower

func NewForgetFollowerFunc(tp elastictransport.Interface) NewForgetFollower {
	_ = "STUB: not implemented"
	return *new(NewForgetFollower)
}

func New(tp elastictransport.Interface) *ForgetFollower { _ = "STUB: not implemented"; return nil }

func (r *ForgetFollower) Raw(raw io.Reader) *ForgetFollower { _ = "STUB: not implemented"; return nil }

func (r *ForgetFollower) Request(req *Request) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ForgetFollower) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ForgetFollower) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ForgetFollower) Header(key, value string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) _index(index string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) Timeout(duration string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) ErrorTrace(errortrace bool) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) FilterPath(filterpaths ...string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) Human(human bool) *ForgetFollower { _ = "STUB: not implemented"; return nil }

func (r *ForgetFollower) Pretty(pretty bool) *ForgetFollower { _ = "STUB: not implemented"; return nil }

func (r *ForgetFollower) FollowerCluster(followercluster string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) FollowerIndex(indexname string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) FollowerIndexUuid(uuid string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (r *ForgetFollower) LeaderRemoteCluster(leaderremotecluster string) *ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}
