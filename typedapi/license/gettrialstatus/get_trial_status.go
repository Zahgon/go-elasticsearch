package gettrialstatus

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTrialStatus struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetTrialStatus func() *GetTrialStatus

func NewGetTrialStatusFunc(tp elastictransport.Interface) NewGetTrialStatus {
	_ = "STUB: not implemented"
	return *new(NewGetTrialStatus)
}

func New(tp elastictransport.Interface) *GetTrialStatus { _ = "STUB: not implemented"; return nil }

func (r *GetTrialStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrialStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrialStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrialStatus) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetTrialStatus) Header(key, value string) *GetTrialStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrialStatus) ErrorTrace(errortrace bool) *GetTrialStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrialStatus) FilterPath(filterpaths ...string) *GetTrialStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrialStatus) Human(human bool) *GetTrialStatus { _ = "STUB: not implemented"; return nil }

func (r *GetTrialStatus) Pretty(pretty bool) *GetTrialStatus { _ = "STUB: not implemented"; return nil }
