package types

type DataframeEvaluationClassification struct {
	ActualField string `json:"actual_field"`

	Metrics *DataframeEvaluationClassificationMetrics `json:"metrics,omitempty"`

	PredictedField *string `json:"predicted_field,omitempty"`

	TopClassesField *string `json:"top_classes_field,omitempty"`
}

func (s *DataframeEvaluationClassification) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationClassification() *DataframeEvaluationClassification {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationClassificationVariant interface {
	DataframeEvaluationClassificationCaster() *DataframeEvaluationClassification
}

func (s *DataframeEvaluationClassification) DataframeEvaluationClassificationCaster() *DataframeEvaluationClassification {
	_ = "STUB: not implemented"
	return nil
}
