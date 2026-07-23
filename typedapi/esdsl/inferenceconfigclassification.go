package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfigClassification struct {
	v *types.InferenceConfigClassification
}

func NewInferenceConfigClassification() *_inferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) NumTopClasses(numtopclasses int) *_inferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_inferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) PredictionFieldType(predictionfieldtype string) *_inferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) ResultsField(field string) *_inferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) TopClassesResultsField(field string) *_inferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) InferenceConfigCaster() *types.InferenceConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigClassification) InferenceConfigClassificationCaster() *types.InferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}
