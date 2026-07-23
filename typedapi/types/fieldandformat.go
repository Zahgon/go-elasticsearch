package types

type FieldAndFormat struct {
	Field string `json:"field"`

	Format          *string `json:"format,omitempty"`
	IncludeUnmapped *bool   `json:"include_unmapped,omitempty"`
}

func (s *FieldAndFormat) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldAndFormat() *FieldAndFormat { _ = "STUB: not implemented"; return nil }

type FieldAndFormatVariant interface {
	FieldAndFormatCaster() *FieldAndFormat
}

func (s *FieldAndFormat) FieldAndFormatCaster() *FieldAndFormat {
	_ = "STUB: not implemented"
	return nil
}
