package types

type DataframeAnalysisContainer struct {
	Classification *DataframeAnalysisClassification `json:"classification,omitempty"`

	OutlierDetection *DataframeAnalysisOutlierDetection `json:"outlier_detection,omitempty"`

	Regression *DataframeAnalysisRegression `json:"regression,omitempty"`
}

func NewDataframeAnalysisContainer() *DataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisContainerVariant interface {
	DataframeAnalysisContainerCaster() *DataframeAnalysisContainer
}

func (s *DataframeAnalysisContainer) DataframeAnalysisContainerCaster() *DataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}
