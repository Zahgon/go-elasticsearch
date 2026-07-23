package types

type DataframeEvaluationClassificationMetricsAucRoc struct {
	ClassName *string `json:"class_name,omitempty"`

	IncludeCurve *bool `json:"include_curve,omitempty"`
}

func (s *DataframeEvaluationClassificationMetricsAucRoc) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationClassificationMetricsAucRoc() *DataframeEvaluationClassificationMetricsAucRoc {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationClassificationMetricsAucRocVariant interface {
	DataframeEvaluationClassificationMetricsAucRocCaster() *DataframeEvaluationClassificationMetricsAucRoc
}

func (s *DataframeEvaluationClassificationMetricsAucRoc) DataframeEvaluationClassificationMetricsAucRocCaster() *DataframeEvaluationClassificationMetricsAucRoc {
	_ = "STUB: not implemented"
	return nil
}
