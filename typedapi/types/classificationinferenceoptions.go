package types

type ClassificationInferenceOptions struct {
	NumTopClasses *int `json:"num_top_classes,omitempty"`

	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`

	PredictionFieldType *string `json:"prediction_field_type,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	TopClassesResultsField *string `json:"top_classes_results_field,omitempty"`
}

func (s *ClassificationInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewClassificationInferenceOptions() *ClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type ClassificationInferenceOptionsVariant interface {
	ClassificationInferenceOptionsCaster() *ClassificationInferenceOptions
}

func (s *ClassificationInferenceOptions) ClassificationInferenceOptionsCaster() *ClassificationInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
