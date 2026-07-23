package types

type TrainedModelInferenceClassImportance struct {
	ClassName  string  `json:"class_name"`
	Importance Float64 `json:"importance"`
}

func (s *TrainedModelInferenceClassImportance) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelInferenceClassImportance() *TrainedModelInferenceClassImportance {
	_ = "STUB: not implemented"
	return nil
}
