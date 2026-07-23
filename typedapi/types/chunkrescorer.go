package types

type ChunkRescorer struct {
	ChunkingSettings *ChunkRescorerChunkingSettings `json:"chunking_settings,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *ChunkRescorer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewChunkRescorer() *ChunkRescorer { _ = "STUB: not implemented"; return nil }

type ChunkRescorerVariant interface {
	ChunkRescorerCaster() *ChunkRescorer
}

func (s *ChunkRescorer) ChunkRescorerCaster() *ChunkRescorer { _ = "STUB: not implemented"; return nil }
