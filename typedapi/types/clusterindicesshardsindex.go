package types

type ClusterIndicesShardsIndex struct {
	Primaries ClusterShardMetrics `json:"primaries"`

	Replication ClusterShardMetrics `json:"replication"`

	Shards ClusterShardMetrics `json:"shards"`
}

func NewClusterIndicesShardsIndex() *ClusterIndicesShardsIndex {
	_ = "STUB: not implemented"
	return nil
}
