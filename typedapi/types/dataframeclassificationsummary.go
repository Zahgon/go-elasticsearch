package types

type DataframeClassificationSummary struct {
	Accuracy *DataframeClassificationSummaryAccuracy `json:"accuracy,omitempty"`

	AucRoc *DataframeEvaluationSummaryAucRoc `json:"auc_roc,omitempty"`

	MulticlassConfusionMatrix *DataframeClassificationSummaryMulticlassConfusionMatrix `json:"multiclass_confusion_matrix,omitempty"`

	Precision *DataframeClassificationSummaryPrecision `json:"precision,omitempty"`

	Recall *DataframeClassificationSummaryRecall `json:"recall,omitempty"`
}

func NewDataframeClassificationSummary() *DataframeClassificationSummary {
	_ = "STUB: not implemented"
	return nil
}
