package index

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Index struct {
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

type NewIndex func(index string) *Index

func NewIndexFunc(tp elastictransport.Interface) NewIndex {
	_ = "STUB: not implemented"
	return *new(NewIndex)
}

func New(tp elastictransport.Interface) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Raw(raw io.Reader) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Request(req any) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Document(document any) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Index) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Index) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Index) Header(key, value string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Id(id string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) _index(index string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) IfPrimaryTerm(ifprimaryterm string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) IfSeqNo(sequencenumber string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) IncludeSourceOnError(includesourceonerror bool) *Index {
	_ = "STUB: not implemented"
	return nil
}

func (r *Index) OpType(optype optype.OpType) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Pipeline(pipeline string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Refresh(refresh refresh.Refresh) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Routing(routings ...string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Timeout(duration string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Version(versionnumber string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) VersionType(versiontype versiontype.VersionType) *Index {
	_ = "STUB: not implemented"
	return nil
}

func (r *Index) WaitForActiveShards(waitforactiveshards string) *Index {
	_ = "STUB: not implemented"
	return nil
}

func (r *Index) RequireAlias(requirealias bool) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) RequireDataStream(requiredatastream bool) *Index {
	_ = "STUB: not implemented"
	return nil
}

func (r *Index) ErrorTrace(errortrace bool) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) FilterPath(filterpaths ...string) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Human(human bool) *Index { _ = "STUB: not implemented"; return nil }

func (r *Index) Pretty(pretty bool) *Index { _ = "STUB: not implemented"; return nil }
