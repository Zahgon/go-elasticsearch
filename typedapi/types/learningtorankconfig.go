package types

import (
	"encoding/json"
)

type LearningToRankConfig struct {
	DefaultParams                 map[string]json.RawMessage         `json:"default_params,omitempty"`
	FeatureExtractors             []map[string]QueryFeatureExtractor `json:"feature_extractors,omitempty"`
	NumTopFeatureImportanceValues int                                `json:"num_top_feature_importance_values"`
}

func (s *LearningToRankConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLearningToRankConfig() *LearningToRankConfig { _ = "STUB: not implemented"; return nil }

type LearningToRankConfigVariant interface {
	LearningToRankConfigCaster() *LearningToRankConfig
}

func (s *LearningToRankConfig) LearningToRankConfigCaster() *LearningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}
