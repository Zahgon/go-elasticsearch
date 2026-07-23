package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _featureExtractor struct {
	v types.FeatureExtractor
}

func NewFeatureExtractor(featureextractor types.QueryFeatureExtractor) *_featureExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (u *_featureExtractor) FeatureExtractorCaster() *types.FeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}
