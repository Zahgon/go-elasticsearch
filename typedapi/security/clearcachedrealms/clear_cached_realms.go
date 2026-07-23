package clearcachedrealms

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	realmsMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearCachedRealms struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	realms string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClearCachedRealms func(realms string) *ClearCachedRealms

func NewClearCachedRealmsFunc(tp elastictransport.Interface) NewClearCachedRealms {
	_ = "STUB: not implemented"
	return *new(NewClearCachedRealms)
}

func New(tp elastictransport.Interface) *ClearCachedRealms { _ = "STUB: not implemented"; return nil }

func (r *ClearCachedRealms) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedRealms) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedRealms) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedRealms) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearCachedRealms) Header(key, value string) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRealms) _realms(realms string) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRealms) Usernames(usernames ...string) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRealms) ErrorTrace(errortrace bool) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRealms) FilterPath(filterpaths ...string) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRealms) Human(human bool) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRealms) Pretty(pretty bool) *ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}
