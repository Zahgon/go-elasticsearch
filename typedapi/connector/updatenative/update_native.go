package updatenative

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

type UpdateNative struct {
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

type NewUpdateNative func(connectorid string) *UpdateNative

func NewUpdateNativeFunc(tp elastictransport.Interface) NewUpdateNative {
	_ = "STUB: not implemented"
	return *new(NewUpdateNative)
}

func New(tp elastictransport.Interface) *UpdateNative { _ = "STUB: not implemented"; return nil }

func (r *UpdateNative) Raw(raw io.Reader) *UpdateNative { _ = "STUB: not implemented"; return nil }

func (r *UpdateNative) Request(req *Request) *UpdateNative { _ = "STUB: not implemented"; return nil }

func (r *UpdateNative) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateNative) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateNative) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateNative) Header(key, value string) *UpdateNative {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateNative) _connectorid(connectorid string) *UpdateNative {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateNative) ErrorTrace(errortrace bool) *UpdateNative {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateNative) FilterPath(filterpaths ...string) *UpdateNative {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateNative) Human(human bool) *UpdateNative { _ = "STUB: not implemented"; return nil }

func (r *UpdateNative) Pretty(pretty bool) *UpdateNative { _ = "STUB: not implemented"; return nil }

func (r *UpdateNative) IsNative(isnative bool) *UpdateNative { _ = "STUB: not implemented"; return nil }
