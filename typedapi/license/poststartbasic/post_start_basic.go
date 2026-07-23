package poststartbasic

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostStartBasic struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostStartBasic func() *PostStartBasic

func NewPostStartBasicFunc(tp elastictransport.Interface) NewPostStartBasic {
	_ = "STUB: not implemented"
	return *new(NewPostStartBasic)
}

func New(tp elastictransport.Interface) *PostStartBasic { _ = "STUB: not implemented"; return nil }

func (r *PostStartBasic) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostStartBasic) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostStartBasic) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostStartBasic) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PostStartBasic) Header(key, value string) *PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartBasic) Acknowledge(acknowledge bool) *PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartBasic) MasterTimeout(duration string) *PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartBasic) Timeout(duration string) *PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartBasic) ErrorTrace(errortrace bool) *PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartBasic) FilterPath(filterpaths ...string) *PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostStartBasic) Human(human bool) *PostStartBasic { _ = "STUB: not implemented"; return nil }

func (r *PostStartBasic) Pretty(pretty bool) *PostStartBasic { _ = "STUB: not implemented"; return nil }
