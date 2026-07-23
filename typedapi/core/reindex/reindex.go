package reindex

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conflicts"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Reindex struct {
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

type NewReindex func() *Reindex

func NewReindexFunc(tp elastictransport.Interface) NewReindex {
	_ = "STUB: not implemented"
	return *new(NewReindex)
}

func New(tp elastictransport.Interface) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Raw(raw io.Reader) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Request(req *Request) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Reindex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Reindex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reindex) Header(key, value string) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Refresh(refresh bool) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) RequestsPerSecond(requestspersecond string) *Reindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reindex) Scroll(duration string) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Slices(slices string) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Timeout(duration string) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) WaitForActiveShards(waitforactiveshards string) *Reindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reindex) WaitForCompletion(waitforcompletion bool) *Reindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reindex) RequireAlias(requirealias bool) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) ErrorTrace(errortrace bool) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) FilterPath(filterpaths ...string) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Human(human bool) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Pretty(pretty bool) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Conflicts(conflicts conflicts.Conflicts) *Reindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reindex) Dest(dest types.ReindexDestinationVariant) *Reindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reindex) MaxDocs(maxdocs int64) *Reindex { _ = "STUB: not implemented"; return nil }

func (r *Reindex) Script(script types.ScriptVariant) *Reindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reindex) Source(source types.ReindexSourceVariant) *Reindex {
	_ = "STUB: not implemented"
	return nil
}
