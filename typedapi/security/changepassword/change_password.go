package changepassword

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	usernameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ChangePassword struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	username string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewChangePassword func() *ChangePassword

func NewChangePasswordFunc(tp elastictransport.Interface) NewChangePassword {
	_ = "STUB: not implemented"
	return *new(NewChangePassword)
}

func New(tp elastictransport.Interface) *ChangePassword { _ = "STUB: not implemented"; return nil }

func (r *ChangePassword) Raw(raw io.Reader) *ChangePassword { _ = "STUB: not implemented"; return nil }

func (r *ChangePassword) Request(req *Request) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ChangePassword) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ChangePassword) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ChangePassword) Header(key, value string) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) Username(username string) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) Refresh(refresh refresh.Refresh) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) ErrorTrace(errortrace bool) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) FilterPath(filterpaths ...string) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) Human(human bool) *ChangePassword { _ = "STUB: not implemented"; return nil }

func (r *ChangePassword) Pretty(pretty bool) *ChangePassword { _ = "STUB: not implemented"; return nil }

func (r *ChangePassword) Password(password string) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChangePassword) PasswordHash(passwordhash string) *ChangePassword {
	_ = "STUB: not implemented"
	return nil
}
