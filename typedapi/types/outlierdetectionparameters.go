package types

type OutlierDetectionParameters struct {
	ComputeFeatureInfluence *bool `json:"compute_feature_influence,omitempty"`

	FeatureInfluenceThreshold *Float64 `json:"feature_influence_threshold,omitempty"`

	Method *string `json:"method,omitempty"`

	NNeighbors *int `json:"n_neighbors,omitempty"`

	OutlierFraction *Float64 `json:"outlier_fraction,omitempty"`

	StandardizationEnabled *bool `json:"standardization_enabled,omitempty"`
}

func (s *OutlierDetectionParameters) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewOutlierDetectionParameters() *OutlierDetectionParameters {
	_ = "STUB: not implemented"
	return nil
}
