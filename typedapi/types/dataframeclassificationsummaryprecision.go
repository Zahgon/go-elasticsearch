package types

type DataframeClassificationSummaryPrecision struct {
	AvgPrecision Float64                    `json:"avg_precision"`
	Classes      []DataframeEvaluationClass `json:"classes"`
}

func (s *DataframeClassificationSummaryPrecision) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeClassificationSummaryPrecision() *DataframeClassificationSummaryPrecision {
	_ = "STUB: not implemented"
	return nil
}
