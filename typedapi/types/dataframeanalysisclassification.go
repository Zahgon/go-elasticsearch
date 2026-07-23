package types

type DataframeAnalysisClassification struct {
	Alpha                    *Float64 `json:"alpha,omitempty"`
	ClassAssignmentObjective *string  `json:"class_assignment_objective,omitempty"`

	DependentVariable string `json:"dependent_variable"`

	DownsampleFactor *Float64 `json:"downsample_factor,omitempty"`

	EarlyStoppingEnabled *bool `json:"early_stopping_enabled,omitempty"`

	Eta *Float64 `json:"eta,omitempty"`

	EtaGrowthRatePerTree *Float64 `json:"eta_growth_rate_per_tree,omitempty"`

	FeatureBagFraction *Float64 `json:"feature_bag_fraction,omitempty"`

	FeatureProcessors []DataframeAnalysisFeatureProcessor `json:"feature_processors,omitempty"`

	Gamma *Float64 `json:"gamma,omitempty"`

	Lambda *Float64 `json:"lambda,omitempty"`

	MaxOptimizationRoundsPerHyperparameter *int `json:"max_optimization_rounds_per_hyperparameter,omitempty"`

	MaxTrees *int `json:"max_trees,omitempty"`

	NumTopClasses *int `json:"num_top_classes,omitempty"`

	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`

	PredictionFieldName *string `json:"prediction_field_name,omitempty"`

	RandomizeSeed *Float64 `json:"randomize_seed,omitempty"`

	SoftTreeDepthLimit *int `json:"soft_tree_depth_limit,omitempty"`

	SoftTreeDepthTolerance *Float64 `json:"soft_tree_depth_tolerance,omitempty"`

	TrainingPercent Percentage `json:"training_percent,omitempty"`
}

func (s *DataframeAnalysisClassification) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalysisClassification() *DataframeAnalysisClassification {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalysisClassificationVariant interface {
	DataframeAnalysisClassificationCaster() *DataframeAnalysisClassification
}

func (s *DataframeAnalysisClassification) DataframeAnalysisClassificationCaster() *DataframeAnalysisClassification {
	_ = "STUB: not implemented"
	return nil
}
