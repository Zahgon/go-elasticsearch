package types

type InferenceConfigClassification struct {
	NumTopClasses *int `json:"num_top_classes,omitempty"`

	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`

	PredictionFieldType *string `json:"prediction_field_type,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	TopClassesResultsField *string `json:"top_classes_results_field,omitempty"`
}

func (s *InferenceConfigClassification) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceConfigClassification() *InferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}

type InferenceConfigClassificationVariant interface {
	InferenceConfigClassificationCaster() *InferenceConfigClassification
}

func (s *InferenceConfigClassification) InferenceConfigClassificationCaster() *InferenceConfigClassification {
	_ = "STUB: not implemented"
	return nil
}
