package types

type MlInferenceDeploymentsTimeMs struct {
	Avg Float64 `json:"avg"`
}

func (s *MlInferenceDeploymentsTimeMs) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMlInferenceDeploymentsTimeMs() *MlInferenceDeploymentsTimeMs {
	_ = "STUB: not implemented"
	return nil
}
