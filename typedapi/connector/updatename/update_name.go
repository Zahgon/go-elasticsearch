package updatename

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

type UpdateName struct {
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

type NewUpdateName func(connectorid string) *UpdateName

func NewUpdateNameFunc(tp elastictransport.Interface) NewUpdateName {
	_ = "STUB: not implemented"
	return *new(NewUpdateName)
}

func New(tp elastictransport.Interface) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) Raw(raw io.Reader) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) Request(req *Request) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateName) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateName) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateName) Header(key, value string) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) _connectorid(connectorid string) *UpdateName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateName) ErrorTrace(errortrace bool) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) FilterPath(filterpaths ...string) *UpdateName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateName) Human(human bool) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) Pretty(pretty bool) *UpdateName { _ = "STUB: not implemented"; return nil }

func (r *UpdateName) Description(description string) *UpdateName {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateName) Name(name string) *UpdateName { _ = "STUB: not implemented"; return nil }
