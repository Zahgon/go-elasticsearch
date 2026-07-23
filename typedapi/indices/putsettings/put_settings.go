package putsettings

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexcheckonstartup"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSettings struct {
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

type NewPutSettings func() *PutSettings

func NewPutSettingsFunc(tp elastictransport.Interface) NewPutSettings {
	_ = "STUB: not implemented"
	return *new(NewPutSettings)
}

func New(tp elastictransport.Interface) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Raw(raw io.Reader) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Request(req *Request) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutSettings) Header(key, value string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Indices(index string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) AllowNoIndices(allownoindices bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) FlatSettings(flatsettings bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) IgnoreUnavailable(ignoreunavailable bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MasterTimeout(duration string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) PreserveExisting(preserveexisting bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Reopen(reopen bool) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Timeout(duration string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) ErrorTrace(errortrace bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) FilterPath(filterpaths ...string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Human(human bool) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Pretty(pretty bool) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Analysis(analysis types.IndexSettingsAnalysisVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Analyze(analyze types.SettingsAnalyzeVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) AutoExpandReplicas(autoexpandreplicas any) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Blocks(blocks types.IndexSettingBlocksVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Codec(codec string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillisVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) CreationDateString(datetime types.DateTimeVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) DefaultPipeline(pipelinename string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) FinalPipeline(pipelinename string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Format(format string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) GcDeletes(duration types.DurationVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Hidden(hidden string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Highlight(highlight types.SettingsHighlightVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Index(index types.IndexSettingsVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) IndexSettings(indexsettings map[string]json.RawMessage) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) AddIndexSetting(key string, value json.RawMessage) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) IndexingPressure(indexingpressure types.IndicesIndexingPressureVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) IndexingSlowlog(indexingslowlog types.IndexingSlowlogSettingsVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Lifecycle(lifecycle types.IndexSettingsLifecycleVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Mapping(mapping types.MappingLimitSettingsVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxInnerResultWindow(maxinnerresultwindow int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxNgramDiff(maxngramdiff int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxRefreshListeners(maxrefreshlisteners int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxRegexLength(maxregexlength int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxRescoreWindow(maxrescorewindow int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxResultWindow(maxresultwindow int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxScriptFields(maxscriptfields int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxShingleDiff(maxshinglediff int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxSlicesPerScroll(maxslicesperscroll int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) MaxTermsCount(maxtermscount int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Merge(merge types.MergeVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Mode(mode string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) NumberOfReplicas(numberofreplicas string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) NumberOfRoutingShards(numberofroutingshards int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) NumberOfShards(numberofshards string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Priority(priority string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) ProvidedName(name string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) Queries(queries types.QueriesVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) QueryString(querystring types.SettingsQueryStringVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) RefreshInterval(duration types.DurationVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Routing(routing types.IndexRoutingVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) RoutingPartitionSize(stringifiedinteger types.StringifiedintegerVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) RoutingPath(routingpaths ...string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Search(search types.SettingsSearchVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Settings(settings types.IndexSettingsVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Similarity(similarity map[string]types.SettingsSimilarity) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) AddSimilarity(key string, value types.SettingsSimilarityVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) SoftDeletes(softdeletes types.SoftDeletesVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Sort(sort types.IndexSegmentSortVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Store(store types.StorageVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) TimeSeries(timeseries types.IndexSettingsTimeSeriesVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) TopMetricsMaxSize(topmetricsmaxsize int) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Translog(translog types.TranslogVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Unassigned(unassigned types.IndexSettingsUnassignedVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Uuid(uuid string) *PutSettings { _ = "STUB: not implemented"; return nil }

func (r *PutSettings) VerifiedBeforeClose(verifiedbeforeclose string) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSettings) Version(version types.IndexVersioningVariant) *PutSettings {
	_ = "STUB: not implemented"
	return nil
}
