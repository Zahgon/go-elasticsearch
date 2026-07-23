package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _chunkRescorer struct {
	v *types.ChunkRescorer
}

func NewChunkRescorer() *_chunkRescorer { _ = "STUB: not implemented"; return nil }

func (s *_chunkRescorer) ChunkingSettings(chunkingsettings types.ChunkRescorerChunkingSettingsVariant) *_chunkRescorer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkRescorer) Size(size int) *_chunkRescorer { _ = "STUB: not implemented"; return nil }

func (s *_chunkRescorer) ChunkRescorerCaster() *types.ChunkRescorer {
	_ = "STUB: not implemented"
	return nil
}
