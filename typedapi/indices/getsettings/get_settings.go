package getsettings

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string
	name  string

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

func (r GetSettings) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetSettings) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSettings) Header(key, value string) *GetSettings { _ = "STUB: not implemented"; return nil }

func (r *GetSettings) Index(index string) *GetSettings { _ = "STUB: not implemented"; return nil }

func (r *GetSettings) Name(name string) *GetSettings { _ = "STUB: not implemented"; return nil }

func (r *GetSettings) AllowNoIndices(allownoindices bool) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) FlatSettings(flatsettings bool) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) IgnoreUnavailable(ignoreunavailable bool) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) IncludeDefaults(includedefaults bool) *GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSettings) Local(local bool) *GetSettings { _ = "STUB: not implemented"; return nil }

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
