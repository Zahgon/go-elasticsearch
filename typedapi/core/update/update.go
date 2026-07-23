package update

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	idMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Update struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdate func(index, id string) *Update

func NewUpdateFunc(tp elastictransport.Interface) NewUpdate {
	_ = "STUB: not implemented"
	return *new(NewUpdate)
}

func New(tp elastictransport.Interface) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Raw(raw io.Reader) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Request(req *Request) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Update) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Update) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Update) Header(key, value string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) _id(id string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) _index(index string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) IfPrimaryTerm(ifprimaryterm string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) IfSeqNo(sequencenumber string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) IncludeSourceOnError(includesourceonerror bool) *Update {
	_ = "STUB: not implemented"
	return nil
}

func (r *Update) Lang(lang string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Refresh(refresh refresh.Refresh) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) RequireAlias(requirealias bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) RetryOnConflict(retryonconflict int) *Update {
	_ = "STUB: not implemented"
	return nil
}

func (r *Update) Routing(routings ...string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Timeout(duration string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) WaitForActiveShards(waitforactiveshards string) *Update {
	_ = "STUB: not implemented"
	return nil
}

func (r *Update) SourceExcludes_(fields ...string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) SourceIncludes_(fields ...string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) ErrorTrace(errortrace bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) FilterPath(filterpaths ...string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Human(human bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Pretty(pretty bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) DetectNoop(detectnoop bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Doc(doc any) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) DocAsUpsert(docasupsert bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Script(script types.ScriptVariant) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) ScriptedUpsert(scriptedupsert bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Source_(sourceconfig types.SourceConfigVariant) *Update {
	_ = "STUB: not implemented"
	return nil
}

func (r *Update) Upsert(upsert any) *Update { _ = "STUB: not implemented"; return nil }
