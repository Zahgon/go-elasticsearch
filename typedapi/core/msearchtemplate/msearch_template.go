package msearchtemplate

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MsearchTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewMsearchTemplate func() *MsearchTemplate

func NewMsearchTemplateFunc(tp elastictransport.Interface) NewMsearchTemplate {
	_ = "STUB: not implemented"
	return *new(NewMsearchTemplate)
}

func New(tp elastictransport.Interface) *MsearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *MsearchTemplate) Raw(raw io.Reader) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) Request(req *Request) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MsearchTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MsearchTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *MsearchTemplate) Header(key, value string) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) Index(index string) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) MaxConcurrentSearches(maxconcurrentsearches string) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) ProjectRouting(projectrouting string) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) SearchType(searchtype searchtype.SearchType) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) RestTotalHitsAsInt(resttotalhitsasint bool) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) TypedKeys(typedkeys bool) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) ErrorTrace(errortrace bool) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) FilterPath(filterpaths ...string) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *MsearchTemplate) Human(human bool) *MsearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *MsearchTemplate) Pretty(pretty bool) *MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}
