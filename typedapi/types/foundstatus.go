package types

type FoundStatus struct {
	Found bool `json:"found"`
}

func (s *FoundStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFoundStatus() *FoundStatus { _ = "STUB: not implemented"; return nil }
