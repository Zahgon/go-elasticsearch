package getsettings

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetSettings func() *GetSettings

func NewGetSettingsFunc(tp elastictransport.Interface) NewGetSettings {
	_ = "STUB: not implemented"
	return *new(NewGetSettings)
}

func New(tp elastictransport.Interface) *GetSettings { _ = "STUB: not implemented"; return nil }

func (r *GetSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSettings) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSettings) Header(key, value string) *GetSettings { _ = "STUB: not implemented"; return nil }

func (r *GetSettings) MasterTimeout(duration string) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) ErrorTrace(errortrace bool) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) FilterPath(filterpaths ...string) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) Human(human bool) *GetSettings { _ = "STUB: not implemented"; return nil }

func (r *GetSettings) Pretty(pretty bool) *GetSettings { _ = "STUB: not implemented"; return nil }
