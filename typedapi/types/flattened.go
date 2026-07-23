package types

type Flattened struct {
	Available  bool `json:"available"`
	Enabled    bool `json:"enabled"`
	FieldCount int  `json:"field_count"`
}

func (s *Flattened) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFlattened() *Flattened { _ = "STUB: not implemented"; return nil }
