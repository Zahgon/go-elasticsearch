package updatescheduling

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

type UpdateScheduling struct {
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

type NewUpdateScheduling func(connectorid string) *UpdateScheduling

func NewUpdateSchedulingFunc(tp elastictransport.Interface) NewUpdateScheduling {
	_ = "STUB: not implemented"
	return *new(NewUpdateScheduling)
}

func New(tp elastictransport.Interface) *UpdateScheduling { _ = "STUB: not implemented"; return nil }

func (r *UpdateScheduling) Raw(raw io.Reader) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) Request(req *Request) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateScheduling) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateScheduling) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateScheduling) Header(key, value string) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) _connectorid(connectorid string) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) ErrorTrace(errortrace bool) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) FilterPath(filterpaths ...string) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) Human(human bool) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) Pretty(pretty bool) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateScheduling) Scheduling(scheduling types.SchedulingConfigurationVariant) *UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}
