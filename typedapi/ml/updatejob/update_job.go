package updatejob

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
)

const (
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateJob struct {
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

type NewUpdateJob func(jobid string) *UpdateJob

func NewUpdateJobFunc(tp elastictransport.Interface) NewUpdateJob {
	_ = "STUB: not implemented"
	return *new(NewUpdateJob)
}

func New(tp elastictransport.Interface) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) Raw(raw io.Reader) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) Request(req *Request) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateJob) Header(key, value string) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) _jobid(jobid string) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) ErrorTrace(errortrace bool) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) FilterPath(filterpaths ...string) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) Human(human bool) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) Pretty(pretty bool) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) AllowLazyOpen(allowlazyopen bool) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) AnalysisLimits(analysislimits types.AnalysisMemoryLimitVariant) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) BackgroundPersistInterval(duration types.DurationVariant) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) CategorizationFilters(categorizationfilters ...string) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) CustomSettings(customsettings map[string]json.RawMessage) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) AddCustomSetting(key string, value json.RawMessage) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) Description(description string) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) Detectors(detectors ...types.DetectorUpdateVariant) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) DetectorsValues(detectorsvalues []types.DetectorUpdate) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) Groups(groups ...string) *UpdateJob { _ = "STUB: not implemented"; return nil }

func (r *UpdateJob) ModelPlotConfig(modelplotconfig types.ModelPlotConfigVariant) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) ModelPruneWindow(duration types.DurationVariant) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) PerPartitionCategorization(perpartitioncategorization types.PerPartitionCategorizationVariant) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) RenormalizationWindowDays(renormalizationwindowdays int64) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateJob) ResultsRetentionDays(resultsretentiondays int64) *UpdateJob {
	_ = "STUB: not implemented"
	return nil
}
