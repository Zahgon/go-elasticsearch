package types

type FieldNamesField struct {
	Enabled bool `json:"enabled"`
}

func (s *FieldNamesField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldNamesField() *FieldNamesField { _ = "STUB: not implemented"; return nil }

type FieldNamesFieldVariant interface {
	FieldNamesFieldCaster() *FieldNamesField
}

func (s *FieldNamesField) FieldNamesFieldCaster() *FieldNamesField {
	_ = "STUB: not implemented"
	return nil
}
