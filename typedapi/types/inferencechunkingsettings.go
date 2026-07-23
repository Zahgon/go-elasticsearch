package types

type InferenceChunkingSettings struct {
	MaxChunkSize *int `json:"max_chunk_size,omitempty"`

	Overlap *int `json:"overlap,omitempty"`

	SentenceOverlap *int `json:"sentence_overlap,omitempty"`

	SeparatorGroup *string `json:"separator_group,omitempty"`

	Separators []string `json:"separators,omitempty"`

	Strategy *string `json:"strategy,omitempty"`
}

func (s *InferenceChunkingSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceChunkingSettings() *InferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

type InferenceChunkingSettingsVariant interface {
	InferenceChunkingSettingsCaster() *InferenceChunkingSettings
}

func (s *InferenceChunkingSettings) InferenceChunkingSettingsCaster() *InferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}
