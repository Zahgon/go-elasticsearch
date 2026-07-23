package unfollow

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

type Unfollow struct {
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

type NewUnfollow func(index string) *Unfollow

func NewUnfollowFunc(tp elastictransport.Interface) NewUnfollow {
	_ = "STUB: not implemented"
	return *new(NewUnfollow)
}

func New(tp elastictransport.Interface) *Unfollow { _ = "STUB: not implemented"; return nil }

func (r *Unfollow) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Unfollow) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Unfollow) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Unfollow) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Unfollow) Header(key, value string) *Unfollow { _ = "STUB: not implemented"; return nil }

func (r *Unfollow) _index(index string) *Unfollow { _ = "STUB: not implemented"; return nil }

func (r *Unfollow) MasterTimeout(duration string) *Unfollow { _ = "STUB: not implemented"; return nil }

func (r *Unfollow) ErrorTrace(errortrace bool) *Unfollow { _ = "STUB: not implemented"; return nil }

func (r *Unfollow) FilterPath(filterpaths ...string) *Unfollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Unfollow) Human(human bool) *Unfollow { _ = "STUB: not implemented"; return nil }

func (r *Unfollow) Pretty(pretty bool) *Unfollow { _ = "STUB: not implemented"; return nil }
