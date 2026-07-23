package getasyncstatus

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetAsyncStatus struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetAsyncStatus func(id string) *GetAsyncStatus

func NewGetAsyncStatusFunc(tp elastictransport.Interface) NewGetAsyncStatus {
	_ = "STUB: not implemented"
	return *new(NewGetAsyncStatus)
}

func New(tp elastictransport.Interface) *GetAsyncStatus { _ = "STUB: not implemented"; return nil }

func (r *GetAsyncStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAsyncStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAsyncStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAsyncStatus) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetAsyncStatus) Header(key, value string) *GetAsyncStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAsyncStatus) _id(id string) *GetAsyncStatus { _ = "STUB: not implemented"; return nil }

func (r *GetAsyncStatus) ErrorTrace(errortrace bool) *GetAsyncStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAsyncStatus) FilterPath(filterpaths ...string) *GetAsyncStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAsyncStatus) Human(human bool) *GetAsyncStatus { _ = "STUB: not implemented"; return nil }

func (r *GetAsyncStatus) Pretty(pretty bool) *GetAsyncStatus { _ = "STUB: not implemented"; return nil }
