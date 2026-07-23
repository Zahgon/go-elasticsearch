package types

import (
	"encoding/json"
)

type DataframeEvaluationOutlierDetectionMetrics struct {
	AucRoc *DataframeEvaluationClassificationMetricsAucRoc `json:"auc_roc,omitempty"`

	ConfusionMatrix map[string]json.RawMessage `json:"confusion_matrix,omitempty"`

	Precision map[string]json.RawMessage `json:"precision,omitempty"`

	Recall map[string]json.RawMessage `json:"recall,omitempty"`
}

func NewDataframeEvaluationOutlierDetectionMetrics() *DataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}

type DataframeEvaluationOutlierDetectionMetricsVariant interface {
	DataframeEvaluationOutlierDetectionMetricsCaster() *DataframeEvaluationOutlierDetectionMetrics
}

func (s *DataframeEvaluationOutlierDetectionMetrics) DataframeEvaluationOutlierDetectionMetricsCaster() *DataframeEvaluationOutlierDetectionMetrics {
	_ = "STUB: not implemented"
	return nil
}
