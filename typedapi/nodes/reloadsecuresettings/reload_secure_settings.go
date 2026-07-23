package reloadsecuresettings

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ReloadSecureSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewReloadSecureSettings func() *ReloadSecureSettings

func NewReloadSecureSettingsFunc(tp elastictransport.Interface) NewReloadSecureSettings {
	_ = "STUB: not implemented"
	return *new(NewReloadSecureSettings)
}

func New(tp elastictransport.Interface) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) Raw(raw io.Reader) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) Request(req *Request) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReloadSecureSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReloadSecureSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReloadSecureSettings) Header(key, value string) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) NodeId(nodeid string) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) Timeout(duration string) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) ErrorTrace(errortrace bool) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) FilterPath(filterpaths ...string) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) Human(human bool) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) Pretty(pretty bool) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSecureSettings) SecureSettingsPassword(password string) *ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}
