package updateconfiguration

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateConfiguration struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateConfiguration func(connectorid string) *UpdateConfiguration

func NewUpdateConfigurationFunc(tp elastictransport.Interface) NewUpdateConfiguration {
	_ = "STUB: not implemented"
	return *new(NewUpdateConfiguration)
}

func New(tp elastictransport.Interface) *UpdateConfiguration { _ = "STUB: not implemented"; return nil }

func (r *UpdateConfiguration) Raw(raw io.Reader) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) Request(req *Request) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateConfiguration) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateConfiguration) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateConfiguration) Header(key, value string) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) _connectorid(connectorid string) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) ErrorTrace(errortrace bool) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) FilterPath(filterpaths ...string) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) Human(human bool) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) Pretty(pretty bool) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) Configuration(connectorconfiguration types.ConnectorConfigurationVariant) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) Values(values map[string]json.RawMessage) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateConfiguration) AddValue(key string, value json.RawMessage) *UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}
