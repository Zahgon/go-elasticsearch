package types

type MappingLimitSettingsTotalFields struct {
	IgnoreDynamicBeyondLimit *string `json:"ignore_dynamic_beyond_limit,omitempty"`

	Limit *string `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsTotalFields) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettingsTotalFields() *MappingLimitSettingsTotalFields {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsTotalFieldsVariant interface {
	MappingLimitSettingsTotalFieldsCaster() *MappingLimitSettingsTotalFields
}

func (s *MappingLimitSettingsTotalFields) MappingLimitSettingsTotalFieldsCaster() *MappingLimitSettingsTotalFields {
	_ = "STUB: not implemented"
	return nil
}
