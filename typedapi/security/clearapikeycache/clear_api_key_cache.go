package clearapikeycache

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idsMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearApiKeyCache struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	ids string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClearApiKeyCache func(ids string) *ClearApiKeyCache

func NewClearApiKeyCacheFunc(tp elastictransport.Interface) NewClearApiKeyCache {
	_ = "STUB: not implemented"
	return *new(NewClearApiKeyCache)
}

func New(tp elastictransport.Interface) *ClearApiKeyCache { _ = "STUB: not implemented"; return nil }

func (r *ClearApiKeyCache) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearApiKeyCache) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearApiKeyCache) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearApiKeyCache) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearApiKeyCache) Header(key, value string) *ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearApiKeyCache) _ids(ids string) *ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearApiKeyCache) ErrorTrace(errortrace bool) *ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearApiKeyCache) FilterPath(filterpaths ...string) *ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearApiKeyCache) Human(human bool) *ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearApiKeyCache) Pretty(pretty bool) *ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}
