package types

type DataframeOutlierDetectionSummary struct {
	AucRoc *DataframeEvaluationSummaryAucRoc `json:"auc_roc,omitempty"`

	ConfusionMatrix map[string]ConfusionMatrixThreshold `json:"confusion_matrix,omitempty"`

	Precision map[string]Float64 `json:"precision,omitempty"`

	Recall map[string]Float64 `json:"recall,omitempty"`
}

func NewDataframeOutlierDetectionSummary() *DataframeOutlierDetectionSummary {
	_ = "STUB: not implemented"
	return nil
}
