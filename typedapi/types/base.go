package types

type Base struct {
	Available bool `json:"available"`
	Enabled   bool `json:"enabled"`
}

func (s *Base) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBase() *Base { _ = "STUB: not implemented"; return nil }
