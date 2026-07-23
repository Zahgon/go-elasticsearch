package activateuserprofile

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/granttype"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ActivateUserProfile struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewActivateUserProfile func() *ActivateUserProfile

func NewActivateUserProfileFunc(tp elastictransport.Interface) NewActivateUserProfile {
	_ = "STUB: not implemented"
	return *new(NewActivateUserProfile)
}

func New(tp elastictransport.Interface) *ActivateUserProfile { _ = "STUB: not implemented"; return nil }

func (r *ActivateUserProfile) Raw(raw io.Reader) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) Request(req *Request) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ActivateUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ActivateUserProfile) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ActivateUserProfile) Header(key, value string) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) ErrorTrace(errortrace bool) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) FilterPath(filterpaths ...string) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) Human(human bool) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) Pretty(pretty bool) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) AccessToken(accesstoken string) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) GrantType(granttype granttype.GrantType) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) Password(password string) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateUserProfile) Username(username string) *ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}
