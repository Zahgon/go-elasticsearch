package health

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/level"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/waitforevents"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Health struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

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

func (r Health) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Health) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Health) Header(key, value string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Index(index string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) Level(level level.Level) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Local(local bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) MasterTimeout(duration string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Timeout(duration string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) WaitForActiveShards(waitforactiveshards string) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) WaitForEvents(waitforevents waitforevents.WaitForEvents) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) WaitForNodes(waitfornodes string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) WaitForNoInitializingShards(waitfornoinitializingshards bool) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) WaitForNoRelocatingShards(waitfornorelocatingshards bool) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) WaitForStatus(waitforstatus healthstatus.HealthStatus) *Health {
	_ = "STUB: not implemented"
	return nil
}

func (r *Health) ErrorTrace(errortrace bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) FilterPath(filterpaths ...string) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Human(human bool) *Health { _ = "STUB: not implemented"; return nil }

func (r *Health) Pretty(pretty bool) *Health { _ = "STUB: not implemented"; return nil }
