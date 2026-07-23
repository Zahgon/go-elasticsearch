package types

type MappingLimitSettingsDimensionFields struct {
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsDimensionFields) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettingsDimensionFields() *MappingLimitSettingsDimensionFields {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsDimensionFieldsVariant interface {
	MappingLimitSettingsDimensionFieldsCaster() *MappingLimitSettingsDimensionFields
}

func (s *MappingLimitSettingsDimensionFields) MappingLimitSettingsDimensionFieldsCaster() *MappingLimitSettingsDimensionFields {
	_ = "STUB: not implemented"
	return nil
}
