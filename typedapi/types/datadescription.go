package types

type DataDescription struct {
	FieldDelimiter *string `json:"field_delimiter,omitempty"`

	Format *string `json:"format,omitempty"`

	TimeField *string `json:"time_field,omitempty"`

	TimeFormat *string `json:"time_format,omitempty"`
}

func (s *DataDescription) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataDescription() *DataDescription { _ = "STUB: not implemented"; return nil }

type DataDescriptionVariant interface {
	DataDescriptionCaster() *DataDescription
}

func (s *DataDescription) DataDescriptionCaster() *DataDescription {
	_ = "STUB: not implemented"
	return nil
}
