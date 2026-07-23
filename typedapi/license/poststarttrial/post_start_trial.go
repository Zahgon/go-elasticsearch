package poststarttrial

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostStartTrial struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostStartTrial func() *PostStartTrial

func NewPostStartTrialFunc(tp elastictransport.Interface) NewPostStartTrial {
	_ = "STUB: not implemented"
	return *new(NewPostStartTrial)
}

func New(tp elastictransport.Interface) *PostStartTrial { _ = "STUB: not implemented"; return nil }

func (r *PostStartTrial) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostStartTrial) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostStartTrial) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostStartTrial) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PostStartTrial) Header(key, value string) *PostStartTrial {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartTrial) Acknowledge(acknowledge bool) *PostStartTrial {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartTrial) Type(type_ string) *PostStartTrial { _ = "STUB: not implemented"; return nil }

func (r *PostStartTrial) MasterTimeout(duration string) *PostStartTrial {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartTrial) ErrorTrace(errortrace bool) *PostStartTrial {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartTrial) FilterPath(filterpaths ...string) *PostStartTrial {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartTrial) Human(human bool) *PostStartTrial { _ = "STUB: not implemented"; return nil }

func (r *PostStartTrial) Pretty(pretty bool) *PostStartTrial { _ = "STUB: not implemented"; return nil }
