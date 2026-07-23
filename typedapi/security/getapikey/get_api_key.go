package getapikey

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetApiKey func() *GetApiKey

func NewGetApiKeyFunc(tp elastictransport.Interface) NewGetApiKey {
	_ = "STUB: not implemented"
	return *new(NewGetApiKey)
}

func New(tp elastictransport.Interface) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetApiKey) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetApiKey) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetApiKey) Header(key, value string) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) Id(id string) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) Name(name string) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) Owner(owner bool) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) RealmName(name string) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) Username(username string) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) WithLimitedBy(withlimitedby bool) *GetApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetApiKey) ActiveOnly(activeonly bool) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) WithProfileUid(withprofileuid bool) *GetApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetApiKey) ErrorTrace(errortrace bool) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) FilterPath(filterpaths ...string) *GetApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetApiKey) Human(human bool) *GetApiKey { _ = "STUB: not implemented"; return nil }

func (r *GetApiKey) Pretty(pretty bool) *GetApiKey { _ = "STUB: not implemented"; return nil }
