package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingsLifecycle struct {
	v *types.IndexSettingsLifecycle
}

func NewIndexSettingsLifecycle() *_indexSettingsLifecycle { _ = "STUB: not implemented"; return nil }

func (s *_indexSettingsLifecycle) IndexingComplete(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) Name(name string) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) OriginationDate(originationdate int64) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) ParseOriginationDate(parseoriginationdate bool) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) PreferIlm(preferilm string) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) RolloverAlias(rolloveralias string) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) Step(step types.IndexSettingsLifecycleStepVariant) *_indexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycle) IndexSettingsLifecycleCaster() *types.IndexSettingsLifecycle {
	_ = "STUB: not implemented"
	return nil
}
