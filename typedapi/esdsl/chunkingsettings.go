package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _chunkingSettings struct {
	v *types.ChunkingSettings
}

func NewChunkingSettings(maxchunksize int, strategy string) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) MaxChunkSize(maxchunksize int) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) Overlap(overlap int) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) SentenceOverlap(sentenceoverlap int) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) SeparatorGroup(separatorgroup string) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) Separators(separators ...string) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) Strategy(strategy string) *_chunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_chunkingSettings) ChunkingSettingsCaster() *types.ChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}
