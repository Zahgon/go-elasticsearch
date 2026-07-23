package updatesettings

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateSettings struct {
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

type NewUpdateSettings func() *UpdateSettings

func NewUpdateSettingsFunc(tp elastictransport.Interface) NewUpdateSettings {
	_ = "STUB: not implemented"
	return *new(NewUpdateSettings)
}

func New(tp elastictransport.Interface) *UpdateSettings { _ = "STUB: not implemented"; return nil }

func (r *UpdateSettings) Raw(raw io.Reader) *UpdateSettings { _ = "STUB: not implemented"; return nil }

func (r *UpdateSettings) Request(req *Request) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateSettings) Header(key, value string) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) MasterTimeout(duration string) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) Timeout(duration string) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) ErrorTrace(errortrace bool) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) FilterPath(filterpaths ...string) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) Human(human bool) *UpdateSettings { _ = "STUB: not implemented"; return nil }

func (r *UpdateSettings) Pretty(pretty bool) *UpdateSettings { _ = "STUB: not implemented"; return nil }

func (r *UpdateSettings) IndexAutoExpandReplicas(indexautoexpandreplicas string) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateSettings) IndexNumberOfReplicas(indexnumberofreplicas int) *UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}
