package types

type Indeterminable struct {
	Reason string `json:"reason"`
}

func (s *Indeterminable) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndeterminable() *Indeterminable { _ = "STUB: not implemented"; return nil }
