package types

type RankFeatureFunctionSaturation struct {
	Pivot *float32 `json:"pivot,omitempty"`
}

func (s *RankFeatureFunctionSaturation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankFeatureFunctionSaturation() *RankFeatureFunctionSaturation {
	_ = "STUB: not implemented"
	return nil
}

type RankFeatureFunctionSaturationVariant interface {
	RankFeatureFunctionSaturationCaster() *RankFeatureFunctionSaturation
}

func (s *RankFeatureFunctionSaturation) RankFeatureFunctionSaturationCaster() *RankFeatureFunctionSaturation {
	_ = "STUB: not implemented"
	return nil
}
