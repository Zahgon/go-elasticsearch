package types

import (
	"encoding/json"
)

type IndicesShardStats struct {
	Bulk            *BulkStats                 `json:"bulk,omitempty"`
	Commit          *ShardCommit               `json:"commit,omitempty"`
	Completion      *CompletionStats           `json:"completion,omitempty"`
	Docs            *DocStats                  `json:"docs,omitempty"`
	Fielddata       *FielddataStats            `json:"fielddata,omitempty"`
	Flush           *FlushStats                `json:"flush,omitempty"`
	Get             *GetStats                  `json:"get,omitempty"`
	Indexing        *IndexingStats             `json:"indexing,omitempty"`
	Indices         *IndicesStats              `json:"indices,omitempty"`
	Mappings        *MappingStats              `json:"mappings,omitempty"`
	Merges          *MergesStats               `json:"merges,omitempty"`
	QueryCache      *ShardQueryCache           `json:"query_cache,omitempty"`
	Recovery        *RecoveryStats             `json:"recovery,omitempty"`
	Refresh         *RefreshStats              `json:"refresh,omitempty"`
	RequestCache    *RequestCacheStats         `json:"request_cache,omitempty"`
	RetentionLeases *ShardRetentionLeases      `json:"retention_leases,omitempty"`
	Routing         *ShardRouting              `json:"routing,omitempty"`
	Search          *SearchStats               `json:"search,omitempty"`
	Segments        *SegmentsStats             `json:"segments,omitempty"`
	SeqNo           *ShardSequenceNumber       `json:"seq_no,omitempty"`
	ShardPath       *ShardPath                 `json:"shard_path,omitempty"`
	ShardStats      *ShardsTotalStats          `json:"shard_stats,omitempty"`
	Shards          map[string]json.RawMessage `json:"shards,omitempty"`
	Store           *StoreStats                `json:"store,omitempty"`
	Translog        *TranslogStats             `json:"translog,omitempty"`
	Warmer          *WarmerStats               `json:"warmer,omitempty"`
}

func NewIndicesShardStats() *IndicesShardStats { _ = "STUB: not implemented"; return nil }
