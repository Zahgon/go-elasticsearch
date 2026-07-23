package updateservicetype

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
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateServiceType struct {
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

type NewUpdateServiceType func(connectorid string) *UpdateServiceType

func NewUpdateServiceTypeFunc(tp elastictransport.Interface) NewUpdateServiceType {
	_ = "STUB: not implemented"
	return *new(NewUpdateServiceType)
}

func New(tp elastictransport.Interface) *UpdateServiceType { _ = "STUB: not implemented"; return nil }

func (r *UpdateServiceType) Raw(raw io.Reader) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) Request(req *Request) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateServiceType) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateServiceType) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateServiceType) Header(key, value string) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) _connectorid(connectorid string) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) ErrorTrace(errortrace bool) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) FilterPath(filterpaths ...string) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) Human(human bool) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) Pretty(pretty bool) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateServiceType) ServiceType(servicetype string) *UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}
