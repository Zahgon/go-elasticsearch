package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankEvalMetricPrecision struct {
	v *types.RankEvalMetricPrecision
}

func NewRankEvalMetricPrecision() *_rankEvalMetricPrecision { _ = "STUB: not implemented"; return nil }

func (s *_rankEvalMetricPrecision) IgnoreUnlabeled(ignoreunlabeled bool) *_rankEvalMetricPrecision {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetricPrecision) K(k int) *_rankEvalMetricPrecision {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetricPrecision) RelevantRatingThreshold(relevantratingthreshold int) *_rankEvalMetricPrecision {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetricPrecision) RankEvalMetricPrecisionCaster() *types.RankEvalMetricPrecision {
	_ = "STUB: not implemented"
	return nil
}
