package nodes

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catnodecolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Nodes struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewNodes func() *Nodes

func NewNodesFunc(tp elastictransport.Interface) NewNodes {
	_ = "STUB: not implemented"
	return *new(NewNodes)
}

func New(tp elastictransport.Interface) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Nodes) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Nodes) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Nodes) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Nodes) Header(key, value string) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) FullId(fullid bool) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) IncludeUnloadedSegments(includeunloadedsegments bool) *Nodes {
	_ = "STUB: not implemented"
	return nil
}

func (r *Nodes) H(catnodecolumns ...catnodecolumn.CatNodeColumn) *Nodes {
	_ = "STUB: not implemented"
	return nil
}

func (r *Nodes) S(names ...string) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) MasterTimeout(duration string) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) Bytes(bytes bytes.Bytes) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) Format(format string) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) Help(help bool) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) Time(time timeunit.TimeUnit) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) V(v bool) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) ErrorTrace(errortrace bool) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) FilterPath(filterpaths ...string) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) Human(human bool) *Nodes { _ = "STUB: not implemented"; return nil }

func (r *Nodes) Pretty(pretty bool) *Nodes { _ = "STUB: not implemented"; return nil }
