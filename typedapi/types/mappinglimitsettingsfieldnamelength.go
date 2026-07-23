package types

type MappingLimitSettingsFieldNameLength struct {
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsFieldNameLength) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettingsFieldNameLength() *MappingLimitSettingsFieldNameLength {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsFieldNameLengthVariant interface {
	MappingLimitSettingsFieldNameLengthCaster() *MappingLimitSettingsFieldNameLength
}

func (s *MappingLimitSettingsFieldNameLength) MappingLimitSettingsFieldNameLengthCaster() *MappingLimitSettingsFieldNameLength {
	_ = "STUB: not implemented"
	return nil
}
