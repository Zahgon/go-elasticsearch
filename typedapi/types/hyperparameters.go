package types

type Hyperparameters struct {
	Alpha *Float64 `json:"alpha,omitempty"`

	DownsampleFactor *Float64 `json:"downsample_factor,omitempty"`

	Eta *Float64 `json:"eta,omitempty"`

	EtaGrowthRatePerTree *Float64 `json:"eta_growth_rate_per_tree,omitempty"`

	FeatureBagFraction *Float64 `json:"feature_bag_fraction,omitempty"`

	Gamma *Float64 `json:"gamma,omitempty"`

	Lambda *Float64 `json:"lambda,omitempty"`

	MaxAttemptsToAddTree *int `json:"max_attempts_to_add_tree,omitempty"`

	MaxOptimizationRoundsPerHyperparameter *int `json:"max_optimization_rounds_per_hyperparameter,omitempty"`

	MaxTrees *int `json:"max_trees,omitempty"`

	NumFolds *int `json:"num_folds,omitempty"`

	NumSplitsPerFeature *int `json:"num_splits_per_feature,omitempty"`

	SoftTreeDepthLimit *int `json:"soft_tree_depth_limit,omitempty"`

	SoftTreeDepthTolerance *Float64 `json:"soft_tree_depth_tolerance,omitempty"`
}

func (s *Hyperparameters) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHyperparameters() *Hyperparameters { _ = "STUB: not implemented"; return nil }
