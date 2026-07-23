package types

type DataframeEvaluationSummaryAucRoc struct {
	Curve []DataframeEvaluationSummaryAucRocCurveItem `json:"curve,omitempty"`
	Value Float64                                     `json:"value"`
}

func (s *DataframeEvaluationSummaryAucRoc) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationSummaryAucRoc() *DataframeEvaluationSummaryAucRoc {
	_ = "STUB: not implemented"
	return nil
}
