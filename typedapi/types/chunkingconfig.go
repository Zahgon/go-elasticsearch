package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/chunkingmode"
)

type ChunkingConfig struct {
	Mode chunkingmode.ChunkingMode `json:"mode"`

	TimeSpan Duration `json:"time_span,omitempty"`
}

func (s *ChunkingConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewChunkingConfig() *ChunkingConfig { _ = "STUB: not implemented"; return nil }

type ChunkingConfigVariant interface {
	ChunkingConfigCaster() *ChunkingConfig
}

func (s *ChunkingConfig) ChunkingConfigCaster() *ChunkingConfig {
	_ = "STUB: not implemented"
	return nil
}
