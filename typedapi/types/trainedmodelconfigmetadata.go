package types

type TrainedModelConfigMetadata struct {
	FeatureImportanceBaseline map[string]string `json:"feature_importance_baseline,omitempty"`

	Hyperparameters []Hyperparameter `json:"hyperparameters,omitempty"`
	ModelAliases    []string         `json:"model_aliases,omitempty"`

	TotalFeatureImportance []TotalFeatureImportance `json:"total_feature_importance,omitempty"`
}

func NewTrainedModelConfigMetadata() *TrainedModelConfigMetadata {
	_ = "STUB: not implemented"
	return nil
}
