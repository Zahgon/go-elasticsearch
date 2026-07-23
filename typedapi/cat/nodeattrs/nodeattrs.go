package nodeattrs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catnodeattrscolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Nodeattrs struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewNodeattrs func() *Nodeattrs

func NewNodeattrsFunc(tp elastictransport.Interface) NewNodeattrs {
	_ = "STUB: not implemented"
	return *new(NewNodeattrs)
}

func New(tp elastictransport.Interface) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Nodeattrs) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Nodeattrs) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Nodeattrs) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Nodeattrs) Header(key, value string) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) H(catnodeattrscolumns ...catnodeattrscolumn.CatNodeattrsColumn) *Nodeattrs {
	_ = "STUB: not implemented"
	return nil
}

func (r *Nodeattrs) S(names ...string) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) Local(local bool) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) MasterTimeout(duration string) *Nodeattrs {
	_ = "STUB: not implemented"
	return nil
}

func (r *Nodeattrs) Bytes(bytes bytes.Bytes) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) Format(format string) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) Help(help bool) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) Time(time timeunit.TimeUnit) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) V(v bool) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) ErrorTrace(errortrace bool) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) FilterPath(filterpaths ...string) *Nodeattrs {
	_ = "STUB: not implemented"
	return nil
}

func (r *Nodeattrs) Human(human bool) *Nodeattrs { _ = "STUB: not implemented"; return nil }

func (r *Nodeattrs) Pretty(pretty bool) *Nodeattrs { _ = "STUB: not implemented"; return nil }
