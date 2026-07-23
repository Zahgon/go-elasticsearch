package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settingsAnalyze struct {
	v *types.SettingsAnalyze
}

func NewSettingsAnalyze() *_settingsAnalyze { _ = "STUB: not implemented"; return nil }

func (s *_settingsAnalyze) MaxTokenCount(stringifiedinteger types.StringifiedintegerVariant) *_settingsAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsAnalyze) SettingsAnalyzeCaster() *types.SettingsAnalyze {
	_ = "STUB: not implemented"
	return nil
}
