package getfeatures

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetFeatures struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetFeatures func() *GetFeatures

func NewGetFeaturesFunc(tp elastictransport.Interface) NewGetFeatures {
	_ = "STUB: not implemented"
	return *new(NewGetFeatures)
}

func New(tp elastictransport.Interface) *GetFeatures { _ = "STUB: not implemented"; return nil }

func (r *GetFeatures) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFeatures) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFeatures) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFeatures) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetFeatures) Header(key, value string) *GetFeatures { _ = "STUB: not implemented"; return nil }

func (r *GetFeatures) MasterTimeout(duration string) *GetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatures) ErrorTrace(errortrace bool) *GetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatures) FilterPath(filterpaths ...string) *GetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatures) Human(human bool) *GetFeatures { _ = "STUB: not implemented"; return nil }

func (r *GetFeatures) Pretty(pretty bool) *GetFeatures { _ = "STUB: not implemented"; return nil }
