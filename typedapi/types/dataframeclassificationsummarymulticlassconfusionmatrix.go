package types

type DataframeClassificationSummaryMulticlassConfusionMatrix struct {
	ConfusionMatrix       []ConfusionMatrixItem `json:"confusion_matrix"`
	OtherActualClassCount int                   `json:"other_actual_class_count"`
}

func (s *DataframeClassificationSummaryMulticlassConfusionMatrix) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeClassificationSummaryMulticlassConfusionMatrix() *DataframeClassificationSummaryMulticlassConfusionMatrix {
	_ = "STUB: not implemented"
	return nil
}
