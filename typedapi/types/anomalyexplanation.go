package types

type AnomalyExplanation struct {
	AnomalyCharacteristicsImpact *int `json:"anomaly_characteristics_impact,omitempty"`

	AnomalyLength *int `json:"anomaly_length,omitempty"`

	AnomalyType *string `json:"anomaly_type,omitempty"`

	HighVariancePenalty *bool `json:"high_variance_penalty,omitempty"`

	IncompleteBucketPenalty *bool `json:"incomplete_bucket_penalty,omitempty"`

	LowerConfidenceBound *Float64 `json:"lower_confidence_bound,omitempty"`

	MultiBucketImpact *int `json:"multi_bucket_impact,omitempty"`

	SingleBucketImpact *int `json:"single_bucket_impact,omitempty"`

	TypicalValue *Float64 `json:"typical_value,omitempty"`

	UpperConfidenceBound *Float64 `json:"upper_confidence_bound,omitempty"`
}

func (s *AnomalyExplanation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAnomalyExplanation() *AnomalyExplanation { _ = "STUB: not implemented"; return nil }
