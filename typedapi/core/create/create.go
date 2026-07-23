package create

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      any
	deferred []func(request any) error
	buf      *gobytes.Buffer

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCreate func(index, id string) *Create

func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	_ = "STUB: not implemented"
	return *new(NewCreate)
}

func New(tp elastictransport.Interface) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Raw(raw io.Reader) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Request(req any) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Document(document any) *Create { _ = "STUB: not implemented"; return nil }

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

func (r *Create) _id(id string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) _index(index string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) IncludeSourceOnError(includesourceonerror bool) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) Pipeline(pipeline string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Refresh(refresh refresh.Refresh) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) RequireAlias(requirealias bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) RequireDataStream(requiredatastream bool) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) Routing(routings ...string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Timeout(duration string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Version(versionnumber string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) VersionType(versiontype versiontype.VersionType) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) WaitForActiveShards(waitforactiveshards string) *Create {
	_ = "STUB: not implemented"
	return nil
}

func (r *Create) ErrorTrace(errortrace bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) FilterPath(filterpaths ...string) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Human(human bool) *Create { _ = "STUB: not implemented"; return nil }

func (r *Create) Pretty(pretty bool) *Create { _ = "STUB: not implemented"; return nil }
