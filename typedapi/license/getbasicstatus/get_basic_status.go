package getbasicstatus

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetBasicStatus struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetBasicStatus func() *GetBasicStatus

func NewGetBasicStatusFunc(tp elastictransport.Interface) NewGetBasicStatus {
	_ = "STUB: not implemented"
	return *new(NewGetBasicStatus)
}

func New(tp elastictransport.Interface) *GetBasicStatus { _ = "STUB: not implemented"; return nil }

func (r *GetBasicStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBasicStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBasicStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBasicStatus) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetBasicStatus) Header(key, value string) *GetBasicStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBasicStatus) ErrorTrace(errortrace bool) *GetBasicStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBasicStatus) FilterPath(filterpaths ...string) *GetBasicStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBasicStatus) Human(human bool) *GetBasicStatus { _ = "STUB: not implemented"; return nil }

func (r *GetBasicStatus) Pretty(pretty bool) *GetBasicStatus { _ = "STUB: not implemented"; return nil }
