package types

type ShardsCapacityIndicatorDetails struct {
	Data   ShardsCapacityIndicatorTierDetail `json:"data"`
	Frozen ShardsCapacityIndicatorTierDetail `json:"frozen"`
}

func NewShardsCapacityIndicatorDetails() *ShardsCapacityIndicatorDetails {
	_ = "STUB: not implemented"
	return nil
}
