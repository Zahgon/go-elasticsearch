package types

type XpackRuntimeFieldTypes struct {
	Available  bool                `json:"available"`
	Enabled    bool                `json:"enabled"`
	FieldTypes []RuntimeFieldsType `json:"field_types"`
}

func (s *XpackRuntimeFieldTypes) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewXpackRuntimeFieldTypes() *XpackRuntimeFieldTypes { _ = "STUB: not implemented"; return nil }
