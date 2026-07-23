package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfig struct {
	v *types.InferenceConfig
}

func NewInferenceConfig() *_inferenceConfig { _ = "STUB: not implemented"; return nil }

func (s *_inferenceConfig) Classification(classification types.InferenceConfigClassificationVariant) *_inferenceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfig) Regression(regression types.InferenceConfigRegressionVariant) *_inferenceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfig) InferenceConfigCaster() *types.InferenceConfig {
	_ = "STUB: not implemented"
	return nil
}
