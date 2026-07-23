package asyncquery

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/esqlformat"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AsyncQuery struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewAsyncQuery func() *AsyncQuery

func NewAsyncQueryFunc(tp elastictransport.Interface) NewAsyncQuery {
	_ = "STUB: not implemented"
	return *new(NewAsyncQuery)
}

func New(tp elastictransport.Interface) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Raw(raw io.Reader) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Request(req *Request) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AsyncQuery) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *AsyncQuery) Header(key, value string) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) AllowPartialResults(allowpartialresults bool) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) Delimiter(delimiter string) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) DropNullColumns(dropnullcolumns bool) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) Format(format esqlformat.EsqlFormat) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) ErrorTrace(errortrace bool) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) FilterPath(filterpaths ...string) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) Human(human bool) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Pretty(pretty bool) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Columnar(columnar bool) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Filter(filter types.QueryVariant) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) IncludeCcsMetadata(includeccsmetadata bool) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) IncludeExecutionMetadata(includeexecutionmetadata bool) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) KeepAlive(duration types.DurationVariant) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) KeepOnCompletion(keeponcompletion bool) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) Locale(locale string) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Params(esqlparams types.ESQLParamsVariant) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) Profile(profile bool) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) ProjectRouting(projectrouting string) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) Query(query string) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) Tables(tables map[string]map[string]types.TableValuesContainer) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *AsyncQuery) TimeZone(timezone string) *AsyncQuery { _ = "STUB: not implemented"; return nil }

func (r *AsyncQuery) WaitForCompletionTimeout(duration types.DurationVariant) *AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}
