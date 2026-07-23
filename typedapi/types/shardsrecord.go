package types

type ShardsRecord struct {
	BulkAvgSizeInBytes *string `json:"bulk.avg_size_in_bytes,omitempty"`

	BulkAvgTime *string `json:"bulk.avg_time,omitempty"`

	BulkTotalOperations *string `json:"bulk.total_operations,omitempty"`

	BulkTotalSizeInBytes *string `json:"bulk.total_size_in_bytes,omitempty"`

	BulkTotalTime *string `json:"bulk.total_time,omitempty"`

	CompletionSize *string `json:"completion.size,omitempty"`

	Dataset *string `json:"dataset,omitempty"`

	Docs *string `json:"docs,omitempty"`

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

	Id *string `json:"id,omitempty"`

	Index *string `json:"index,omitempty"`

	IndexingDeleteCurrent *string `json:"indexing.delete_current,omitempty"`

	IndexingDeleteTime *string `json:"indexing.delete_time,omitempty"`

	IndexingDeleteTotal *string `json:"indexing.delete_total,omitempty"`

	IndexingIndexCurrent *string `json:"indexing.index_current,omitempty"`

	IndexingIndexFailed *string `json:"indexing.index_failed,omitempty"`

	IndexingIndexTime *string `json:"indexing.index_time,omitempty"`

	IndexingIndexTotal *string `json:"indexing.index_total,omitempty"`

	Ip *string `json:"ip,omitempty"`

	MergesCurrent *string `json:"merges.current,omitempty"`

	MergesCurrentDocs *string `json:"merges.current_docs,omitempty"`

	MergesCurrentSize *string `json:"merges.current_size,omitempty"`

	MergesTotal *string `json:"merges.total,omitempty"`

	MergesTotalDocs *string `json:"merges.total_docs,omitempty"`

	MergesTotalSize *string `json:"merges.total_size,omitempty"`

	MergesTotalTime *string `json:"merges.total_time,omitempty"`

	Node *string `json:"node,omitempty"`

	PathData *string `json:"path.data,omitempty"`

	PathState *string `json:"path.state,omitempty"`

	Prirep *string `json:"prirep,omitempty"`

	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`

	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`

	RecoverysourceType *string `json:"recoverysource.type,omitempty"`

	RefreshExternalTime *string `json:"refresh.external_time,omitempty"`

	RefreshExternalTotal *string `json:"refresh.external_total,omitempty"`

	RefreshListeners *string `json:"refresh.listeners,omitempty"`

	RefreshTime *string `json:"refresh.time,omitempty"`

	RefreshTotal *string `json:"refresh.total,omitempty"`

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

	SegmentsCount *string `json:"segments.count,omitempty"`

	SegmentsFixedBitsetMemory *string `json:"segments.fixed_bitset_memory,omitempty"`

	SegmentsIndexWriterMemory *string `json:"segments.index_writer_memory,omitempty"`

	SegmentsMemory *string `json:"segments.memory,omitempty"`

	SegmentsVersionMapMemory *string `json:"segments.version_map_memory,omitempty"`

	SeqNoGlobalCheckpoint *string `json:"seq_no.global_checkpoint,omitempty"`

	SeqNoLocalCheckpoint *string `json:"seq_no.local_checkpoint,omitempty"`

	SeqNoMax *string `json:"seq_no.max,omitempty"`

	Shard *string `json:"shard,omitempty"`

	State *string `json:"state,omitempty"`

	Store *string `json:"store,omitempty"`

	SyncId *string `json:"sync_id,omitempty"`

	UnassignedAt *string `json:"unassigned.at,omitempty"`

	UnassignedDetails *string `json:"unassigned.details,omitempty"`

	UnassignedFor *string `json:"unassigned.for,omitempty"`

	UnassignedReason *string `json:"unassigned.reason,omitempty"`

	WarmerCurrent *string `json:"warmer.current,omitempty"`

	WarmerTotal *string `json:"warmer.total,omitempty"`

	WarmerTotalTime *string `json:"warmer.total_time,omitempty"`
}

func (s *ShardsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardsRecord() *ShardsRecord { _ = "STUB: not implemented"; return nil }
