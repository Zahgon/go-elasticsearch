package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _queryFeatureExtractor struct {
	v *types.QueryFeatureExtractor
}

func NewQueryFeatureExtractor(featurename string, query types.QueryVariant) *_queryFeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryFeatureExtractor) DefaultScore(defaultscore float32) *_queryFeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryFeatureExtractor) FeatureName(featurename string) *_queryFeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryFeatureExtractor) Query(query types.QueryVariant) *_queryFeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryFeatureExtractor) QueryFeatureExtractorCaster() *types.QueryFeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}
