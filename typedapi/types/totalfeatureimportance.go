package types

type TotalFeatureImportance struct {
	Classes []TotalFeatureImportanceClass `json:"classes"`

	FeatureName string `json:"feature_name"`

	Importance []TotalFeatureImportanceStatistics `json:"importance"`
}

func (s *TotalFeatureImportance) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTotalFeatureImportance() *TotalFeatureImportance { _ = "STUB: not implemented"; return nil }
