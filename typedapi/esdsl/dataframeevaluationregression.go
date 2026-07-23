package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeEvaluationRegression struct {
	v *types.DataframeEvaluationRegression
}

func NewDataframeEvaluationRegression() *_dataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegression) ActualField(field string) *_dataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegression) Metrics(metrics types.DataframeEvaluationRegressionMetricsVariant) *_dataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegression) PredictedField(field string) *_dataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegression) DataframeEvaluationContainerCaster() *types.DataframeEvaluationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegression) DataframeEvaluationRegressionCaster() *types.DataframeEvaluationRegression {
	_ = "STUB: not implemented"
	return nil
}
