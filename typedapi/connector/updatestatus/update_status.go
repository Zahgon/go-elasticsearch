package updatestatus

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectorstatus"
)

const (
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateStatus struct {
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

type NewUpdateStatus func(connectorid string) *UpdateStatus

func NewUpdateStatusFunc(tp elastictransport.Interface) NewUpdateStatus {
	_ = "STUB: not implemented"
	return *new(NewUpdateStatus)
}

func New(tp elastictransport.Interface) *UpdateStatus { _ = "STUB: not implemented"; return nil }

func (r *UpdateStatus) Raw(raw io.Reader) *UpdateStatus { _ = "STUB: not implemented"; return nil }

func (r *UpdateStatus) Request(req *Request) *UpdateStatus { _ = "STUB: not implemented"; return nil }

func (r *UpdateStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateStatus) Header(key, value string) *UpdateStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateStatus) _connectorid(connectorid string) *UpdateStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateStatus) ErrorTrace(errortrace bool) *UpdateStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateStatus) FilterPath(filterpaths ...string) *UpdateStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateStatus) Human(human bool) *UpdateStatus { _ = "STUB: not implemented"; return nil }

func (r *UpdateStatus) Pretty(pretty bool) *UpdateStatus { _ = "STUB: not implemented"; return nil }

func (r *UpdateStatus) Status(status connectorstatus.ConnectorStatus) *UpdateStatus {
	_ = "STUB: not implemented"
	return nil
}
