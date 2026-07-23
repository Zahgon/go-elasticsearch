package types

type DataframeRegressionSummary struct {
	Huber *DataframeEvaluationValue `json:"huber,omitempty"`

	Mse *DataframeEvaluationValue `json:"mse,omitempty"`

	Msle *DataframeEvaluationValue `json:"msle,omitempty"`

	RSquared *DataframeEvaluationValue `json:"r_squared,omitempty"`
}

func NewDataframeRegressionSummary() *DataframeRegressionSummary {
	_ = "STUB: not implemented"
	return nil
}
