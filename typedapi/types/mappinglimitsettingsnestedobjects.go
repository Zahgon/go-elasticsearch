package types

type MappingLimitSettingsNestedObjects struct {
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsNestedObjects) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettingsNestedObjects() *MappingLimitSettingsNestedObjects {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsNestedObjectsVariant interface {
	MappingLimitSettingsNestedObjectsCaster() *MappingLimitSettingsNestedObjects
}

func (s *MappingLimitSettingsNestedObjects) MappingLimitSettingsNestedObjectsCaster() *MappingLimitSettingsNestedObjects {
	_ = "STUB: not implemented"
	return nil
}
