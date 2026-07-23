package types

type InferenceResponseResult struct {
	Entities []TrainedModelEntities `json:"entities,omitempty"`

	FeatureImportance []TrainedModelInferenceFeatureImportance `json:"feature_importance,omitempty"`

	IsTruncated *bool `json:"is_truncated,omitempty"`

	PredictedValue [][]ScalarValue `json:"predicted_value,omitempty"`

	PredictedValueSequence *string `json:"predicted_value_sequence,omitempty"`

	PredictionProbability *Float64 `json:"prediction_probability,omitempty"`

	PredictionScore *Float64 `json:"prediction_score,omitempty"`

	TopClasses []TopClassEntry `json:"top_classes,omitempty"`

	Warning *string `json:"warning,omitempty"`
}

func (s *InferenceResponseResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceResponseResult() *InferenceResponseResult { _ = "STUB: not implemented"; return nil }
