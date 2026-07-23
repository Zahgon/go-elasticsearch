package types

type DataframeEvaluationOutlierDetection struct {
	ActualField string `json:"actual_field"`

	Metrics *DataframeEvaluationOutlierDetectionMetrics `json:"metrics,omitempty"`

	PredictedProbabilityField string `json:"predicted_probability_field"`
}

func (s *DataframeEvaluationOutlierDetection) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationOutlierDetection() *DataframeEvaluationOutlierDetection {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationOutlierDetectionVariant interface {
	DataframeEvaluationOutlierDetectionCaster() *DataframeEvaluationOutlierDetection
}

func (s *DataframeEvaluationOutlierDetection) DataframeEvaluationOutlierDetectionCaster() *DataframeEvaluationOutlierDetection {
	_ = "STUB: not implemented"
	return nil
}
