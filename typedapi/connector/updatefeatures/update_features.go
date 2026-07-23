package updatefeatures

import (
	gobytes "bytes"
	"context"
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

type UpdateFeatures struct {
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

type NewUpdateFeatures func(connectorid string) *UpdateFeatures

func NewUpdateFeaturesFunc(tp elastictransport.Interface) NewUpdateFeatures {
	_ = "STUB: not implemented"
	return *new(NewUpdateFeatures)
}

func New(tp elastictransport.Interface) *UpdateFeatures { _ = "STUB: not implemented"; return nil }

func (r *UpdateFeatures) Raw(raw io.Reader) *UpdateFeatures { _ = "STUB: not implemented"; return nil }

func (r *UpdateFeatures) Request(req *Request) *UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFeatures) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFeatures) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFeatures) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateFeatures) Header(key, value string) *UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFeatures) _connectorid(connectorid string) *UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFeatures) ErrorTrace(errortrace bool) *UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFeatures) FilterPath(filterpaths ...string) *UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFeatures) Human(human bool) *UpdateFeatures { _ = "STUB: not implemented"; return nil }

func (r *UpdateFeatures) Pretty(pretty bool) *UpdateFeatures { _ = "STUB: not implemented"; return nil }

func (r *UpdateFeatures) Features(features types.ConnectorFeaturesVariant) *UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}
