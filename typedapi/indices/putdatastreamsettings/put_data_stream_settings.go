package putdatastreamsettings

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexcheckonstartup"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataStreamSettings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutDataStreamSettings func(name string) *PutDataStreamSettings

func NewPutDataStreamSettingsFunc(tp elastictransport.Interface) NewPutDataStreamSettings {
	_ = "STUB: not implemented"
	return *new(NewPutDataStreamSettings)
}

func New(tp elastictransport.Interface) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Raw(raw io.Reader) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Request(req *Request) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataStreamSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataStreamSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDataStreamSettings) Header(key, value string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) _name(name string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) DryRun(dryrun bool) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MasterTimeout(duration string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Timeout(duration string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) ErrorTrace(errortrace bool) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) FilterPath(filterpaths ...string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Human(human bool) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Pretty(pretty bool) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Analysis(analysis types.IndexSettingsAnalysisVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Analyze(analyze types.SettingsAnalyzeVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) AutoExpandReplicas(autoexpandreplicas any) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Blocks(blocks types.IndexSettingBlocksVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Codec(codec string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillisVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) CreationDateString(datetime types.DateTimeVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) DefaultPipeline(pipelinename string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) FinalPipeline(pipelinename string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Format(format string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) GcDeletes(duration types.DurationVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Hidden(hidden string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Highlight(highlight types.SettingsHighlightVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Index(index types.IndexSettingsVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) IndexSettings(indexsettings map[string]json.RawMessage) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) AddIndexSetting(key string, value json.RawMessage) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) IndexingPressure(indexingpressure types.IndicesIndexingPressureVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) IndexingSlowlog(indexingslowlog types.IndexingSlowlogSettingsVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Lifecycle(lifecycle types.IndexSettingsLifecycleVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Mapping(mapping types.MappingLimitSettingsVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxInnerResultWindow(maxinnerresultwindow int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxNgramDiff(maxngramdiff int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxRefreshListeners(maxrefreshlisteners int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxRegexLength(maxregexlength int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxRescoreWindow(maxrescorewindow int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxResultWindow(maxresultwindow int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxScriptFields(maxscriptfields int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxShingleDiff(maxshinglediff int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxSlicesPerScroll(maxslicesperscroll int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) MaxTermsCount(maxtermscount int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Merge(merge types.MergeVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Mode(mode string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) NumberOfReplicas(numberofreplicas string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) NumberOfRoutingShards(numberofroutingshards int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) NumberOfShards(numberofshards string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Priority(priority string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) ProvidedName(name string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Queries(queries types.QueriesVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) QueryString(querystring types.SettingsQueryStringVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) RefreshInterval(duration types.DurationVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Routing(routing types.IndexRoutingVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) RoutingPartitionSize(stringifiedinteger types.StringifiedintegerVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) RoutingPath(routingpaths ...string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Search(search types.SettingsSearchVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Settings(settings types.IndexSettingsVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Similarity(similarity map[string]types.SettingsSimilarity) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) AddSimilarity(key string, value types.SettingsSimilarityVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) SoftDeletes(softdeletes types.SoftDeletesVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Sort(sort types.IndexSegmentSortVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Store(store types.StorageVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) TimeSeries(timeseries types.IndexSettingsTimeSeriesVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) TopMetricsMaxSize(topmetricsmaxsize int) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Translog(translog types.TranslogVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Unassigned(unassigned types.IndexSettingsUnassignedVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Uuid(uuid string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) VerifiedBeforeClose(verifiedbeforeclose string) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamSettings) Version(version types.IndexVersioningVariant) *PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}
