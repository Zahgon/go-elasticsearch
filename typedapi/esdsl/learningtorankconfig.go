package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _learningToRankConfig struct {
	v *types.LearningToRankConfig
}

func NewLearningToRankConfig(numtopfeatureimportancevalues int) *_learningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRankConfig) DefaultParams(defaultparams map[string]json.RawMessage) *_learningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRankConfig) AddDefaultParam(key string, value json.RawMessage) *_learningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRankConfig) FeatureExtractors(featureextractors []map[string]types.QueryFeatureExtractor) *_learningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRankConfig) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_learningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRankConfig) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRankConfig) LearningToRankConfigCaster() *types.LearningToRankConfig {
	_ = "STUB: not implemented"
	return nil
}
