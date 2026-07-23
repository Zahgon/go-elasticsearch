package types

type InferenceFeatureImportance struct {
	Classes     []InferenceClassImportance `json:"classes,omitempty"`
	FeatureName string                     `json:"feature_name"`
	Importance  *Float64                   `json:"importance,omitempty"`
}

func (s *InferenceFeatureImportance) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceFeatureImportance() *InferenceFeatureImportance {
	_ = "STUB: not implemented"
	return nil
}
