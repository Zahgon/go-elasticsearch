package types

type Ilm struct {
	PolicyCount int                   `json:"policy_count"`
	PolicyStats []IlmPolicyStatistics `json:"policy_stats"`
}

func (s *Ilm) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIlm() *Ilm { _ = "STUB: not implemented"; return nil }
