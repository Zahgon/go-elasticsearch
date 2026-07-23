package types

type UpdatedDataStreamMappings struct {
	AppliedToDataStream bool `json:"applied_to_data_stream"`

	EffectiveMappings *TypeMapping `json:"effective_mappings,omitempty"`

	Error *string `json:"error,omitempty"`

	Mappings *TypeMapping `json:"mappings,omitempty"`

	Name string `json:"name"`
}

func (s *UpdatedDataStreamMappings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewUpdatedDataStreamMappings() *UpdatedDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}
