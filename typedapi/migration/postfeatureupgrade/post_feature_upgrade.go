package postfeatureupgrade

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostFeatureUpgrade struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostFeatureUpgrade func() *PostFeatureUpgrade

func NewPostFeatureUpgradeFunc(tp elastictransport.Interface) NewPostFeatureUpgrade {
	_ = "STUB: not implemented"
	return *new(NewPostFeatureUpgrade)
}

func New(tp elastictransport.Interface) *PostFeatureUpgrade { _ = "STUB: not implemented"; return nil }

func (r *PostFeatureUpgrade) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostFeatureUpgrade) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostFeatureUpgrade) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostFeatureUpgrade) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PostFeatureUpgrade) Header(key, value string) *PostFeatureUpgrade {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostFeatureUpgrade) ErrorTrace(errortrace bool) *PostFeatureUpgrade {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostFeatureUpgrade) FilterPath(filterpaths ...string) *PostFeatureUpgrade {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostFeatureUpgrade) Human(human bool) *PostFeatureUpgrade {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostFeatureUpgrade) Pretty(pretty bool) *PostFeatureUpgrade {
	_ = "STUB: not implemented"
	return nil
}
