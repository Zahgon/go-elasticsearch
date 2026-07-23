package validate

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

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Validate struct {
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

type NewValidate func() *Validate

func NewValidateFunc(tp elastictransport.Interface) NewValidate {
	_ = "STUB: not implemented"
	return *new(NewValidate)
}

func New(tp elastictransport.Interface) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) Raw(raw io.Reader) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) Request(req *Request) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Validate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Validate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Validate) Header(key, value string) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) ErrorTrace(errortrace bool) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) FilterPath(filterpaths ...string) *Validate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Validate) Human(human bool) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) Pretty(pretty bool) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *Validate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Validate) AnalysisLimits(analysislimits types.AnalysisLimitsVariant) *Validate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Validate) DataDescription(datadescription types.DataDescriptionVariant) *Validate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Validate) Description(description string) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) JobId(id string) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) ModelPlot(modelplot types.ModelPlotConfigVariant) *Validate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Validate) ModelSnapshotId(id string) *Validate { _ = "STUB: not implemented"; return nil }

func (r *Validate) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *Validate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Validate) ResultsIndexName(indexname string) *Validate {
	_ = "STUB: not implemented"
	return nil
}
