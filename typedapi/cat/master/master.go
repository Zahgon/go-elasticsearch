package master

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catmastercolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Master struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewMaster func() *Master

func NewMasterFunc(tp elastictransport.Interface) NewMaster {
	_ = "STUB: not implemented"
	return *new(NewMaster)
}

func New(tp elastictransport.Interface) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Master) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Master) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Master) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Master) Header(key, value string) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) H(catmastercolumns ...catmastercolumn.CatMasterColumn) *Master {
	_ = "STUB: not implemented"
	return nil
}

func (r *Master) S(names ...string) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Local(local bool) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) MasterTimeout(duration string) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Bytes(bytes bytes.Bytes) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Format(format string) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Help(help bool) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Time(time timeunit.TimeUnit) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) V(v bool) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) ErrorTrace(errortrace bool) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) FilterPath(filterpaths ...string) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Human(human bool) *Master { _ = "STUB: not implemented"; return nil }

func (r *Master) Pretty(pretty bool) *Master { _ = "STUB: not implemented"; return nil }
