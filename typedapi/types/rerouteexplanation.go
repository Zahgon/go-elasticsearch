package types

type RerouteExplanation struct {
	Command    string            `json:"command"`
	Decisions  []RerouteDecision `json:"decisions"`
	Parameters RerouteParameters `json:"parameters"`
}

func (s *RerouteExplanation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRerouteExplanation() *RerouteExplanation { _ = "STUB: not implemented"; return nil }
