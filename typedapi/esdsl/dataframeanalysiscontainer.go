package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisContainer struct {
	v *types.DataframeAnalysisContainer
}

func NewDataframeAnalysisContainer() *_dataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisContainer) Classification(classification types.DataframeAnalysisClassificationVariant) *_dataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisContainer) OutlierDetection(outlierdetection types.DataframeAnalysisOutlierDetectionVariant) *_dataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisContainer) Regression(regression types.DataframeAnalysisRegressionVariant) *_dataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisContainer) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}
