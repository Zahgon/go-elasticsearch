package types

import (
	"encoding/json"
)

type DataframeEvaluationClassificationMetrics struct {
	Accuracy map[string]json.RawMessage `json:"accuracy,omitempty"`

	AucRoc *DataframeEvaluationClassificationMetricsAucRoc `json:"auc_roc,omitempty"`

	MulticlassConfusionMatrix map[string]json.RawMessage `json:"multiclass_confusion_matrix,omitempty"`

	Precision map[string]json.RawMessage `json:"precision,omitempty"`

	Recall map[string]json.RawMessage `json:"recall,omitempty"`
}

func NewDataframeEvaluationClassificationMetrics() *DataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationClassificationMetricsVariant interface {
	DataframeEvaluationClassificationMetricsCaster() *DataframeEvaluationClassificationMetrics
}

func (s *DataframeEvaluationClassificationMetrics) DataframeEvaluationClassificationMetricsCaster() *DataframeEvaluationClassificationMetrics {
	_ = "STUB: not implemented"
	return nil
}
