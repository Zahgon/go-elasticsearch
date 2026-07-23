package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _filteringConfig struct {
	v *types.FilteringConfig
}

func NewFilteringConfig(active types.FilteringRulesVariant, draft types.FilteringRulesVariant) *_filteringConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringConfig) Active(active types.FilteringRulesVariant) *_filteringConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringConfig) Domain(domain string) *_filteringConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringConfig) Draft(draft types.FilteringRulesVariant) *_filteringConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringConfig) FilteringConfigCaster() *types.FilteringConfig {
	_ = "STUB: not implemented"
	return nil
}
