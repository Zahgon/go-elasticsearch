package types

type RankEvalMetricPrecision struct {
	IgnoreUnlabeled *bool `json:"ignore_unlabeled,omitempty"`

	K *int `json:"k,omitempty"`

	RelevantRatingThreshold *int `json:"relevant_rating_threshold,omitempty"`
}

func (s *RankEvalMetricPrecision) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalMetricPrecision() *RankEvalMetricPrecision { _ = "STUB: not implemented"; return nil }

type RankEvalMetricPrecisionVariant interface {
	RankEvalMetricPrecisionCaster() *RankEvalMetricPrecision
}

func (s *RankEvalMetricPrecision) RankEvalMetricPrecisionCaster() *RankEvalMetricPrecision {
	_ = "STUB: not implemented"
	return nil
}
