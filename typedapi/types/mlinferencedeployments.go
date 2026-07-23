package types

type MlInferenceDeployments struct {
	Count           int                          `json:"count"`
	InferenceCounts JobStatistics                `json:"inference_counts"`
	ModelSizesBytes JobStatistics                `json:"model_sizes_bytes"`
	TimeMs          MlInferenceDeploymentsTimeMs `json:"time_ms"`
}

func (s *MlInferenceDeployments) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMlInferenceDeployments() *MlInferenceDeployments { _ = "STUB: not implemented"; return nil }
