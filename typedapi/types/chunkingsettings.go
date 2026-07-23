package types

type ChunkingSettings struct {
	MaxChunkSize int `json:"max_chunk_size"`

	Overlap *int `json:"overlap,omitempty"`

	SentenceOverlap *int `json:"sentence_overlap,omitempty"`

	SeparatorGroup *string `json:"separator_group,omitempty"`

	Separators []string `json:"separators,omitempty"`

	Strategy string `json:"strategy"`
}

func (s *ChunkingSettings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewChunkingSettings() *ChunkingSettings { _ = "STUB: not implemented"; return nil }

type ChunkingSettingsVariant interface {
	ChunkingSettingsCaster() *ChunkingSettings
}

func (s *ChunkingSettings) ChunkingSettingsCaster() *ChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}
