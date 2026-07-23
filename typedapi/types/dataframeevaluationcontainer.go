package types

type DataframeEvaluationContainer struct {
	Classification *DataframeEvaluationClassification `json:"classification,omitempty"`

	OutlierDetection *DataframeEvaluationOutlierDetection `json:"outlier_detection,omitempty"`

	Regression *DataframeEvaluationRegression `json:"regression,omitempty"`
}

func NewDataframeEvaluationContainer() *DataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationContainerVariant interface {
	DataframeEvaluationContainerCaster() *DataframeEvaluationContainer
}

func (s *DataframeEvaluationContainer) DataframeEvaluationContainerCaster() *DataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}
