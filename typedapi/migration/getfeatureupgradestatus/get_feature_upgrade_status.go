package getfeatureupgradestatus

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetFeatureUpgradeStatus struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetFeatureUpgradeStatus func() *GetFeatureUpgradeStatus

func NewGetFeatureUpgradeStatusFunc(tp elastictransport.Interface) NewGetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return *new(NewGetFeatureUpgradeStatus)
}

func New(tp elastictransport.Interface) *GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatureUpgradeStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFeatureUpgradeStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFeatureUpgradeStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFeatureUpgradeStatus) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetFeatureUpgradeStatus) Header(key, value string) *GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatureUpgradeStatus) ErrorTrace(errortrace bool) *GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatureUpgradeStatus) FilterPath(filterpaths ...string) *GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatureUpgradeStatus) Human(human bool) *GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFeatureUpgradeStatus) Pretty(pretty bool) *GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}
