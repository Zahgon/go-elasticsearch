package types

type DataframeClassificationSummaryAccuracy struct {
	Classes         []DataframeEvaluationClass `json:"classes"`
	OverallAccuracy Float64                    `json:"overall_accuracy"`
}

func (s *DataframeClassificationSummaryAccuracy) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeClassificationSummaryAccuracy() *DataframeClassificationSummaryAccuracy {
	_ = "STUB: not implemented"
	return nil
}
