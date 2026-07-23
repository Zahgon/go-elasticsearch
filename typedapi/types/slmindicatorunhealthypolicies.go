package types

type SlmIndicatorUnhealthyPolicies struct {
	Count                       int64            `json:"count"`
	InvocationsSinceLastSuccess map[string]int64 `json:"invocations_since_last_success,omitempty"`
}

func (s *SlmIndicatorUnhealthyPolicies) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSlmIndicatorUnhealthyPolicies() *SlmIndicatorUnhealthyPolicies {
	_ = "STUB: not implemented"
	return nil
}
