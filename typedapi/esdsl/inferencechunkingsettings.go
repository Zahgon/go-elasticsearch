package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceChunkingSettings struct {
	v *types.InferenceChunkingSettings
}

func NewInferenceChunkingSettings() *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) MaxChunkSize(maxchunksize int) *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) Overlap(overlap int) *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) SentenceOverlap(sentenceoverlap int) *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) SeparatorGroup(separatorgroup string) *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) Separators(separators ...string) *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) Strategy(strategy string) *_inferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceChunkingSettings) InferenceChunkingSettingsCaster() *types.InferenceChunkingSettings {
	_ = "STUB: not implemented"
	return nil
}
