package health

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cathealthcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Health struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewHealth func() *Health

func NewHealthFunc(tp elastictransport.Interface) NewHealth {
	_ = "STUB: not implemented"
	return *new(NewHealth)
}

func New(tp elastictransport.Interface) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Health) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Health) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Health) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Health) Header(key, value string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Ts(ts bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) H(cathealthcolumns ...cathealthcolumn.CatHealthColumn) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) S(names ...string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Bytes(bytes bytes.Bytes) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Format(format string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Help(help bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Time(time timeunit.TimeUnit) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) V(v bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) ErrorTrace(errortrace bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) FilterPath(filterpaths ...string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Human(human bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Pretty(pretty bool) *Health { _ = "STUB: not implemented"; return nil }
