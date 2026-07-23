package types

type ConfusionMatrixItem struct {
	ActualClass                 string                      `json:"actual_class"`
	ActualClassDocCount         int                         `json:"actual_class_doc_count"`
	OtherPredictedClassDocCount int                         `json:"other_predicted_class_doc_count"`
	PredictedClasses            []ConfusionMatrixPrediction `json:"predicted_classes"`
}

func (s *ConfusionMatrixItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConfusionMatrixItem() *ConfusionMatrixItem { _ = "STUB: not implemented"; return nil }
