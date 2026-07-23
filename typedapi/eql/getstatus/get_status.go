package getstatus

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

type GetStatus struct {
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

type NewGetStatus func(id string) *GetStatus

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

func (r *GetStatus) _id(id string) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) ErrorTrace(errortrace bool) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) FilterPath(filterpaths ...string) *GetStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetStatus) Human(human bool) *GetStatus { _ = "STUB: not implemented"; return nil }

func (r *GetStatus) Pretty(pretty bool) *GetStatus { _ = "STUB: not implemented"; return nil }
