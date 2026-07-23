package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeEvaluationClassificationMetrics struct {
	v *types.DataframeEvaluationClassificationMetrics
}

func NewDataframeEvaluationClassificationMetrics() *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) Accuracy(accuracy map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) AddAccuracy(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) MulticlassConfusionMatrix(multiclassconfusionmatrix map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) AddMulticlassConfusionMatrix(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) AucRoc(aucroc types.DataframeEvaluationClassificationMetricsAucRocVariant) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) Precision(precision map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) AddPrecision(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) Recall(recall map[string]json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) AddRecall(key string, value json.RawMessage) *_dataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeEvaluationClassificationMetrics) DataframeEvaluationClassificationMetricsCaster() *types.DataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}
