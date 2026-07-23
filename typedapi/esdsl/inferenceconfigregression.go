package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfigRegression struct {
	v *types.InferenceConfigRegression
}

func NewInferenceConfigRegression() *_inferenceConfigRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigRegression) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_inferenceConfigRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigRegression) ResultsField(field string) *_inferenceConfigRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigRegression) InferenceConfigCaster() *types.InferenceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigRegression) InferenceConfigRegressionCaster() *types.InferenceConfigRegression {
	_ = "STUB: not implemented"
	return nil
}
