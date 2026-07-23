package types

type DataframeEvaluationSummaryAucRocCurveItem struct {
	Fpr       Float64 `json:"fpr"`
	Threshold Float64 `json:"threshold"`
	Tpr       Float64 `json:"tpr"`
}

func (s *DataframeEvaluationSummaryAucRocCurveItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationSummaryAucRocCurveItem() *DataframeEvaluationSummaryAucRocCurveItem {
	_ = "STUB: not implemented"
	return nil
}
