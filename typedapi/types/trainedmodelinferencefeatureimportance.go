package types

type TrainedModelInferenceFeatureImportance struct {
	Classes     []TrainedModelInferenceClassImportance `json:"classes,omitempty"`
	FeatureName string                                 `json:"feature_name"`
	Importance  *Float64                               `json:"importance,omitempty"`
}

func (s *TrainedModelInferenceFeatureImportance) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelInferenceFeatureImportance() *TrainedModelInferenceFeatureImportance {
	_ = "STUB: not implemented"
	return nil
}
