package datastreamsstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	BackingIndices int `json:"backing_indices"`

	DataStreamCount int `json:"data_stream_count"`

	DataStreams []types.DataStreamsStatsItem `json:"data_streams"`

	Shards_ types.ShardStatistics `json:"_shards"`

	TotalStoreSizeBytes int64 `json:"total_store_size_bytes"`

	TotalStoreSizes types.ByteSize `json:"total_store_sizes,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
