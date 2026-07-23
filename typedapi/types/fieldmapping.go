package types

type FieldMapping struct {
	FullName string              `json:"full_name"`
	Mapping  map[string]Property `json:"mapping"`
}

func (s *FieldMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldMapping() *FieldMapping { _ = "STUB: not implemented"; return nil }
