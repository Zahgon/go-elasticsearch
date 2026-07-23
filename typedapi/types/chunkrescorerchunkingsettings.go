package types

type ChunkRescorerChunkingSettings struct {
	MaxChunkSize int `json:"max_chunk_size"`

	Overlap *int `json:"overlap,omitempty"`

	SentenceOverlap *int `json:"sentence_overlap,omitempty"`

	SeparatorGroup *string `json:"separator_group,omitempty"`

	Separators []string `json:"separators,omitempty"`

	Strategy *string `json:"strategy,omitempty"`
}

func (s *ChunkRescorerChunkingSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewChunkRescorerChunkingSettings() *ChunkRescorerChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

type ChunkRescorerChunkingSettingsVariant interface {
	ChunkRescorerChunkingSettingsCaster() *ChunkRescorerChunkingSettings
}

func (s *ChunkRescorerChunkingSettings) ChunkRescorerChunkingSettingsCaster() *ChunkRescorerChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}
