package updateerror

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

type UpdateError struct {
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

type NewUpdateError func(connectorid string) *UpdateError

func NewUpdateErrorFunc(tp elastictransport.Interface) NewUpdateError {
	_ = "STUB: not implemented"
	return *new(NewUpdateError)
}

func New(tp elastictransport.Interface) *UpdateError { _ = "STUB: not implemented"; return nil }

func (r *UpdateError) Raw(raw io.Reader) *UpdateError { _ = "STUB: not implemented"; return nil }

func (r *UpdateError) Request(req *Request) *UpdateError { _ = "STUB: not implemented"; return nil }

func (r *UpdateError) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateError) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateError) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateError) Header(key, value string) *UpdateError { _ = "STUB: not implemented"; return nil }

func (r *UpdateError) _connectorid(connectorid string) *UpdateError {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateError) ErrorTrace(errortrace bool) *UpdateError {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateError) FilterPath(filterpaths ...string) *UpdateError {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateError) Human(human bool) *UpdateError { _ = "STUB: not implemented"; return nil }

func (r *UpdateError) Pretty(pretty bool) *UpdateError { _ = "STUB: not implemented"; return nil }

func (r *UpdateError) Error(error any) *UpdateError { _ = "STUB: not implemented"; return nil }
