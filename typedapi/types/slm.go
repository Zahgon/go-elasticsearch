package types

type Slm struct {
	Available   bool        `json:"available"`
	Enabled     bool        `json:"enabled"`
	PolicyCount *int        `json:"policy_count,omitempty"`
	PolicyStats *Statistics `json:"policy_stats,omitempty"`
}

func (s *Slm) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlm() *Slm { _ = "STUB: not implemented"; return nil }
