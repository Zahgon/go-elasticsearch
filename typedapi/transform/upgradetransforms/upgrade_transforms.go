package upgradetransforms

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpgradeTransforms struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpgradeTransforms func() *UpgradeTransforms

func NewUpgradeTransformsFunc(tp elastictransport.Interface) NewUpgradeTransforms {
	_ = "STUB: not implemented"
	return *new(NewUpgradeTransforms)
}

func New(tp elastictransport.Interface) *UpgradeTransforms { _ = "STUB: not implemented"; return nil }

func (r *UpgradeTransforms) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpgradeTransforms) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpgradeTransforms) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpgradeTransforms) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *UpgradeTransforms) Header(key, value string) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeTransforms) DryRun(dryrun bool) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeTransforms) Timeout(duration string) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeTransforms) ErrorTrace(errortrace bool) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeTransforms) FilterPath(filterpaths ...string) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeTransforms) Human(human bool) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeTransforms) Pretty(pretty bool) *UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}
