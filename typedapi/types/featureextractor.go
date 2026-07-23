package types

type FeatureExtractor QueryFeatureExtractor

type FeatureExtractorVariant interface {
	FeatureExtractorCaster() *FeatureExtractor
}
