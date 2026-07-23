package types

type QueryFeatureExtractor struct {
	DefaultScore *float32 `json:"default_score,omitempty"`
	FeatureName  string   `json:"feature_name"`
	Query        Query    `json:"query"`
}

func (s *QueryFeatureExtractor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewQueryFeatureExtractor() *QueryFeatureExtractor { _ = "STUB: not implemented"; return nil }

type QueryFeatureExtractorVariant interface {
	QueryFeatureExtractorCaster() *QueryFeatureExtractor
}

func (s *QueryFeatureExtractor) QueryFeatureExtractorCaster() *QueryFeatureExtractor {
	_ = "STUB: not implemented"
	return nil
}
