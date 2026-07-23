package enableuserprofile

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	uidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type EnableUserProfile struct {
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

type NewEnableUserProfile func(uid string) *EnableUserProfile

func NewEnableUserProfileFunc(tp elastictransport.Interface) NewEnableUserProfile {
	_ = "STUB: not implemented"
	return *new(NewEnableUserProfile)
}

func New(tp elastictransport.Interface) *EnableUserProfile { _ = "STUB: not implemented"; return nil }

func (r *EnableUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnableUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnableUserProfile) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnableUserProfile) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *EnableUserProfile) Header(key, value string) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUserProfile) _uid(uid string) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUserProfile) Refresh(refresh refresh.Refresh) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUserProfile) ErrorTrace(errortrace bool) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUserProfile) FilterPath(filterpaths ...string) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUserProfile) Human(human bool) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUserProfile) Pretty(pretty bool) *EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}
