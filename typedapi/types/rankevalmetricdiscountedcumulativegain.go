package types

type RankEvalMetricDiscountedCumulativeGain struct {
	K *int `json:"k,omitempty"`

	Normalize *bool `json:"normalize,omitempty"`
}

func (s *RankEvalMetricDiscountedCumulativeGain) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalMetricDiscountedCumulativeGain() *RankEvalMetricDiscountedCumulativeGain {
	_ = "STUB: not implemented"
	return nil
}

type RankEvalMetricDiscountedCumulativeGainVariant interface {
	RankEvalMetricDiscountedCumulativeGainCaster() *RankEvalMetricDiscountedCumulativeGain
}

func (s *RankEvalMetricDiscountedCumulativeGain) RankEvalMetricDiscountedCumulativeGainCaster() *RankEvalMetricDiscountedCumulativeGain {
	_ = "STUB: not implemented"
	return nil
}
