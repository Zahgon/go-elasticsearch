package componenttemplates

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catcomponentcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ComponentTemplates struct {
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

type NewComponentTemplates func() *ComponentTemplates

func NewComponentTemplatesFunc(tp elastictransport.Interface) NewComponentTemplates {
	_ = "STUB: not implemented"
	return *new(NewComponentTemplates)
}

func New(tp elastictransport.Interface) *ComponentTemplates { _ = "STUB: not implemented"; return nil }

func (r *ComponentTemplates) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ComponentTemplates) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ComponentTemplates) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r ComponentTemplates) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ComponentTemplates) Header(key, value string) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Name(name string) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) H(catcomponentcolumns ...catcomponentcolumn.CatComponentColumn) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) S(names ...string) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Local(local bool) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) MasterTimeout(duration string) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Bytes(bytes bytes.Bytes) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Format(format string) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Help(help bool) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Time(time timeunit.TimeUnit) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) V(v bool) *ComponentTemplates { _ = "STUB: not implemented"; return nil }

func (r *ComponentTemplates) ErrorTrace(errortrace bool) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) FilterPath(filterpaths ...string) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Human(human bool) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (r *ComponentTemplates) Pretty(pretty bool) *ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}
