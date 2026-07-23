package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _customTaskSettings struct {
	v *types.CustomTaskSettings
}

func NewCustomTaskSettings() *_customTaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_customTaskSettings) Parameters(parameters map[string]types.CustomTaskParameter) *_customTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customTaskSettings) AddParameter(key string, value types.CustomTaskParameterVariant) *_customTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customTaskSettings) CustomTaskSettingsCaster() *types.CustomTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
