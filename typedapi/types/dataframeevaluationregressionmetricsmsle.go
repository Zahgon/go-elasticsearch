package types

type DataframeEvaluationRegressionMetricsMsle struct {
	Offset *Float64 `json:"offset,omitempty"`
}

func (s *DataframeEvaluationRegressionMetricsMsle) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationRegressionMetricsMsle() *DataframeEvaluationRegressionMetricsMsle {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationRegressionMetricsMsleVariant interface {
	DataframeEvaluationRegressionMetricsMsleCaster() *DataframeEvaluationRegressionMetricsMsle
}

func (s *DataframeEvaluationRegressionMetricsMsle) DataframeEvaluationRegressionMetricsMsleCaster() *DataframeEvaluationRegressionMetricsMsle {
	_ = "STUB: not implemented"
	return nil
}
