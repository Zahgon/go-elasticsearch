package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeEvaluationContainer struct {
	v *types.DataframeEvaluationContainer
}

func NewDataframeEvaluationContainer() *_dataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationContainer) Classification(classification types.DataframeEvaluationClassificationVariant) *_dataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationContainer) OutlierDetection(outlierdetection types.DataframeEvaluationOutlierDetectionVariant) *_dataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationContainer) Regression(regression types.DataframeEvaluationRegressionVariant) *_dataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationContainer) DataframeEvaluationContainerCaster() *types.DataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}
