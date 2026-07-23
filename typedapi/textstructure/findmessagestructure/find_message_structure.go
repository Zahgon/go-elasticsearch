package findmessagestructure

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ecscompatibilitytype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/formattype"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FindMessageStructure struct {
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

type NewFindMessageStructure func() *FindMessageStructure

func NewFindMessageStructureFunc(tp elastictransport.Interface) NewFindMessageStructure {
	_ = "STUB: not implemented"
	return *new(NewFindMessageStructure)
}

func New(tp elastictransport.Interface) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Raw(raw io.Reader) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Request(req *Request) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindMessageStructure) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindMessageStructure) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *FindMessageStructure) Header(key, value string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) ColumnNames(columnnames ...string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Delimiter(delimiter string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) EcsCompatibility(ecscompatibility ecscompatibilitytype.EcsCompatibilityType) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Explain(explain bool) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Format(format formattype.FormatType) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) GrokPattern(grokpattern string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Quote(quote string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) ShouldTrimFields(shouldtrimfields bool) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) ShouldParseRecursively(shouldparserecursively bool) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Timeout(duration string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) TimestampField(field string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) TimestampFormat(timestampformat string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) ErrorTrace(errortrace bool) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) FilterPath(filterpaths ...string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Human(human bool) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Pretty(pretty bool) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindMessageStructure) Messages(messages ...string) *FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}
