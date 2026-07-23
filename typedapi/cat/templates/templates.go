package templates

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cattemplatescolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Templates struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTemplates func() *Templates

func NewTemplatesFunc(tp elastictransport.Interface) NewTemplates {
	_ = "STUB: not implemented"
	return *new(NewTemplates)
}

func New(tp elastictransport.Interface) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Templates) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Templates) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Templates) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Templates) Header(key, value string) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) Name(name string) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) H(cattemplatescolumns ...cattemplatescolumn.CatTemplatesColumn) *Templates {
	_ = "STUB: not implemented"
	return nil
}

func (r *Templates) S(names ...string) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) Local(local bool) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) MasterTimeout(duration string) *Templates {
	_ = "STUB: not implemented"
	return nil
}

func (r *Templates) Bytes(bytes bytes.Bytes) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) Format(format string) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) Help(help bool) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) Time(time timeunit.TimeUnit) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) V(v bool) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) ErrorTrace(errortrace bool) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) FilterPath(filterpaths ...string) *Templates {
	_ = "STUB: not implemented"
	return nil
}

func (r *Templates) Human(human bool) *Templates { _ = "STUB: not implemented"; return nil }

func (r *Templates) Pretty(pretty bool) *Templates { _ = "STUB: not implemented"; return nil }
