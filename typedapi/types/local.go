package types

type Local struct {
	Type string `json:"type"`
}

func (s *Local) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLocal() *Local { _ = "STUB: not implemented"; return nil }
