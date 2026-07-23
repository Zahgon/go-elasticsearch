package resetfeatures

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResetFeatures struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewResetFeatures func() *ResetFeatures

func NewResetFeaturesFunc(tp elastictransport.Interface) NewResetFeatures {
	_ = "STUB: not implemented"
	return *new(NewResetFeatures)
}

func New(tp elastictransport.Interface) *ResetFeatures { _ = "STUB: not implemented"; return nil }

func (r *ResetFeatures) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetFeatures) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetFeatures) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetFeatures) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ResetFeatures) Header(key, value string) *ResetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetFeatures) MasterTimeout(duration string) *ResetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetFeatures) ErrorTrace(errortrace bool) *ResetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetFeatures) FilterPath(filterpaths ...string) *ResetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetFeatures) Human(human bool) *ResetFeatures { _ = "STUB: not implemented"; return nil }

func (r *ResetFeatures) Pretty(pretty bool) *ResetFeatures { _ = "STUB: not implemented"; return nil }
