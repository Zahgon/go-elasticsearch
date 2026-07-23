package types

type NodesRecord struct {
	AvailableProcessors *string `json:"available_processors,omitempty"`

	Build *string `json:"build,omitempty"`

	BulkAvgSizeInBytes *string `json:"bulk.avg_size_in_bytes,omitempty"`

	BulkAvgTime *string `json:"bulk.avg_time,omitempty"`

	BulkTotalOperations *string `json:"bulk.total_operations,omitempty"`

	BulkTotalSizeInBytes *string `json:"bulk.total_size_in_bytes,omitempty"`

	BulkTotalTime *string `json:"bulk.total_time,omitempty"`

	CompletionSize *string `json:"completion.size,omitempty"`

	Cpu *string `json:"cpu,omitempty"`

	DiskAvail ByteSize `json:"disk.avail,omitempty"`

	DiskTotal ByteSize `json:"disk.total,omitempty"`

	DiskUsed ByteSize `json:"disk.used,omitempty"`

	DiskUsedPercent Percentage `json:"disk.used_percent,omitempty"`

	FielddataEvictions *string `json:"fielddata.evictions,omitempty"`

	FielddataMemorySize *string `json:"fielddata.memory_size,omitempty"`

	FileDescCurrent *string `json:"file_desc.current,omitempty"`

	FileDescMax *string `json:"file_desc.max,omitempty"`

	FileDescPercent Percentage `json:"file_desc.percent,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	FlushTotal *string `json:"flush.total,omitempty"`

	FlushTotalTime *string `json:"flush.total_time,omitempty"`

	GetCurrent *string `json:"get.current,omitempty"`

	GetExistsTime *string `json:"get.exists_time,omitempty"`

	GetExistsTotal *string `json:"get.exists_total,omitempty"`

	GetMissingTime *string `json:"get.missing_time,omitempty"`

	GetMissingTotal *string `json:"get.missing_total,omitempty"`

	GetTime *string `json:"get.time,omitempty"`

	GetTotal *string `json:"get.total,omitempty"`

	HeapCurrent *string `json:"heap.current,omitempty"`

	HeapMax *string `json:"heap.max,omitempty"`

	HeapPercent Percentage `json:"heap.percent,omitempty"`

	HttpAddress *string `json:"http_address,omitempty"`

	Id *string `json:"id,omitempty"`

	IndexingDeleteCurrent *string `json:"indexing.delete_current,omitempty"`

	IndexingDeleteTime *string `json:"indexing.delete_time,omitempty"`

	IndexingDeleteTotal *string `json:"indexing.delete_total,omitempty"`

	IndexingIndexCurrent *string `json:"indexing.index_current,omitempty"`

	IndexingIndexFailed *string `json:"indexing.index_failed,omitempty"`

	IndexingIndexTime *string `json:"indexing.index_time,omitempty"`

	IndexingIndexTotal *string `json:"indexing.index_total,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Jdk *string `json:"jdk,omitempty"`

	Load15M *string `json:"load_15m,omitempty"`

	Load1M *string `json:"load_1m,omitempty"`

	Load5M *string `json:"load_5m,omitempty"`

	Master *string `json:"master,omitempty"`

	MergesCurrent *string `json:"merges.current,omitempty"`

	MergesCurrentDocs *string `json:"merges.current_docs,omitempty"`

	MergesCurrentSize *string `json:"merges.current_size,omitempty"`

	MergesTotal *string `json:"merges.total,omitempty"`

	MergesTotalDocs *string `json:"merges.total_docs,omitempty"`

	MergesTotalSize *string `json:"merges.total_size,omitempty"`

	MergesTotalTime *string `json:"merges.total_time,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeRole *string `json:"node.role,omitempty"`

	Pid *string `json:"pid,omitempty"`

	Port *string `json:"port,omitempty"`

	QueryCacheEvictions *string `json:"query_cache.evictions,omitempty"`

	QueryCacheHitCount *string `json:"query_cache.hit_count,omitempty"`

	QueryCacheMemorySize *string `json:"query_cache.memory_size,omitempty"`

	QueryCacheMissCount *string `json:"query_cache.miss_count,omitempty"`

	RamCurrent *string `json:"ram.current,omitempty"`

	RamMax *string `json:"ram.max,omitempty"`

	RamPercent Percentage `json:"ram.percent,omitempty"`

	RefreshExternalTime *string `json:"refresh.external_time,omitempty"`

	RefreshExternalTotal *string `json:"refresh.external_total,omitempty"`

	RefreshListeners *string `json:"refresh.listeners,omitempty"`

	RefreshTime *string `json:"refresh.time,omitempty"`

	RefreshTotal *string `json:"refresh.total,omitempty"`

	RequestCacheEvictions *string `json:"request_cache.evictions,omitempty"`

	RequestCacheHitCount *string `json:"request_cache.hit_count,omitempty"`

	RequestCacheMemorySize *string `json:"request_cache.memory_size,omitempty"`

	RequestCacheMissCount *string `json:"request_cache.miss_count,omitempty"`

	ScriptCacheEvictions *string `json:"script.cache_evictions,omitempty"`

	ScriptCompilationLimitTriggered *string `json:"script.compilation_limit_triggered,omitempty"`

	ScriptCompilations *string `json:"script.compilations,omitempty"`

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

	SuggestCurrent *string `json:"suggest.current,omitempty"`

	SuggestTime *string `json:"suggest.time,omitempty"`

	SuggestTotal *string `json:"suggest.total,omitempty"`

	Type *string `json:"type,omitempty"`

	Uptime *string `json:"uptime,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *NodesRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodesRecord() *NodesRecord { _ = "STUB: not implemented"; return nil }
