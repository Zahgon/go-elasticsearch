package termsenum

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type TermsEnum struct {
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

type NewTermsEnum func(index string) *TermsEnum

func NewTermsEnumFunc(tp elastictransport.Interface) NewTermsEnum {
	_ = "STUB: not implemented"
	return *new(NewTermsEnum)
}

func New(tp elastictransport.Interface) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) Raw(raw io.Reader) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) Request(req *Request) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TermsEnum) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TermsEnum) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TermsEnum) Header(key, value string) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) _index(index string) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) ErrorTrace(errortrace bool) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) FilterPath(filterpaths ...string) *TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (r *TermsEnum) Human(human bool) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) Pretty(pretty bool) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) CaseInsensitive(caseinsensitive bool) *TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (r *TermsEnum) Field(field string) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) IndexFilter(indexfilter types.QueryVariant) *TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (r *TermsEnum) ProjectRouting(projectrouting string) *TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (r *TermsEnum) SearchAfter(searchafter string) *TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (r *TermsEnum) Size(size int) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) String(string string) *TermsEnum { _ = "STUB: not implemented"; return nil }

func (r *TermsEnum) Timeout(duration types.DurationVariant) *TermsEnum {
	_ = "STUB: not implemented"
	return nil
}
