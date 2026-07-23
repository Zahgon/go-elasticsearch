package types

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *Column) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewColumn() *Column { _ = "STUB: not implemented"; return nil }
