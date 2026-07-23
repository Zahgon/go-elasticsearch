package types

type IndexStats struct {
	Bulk *BulkStats `json:"bulk,omitempty"`

	Completion *CompletionStats `json:"completion,omitempty"`

	Docs *DocStats `json:"docs,omitempty"`

	Fielddata *FielddataStats `json:"fielddata,omitempty"`

	Flush *FlushStats `json:"flush,omitempty"`

	Get *GetStats `json:"get,omitempty"`

	Indexing *IndexingStats `json:"indexing,omitempty"`

	Indices *IndicesStats `json:"indices,omitempty"`

	Merges *MergesStats `json:"merges,omitempty"`

	QueryCache *QueryCacheStats `json:"query_cache,omitempty"`

	Recovery *RecoveryStats `json:"recovery,omitempty"`

	Refresh *RefreshStats `json:"refresh,omitempty"`

	RequestCache *RequestCacheStats `json:"request_cache,omitempty"`

	Search *SearchStats `json:"search,omitempty"`

	Segments   *SegmentsStats    `json:"segments,omitempty"`
	ShardStats *ShardsTotalStats `json:"shard_stats,omitempty"`

	Store *StoreStats `json:"store,omitempty"`

	Translog *TranslogStats `json:"translog,omitempty"`

	Warmer *WarmerStats `json:"warmer,omitempty"`
}

func NewIndexStats() *IndexStats { _ = "STUB: not implemented"; return nil }
