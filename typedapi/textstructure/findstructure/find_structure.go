package findstructure

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/findstructureformat"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FindStructure struct {
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

type NewFindStructure func() *FindStructure

func NewFindStructureFunc(tp elastictransport.Interface) NewFindStructure {
	_ = "STUB: not implemented"
	return *new(NewFindStructure)
}

func New(tp elastictransport.Interface) *FindStructure { _ = "STUB: not implemented"; return nil }

func (r *FindStructure) Raw(raw io.Reader) *FindStructure { _ = "STUB: not implemented"; return nil }

func (r *FindStructure) Request(req *Request) *FindStructure { _ = "STUB: not implemented"; return nil }

func (r *FindStructure) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindStructure) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FindStructure) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *FindStructure) Header(key, value string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) Charset(charset string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) ColumnNames(columnnames ...string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) Delimiter(delimiter string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) EcsCompatibility(ecscompatibility string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) Explain(explain bool) *FindStructure { _ = "STUB: not implemented"; return nil }

func (r *FindStructure) Format(format findstructureformat.FindStructureFormat) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) GrokPattern(grokpattern string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) HasHeaderRow(hasheaderrow bool) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) LineMergeSizeLimit(linemergesizelimit string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) LinesToSample(linestosample string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) Quote(quote string) *FindStructure { _ = "STUB: not implemented"; return nil }

func (r *FindStructure) ShouldTrimFields(shouldtrimfields bool) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) ShouldParseRecursively(shouldparserecursively bool) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) Timeout(duration string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) TimestampField(field string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (r *FindStructure) TimestampFormat(timestampformat string) *FindStructure {
	_ = "STUB: not implemented"
	return nil
}
