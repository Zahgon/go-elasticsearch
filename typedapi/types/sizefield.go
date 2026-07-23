package types

type SizeField struct {
	Enabled bool `json:"enabled"`
}

func (s *SizeField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSizeField() *SizeField { _ = "STUB: not implemented"; return nil }

type SizeFieldVariant interface {
	SizeFieldCaster() *SizeField
}

func (s *SizeField) SizeFieldCaster() *SizeField { _ = "STUB: not implemented"; return nil }
