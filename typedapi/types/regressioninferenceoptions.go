package types

type RegressionInferenceOptions struct {
	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`
}

func (s *RegressionInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRegressionInferenceOptions() *RegressionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type RegressionInferenceOptionsVariant interface {
	RegressionInferenceOptionsCaster() *RegressionInferenceOptions
}

func (s *RegressionInferenceOptions) RegressionInferenceOptionsCaster() *RegressionInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
