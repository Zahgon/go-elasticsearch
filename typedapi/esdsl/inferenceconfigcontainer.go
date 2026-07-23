package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfigContainer struct {
	v *types.InferenceConfigContainer
}

func NewInferenceConfigContainer() *_inferenceConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigContainer) Classification(classification types.ClassificationInferenceOptionsVariant) *_inferenceConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigContainer) Regression(regression types.RegressionInferenceOptionsVariant) *_inferenceConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigContainer) InferenceConfigContainerCaster() *types.InferenceConfigContainer {
	_ = "STUB: not implemented"
	return nil
}
