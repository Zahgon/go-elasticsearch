package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexcheckonstartup"
)

type _indexSettings struct {
	v *types.IndexSettings
}

func NewIndexSettings() *_indexSettings { _ = "STUB: not implemented"; return nil }

func (s *_indexSettings) Analysis(analysis types.IndexSettingsAnalysisVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Analyze(analyze types.SettingsAnalyzeVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) AutoExpandReplicas(autoexpandreplicas any) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Blocks(blocks types.IndexSettingBlocksVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Codec(codec string) *_indexSettings { _ = "STUB: not implemented"; return nil }

func (s *_indexSettings) CreationDate(stringifiedepochtimeunitmillis types.StringifiedEpochTimeUnitMillisVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) CreationDateString(datetime types.DateTimeVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) DefaultPipeline(pipelinename string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) FinalPipeline(pipelinename string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Format(format string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) GcDeletes(duration types.DurationVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Hidden(hidden string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Highlight(highlight types.SettingsHighlightVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Index(index types.IndexSettingsVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) IndexSettings(indexsettings map[string]json.RawMessage) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) AddIndexSetting(key string, value json.RawMessage) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) IndexingPressure(indexingpressure types.IndicesIndexingPressureVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) IndexingSlowlog(indexingslowlog types.IndexingSlowlogSettingsVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Lifecycle(lifecycle types.IndexSettingsLifecycleVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Mapping(mapping types.MappingLimitSettingsVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxInnerResultWindow(maxinnerresultwindow int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxNgramDiff(maxngramdiff int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxRefreshListeners(maxrefreshlisteners int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxRegexLength(maxregexlength int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxRescoreWindow(maxrescorewindow int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxResultWindow(maxresultwindow int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxScriptFields(maxscriptfields int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxShingleDiff(maxshinglediff int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxSlicesPerScroll(maxslicesperscroll int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) MaxTermsCount(maxtermscount int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Merge(merge types.MergeVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Mode(mode string) *_indexSettings { _ = "STUB: not implemented"; return nil }

func (s *_indexSettings) NumberOfReplicas(numberofreplicas string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) NumberOfRoutingShards(numberofroutingshards int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) NumberOfShards(numberofshards string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Priority(priority string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) ProvidedName(name string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Queries(queries types.QueriesVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) QueryString(querystring types.SettingsQueryStringVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) RefreshInterval(duration types.DurationVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Routing(routing types.IndexRoutingVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) RoutingPartitionSize(stringifiedinteger types.StringifiedintegerVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) RoutingPath(routingpaths ...string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Search(search types.SettingsSearchVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Settings(settings types.IndexSettingsVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Similarity(similarity map[string]types.SettingsSimilarity) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) AddSimilarity(key string, value types.SettingsSimilarityVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) SoftDeletes(softdeletes types.SoftDeletesVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Sort(sort types.IndexSegmentSortVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Store(store types.StorageVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) TimeSeries(timeseries types.IndexSettingsTimeSeriesVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) TopMetricsMaxSize(topmetricsmaxsize int) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Translog(translog types.TranslogVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Unassigned(unassigned types.IndexSettingsUnassignedVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Uuid(uuid string) *_indexSettings { _ = "STUB: not implemented"; return nil }

func (s *_indexSettings) VerifiedBeforeClose(verifiedbeforeclose string) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) Version(version types.IndexVersioningVariant) *_indexSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettings) IndexSettingsCaster() *types.IndexSettings {
	_ = "STUB: not implemented"
	return nil
}
