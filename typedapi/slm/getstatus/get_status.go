package getstatus

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetStatus struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetStatus func() *GetStatus

func NewGetStatusFunc(tp elastictransport.Interface) NewGetStatus {
	_ = "STUB: not implemented"
	return *new(NewGetStatus)
}

func New(tp elastictransport.Interface) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetStatus) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetStatus) Header(key, value string) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) MasterTimeout(duration string) *GetStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetStatus) Timeout(duration string) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) ErrorTrace(errortrace bool) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) FilterPath(filterpaths ...string) *GetStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetStatus) Human(human bool) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) Pretty(pretty bool) *GetStatus { _ = "STUB: not implemented"; return nil }
