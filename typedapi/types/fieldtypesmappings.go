package types

type FieldTypesMappings struct {
	FieldTypes []FieldTypes `json:"field_types"`

	RuntimeFieldTypes []ClusterRuntimeFieldTypes `json:"runtime_field_types"`

	SourceModes map[string]int `json:"source_modes"`

	TotalDeduplicatedFieldCount *int64 `json:"total_deduplicated_field_count,omitempty"`

	TotalDeduplicatedMappingSize ByteSize `json:"total_deduplicated_mapping_size,omitempty"`

	TotalDeduplicatedMappingSizeInBytes *int64 `json:"total_deduplicated_mapping_size_in_bytes,omitempty"`

	TotalFieldCount *int64 `json:"total_field_count,omitempty"`
}

func (s *FieldTypesMappings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFieldTypesMappings() *FieldTypesMappings { _ = "STUB: not implemented"; return nil }
