package types

type RankEvalMetricExpectedReciprocalRank struct {
	K *int `json:"k,omitempty"`

	MaximumRelevance int `json:"maximum_relevance"`
}

func (s *RankEvalMetricExpectedReciprocalRank) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRankEvalMetricExpectedReciprocalRank() *RankEvalMetricExpectedReciprocalRank {
	_ = "STUB: not implemented"
	return nil
}

type RankEvalMetricExpectedReciprocalRankVariant interface {
	RankEvalMetricExpectedReciprocalRankCaster() *RankEvalMetricExpectedReciprocalRank
}

func (s *RankEvalMetricExpectedReciprocalRank) RankEvalMetricExpectedReciprocalRankCaster() *RankEvalMetricExpectedReciprocalRank {
	_ = "STUB: not implemented"
	return nil
}
