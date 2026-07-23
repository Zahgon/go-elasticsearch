package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingsLifecycleStep struct {
	v *types.IndexSettingsLifecycleStep
}

func NewIndexSettingsLifecycleStep() *_indexSettingsLifecycleStep {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycleStep) WaitTimeThreshold(duration types.DurationVariant) *_indexSettingsLifecycleStep {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsLifecycleStep) IndexSettingsLifecycleStepCaster() *types.IndexSettingsLifecycleStep {
	_ = "STUB: not implemented"
	return nil
}
