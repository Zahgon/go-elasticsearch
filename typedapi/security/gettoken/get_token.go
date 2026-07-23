package gettoken

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/accesstokengranttype"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetToken struct {
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

type NewGetToken func() *GetToken

func NewGetTokenFunc(tp elastictransport.Interface) NewGetToken {
	_ = "STUB: not implemented"
	return *new(NewGetToken)
}

func New(tp elastictransport.Interface) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) Raw(raw io.Reader) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) Request(req *Request) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetToken) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetToken) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetToken) Header(key, value string) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) ErrorTrace(errortrace bool) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) FilterPath(filterpaths ...string) *GetToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetToken) Human(human bool) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) Pretty(pretty bool) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) GrantType(granttype accesstokengranttype.AccessTokenGrantType) *GetToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetToken) KerberosTicket(kerberosticket string) *GetToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetToken) Password(password string) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) RefreshToken(refreshtoken string) *GetToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetToken) Scope(scope string) *GetToken { _ = "STUB: not implemented"; return nil }

func (r *GetToken) Username(username string) *GetToken { _ = "STUB: not implemented"; return nil }
