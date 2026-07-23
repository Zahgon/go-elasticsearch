package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _thinkingConfig struct {
	v *types.ThinkingConfig
}

func NewThinkingConfig() *_thinkingConfig { _ = "STUB: not implemented"; return nil }

func (s *_thinkingConfig) ThinkingBudget(thinkingbudget int) *_thinkingConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_thinkingConfig) ThinkingConfigCaster() *types.ThinkingConfig {
	_ = "STUB: not implemented"
	return nil
}
