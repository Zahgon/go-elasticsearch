package types

type RankEvalMetricRecall struct {
	K *int `json:"k,omitempty"`

	RelevantRatingThreshold *int `json:"relevant_rating_threshold,omitempty"`
}

func (s *RankEvalMetricRecall) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalMetricRecall() *RankEvalMetricRecall { _ = "STUB: not implemented"; return nil }

type RankEvalMetricRecallVariant interface {
	RankEvalMetricRecallCaster() *RankEvalMetricRecall
}

func (s *RankEvalMetricRecall) RankEvalMetricRecallCaster() *RankEvalMetricRecall {
	_ = "STUB: not implemented"
	return nil
}
