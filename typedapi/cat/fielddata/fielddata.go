package fielddata

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catfielddatacolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	fieldsMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Fielddata struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	fields string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewFielddata func() *Fielddata

func NewFielddataFunc(tp elastictransport.Interface) NewFielddata {
	_ = "STUB: not implemented"
	return *new(NewFielddata)
}

func New(tp elastictransport.Interface) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Fielddata) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Fielddata) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Fielddata) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Fielddata) Header(key, value string) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) Fields(fields string) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) H(catfielddatacolumns ...catfielddatacolumn.CatFieldDataColumn) *Fielddata {
	_ = "STUB: not implemented"
	return nil
}

func (r *Fielddata) S(names ...string) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) Bytes(bytes bytes.Bytes) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) Format(format string) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) Help(help bool) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) Time(time timeunit.TimeUnit) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) V(v bool) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) ErrorTrace(errortrace bool) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) FilterPath(filterpaths ...string) *Fielddata {
	_ = "STUB: not implemented"
	return nil
}

func (r *Fielddata) Human(human bool) *Fielddata { _ = "STUB: not implemented"; return nil }

func (r *Fielddata) Pretty(pretty bool) *Fielddata { _ = "STUB: not implemented"; return nil }
