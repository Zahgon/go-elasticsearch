package types

type DataframeClassificationSummaryRecall struct {
	AvgRecall Float64                    `json:"avg_recall"`
	Classes   []DataframeEvaluationClass `json:"classes"`
}

func (s *DataframeClassificationSummaryRecall) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeClassificationSummaryRecall() *DataframeClassificationSummaryRecall {
	_ = "STUB: not implemented"
	return nil
}
