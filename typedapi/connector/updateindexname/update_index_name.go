package updateindexname

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

type UpdateIndexName struct {
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

type NewUpdateIndexName func(connectorid string) *UpdateIndexName

func NewUpdateIndexNameFunc(tp elastictransport.Interface) NewUpdateIndexName {
	_ = "STUB: not implemented"
	return *new(NewUpdateIndexName)
}

func New(tp elastictransport.Interface) *UpdateIndexName { _ = "STUB: not implemented"; return nil }

func (r *UpdateIndexName) Raw(raw io.Reader) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) Request(req *Request) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateIndexName) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateIndexName) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateIndexName) Header(key, value string) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) _connectorid(connectorid string) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) ErrorTrace(errortrace bool) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) FilterPath(filterpaths ...string) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) Human(human bool) *UpdateIndexName { _ = "STUB: not implemented"; return nil }

func (r *UpdateIndexName) Pretty(pretty bool) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateIndexName) IndexName(indexname any) *UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}
