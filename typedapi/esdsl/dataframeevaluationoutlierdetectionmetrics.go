package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeEvaluationOutlierDetectionMetrics struct {
	v *types.DataframeEvaluationOutlierDetectionMetrics
}

func NewDataframeEvaluationOutlierDetectionMetrics() *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) ConfusionMatrix(confusionmatrix map[string]json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AddConfusionMatrix(key string, value json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AucRoc(aucroc types.DataframeEvaluationClassificationMetricsAucRocVariant) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) Precision(precision map[string]json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AddPrecision(key string, value json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) Recall(recall map[string]json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) AddRecall(key string, value json.RawMessage) *_dataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationOutlierDetectionMetrics) DataframeEvaluationOutlierDetectionMetricsCaster() *types.DataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}
