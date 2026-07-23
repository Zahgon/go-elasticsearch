package enrollkibana

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type EnrollKibana struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewEnrollKibana func() *EnrollKibana

func NewEnrollKibanaFunc(tp elastictransport.Interface) NewEnrollKibana {
	_ = "STUB: not implemented"
	return *new(NewEnrollKibana)
}

func New(tp elastictransport.Interface) *EnrollKibana { _ = "STUB: not implemented"; return nil }

func (r *EnrollKibana) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnrollKibana) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnrollKibana) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnrollKibana) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *EnrollKibana) Header(key, value string) *EnrollKibana {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnrollKibana) ErrorTrace(errortrace bool) *EnrollKibana {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnrollKibana) FilterPath(filterpaths ...string) *EnrollKibana {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnrollKibana) Human(human bool) *EnrollKibana { _ = "STUB: not implemented"; return nil }

func (r *EnrollKibana) Pretty(pretty bool) *EnrollKibana { _ = "STUB: not implemented"; return nil }
