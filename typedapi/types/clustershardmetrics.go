package types

type ClusterShardMetrics struct {
	Avg Float64 `json:"avg"`

	Max Float64 `json:"max"`

	Min Float64 `json:"min"`
}

func (s *ClusterShardMetrics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClusterShardMetrics() *ClusterShardMetrics { _ = "STUB: not implemented"; return nil }
