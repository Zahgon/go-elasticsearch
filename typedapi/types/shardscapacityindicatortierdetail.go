package types

type ShardsCapacityIndicatorTierDetail struct {
	CurrentUsedShards  *int `json:"current_used_shards,omitempty"`
	MaxShardsInCluster int  `json:"max_shards_in_cluster"`
}

func (s *ShardsCapacityIndicatorTierDetail) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardsCapacityIndicatorTierDetail() *ShardsCapacityIndicatorTierDetail {
	_ = "STUB: not implemented"
	return nil
}
