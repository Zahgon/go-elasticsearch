package capabilities

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/restmethod"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Capabilities struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCapabilities func() *Capabilities

func NewCapabilitiesFunc(tp elastictransport.Interface) NewCapabilities {
	_ = "STUB: not implemented"
	return *new(NewCapabilities)
}

func New(tp elastictransport.Interface) *Capabilities { _ = "STUB: not implemented"; return nil }

func (r *Capabilities) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Capabilities) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Capabilities) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Capabilities) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Capabilities) Header(key, value string) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) Method(method restmethod.RestMethod) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) Path(path string) *Capabilities { _ = "STUB: not implemented"; return nil }

func (r *Capabilities) Parameters(parameters ...string) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) Capabilities(capabilities ...string) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) LocalOnly(localonly bool) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) Timeout(duration string) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) ErrorTrace(errortrace bool) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) FilterPath(filterpaths ...string) *Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (r *Capabilities) Human(human bool) *Capabilities { _ = "STUB: not implemented"; return nil }

func (r *Capabilities) Pretty(pretty bool) *Capabilities { _ = "STUB: not implemented"; return nil }
