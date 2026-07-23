package types

type DataframeEvaluationRegression struct {
	ActualField string `json:"actual_field"`

	Metrics *DataframeEvaluationRegressionMetrics `json:"metrics,omitempty"`

	PredictedField string `json:"predicted_field"`
}

func (s *DataframeEvaluationRegression) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationRegression() *DataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationRegressionVariant interface {
	DataframeEvaluationRegressionCaster() *DataframeEvaluationRegression
}

func (s *DataframeEvaluationRegression) DataframeEvaluationRegressionCaster() *DataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}
