package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/healthstatus"
)

type RemoteClusterInfo struct {
	ClusterUuid string `json:"cluster_uuid"`

	IndicesCount int `json:"indices_count"`

	IndicesTotalSize *string `json:"indices_total_size,omitempty"`

	IndicesTotalSizeInBytes int64 `json:"indices_total_size_in_bytes"`

	MaxHeap *string `json:"max_heap,omitempty"`

	MaxHeapInBytes int64 `json:"max_heap_in_bytes"`

	MemTotal *string `json:"mem_total,omitempty"`

	MemTotalInBytes int64 `json:"mem_total_in_bytes"`

	Mode string `json:"mode"`

	NodesCount int `json:"nodes_count"`

	ShardsCount int `json:"shards_count"`

	SkipUnavailable bool `json:"skip_unavailable"`

	Status healthstatus.HealthStatus `json:"status"`

	TransportCompress string `json:"transport.compress"`

	Version []string `json:"version"`
}

func (s *RemoteClusterInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRemoteClusterInfo() *RemoteClusterInfo { _ = "STUB: not implemented"; return nil }
