package plugins

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catpluginscolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Plugins struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPlugins func() *Plugins

func NewPluginsFunc(tp elastictransport.Interface) NewPlugins {
	_ = "STUB: not implemented"
	return *new(NewPlugins)
}

func New(tp elastictransport.Interface) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Plugins) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Plugins) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Plugins) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Plugins) Header(key, value string) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) H(catpluginscolumns ...catpluginscolumn.CatPluginsColumn) *Plugins {
	_ = "STUB: not implemented"
	return nil
}

func (r *Plugins) S(names ...string) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) IncludeBootstrap(includebootstrap bool) *Plugins {
	_ = "STUB: not implemented"
	return nil
}

func (r *Plugins) Local(local bool) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) MasterTimeout(duration string) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) Bytes(bytes bytes.Bytes) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) Format(format string) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) Help(help bool) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) Time(time timeunit.TimeUnit) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) V(v bool) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) ErrorTrace(errortrace bool) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) FilterPath(filterpaths ...string) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) Human(human bool) *Plugins { _ = "STUB: not implemented"; return nil }

func (r *Plugins) Pretty(pretty bool) *Plugins { _ = "STUB: not implemented"; return nil }
