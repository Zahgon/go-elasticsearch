package putjob

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutJob func(jobid string) *PutJob

func NewPutJobFunc(tp elastictransport.Interface) NewPutJob {
	_ = "STUB: not implemented"
	return *new(NewPutJob)
}

func New(tp elastictransport.Interface) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Raw(raw io.Reader) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Request(req *Request) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutJob) Header(key, value string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) _jobid(jobid string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) AllowNoIndices(allownoindices bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) IgnoreThrottled(ignorethrottled bool) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) IgnoreUnavailable(ignoreunavailable bool) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) ErrorTrace(errortrace bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) FilterPath(filterpaths ...string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Human(human bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Pretty(pretty bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) AllowLazyOpen(allowlazyopen bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) AnalysisLimits(analysislimits types.AnalysisLimitsVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) BackgroundPersistInterval(duration types.DurationVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) CustomSettings(customsettings json.RawMessage) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) DataDescription(datadescription types.DataDescriptionVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) DatafeedConfig(datafeedconfig types.DatafeedConfigVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) Description(description string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Groups(groups ...string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) ModelPlotConfig(modelplotconfig types.ModelPlotConfigVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) RenormalizationWindowDays(renormalizationwindowdays int64) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) ResultsIndexName(indexname string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) ResultsRetentionDays(resultsretentiondays int64) *PutJob {
	_ = "STUB: not implemented"
	return nil
}
