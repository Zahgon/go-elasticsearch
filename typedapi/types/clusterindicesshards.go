package types

type ClusterIndicesShards struct {
	Index *ClusterIndicesShardsIndex `json:"index,omitempty"`

	Primaries *Float64 `json:"primaries,omitempty"`

	Replication *Float64 `json:"replication,omitempty"`

	Total *Float64 `json:"total,omitempty"`
}

func (s *ClusterIndicesShards) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterIndicesShards() *ClusterIndicesShards { _ = "STUB: not implemented"; return nil }
