package types

type InferenceConfigRegression struct {
	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`
}

func (s *InferenceConfigRegression) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceConfigRegression() *InferenceConfigRegression {
	_ = "STUB: not implemented"
	return nil
}

type InferenceConfigRegressionVariant interface {
	InferenceConfigRegressionCaster() *InferenceConfigRegression
}

func (s *InferenceConfigRegression) InferenceConfigRegressionCaster() *InferenceConfigRegression {
	_ = "STUB: not implemented"
	return nil
}
