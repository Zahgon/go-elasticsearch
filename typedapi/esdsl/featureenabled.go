package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _featureEnabled struct {
	v *types.FeatureEnabled
}

func NewFeatureEnabled(enabled bool) *_featureEnabled { _ = "STUB: not implemented"; return nil }

func (s *_featureEnabled) Enabled(enabled bool) *_featureEnabled {
	_ = "STUB: not implemented"
	return nil
}

func (s *_featureEnabled) FeatureEnabledCaster() *types.FeatureEnabled {
	_ = "STUB: not implemented"
	return nil
}
