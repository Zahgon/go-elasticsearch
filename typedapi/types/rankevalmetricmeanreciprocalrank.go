package types

type RankEvalMetricMeanReciprocalRank struct {
	K *int `json:"k,omitempty"`

	RelevantRatingThreshold *int `json:"relevant_rating_threshold,omitempty"`
}

func (s *RankEvalMetricMeanReciprocalRank) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalMetricMeanReciprocalRank() *RankEvalMetricMeanReciprocalRank {
	_ = "STUB: not implemented"
	return nil
}

type RankEvalMetricMeanReciprocalRankVariant interface {
	RankEvalMetricMeanReciprocalRankCaster() *RankEvalMetricMeanReciprocalRank
}

func (s *RankEvalMetricMeanReciprocalRank) RankEvalMetricMeanReciprocalRankCaster() *RankEvalMetricMeanReciprocalRank {
	_ = "STUB: not implemented"
	return nil
}
