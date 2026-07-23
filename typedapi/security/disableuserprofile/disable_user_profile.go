package disableuserprofile

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

type DisableUserProfile struct {
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

type NewDisableUserProfile func(uid string) *DisableUserProfile

func NewDisableUserProfileFunc(tp elastictransport.Interface) NewDisableUserProfile {
	_ = "STUB: not implemented"
	return *new(NewDisableUserProfile)
}

func New(tp elastictransport.Interface) *DisableUserProfile { _ = "STUB: not implemented"; return nil }

func (r *DisableUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DisableUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DisableUserProfile) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DisableUserProfile) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DisableUserProfile) Header(key, value string) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUserProfile) _uid(uid string) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUserProfile) Refresh(refresh refresh.Refresh) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUserProfile) ErrorTrace(errortrace bool) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUserProfile) FilterPath(filterpaths ...string) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUserProfile) Human(human bool) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUserProfile) Pretty(pretty bool) *DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}
