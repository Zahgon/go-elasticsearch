package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankFeatureQuery struct {
	v *types.RankFeatureQuery
}

func NewRankFeatureQuery() *_rankFeatureQuery { _ = "STUB: not implemented"; return nil }

func (s *_rankFeatureQuery) Field(field string) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) Linear(linear types.RankFeatureFunctionLinearVariant) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) Log(log types.RankFeatureFunctionLogarithmVariant) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) Saturation(saturation types.RankFeatureFunctionSaturationVariant) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) Sigmoid(sigmoid types.RankFeatureFunctionSigmoidVariant) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) Boost(boost float32) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) QueryName_(queryname_ string) *_rankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rankFeatureQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_rankFeatureQuery) RankFeatureQueryCaster() *types.RankFeatureQuery {
	_ = "STUB: not implemented"
	return nil
}
