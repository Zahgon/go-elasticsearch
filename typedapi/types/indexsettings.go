package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexcheckonstartup"
)

type IndexSettings struct {
	Analysis *IndexSettingsAnalysis `json:"analysis,omitempty"`

	Analyze            *SettingsAnalyze                         `json:"analyze,omitempty"`
	AutoExpandReplicas any                                      `json:"auto_expand_replicas,omitempty"`
	Blocks             *IndexSettingBlocks                      `json:"blocks,omitempty"`
	CheckOnStartup     *indexcheckonstartup.IndexCheckOnStartup `json:"check_on_startup,omitempty"`
	Codec              *string                                  `json:"codec,omitempty"`
	CreationDate       StringifiedEpochTimeUnitMillis           `json:"creation_date,omitempty"`
	CreationDateString DateTime                                 `json:"creation_date_string,omitempty"`
	DefaultPipeline    *string                                  `json:"default_pipeline,omitempty"`
	FinalPipeline      *string                                  `json:"final_pipeline,omitempty"`
	Format             *string                                  `json:"format,omitempty"`
	GcDeletes          Duration                                 `json:"gc_deletes,omitempty"`
	Hidden             *string                                  `json:"hidden,omitempty"`
	Highlight          *SettingsHighlight                       `json:"highlight,omitempty"`
	Index              *IndexSettings                           `json:"index,omitempty"`
	IndexSettings      map[string]json.RawMessage               `json:"-"`

	IndexingPressure              *IndicesIndexingPressure `json:"indexing_pressure,omitempty"`
	IndexingSlowlog               *IndexingSlowlogSettings `json:"indexing.slowlog,omitempty"`
	Lifecycle                     *IndexSettingsLifecycle  `json:"lifecycle,omitempty"`
	LoadFixedBitsetFiltersEagerly *bool                    `json:"load_fixed_bitset_filters_eagerly,omitempty"`

	Mapping                 *MappingLimitSettings `json:"mapping,omitempty"`
	MaxDocvalueFieldsSearch *int                  `json:"max_docvalue_fields_search,omitempty"`
	MaxInnerResultWindow    *int                  `json:"max_inner_result_window,omitempty"`
	MaxNgramDiff            *int                  `json:"max_ngram_diff,omitempty"`
	MaxRefreshListeners     *int                  `json:"max_refresh_listeners,omitempty"`
	MaxRegexLength          *int                  `json:"max_regex_length,omitempty"`
	MaxRescoreWindow        *int                  `json:"max_rescore_window,omitempty"`
	MaxResultWindow         *int                  `json:"max_result_window,omitempty"`
	MaxScriptFields         *int                  `json:"max_script_fields,omitempty"`
	MaxShingleDiff          *int                  `json:"max_shingle_diff,omitempty"`
	MaxSlicesPerScroll      *int                  `json:"max_slices_per_scroll,omitempty"`
	MaxTermsCount           *int                  `json:"max_terms_count,omitempty"`
	Merge                   *Merge                `json:"merge,omitempty"`
	Mode                    *string               `json:"mode,omitempty"`
	NumberOfReplicas        *string               `json:"number_of_replicas,omitempty"`
	NumberOfRoutingShards   *int                  `json:"number_of_routing_shards,omitempty"`
	NumberOfShards          *string               `json:"number_of_shards,omitempty"`
	Priority                *string               `json:"priority,omitempty"`
	ProvidedName            *string               `json:"provided_name,omitempty"`
	Queries                 *Queries              `json:"queries,omitempty"`
	QueryString             *SettingsQueryString  `json:"query_string,omitempty"`
	RefreshInterval         Duration              `json:"refresh_interval,omitempty"`
	Routing                 *IndexRouting         `json:"routing,omitempty"`
	RoutingPartitionSize    Stringifiedinteger    `json:"routing_partition_size,omitempty"`
	RoutingPath             []string              `json:"routing_path,omitempty"`
	Search                  *SettingsSearch       `json:"search,omitempty"`
	Settings                *IndexSettings        `json:"settings,omitempty"`

	Similarity  map[string]SettingsSimilarity `json:"similarity,omitempty"`
	SoftDeletes *SoftDeletes                  `json:"soft_deletes,omitempty"`
	Sort        *IndexSegmentSort             `json:"sort,omitempty"`

	Store               *Storage                 `json:"store,omitempty"`
	TimeSeries          *IndexSettingsTimeSeries `json:"time_series,omitempty"`
	TopMetricsMaxSize   *int                     `json:"top_metrics_max_size,omitempty"`
	Translog            *Translog                `json:"translog,omitempty"`
	Unassigned          *IndexSettingsUnassigned `json:"unassigned,omitempty"`
	Uuid                *string                  `json:"uuid,omitempty"`
	VerifiedBeforeClose *string                  `json:"verified_before_close,omitempty"`
	Version             *IndexVersioning         `json:"version,omitempty"`
}

func (s *IndexSettings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s IndexSettings) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIndexSettings() *IndexSettings { _ = "STUB: not implemented"; return nil }

type IndexSettingsVariant interface {
	IndexSettingsCaster() *IndexSettings
}

func (s *IndexSettings) IndexSettingsCaster() *IndexSettings { _ = "STUB: not implemented"; return nil }
