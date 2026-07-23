package types

type IndicesRecord struct {
	BulkAvgSizeInBytes *string `json:"bulk.avg_size_in_bytes,omitempty"`

	BulkAvgTime *string `json:"bulk.avg_time,omitempty"`

	BulkTotalOperations *string `json:"bulk.total_operations,omitempty"`

	BulkTotalSizeInBytes *string `json:"bulk.total_size_in_bytes,omitempty"`

	BulkTotalTime *string `json:"bulk.total_time,omitempty"`

	CompletionSize *string `json:"completion.size,omitempty"`

	CreationDate *string `json:"creation.date,omitempty"`

	CreationDateString *string `json:"creation.date.string,omitempty"`

	DatasetSize *string `json:"dataset.size,omitempty"`

	DocsCount *string `json:"docs.count,omitempty"`

	DocsDeleted *string `json:"docs.deleted,omitempty"`

	FielddataEvictions *string `json:"fielddata.evictions,omitempty"`

	FielddataMemorySize *string `json:"fielddata.memory_size,omitempty"`

	FlushTotal *string `json:"flush.total,omitempty"`

	FlushTotalTime *string `json:"flush.total_time,omitempty"`

	GetCurrent *string `json:"get.current,omitempty"`

	GetExistsTime *string `json:"get.exists_time,omitempty"`

	GetExistsTotal *string `json:"get.exists_total,omitempty"`

	GetMissingTime *string `json:"get.missing_time,omitempty"`

	GetMissingTotal *string `json:"get.missing_total,omitempty"`

	GetTime *string `json:"get.time,omitempty"`

	GetTotal *string `json:"get.total,omitempty"`

	Health *string `json:"health,omitempty"`

	Index *string `json:"index,omitempty"`

	IndexingDeleteCurrent *string `json:"indexing.delete_current,omitempty"`

	IndexingDeleteTime *string `json:"indexing.delete_time,omitempty"`

	IndexingDeleteTotal *string `json:"indexing.delete_total,omitempty"`

	IndexingIndexCurrent *string `json:"indexing.index_current,omitempty"`

	IndexingIndexFailed *string `json:"indexing.index_failed,omitempty"`

	IndexingIndexTime *string `json:"indexing.index_time,omitempty"`

	IndexingIndexTotal *string `json:"indexing.index_total,omitempty"`

	MemoryTotal *string `json:"memory.total,omitempty"`

	MergesCurrent *string `json:"merges.current,omitempty"`

	MergesCurrentDocs *string `json:"merges.current_docs,omitempty"`

	MergesCurrentSize *string `json:"merges.current_size,omitempty"`

	MergesTotal *string `json:"merges.total,omitempty"`

	MergesTotalDocs *string `json:"merges.total_docs,omitempty"`

	MergesTotalSize *string `json:"merges.total_size,omitempty"`

	MergesTotalTime *string `json:"merges.total_time,omitempty"`

	Pri *string `json:"pri,omitempty"`

	PriBulkAvgSizeInBytes *string `json:"pri.bulk.avg_size_in_bytes,omitempty"`

	PriBulkAvgTime *string `json:"pri.bulk.avg_time,omitempty"`

	PriBulkTotalOperations *string `json:"pri.bulk.total_operations,omitempty"`

	PriBulkTotalSizeInBytes *string `json:"pri.bulk.total_size_in_bytes,omitempty"`

	PriBulkTotalTime *string `json:"pri.bulk.total_time,omitempty"`

	PriCompletionSize *string `json:"pri.completion.size,omitempty"`

	PriFielddataEvictions *string `json:"pri.fielddata.evictions,omitempty"`

	PriFielddataMemorySize *string `json:"pri.fielddata.memory_size,omitempty"`

	PriFlushTotal *string `json:"pri.flush.total,omitempty"`

	PriFlushTotalTime *string `json:"pri.flush.total_time,omitempty"`

	PriGetCurrent *string `json:"pri.get.current,omitempty"`

	PriGetExistsTime *string `json:"pri.get.exists_time,omitempty"`

	PriGetExistsTotal *string `json:"pri.get.exists_total,omitempty"`

	PriGetMissingTime *string `json:"pri.get.missing_time,omitempty"`

	PriGetMissingTotal *string `json:"pri.get.missing_total,omitempty"`

	PriGetTime *string `json:"pri.get.time,omitempty"`

	PriGetTotal *string `json:"pri.get.total,omitempty"`

	PriIndexingDeleteCurrent *string `json:"pri.indexing.delete_current,omitempty"`

	PriIndexingDeleteTime *string `json:"pri.indexing.delete_time,omitempty"`

	PriIndexingDeleteTotal *string `json:"pri.indexing.delete_total,omitempty"`

	PriIndexingIndexCurrent *string `json:"pri.indexing.index_current,omitempty"`

	PriIndexingIndexFailed *string `json:"pri.indexing.index_failed,omitempty"`

	PriIndexingIndexTime *string `json:"pri.indexing.index_time,omitempty"`

	PriIndexingIndexTotal *string `json:"pri.indexing.index_total,omitempty"`

	PriMemoryTotal *string `json:"pri.memory.total,omitempty"`

	PriMergesCurrent *string `json:"pri.merges.current,omitempty"`

	PriMergesCurrentDocs *string `json:"pri.merges.current_docs,omitempty"`

	PriMergesCurrentSize *string `json:"pri.merges.current_size,omitempty"`

	PriMergesTotal *string `json:"pri.merges.total,omitempty"`

	PriMergesTotalDocs *string `json:"pri.merges.total_docs,omitempty"`

	PriMergesTotalSize *string `json:"pri.merges.total_size,omitempty"`

	PriMergesTotalTime *string `json:"pri.merges.total_time,omitempty"`

	PriQueryCacheEvictions *string `json:"pri.query_cache.evictions,omitempty"`

	PriQueryCacheMemorySize *string `json:"pri.query_cache.memory_size,omitempty"`

	PriRefreshExternalTime *string `json:"pri.refresh.external_time,omitempty"`

	PriRefreshExternalTotal *string `json:"pri.refresh.external_total,omitempty"`

	PriRefreshListeners *string `json:"pri.refresh.listeners,omitempty"`

	PriRefreshTime *string `json:"pri.refresh.time,omitempty"`

	PriRefreshTotal *string `json:"pri.refresh.total,omitempty"`

	PriRequestCacheEvictions *string `json:"pri.request_cache.evictions,omitempty"`

	PriRequestCacheHitCount *string `json:"pri.request_cache.hit_count,omitempty"`

	PriRequestCacheMemorySize *string `json:"pri.request_cache.memory_size,omitempty"`

	PriRequestCacheMissCount *string `json:"pri.request_cache.miss_count,omitempty"`

	PriSearchFetchCurrent *string `json:"pri.search.fetch_current,omitempty"`

	PriSearchFetchTime *string `json:"pri.search.fetch_time,omitempty"`

	PriSearchFetchTotal *string `json:"pri.search.fetch_total,omitempty"`

	PriSearchOpenContexts *string `json:"pri.search.open_contexts,omitempty"`

	PriSearchQueryCurrent *string `json:"pri.search.query_current,omitempty"`

	PriSearchQueryTime *string `json:"pri.search.query_time,omitempty"`

	PriSearchQueryTotal *string `json:"pri.search.query_total,omitempty"`

	PriSearchScrollCurrent *string `json:"pri.search.scroll_current,omitempty"`

	PriSearchScrollTime *string `json:"pri.search.scroll_time,omitempty"`

	PriSearchScrollTotal *string `json:"pri.search.scroll_total,omitempty"`

	PriSegmentsCount *string `json:"pri.segments.count,omitempty"`

	PriSegmentsFixedBitsetMemory *string `json:"pri.segments.fixed_bitset_memory,omitempty"`

	PriSegmentsIndexWriterMemory *string `json:"pri.segments.index_writer_memory,omitempty"`

	PriSegmentsMemory *string `json:"pri.segments.memory,omitempty"`

	PriSegmentsVersionMapMemory *string `json:"pri.segments.version_map_memory,omitempty"`

	PriStoreSize *string `json:"pri.store.size,omitempty"`

	PriSuggestCurrent *string `json:"pri.suggest.current,omitempty"`

	PriSuggestTime *string `json:"pri.suggest.time,omitempty"`

	PriSuggestTotal *string `json:"pri.suggest.total,omitempty"`

	PriWarmerCurrent *string `json:"pri.warmer.current,omitempty"`

	PriWarmerTotal *string `json:"pri.warmer.total,omitempty"`

	PriWarmerTotalTime *string `json:"pri.warmer.total_time,omitempty"`

	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`

	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`

	RefreshExternalTime *string `json:"refresh.external_time,omitempty"`

	RefreshExternalTotal *string `json:"refresh.external_total,omitempty"`

	RefreshListeners *string `json:"refresh.listeners,omitempty"`

	RefreshTime *string `json:"refresh.time,omitempty"`

	RefreshTotal *string `json:"refresh.total,omitempty"`

	Rep *string `json:"rep,omitempty"`

	RequestCacheEvictions *string `json:"request_cache.evictions,omitempty"`

	RequestCacheHitCount *string `json:"request_cache.hit_count,omitempty"`

	RequestCacheMemorySize *string `json:"request_cache.memory_size,omitempty"`

	RequestCacheMissCount *string `json:"request_cache.miss_count,omitempty"`

	SearchFetchCurrent *string `json:"search.fetch_current,omitempty"`

	SearchFetchTime *string `json:"search.fetch_time,omitempty"`

	SearchFetchTotal *string `json:"search.fetch_total,omitempty"`

	SearchOpenContexts *string `json:"search.open_contexts,omitempty"`

	SearchQueryCurrent *string `json:"search.query_current,omitempty"`

	SearchQueryTime *string `json:"search.query_time,omitempty"`

	SearchQueryTotal *string `json:"search.query_total,omitempty"`

	SearchScrollCurrent *string `json:"search.scroll_current,omitempty"`

	SearchScrollTime *string `json:"search.scroll_time,omitempty"`

	SearchScrollTotal *string `json:"search.scroll_total,omitempty"`

	SearchThrottled *string `json:"search.throttled,omitempty"`

	SegmentsCount *string `json:"segments.count,omitempty"`

	SegmentsFixedBitsetMemory *string `json:"segments.fixed_bitset_memory,omitempty"`

	SegmentsIndexWriterMemory *string `json:"segments.index_writer_memory,omitempty"`

	SegmentsMemory *string `json:"segments.memory,omitempty"`

	SegmentsVersionMapMemory *string `json:"segments.version_map_memory,omitempty"`

	Status *string `json:"status,omitempty"`

	StoreSize *string `json:"store.size,omitempty"`

	SuggestCurrent *string `json:"suggest.current,omitempty"`

	SuggestTime *string `json:"suggest.time,omitempty"`

	SuggestTotal *string `json:"suggest.total,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	WarmerCurrent *string `json:"warmer.current,omitempty"`

	WarmerTotal *string `json:"warmer.total,omitempty"`

	WarmerTotalTime *string `json:"warmer.total_time,omitempty"`
}

func (s *IndicesRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndicesRecord() *IndicesRecord { _ = "STUB: not implemented"; return nil }
