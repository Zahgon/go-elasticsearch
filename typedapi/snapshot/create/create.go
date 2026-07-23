package create

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCreate func(repository, snapshot string) *Create

func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	_ = "STUB: not implemented"
	return *new(NewCreate)
}

func New(tp elastictransport.Interface) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Raw(raw io.Reader) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Request(req *Request) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Create) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Create) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Create) Header(key, value string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) _repository(repository string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) _snapshot(snapshot string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) MasterTimeout(duration string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) WaitForCompletion(waitforcompletion bool) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) ErrorTrace(errortrace bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) FilterPath(filterpaths ...string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Human(human bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Pretty(pretty bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) FeatureStates(featurestates ...string) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) IgnoreUnavailable(ignoreunavailable bool) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) IncludeGlobalState(includeglobalstate bool) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) Indices(indices ...string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Metadata(metadata types.MetadataVariant) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) Partial(partial bool) *Create { _ = "STUB: not implemented"; return nil }
