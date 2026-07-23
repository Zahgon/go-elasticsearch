package types

type ClusterInfo struct {
	Nodes             map[string]NodeDiskUsage `json:"nodes"`
	ReservedSizes     []ReservedSize           `json:"reserved_sizes"`
	ShardDataSetSizes map[string]string        `json:"shard_data_set_sizes,omitempty"`
	ShardPaths        map[string]string        `json:"shard_paths"`
	ShardSizes        map[string]int64         `json:"shard_sizes"`
}

func NewClusterInfo() *ClusterInfo { _ = "STUB: not implemented"; return nil }
