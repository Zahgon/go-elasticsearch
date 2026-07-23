package types

import (
	"encoding/json"
)

type DataframeEvaluationRegressionMetrics struct {
	Huber *DataframeEvaluationRegressionMetricsHuber `json:"huber,omitempty"`

	Mse map[string]json.RawMessage `json:"mse,omitempty"`

	Msle *DataframeEvaluationRegressionMetricsMsle `json:"msle,omitempty"`

	RSquared map[string]json.RawMessage `json:"r_squared,omitempty"`
}

func NewDataframeEvaluationRegressionMetrics() *DataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationRegressionMetricsVariant interface {
	DataframeEvaluationRegressionMetricsCaster() *DataframeEvaluationRegressionMetrics
}

func (s *DataframeEvaluationRegressionMetrics) DataframeEvaluationRegressionMetricsCaster() *DataframeEvaluationRegressionMetrics {
	_ = "STUB: not implemented"
	return nil
}
