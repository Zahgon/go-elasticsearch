package types

type MappingLimitSettings struct {
	Coerce          *bool                                `json:"coerce,omitempty"`
	Depth           *MappingLimitSettingsDepth           `json:"depth,omitempty"`
	DimensionFields *MappingLimitSettingsDimensionFields `json:"dimension_fields,omitempty"`
	FieldNameLength *MappingLimitSettingsFieldNameLength `json:"field_name_length,omitempty"`
	IgnoreMalformed *string                              `json:"ignore_malformed,omitempty"`
	NestedFields    *MappingLimitSettingsNestedFields    `json:"nested_fields,omitempty"`
	NestedObjects   *MappingLimitSettingsNestedObjects   `json:"nested_objects,omitempty"`
	Source          *MappingLimitSettingsSourceFields    `json:"source,omitempty"`
	TotalFields     *MappingLimitSettingsTotalFields     `json:"total_fields,omitempty"`
}

func (s *MappingLimitSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettings() *MappingLimitSettings { _ = "STUB: not implemented"; return nil }

type MappingLimitSettingsVariant interface {
	MappingLimitSettingsCaster() *MappingLimitSettings
}

func (s *MappingLimitSettings) MappingLimitSettingsCaster() *MappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}
