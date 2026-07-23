package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _syncRulesFeature struct {
	v *types.SyncRulesFeature
}

func NewSyncRulesFeature() *_syncRulesFeature { _ = "STUB: not implemented"; return nil }

func (s *_syncRulesFeature) Advanced(advanced types.FeatureEnabledVariant) *_syncRulesFeature {
	_ = "STUB: not implemented"
	return nil
}

func (s *_syncRulesFeature) Basic(basic types.FeatureEnabledVariant) *_syncRulesFeature {
	_ = "STUB: not implemented"
	return nil
}

func (s *_syncRulesFeature) SyncRulesFeatureCaster() *types.SyncRulesFeature {
	_ = "STUB: not implemented"
	return nil
}
