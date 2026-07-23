package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/chunkingmode"
)

type _chunkingConfig struct {
	v *types.ChunkingConfig
}

func NewChunkingConfig(mode chunkingmode.ChunkingMode) *_chunkingConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingConfig) Mode(mode chunkingmode.ChunkingMode) *_chunkingConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingConfig) TimeSpan(duration types.DurationVariant) *_chunkingConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingConfig) ChunkingConfigCaster() *types.ChunkingConfig {
	_ = "STUB: not implemented"
	return nil
}
