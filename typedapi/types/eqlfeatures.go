package types

type EqlFeatures struct {
	Event     uint                 `json:"event"`
	Join      uint                 `json:"join"`
	Joins     EqlFeaturesJoin      `json:"joins"`
	Keys      EqlFeaturesKeys      `json:"keys"`
	Pipes     EqlFeaturesPipes     `json:"pipes"`
	Sequence  uint                 `json:"sequence"`
	Sequences EqlFeaturesSequences `json:"sequences"`
}

func NewEqlFeatures() *EqlFeatures { _ = "STUB: not implemented"; return nil }
