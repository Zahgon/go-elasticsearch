package getuser

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	usernameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetUser struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	username string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetUser func() *GetUser

func NewGetUserFunc(tp elastictransport.Interface) NewGetUser {
	_ = "STUB: not implemented"
	return *new(NewGetUser)
}

func New(tp elastictransport.Interface) *GetUser { _ = "STUB: not implemented"; return nil }

func (r *GetUser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUser) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetUser) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetUser) Header(key, value string) *GetUser { _ = "STUB: not implemented"; return nil }

func (r *GetUser) Username(usernames ...string) *GetUser { _ = "STUB: not implemented"; return nil }

func (r *GetUser) WithProfileUid(withprofileuid bool) *GetUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUser) ErrorTrace(errortrace bool) *GetUser { _ = "STUB: not implemented"; return nil }

func (r *GetUser) FilterPath(filterpaths ...string) *GetUser { _ = "STUB: not implemented"; return nil }

func (r *GetUser) Human(human bool) *GetUser { _ = "STUB: not implemented"; return nil }

func (r *GetUser) Pretty(pretty bool) *GetUser { _ = "STUB: not implemented"; return nil }
