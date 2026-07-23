package getuserprofile

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	uidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetUserProfile struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	uid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetUserProfile func(uid string) *GetUserProfile

func NewGetUserProfileFunc(tp elastictransport.Interface) NewGetUserProfile {
	_ = "STUB: not implemented"
	return *new(NewGetUserProfile)
}

func New(tp elastictransport.Interface) *GetUserProfile { _ = "STUB: not implemented"; return nil }

func (r *GetUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUserProfile) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUserProfile) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetUserProfile) Header(key, value string) *GetUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserProfile) _uid(uids ...string) *GetUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserProfile) Data(data ...string) *GetUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserProfile) ErrorTrace(errortrace bool) *GetUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserProfile) FilterPath(filterpaths ...string) *GetUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserProfile) Human(human bool) *GetUserProfile { _ = "STUB: not implemented"; return nil }

func (r *GetUserProfile) Pretty(pretty bool) *GetUserProfile { _ = "STUB: not implemented"; return nil }
