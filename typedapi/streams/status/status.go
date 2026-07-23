package status

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Status struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStatus func() *Status

func NewStatusFunc(tp elastictransport.Interface) NewStatus {
	_ = "STUB: not implemented"
	return *new(NewStatus)
}

func New(tp elastictransport.Interface) *Status { _ = "STUB: not implemented"; return nil }

func (r *Status) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Status) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Status) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Status) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Status) Header(key, value string) *Status { _ = "STUB: not implemented"; return nil }

func (r *Status) MasterTimeout(duration string) *Status { _ = "STUB: not implemented"; return nil }

func (r *Status) ErrorTrace(errortrace bool) *Status { _ = "STUB: not implemented"; return nil }

func (r *Status) FilterPath(filterpaths ...string) *Status { _ = "STUB: not implemented"; return nil }

func (r *Status) Human(human bool) *Status { _ = "STUB: not implemented"; return nil }

func (r *Status) Pretty(pretty bool) *Status { _ = "STUB: not implemented"; return nil }
