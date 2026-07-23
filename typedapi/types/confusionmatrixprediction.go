package types

type ConfusionMatrixPrediction struct {
	Count          int    `json:"count"`
	PredictedClass string `json:"predicted_class"`
}

func (s *ConfusionMatrixPrediction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConfusionMatrixPrediction() *ConfusionMatrixPrediction {
	_ = "STUB: not implemented"
	return nil
}
