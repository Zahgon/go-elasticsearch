package findfieldstructure

import (
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

type FindFieldStructure struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewFindFieldStructure func() *FindFieldStructure

func NewFindFieldStructureFunc(tp elastictransport.Interface) NewFindFieldStructure {
	_ = "STUB: not implemented"
	return *new(NewFindFieldStructure)
}

func New(tp elastictransport.Interface) *FindFieldStructure { _ = "STUB: not implemented"; return nil }

func (r *FindFieldStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindFieldStructure) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindFieldStructure) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindFieldStructure) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *FindFieldStructure) Header(key, value string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) ColumnNames(columnnames ...string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Delimiter(delimiter string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) DocumentsToSample(documentstosample string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) EcsCompatibility(ecscompatibility ecscompatibilitytype.EcsCompatibilityType) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Explain(explain bool) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Field(field string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Format(format formattype.FormatType) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) GrokPattern(grokpattern string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Index(indexname string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Quote(quote string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) ShouldTrimFields(shouldtrimfields bool) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) ShouldParseRecursively(shouldparserecursively bool) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Timeout(duration string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) TimestampField(field string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) TimestampFormat(timestampformat string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) ErrorTrace(errortrace bool) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) FilterPath(filterpaths ...string) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Human(human bool) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindFieldStructure) Pretty(pretty bool) *FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}
