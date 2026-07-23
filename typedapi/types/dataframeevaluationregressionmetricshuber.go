package types

type DataframeEvaluationRegressionMetricsHuber struct {
	Delta *Float64 `json:"delta,omitempty"`
}

func (s *DataframeEvaluationRegressionMetricsHuber) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationRegressionMetricsHuber() *DataframeEvaluationRegressionMetricsHuber {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationRegressionMetricsHuberVariant interface {
	DataframeEvaluationRegressionMetricsHuberCaster() *DataframeEvaluationRegressionMetricsHuber
}

func (s *DataframeEvaluationRegressionMetricsHuber) DataframeEvaluationRegressionMetricsHuberCaster() *DataframeEvaluationRegressionMetricsHuber {
	_ = "STUB: not implemented"
	return nil
}
