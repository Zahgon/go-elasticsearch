package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankEvalMetric struct {
	v *types.RankEvalMetric
}

func NewRankEvalMetric() *_rankEvalMetric { _ = "STUB: not implemented"; return nil }

func (s *_rankEvalMetric) Dcg(dcg types.RankEvalMetricDiscountedCumulativeGainVariant) *_rankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetric) ExpectedReciprocalRank(expectedreciprocalrank types.RankEvalMetricExpectedReciprocalRankVariant) *_rankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetric) MeanReciprocalRank(meanreciprocalrank types.RankEvalMetricMeanReciprocalRankVariant) *_rankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetric) Precision(precision types.RankEvalMetricPrecisionVariant) *_rankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetric) Recall(recall types.RankEvalMetricRecallVariant) *_rankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetric) RankEvalMetricCaster() *types.RankEvalMetric {
	_ = "STUB: not implemented"
	return nil
}
