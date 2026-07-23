package types

type NestedIdentity struct {
	Field   string          `json:"field"`
	Nested_ *NestedIdentity `json:"_nested,omitempty"`
	Offset  int             `json:"offset"`
}

func (s *NestedIdentity) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNestedIdentity() *NestedIdentity { _ = "STUB: not implemented"; return nil }
