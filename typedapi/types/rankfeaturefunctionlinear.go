package types

type RankFeatureFunctionLinear struct {
}

func NewRankFeatureFunctionLinear() *RankFeatureFunctionLinear {
	_ = "STUB: not implemented"
	return nil
}

type RankFeatureFunctionLinearVariant interface {
	RankFeatureFunctionLinearCaster() *RankFeatureFunctionLinear
}

func (s *RankFeatureFunctionLinear) RankFeatureFunctionLinearCaster() *RankFeatureFunctionLinear {
	_ = "STUB: not implemented"
	return nil
}
