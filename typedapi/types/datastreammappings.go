package types

type DataStreamMappings struct {
	EffectiveMappings TypeMapping `json:"effective_mappings"`

	Mappings TypeMapping `json:"mappings"`

	Name string `json:"name"`
}

func (s *DataStreamMappings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamMappings() *DataStreamMappings { _ = "STUB: not implemented"; return nil }
