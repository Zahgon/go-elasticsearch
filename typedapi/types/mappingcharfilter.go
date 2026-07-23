package types

type MappingCharFilter struct {
	Mappings     []string `json:"mappings,omitempty"`
	MappingsPath *string  `json:"mappings_path,omitempty"`
	Type         string   `json:"type,omitempty"`
	Version      *string  `json:"version,omitempty"`
}

func (s *MappingCharFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s MappingCharFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMappingCharFilter() *MappingCharFilter { _ = "STUB: not implemented"; return nil }

type MappingCharFilterVariant interface {
	MappingCharFilterCaster() *MappingCharFilter
}

func (s *MappingCharFilter) MappingCharFilterCaster() *MappingCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *MappingCharFilter) CharFilterDefinitionCaster() *CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
