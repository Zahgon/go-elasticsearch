package types

type IlmPolicyStatistics struct {
	IndicesManaged int         `json:"indices_managed"`
	Phases         UsagePhases `json:"phases"`
}

func (s *IlmPolicyStatistics) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIlmPolicyStatistics() *IlmPolicyStatistics { _ = "STUB: not implemented"; return nil }
