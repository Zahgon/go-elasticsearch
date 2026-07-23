package setupgrademode

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SetUpgradeMode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSetUpgradeMode func() *SetUpgradeMode

func NewSetUpgradeModeFunc(tp elastictransport.Interface) NewSetUpgradeMode {
	_ = "STUB: not implemented"
	return *new(NewSetUpgradeMode)
}

func New(tp elastictransport.Interface) *SetUpgradeMode { _ = "STUB: not implemented"; return nil }

func (r *SetUpgradeMode) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SetUpgradeMode) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SetUpgradeMode) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SetUpgradeMode) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SetUpgradeMode) Header(key, value string) *SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (r *SetUpgradeMode) Enabled(enabled bool) *SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (r *SetUpgradeMode) Timeout(duration string) *SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (r *SetUpgradeMode) ErrorTrace(errortrace bool) *SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (r *SetUpgradeMode) FilterPath(filterpaths ...string) *SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (r *SetUpgradeMode) Human(human bool) *SetUpgradeMode { _ = "STUB: not implemented"; return nil }

func (r *SetUpgradeMode) Pretty(pretty bool) *SetUpgradeMode { _ = "STUB: not implemented"; return nil }
