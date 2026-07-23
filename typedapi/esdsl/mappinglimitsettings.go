package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mappingLimitSettings struct {
	v *types.MappingLimitSettings
}

func NewMappingLimitSettings() *_mappingLimitSettings { _ = "STUB: not implemented"; return nil }

func (s *_mappingLimitSettings) Coerce(coerce bool) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) Depth(depth types.MappingLimitSettingsDepthVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) DimensionFields(dimensionfields types.MappingLimitSettingsDimensionFieldsVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) FieldNameLength(fieldnamelength types.MappingLimitSettingsFieldNameLengthVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) IgnoreMalformed(ignoremalformed string) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) NestedFields(nestedfields types.MappingLimitSettingsNestedFieldsVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) NestedObjects(nestedobjects types.MappingLimitSettingsNestedObjectsVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) Source(source types.MappingLimitSettingsSourceFieldsVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) TotalFields(totalfields types.MappingLimitSettingsTotalFieldsVariant) *_mappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettings) MappingLimitSettingsCaster() *types.MappingLimitSettings {
	_ = "STUB: not implemented"
	return nil
}
