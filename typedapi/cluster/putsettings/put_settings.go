package putsettings

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSettings struct {
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

type NewPutSettings func() *PutSettings

func NewPutSettingsFunc(tp elastictransport.Interface) NewPutSettings {
	_ = "STUB: not implemented"
	return *new(NewPutSettings)
}

func New(tp elastictransport.Interface) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Raw(raw io.Reader) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Request(req *Request) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutSettings) Header(key, value string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) FlatSettings(flatsettings bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MasterTimeout(duration string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Timeout(duration string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) ErrorTrace(errortrace bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) FilterPath(filterpaths ...string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Human(human bool) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Pretty(pretty bool) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Persistent(persistent map[string]json.RawMessage) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) AddPersistent(key string, value json.RawMessage) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Transient(transient map[string]json.RawMessage) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) AddTransient(key string, value json.RawMessage) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}
