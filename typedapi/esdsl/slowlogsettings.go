package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slowlogSettings struct {
	v *types.SlowlogSettings
}

func NewSlowlogSettings() *_slowlogSettings { _ = "STUB: not implemented"; return nil }

func (s *_slowlogSettings) Level(level string) *_slowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogSettings) Reformat(reformat bool) *_slowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogSettings) Source(source int) *_slowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogSettings) Threshold(threshold types.SlowlogTresholdsVariant) *_slowlogSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slowlogSettings) SlowlogSettingsCaster() *types.SlowlogSettings {
	_ = "STUB: not implemented"
	return nil
}
