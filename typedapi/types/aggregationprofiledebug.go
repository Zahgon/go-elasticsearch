package types

type AggregationProfileDebug struct {
	BruteForceUsed                    *int                                    `json:"brute_force_used,omitempty"`
	BuiltBuckets                      *int                                    `json:"built_buckets,omitempty"`
	CharsFetched                      *int                                    `json:"chars_fetched,omitempty"`
	CollectAnalyzedCount              *int                                    `json:"collect_analyzed_count,omitempty"`
	CollectAnalyzedNs                 *int                                    `json:"collect_analyzed_ns,omitempty"`
	CollectionStrategy                *string                                 `json:"collection_strategy,omitempty"`
	DeferredAggregators               []string                                `json:"deferred_aggregators,omitempty"`
	Delegate                          *string                                 `json:"delegate,omitempty"`
	DelegateDebug                     *AggregationProfileDebug                `json:"delegate_debug,omitempty"`
	DynamicPruningAttempted           *int                                    `json:"dynamic_pruning_attempted,omitempty"`
	DynamicPruningUsed                *int                                    `json:"dynamic_pruning_used,omitempty"`
	EmptyCollectorsUsed               *int                                    `json:"empty_collectors_used,omitempty"`
	ExtractCount                      *int                                    `json:"extract_count,omitempty"`
	ExtractNs                         *int                                    `json:"extract_ns,omitempty"`
	Filters                           []AggregationProfileDelegateDebugFilter `json:"filters,omitempty"`
	HasFilter                         *bool                                   `json:"has_filter,omitempty"`
	MapReducer                        *string                                 `json:"map_reducer,omitempty"`
	NumericCollectorsUsed             *int                                    `json:"numeric_collectors_used,omitempty"`
	OrdinalsCollectorsOverheadTooHigh *int                                    `json:"ordinals_collectors_overhead_too_high,omitempty"`
	OrdinalsCollectorsUsed            *int                                    `json:"ordinals_collectors_used,omitempty"`
	ResultStrategy                    *string                                 `json:"result_strategy,omitempty"`
	SegmentsCollected                 *int                                    `json:"segments_collected,omitempty"`
	SegmentsCounted                   *int                                    `json:"segments_counted,omitempty"`
	SegmentsWithDeletedDocs           *int                                    `json:"segments_with_deleted_docs,omitempty"`
	SegmentsWithDocCountField         *int                                    `json:"segments_with_doc_count_field,omitempty"`
	SegmentsWithMultiValuedOrds       *int                                    `json:"segments_with_multi_valued_ords,omitempty"`
	SegmentsWithSingleValuedOrds      *int                                    `json:"segments_with_single_valued_ords,omitempty"`
	SkippedDueToNoData                *int                                    `json:"skipped_due_to_no_data,omitempty"`
	StringHashingCollectorsUsed       *int                                    `json:"string_hashing_collectors_used,omitempty"`
	SurvivingBuckets                  *int                                    `json:"surviving_buckets,omitempty"`
	TotalBuckets                      *int                                    `json:"total_buckets,omitempty"`
	ValuesFetched                     *int                                    `json:"values_fetched,omitempty"`
}

func (s *AggregationProfileDebug) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAggregationProfileDebug() *AggregationProfileDebug { _ = "STUB: not implemented"; return nil }
