package types

type RankFeatureFunctionSigmoid struct {
	Exponent float32 `json:"exponent"`

	Pivot float32 `json:"pivot"`
}

func (s *RankFeatureFunctionSigmoid) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankFeatureFunctionSigmoid() *RankFeatureFunctionSigmoid {
	_ = "STUB: not implemented"
	return nil
}

type RankFeatureFunctionSigmoidVariant interface {
	RankFeatureFunctionSigmoidCaster() *RankFeatureFunctionSigmoid
}

func (s *RankFeatureFunctionSigmoid) RankFeatureFunctionSigmoidCaster() *RankFeatureFunctionSigmoid {
	_ = "STUB: not implemented"
	return nil
}
