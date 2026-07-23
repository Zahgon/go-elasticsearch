package types

type RankEvalMetric struct {
	Dcg                    *RankEvalMetricDiscountedCumulativeGain `json:"dcg,omitempty"`
	ExpectedReciprocalRank *RankEvalMetricExpectedReciprocalRank   `json:"expected_reciprocal_rank,omitempty"`
	MeanReciprocalRank     *RankEvalMetricMeanReciprocalRank       `json:"mean_reciprocal_rank,omitempty"`
	Precision              *RankEvalMetricPrecision                `json:"precision,omitempty"`
	Recall                 *RankEvalMetricRecall                   `json:"recall,omitempty"`
}

func NewRankEvalMetric() *RankEvalMetric { _ = "STUB: not implemented"; return nil }

type RankEvalMetricVariant interface {
	RankEvalMetricCaster() *RankEvalMetric
}

func (s *RankEvalMetric) RankEvalMetricCaster() *RankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}
