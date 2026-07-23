package grantapikey

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/apikeygranttype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GrantApiKey struct {
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

type NewGrantApiKey func() *GrantApiKey

func NewGrantApiKeyFunc(tp elastictransport.Interface) NewGrantApiKey {
	_ = "STUB: not implemented"
	return *new(NewGrantApiKey)
}

func New(tp elastictransport.Interface) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) Raw(raw io.Reader) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) Request(req *Request) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GrantApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GrantApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GrantApiKey) Header(key, value string) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) Refresh(refresh refresh.Refresh) *GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GrantApiKey) ErrorTrace(errortrace bool) *GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GrantApiKey) FilterPath(filterpaths ...string) *GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GrantApiKey) Human(human bool) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) Pretty(pretty bool) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) AccessToken(accesstoken string) *GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GrantApiKey) ApiKey(apikey types.GrantApiKeyVariant) *GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GrantApiKey) GrantType(granttype apikeygranttype.ApiKeyGrantType) *GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GrantApiKey) Password(password string) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) RunAs(username string) *GrantApiKey { _ = "STUB: not implemented"; return nil }

func (r *GrantApiKey) Username(username string) *GrantApiKey { _ = "STUB: not implemented"; return nil }
