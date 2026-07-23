package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankEvalMetricRecall struct {
	v *types.RankEvalMetricRecall
}

func NewRankEvalMetricRecall() *_rankEvalMetricRecall { _ = "STUB: not implemented"; return nil }

func (s *_rankEvalMetricRecall) K(k int) *_rankEvalMetricRecall {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetricRecall) RelevantRatingThreshold(relevantratingthreshold int) *_rankEvalMetricRecall {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankEvalMetricRecall) RankEvalMetricRecallCaster() *types.RankEvalMetricRecall {
	_ = "STUB: not implemented"
	return nil
}
