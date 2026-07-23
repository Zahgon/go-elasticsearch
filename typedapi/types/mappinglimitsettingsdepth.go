package types

type MappingLimitSettingsDepth struct {
	Limit *int64 `json:"limit,omitempty"`
}

func (s *MappingLimitSettingsDepth) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMappingLimitSettingsDepth() *MappingLimitSettingsDepth {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsDepthVariant interface {
	MappingLimitSettingsDepthCaster() *MappingLimitSettingsDepth
}

func (s *MappingLimitSettingsDepth) MappingLimitSettingsDepthCaster() *MappingLimitSettingsDepth {
	_ = "STUB: not implemented"
	return nil
}
