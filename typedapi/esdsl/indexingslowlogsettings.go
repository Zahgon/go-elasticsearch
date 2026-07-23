package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexingSlowlogSettings struct {
	v *types.IndexingSlowlogSettings
}

func NewIndexingSlowlogSettings() *_indexingSlowlogSettings { _ = "STUB: not implemented"; return nil }

func (s *_indexingSlowlogSettings) Level(level string) *_indexingSlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexingSlowlogSettings) Reformat(reformat bool) *_indexingSlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexingSlowlogSettings) Source(source int) *_indexingSlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexingSlowlogSettings) Threshold(threshold types.IndexingSlowlogTresholdsVariant) *_indexingSlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexingSlowlogSettings) IndexingSlowlogSettingsCaster() *types.IndexingSlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}
