package types

type RerouteDecision struct {
	Decider     string `json:"decider"`
	Decision    string `json:"decision"`
	Explanation string `json:"explanation"`
}

func (s *RerouteDecision) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRerouteDecision() *RerouteDecision { _ = "STUB: not implemented"; return nil }
