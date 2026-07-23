package types

type RankFeatureFunctionLogarithm struct {
	ScalingFactor float32 `json:"scaling_factor"`
}

func (s *RankFeatureFunctionLogarithm) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankFeatureFunctionLogarithm() *RankFeatureFunctionLogarithm {
	_ = "STUB: not implemented"
	return nil
}

type RankFeatureFunctionLogarithmVariant interface {
	RankFeatureFunctionLogarithmCaster() *RankFeatureFunctionLogarithm
}

func (s *RankFeatureFunctionLogarithm) RankFeatureFunctionLogarithmCaster() *RankFeatureFunctionLogarithm {
	_ = "STUB: not implemented"
	return nil
}
