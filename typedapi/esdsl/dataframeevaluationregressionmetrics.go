package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeEvaluationRegressionMetrics struct {
	v *types.DataframeEvaluationRegressionMetrics
}

func NewDataframeEvaluationRegressionMetrics() *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) Huber(huber types.DataframeEvaluationRegressionMetricsHuberVariant) *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) Mse(mse map[string]json.RawMessage) *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) AddMse(key string, value json.RawMessage) *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) Msle(msle types.DataframeEvaluationRegressionMetricsMsleVariant) *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) RSquared(rsquared map[string]json.RawMessage) *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) AddRSquared(key string, value json.RawMessage) *_dataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationRegressionMetrics) DataframeEvaluationRegressionMetricsCaster() *types.DataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}
