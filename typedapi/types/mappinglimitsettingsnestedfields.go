package types

type MappingLimitSettingsNestedFields struct {
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsNestedFields) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettingsNestedFields() *MappingLimitSettingsNestedFields {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsNestedFieldsVariant interface {
	MappingLimitSettingsNestedFieldsCaster() *MappingLimitSettingsNestedFields
}

func (s *MappingLimitSettingsNestedFields) MappingLimitSettingsNestedFieldsCaster() *MappingLimitSettingsNestedFields {
	_ = "STUB: not implemented"
	return nil
}
